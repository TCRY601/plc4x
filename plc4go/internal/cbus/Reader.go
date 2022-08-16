/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package cbus

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/apache/plc4x/plc4go/pkg/api/model"
	"github.com/apache/plc4x/plc4go/pkg/api/values"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi"
	spiModel "github.com/apache/plc4x/plc4go/spi/model"
	spiValues "github.com/apache/plc4x/plc4go/spi/values"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Reader struct {
	alphaGenerator *AlphaGenerator
	messageCodec   spi.MessageCodec
	tm             *spi.RequestTransactionManager
}

func NewReader(tpduGenerator *AlphaGenerator, messageCodec spi.MessageCodec, tm *spi.RequestTransactionManager) *Reader {
	return &Reader{
		alphaGenerator: tpduGenerator,
		messageCodec:   messageCodec,
		tm:             tm,
	}
}

func (m *Reader) Read(ctx context.Context, readRequest model.PlcReadRequest) <-chan model.PlcReadRequestResult {
	// TODO: handle ctx
	log.Trace().Msg("Reading")
	result := make(chan model.PlcReadRequestResult)
	go func() {
		numFields := len(readRequest.GetFieldNames())
		if numFields > 20 {
			result <- &spiModel.DefaultPlcReadRequestResult{
				Request:  readRequest,
				Response: nil,
				Err:      errors.New("Only 20 fields can be handled at once"),
			}
			return
		}
		messages := make(map[string]readWriteModel.CBusMessage)
		for _, fieldName := range readRequest.GetFieldNames() {
			field := readRequest.GetField(fieldName)
			message, err := m.fieldToCBusMessage(field)
			if err != nil {
				result <- &spiModel.DefaultPlcReadRequestResult{
					Request:  readRequest,
					Response: nil,
					Err:      errors.Wrapf(err, "Error encoding cbus message for field %s", fieldName),
				}
				return
			}
			messages[fieldName] = message
		}
		responseMu := sync.Mutex{}
		responseCodes := map[string]model.PlcResponseCode{}
		addResponseCode := func(name string, responseCode model.PlcResponseCode) {
			responseMu.Lock()
			defer responseMu.Unlock()
			responseCodes[name] = responseCode
		}
		valueMu := sync.Mutex{}
		plcValues := map[string]values.PlcValue{}
		addPlcValue := func(name string, plcValue values.PlcValue) {
			valueMu.Lock()
			defer valueMu.Unlock()
			plcValues[name] = plcValue
		}
		for fieldName, messageToSend := range messages {
			fieldNameCopy := fieldName
			// Start a new request-transaction (Is ended in the response-handler)
			requestWasOk := make(chan bool)
			transaction := m.tm.StartTransaction()
			transaction.Submit(func() {

				// Send the  over the wire
				log.Trace().Msg("Send ")
				if err := m.messageCodec.SendRequest(ctx, messageToSend, func(receivedMessage spi.Message) bool {
					cbusMessage, ok := receivedMessage.(readWriteModel.CBusMessageExactly)
					if !ok {
						return false
					}
					messageToClient, ok := cbusMessage.(readWriteModel.CBusMessageToClientExactly)
					if !ok {
						return false
					}
					// Check if this errored
					if _, ok = messageToClient.GetReply().(readWriteModel.ServerErrorReplyExactly); ok {
						// This means we must handle this below
						return true
					}

					confirmation, ok := messageToClient.GetReply().(readWriteModel.ReplyOrConfirmationConfirmationExactly)
					if !ok {
						return false
					}
					return confirmation.GetConfirmation().GetAlpha().GetCharacter() == messageToSend.(readWriteModel.CBusMessageToServer).GetRequest().(readWriteModel.RequestCommand).GetAlpha().GetCharacter()
				}, func(receivedMessage spi.Message) error {
					// Convert the response into an
					log.Trace().Msg("convert response to ")
					cbusMessage := receivedMessage.(readWriteModel.CBusMessage)
					messageToClient := cbusMessage.(readWriteModel.CBusMessageToClient)
					if _, ok := messageToClient.GetReply().(readWriteModel.ServerErrorReplyExactly); ok {
						log.Trace().Msg("We got a server failure")
						addResponseCode(fieldNameCopy, model.PlcResponseCode_INVALID_DATA)
						requestWasOk <- false
						return transaction.EndRequest()
					}
					replyOrConfirmationConfirmation := messageToClient.GetReply().(readWriteModel.ReplyOrConfirmationConfirmationExactly)
					if !replyOrConfirmationConfirmation.GetConfirmation().GetIsSuccess() {
						var responseCode model.PlcResponseCode
						switch replyOrConfirmationConfirmation.GetConfirmation().GetConfirmationType() {
						case readWriteModel.ConfirmationType_NOT_TRANSMITTED_TO_MANY_RE_TRANSMISSIONS:
							responseCode = model.PlcResponseCode_REMOTE_ERROR
						case readWriteModel.ConfirmationType_NOT_TRANSMITTED_CORRUPTION:
							responseCode = model.PlcResponseCode_INVALID_DATA
						case readWriteModel.ConfirmationType_NOT_TRANSMITTED_SYNC_LOSS:
							responseCode = model.PlcResponseCode_REMOTE_BUSY
						case readWriteModel.ConfirmationType_NOT_TRANSMITTED_TOO_LONG:
							responseCode = model.PlcResponseCode_INVALID_DATA
						default:
							panic("Every code should be mapped here")
						}
						log.Trace().Msgf("Was no success %s:%v", fieldNameCopy, responseCode)
						addResponseCode(fieldNameCopy, responseCode)
						requestWasOk <- true
						return transaction.EndRequest()
					}

					alpha := replyOrConfirmationConfirmation.GetConfirmation().GetAlpha()
					// TODO: it could be double confirmed but this is not implemented yet
					embeddedReply, ok := replyOrConfirmationConfirmation.GetEmbeddedReply().(readWriteModel.ReplyOrConfirmationReplyExactly)
					if !ok {
						log.Trace().Msgf("Is a confirm only, no data. Alpha: %c", alpha.GetCharacter())
						addResponseCode(fieldNameCopy, model.PlcResponseCode_NOT_FOUND)
						requestWasOk <- true
						return transaction.EndRequest()
					}

					log.Trace().Msg("Handling confirmed data")
					// TODO: check if we can use a plcValueSerializer
					switch reply := embeddedReply.GetReply().(readWriteModel.ReplyEncodedReply).GetEncodedReply().(type) {
					case readWriteModel.EncodedReplyCALReplyExactly:
						calData := reply.GetCalReply().GetCalData()
						addResponseCode(fieldNameCopy, model.PlcResponseCode_OK)
						switch calData := calData.(type) {
						case readWriteModel.CALDataStatusExactly:
							application := calData.GetApplication()
							// TODO: verify application... this should be the same
							_ = application
							blockStart := calData.GetBlockStart()
							// TODO: verify application... this should be the same
							_ = blockStart
							statusBytes := calData.GetStatusBytes()
							addResponseCode(fieldNameCopy, model.PlcResponseCode_OK)
							plcListValues := make([]values.PlcValue, len(statusBytes)*4)
							for i, statusByte := range statusBytes {
								plcListValues[i*4+0] = spiValues.NewPlcSTRING(statusByte.GetGav0().String())
								plcListValues[i*4+1] = spiValues.NewPlcSTRING(statusByte.GetGav1().String())
								plcListValues[i*4+2] = spiValues.NewPlcSTRING(statusByte.GetGav2().String())
								plcListValues[i*4+3] = spiValues.NewPlcSTRING(statusByte.GetGav3().String())
							}
							addPlcValue(fieldNameCopy, spiValues.NewPlcList(plcListValues))
						case readWriteModel.CALDataStatusExtendedExactly:
							coding := calData.GetCoding()
							// TODO: verify coding... this should be the same
							_ = coding
							application := calData.GetApplication()
							// TODO: verify application... this should be the same
							_ = application
							blockStart := calData.GetBlockStart()
							// TODO: verify application... this should be the same
							_ = blockStart
							switch coding {
							case readWriteModel.StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE:
								fallthrough
							case readWriteModel.StatusCoding_BINARY_BY_ELSEWHERE:
								statusBytes := calData.GetStatusBytes()
								addResponseCode(fieldNameCopy, model.PlcResponseCode_OK)
								plcListValues := make([]values.PlcValue, len(statusBytes)*4)
								for i, statusByte := range statusBytes {
									plcListValues[i*4+0] = spiValues.NewPlcSTRING(statusByte.GetGav0().String())
									plcListValues[i*4+1] = spiValues.NewPlcSTRING(statusByte.GetGav1().String())
									plcListValues[i*4+2] = spiValues.NewPlcSTRING(statusByte.GetGav2().String())
									plcListValues[i*4+3] = spiValues.NewPlcSTRING(statusByte.GetGav3().String())
								}
								addPlcValue(fieldNameCopy, spiValues.NewPlcList(plcListValues))
							case readWriteModel.StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE:
								fallthrough
							case readWriteModel.StatusCoding_LEVEL_BY_ELSEWHERE:
								levelInformation := calData.GetLevelInformation()
								addResponseCode(fieldNameCopy, model.PlcResponseCode_OK)
								plcListValues := make([]values.PlcValue, len(levelInformation))
								for i, levelInformation := range levelInformation {
									switch levelInformation := levelInformation.(type) {
									case readWriteModel.LevelInformationAbsentExactly:
										plcListValues[i] = spiValues.NewPlcSTRING("is absent")
									case readWriteModel.LevelInformationCorruptedExactly:
										plcListValues[i] = spiValues.NewPlcSTRING("corrupted")
									case readWriteModel.LevelInformationNormalExactly:
										plcListValues[i] = spiValues.NewPlcUSINT(levelInformation.GetActualLevel())
									default:
										panic("Impossible case")
									}
								}
								addPlcValue(fieldNameCopy, spiValues.NewPlcList(plcListValues))
							}
						case readWriteModel.CALDataIdentifyReplyExactly:
							switch identifyReplyCommand := calData.GetIdentifyReplyCommand().(type) {
							case readWriteModel.IdentifyReplyCommandCurrentSenseLevelsExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetCurrentSenseLevels()))
							case readWriteModel.IdentifyReplyCommandDelaysExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcStruct(map[string]values.PlcValue{
									"ReStrikeDelay": spiValues.NewPlcUINT(uint16(identifyReplyCommand.GetReStrikeDelay())),
									"TerminalLevel": spiValues.NewPlcByteArray(identifyReplyCommand.GetTerminalLevels()),
								}))
							case readWriteModel.IdentifyReplyCommandDSIStatusExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcStruct(map[string]values.PlcValue{
									"ChannelStatus1":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus1().String()),
									"ChannelStatus2":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus2().String()),
									"ChannelStatus3":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus3().String()),
									"ChannelStatus4":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus4().String()),
									"ChannelStatus5":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus5().String()),
									"ChannelStatus6":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus6().String()),
									"ChannelStatus7":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus7().String()),
									"ChannelStatus8":          spiValues.NewPlcSTRING(identifyReplyCommand.GetChannelStatus8().String()),
									"UnitStatus":              spiValues.NewPlcSTRING(identifyReplyCommand.GetUnitStatus().String()),
									"DimmingUCRevisionNumber": spiValues.NewPlcUINT(uint16(identifyReplyCommand.GetDimmingUCRevisionNumber())),
								}))
							case readWriteModel.IdentifyReplyCommandExtendedDiagnosticSummaryExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcStruct(map[string]values.PlcValue{
									"LowApplication":         spiValues.NewPlcSTRING(identifyReplyCommand.GetLowApplication().String()),
									"HighApplication":        spiValues.NewPlcSTRING(identifyReplyCommand.GetHighApplication().String()),
									"Area":                   spiValues.NewPlcUINT(uint16(identifyReplyCommand.GetArea())),
									"Crc":                    spiValues.NewPlcUINT(identifyReplyCommand.GetCrc()),
									"SerialNumber":           spiValues.NewPlcUDINT(identifyReplyCommand.GetSerialNumber()),
									"NetworkVoltage":         spiValues.NewPlcUINT(uint16(identifyReplyCommand.GetNetworkVoltage())),
									"UnitInLearnMode":        spiValues.NewPlcBOOL(identifyReplyCommand.GetUnitInLearnMode()),
									"NetworkVoltageLow":      spiValues.NewPlcBOOL(identifyReplyCommand.GetNetworkVoltageLow()),
									"NetworkVoltageMarginal": spiValues.NewPlcBOOL(identifyReplyCommand.GetNetworkVoltageMarginal()),
									"EnableChecksumAlarm":    spiValues.NewPlcBOOL(identifyReplyCommand.GetEnableChecksumAlarm()),
									"OutputUnit":             spiValues.NewPlcBOOL(identifyReplyCommand.GetOutputUnit()),
									"InstallationMMIError":   spiValues.NewPlcBOOL(identifyReplyCommand.GetInstallationMMIError()),
									"EEWriteError":           spiValues.NewPlcBOOL(identifyReplyCommand.GetEEWriteError()),
									"EEChecksumError":        spiValues.NewPlcBOOL(identifyReplyCommand.GetEEChecksumError()),
									"EEDataError":            spiValues.NewPlcBOOL(identifyReplyCommand.GetEEDataError()),
									"MicroReset":             spiValues.NewPlcBOOL(identifyReplyCommand.GetMicroReset()),
									"CommsTxError":           spiValues.NewPlcBOOL(identifyReplyCommand.GetCommsTxError()),
									"InternalStackOverflow":  spiValues.NewPlcBOOL(identifyReplyCommand.GetInternalStackOverflow()),
									"MicroPowerReset":        spiValues.NewPlcBOOL(identifyReplyCommand.GetMicroPowerReset()),
								}))
							case readWriteModel.IdentifyReplyCommandSummaryExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcStruct(map[string]values.PlcValue{
									"PartName":        spiValues.NewPlcSTRING(identifyReplyCommand.GetPartName()),
									"UnitServiceType": spiValues.NewPlcUINT(uint16(identifyReplyCommand.GetUnitServiceType())),
									"Version":         spiValues.NewPlcSTRING(identifyReplyCommand.GetVersion()),
								}))
							case readWriteModel.IdentifyReplyCommandFirmwareVersionExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcSTRING(identifyReplyCommand.GetFirmwareVersion()))
							case readWriteModel.IdentifyReplyCommandGAVPhysicalAddressesExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetValues()))
							case readWriteModel.IdentifyReplyCommandGAVValuesCurrentExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetValues()))
							case readWriteModel.IdentifyReplyCommandGAVValuesStoredExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetValues()))
							case readWriteModel.IdentifyReplyCommandLogicalAssignmentExactly:
								var plcValues []values.PlcValue
								for _, logicAssigment := range identifyReplyCommand.GetLogicAssigment() {
									plcValues = append(plcValues, spiValues.NewPlcStruct(map[string]values.PlcValue{
										"GreaterOfOrLogic": spiValues.NewPlcBOOL(logicAssigment.GetGreaterOfOrLogic()),
										"ReStrikeDelay":    spiValues.NewPlcBOOL(logicAssigment.GetReStrikeDelay()),
										"AssignedToGav16":  spiValues.NewPlcBOOL(logicAssigment.GetAssignedToGav16()),
										"AssignedToGav15":  spiValues.NewPlcBOOL(logicAssigment.GetAssignedToGav15()),
										"AssignedToGav14":  spiValues.NewPlcBOOL(logicAssigment.GetAssignedToGav14()),
										"AssignedToGav13":  spiValues.NewPlcBOOL(logicAssigment.GetAssignedToGav13()),
									}))
								}
								addPlcValue(fieldNameCopy, spiValues.NewPlcList(plcValues))
							case readWriteModel.IdentifyReplyCommandManufacturerExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcSTRING(identifyReplyCommand.GetManufacturerName()))
							case readWriteModel.IdentifyReplyCommandMaximumLevelsExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetMaximumLevels()))
							case readWriteModel.IdentifyReplyCommandMinimumLevelsExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetMinimumLevels()))
							case readWriteModel.IdentifyReplyCommandNetworkTerminalLevelsExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetNetworkTerminalLevels()))
							case readWriteModel.IdentifyReplyCommandNetworkVoltageExactly:
								volts := identifyReplyCommand.GetVolts()
								voltsFloat, err := strconv.ParseFloat(volts, 0)
								if err != nil {
									return errors.Wrap(err, "Error parsing volts")
								}
								voltsDecimalPlace := identifyReplyCommand.GetVoltsDecimalPlace()
								voltsDecimalPlaceFloat, err := strconv.ParseFloat(voltsDecimalPlace, 0)
								if err != nil {
									return errors.Wrap(err, "Error parsing volts decimal place")
								}
								voltsFloat += voltsDecimalPlaceFloat / 10
								addPlcValue(fieldNameCopy, spiValues.NewPlcLREAL(voltsFloat))
							case readWriteModel.IdentifyReplyCommandOutputUnitSummaryExactly:
								unitFlags := identifyReplyCommand.GetUnitFlags()
								structContent := map[string]values.PlcValue{
									"UnitFlags": spiValues.NewPlcStruct(map[string]values.PlcValue{
										"AssertingNetworkBurden": spiValues.NewPlcBOOL(unitFlags.GetAssertingNetworkBurden()),
										"RestrikeTimingActive":   spiValues.NewPlcBOOL(unitFlags.GetRestrikeTimingActive()),
										"RemoteOFFInputAsserted": spiValues.NewPlcBOOL(unitFlags.GetRemoteOFFInputAsserted()),
										"RemoteONInputAsserted":  spiValues.NewPlcBOOL(unitFlags.GetRemoteONInputAsserted()),
										"LocalToggleEnabled":     spiValues.NewPlcBOOL(unitFlags.GetLocalToggleEnabled()),
										"LocalToggleActiveState": spiValues.NewPlcBOOL(unitFlags.GetLocalToggleActiveState()),
										"ClockGenerationEnabled": spiValues.NewPlcBOOL(unitFlags.GetClockGenerationEnabled()),
										"UnitGeneratingClock":    spiValues.NewPlcBOOL(unitFlags.GetUnitGeneratingClock()),
									}),
									"TimeFromLastRecoverOfMainsInSeconds": spiValues.NewPlcUINT(uint16(identifyReplyCommand.GetTimeFromLastRecoverOfMainsInSeconds())),
								}
								if gavStoreEnabledByte1 := identifyReplyCommand.GetGavStoreEnabledByte1(); gavStoreEnabledByte1 != nil {
									structContent["GavStoreEnabledByte1"] = spiValues.NewPlcUINT(uint16(*gavStoreEnabledByte1))
								}
								if gavStoreEnabledByte2 := identifyReplyCommand.GetGavStoreEnabledByte2(); gavStoreEnabledByte2 != nil {
									structContent["GavStoreEnabledByte2"] = spiValues.NewPlcUINT(uint16(*gavStoreEnabledByte2))
								}
								addPlcValue(fieldNameCopy, spiValues.NewPlcStruct(structContent))
							case readWriteModel.IdentifyReplyCommandTerminalLevelsExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcByteArray(identifyReplyCommand.GetTerminalLevels()))
							case readWriteModel.IdentifyReplyCommandTypeExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcSTRING(identifyReplyCommand.GetUnitType()))
							case readWriteModel.IdentifyReplyCommandUnitSummaryExactly:
								addPlcValue(fieldNameCopy, spiValues.NewPlcStruct(map[string]values.PlcValue{
									"AssertingNetworkBurden": spiValues.NewPlcBOOL(identifyReplyCommand.GetAssertingNetworkBurden()),
									"RestrikeTimingActive":   spiValues.NewPlcBOOL(identifyReplyCommand.GetRestrikeTimingActive()),
									"RemoteOFFInputAsserted": spiValues.NewPlcBOOL(identifyReplyCommand.GetRemoteOFFInputAsserted()),
									"RemoteONInputAsserted":  spiValues.NewPlcBOOL(identifyReplyCommand.GetRemoteONInputAsserted()),
									"LocalToggleEnabled":     spiValues.NewPlcBOOL(identifyReplyCommand.GetLocalToggleEnabled()),
									"LocalToggleActiveState": spiValues.NewPlcBOOL(identifyReplyCommand.GetLocalToggleActiveState()),
									"ClockGenerationEnabled": spiValues.NewPlcBOOL(identifyReplyCommand.GetClockGenerationEnabled()),
									"UnitGeneratingClock":    spiValues.NewPlcBOOL(identifyReplyCommand.GetUnitGeneratingClock()),
								}))
							default:
								log.Error().Msgf("Unmapped type %T", identifyReplyCommand)
								requestWasOk <- false
								return transaction.EndRequest()
							}
						default:
							log.Warn().Msgf("Unmapped cal data type %T. Returning raw to string", calData)
							addPlcValue(fieldNameCopy, spiValues.NewPlcSTRING(fmt.Sprintf("%s", calData)))
						}
					default:
						panic(fmt.Sprintf("All types should be mapped here. Not mapped: %T", reply))
					}
					requestWasOk <- true
					return transaction.EndRequest()
				}, func(err error) error {
					log.Debug().Msgf("Error waiting for field %s", fieldNameCopy)
					addResponseCode(fieldNameCopy, model.PlcResponseCode_REQUEST_TIMEOUT)
					// TODO: ok or not ok?
					requestWasOk <- true
					return transaction.EndRequest()
				}, time.Second*1); err != nil {
					log.Debug().Err(err).Msgf("Error sending message for field %s", fieldNameCopy)
					addResponseCode(fieldNameCopy, model.PlcResponseCode_INTERNAL_ERROR)
					_ = transaction.EndRequest()
					requestWasOk <- false
				}
			})
			if !<-requestWasOk {
				// TODO: if we found a error we can abort
				break
			}
		}
		readResponse := spiModel.NewDefaultPlcReadResponse(readRequest, responseCodes, plcValues)
		result <- &spiModel.DefaultPlcReadRequestResult{
			Request:  readRequest,
			Response: readResponse,
		}
	}()
	return result
}

func (m *Reader) fieldToCBusMessage(field model.PlcField) (readWriteModel.CBusMessage, error) {
	cbusOptions := m.messageCodec.(*MessageCodec).cbusOptions
	requestContext := m.messageCodec.(*MessageCodec).requestContext
	switch field := field.(type) {
	case *statusField:
		var statusRequest readWriteModel.StatusRequest
		switch field.statusRequestType {
		case StatusRequestTypeBinaryState:
			statusRequest = readWriteModel.NewStatusRequestBinaryState(field.application, 0x7A)
		case StatusRequestTypeLevel:
			statusRequest = readWriteModel.NewStatusRequestLevel(field.application, *field.startingGroupAddressLabel, 0x73)
		}
		command := readWriteModel.NewCBusPointToMultiPointCommandStatus(statusRequest, byte(field.application), cbusOptions)
		header := readWriteModel.NewCBusHeader(readWriteModel.PriorityClass_Class4, false, 0, readWriteModel.DestinationAddressType_PointToMultiPoint)
		cbusCommand := readWriteModel.NewCBusCommandPointToMultiPoint(command, header, cbusOptions)
		request := readWriteModel.NewRequestCommand(cbusCommand, nil, readWriteModel.NewAlpha(m.alphaGenerator.getAndIncrement()), readWriteModel.RequestType_REQUEST_COMMAND, nil, nil, readWriteModel.RequestType_EMPTY, readWriteModel.NewRequestTermination(), cbusOptions)
		return readWriteModel.NewCBusMessageToServer(request, requestContext, cbusOptions), nil
	case *calRecallField:
		calData := readWriteModel.NewCALDataRecall(field.parameter, field.count, readWriteModel.CALCommandTypeContainer_CALCommandRecall, nil, requestContext)
		//TODO: we need support for bridged commands
		command := readWriteModel.NewCBusPointToPointCommandDirect(field.unitAddress, 0x0000, calData, cbusOptions)
		header := readWriteModel.NewCBusHeader(readWriteModel.PriorityClass_Class4, false, 0, readWriteModel.DestinationAddressType_PointToPoint)
		cbusCommand := readWriteModel.NewCBusCommandPointToPoint(command, header, cbusOptions)
		request := readWriteModel.NewRequestCommand(cbusCommand, nil, readWriteModel.NewAlpha(m.alphaGenerator.getAndIncrement()), readWriteModel.RequestType_REQUEST_COMMAND, nil, nil, readWriteModel.RequestType_EMPTY, readWriteModel.NewRequestTermination(), cbusOptions)
		return readWriteModel.NewCBusMessageToServer(request, requestContext, cbusOptions), nil
	case *calIdentifyField:
		calData := readWriteModel.NewCALDataIdentify(field.attribute, readWriteModel.CALCommandTypeContainer_CALCommandIdentify, nil, requestContext)
		//TODO: we need support for bridged commands
		command := readWriteModel.NewCBusPointToPointCommandDirect(field.unitAddress, 0x0000, calData, cbusOptions)
		header := readWriteModel.NewCBusHeader(readWriteModel.PriorityClass_Class4, false, 0, readWriteModel.DestinationAddressType_PointToPoint)
		cbusCommand := readWriteModel.NewCBusCommandPointToPoint(command, header, cbusOptions)
		request := readWriteModel.NewRequestCommand(cbusCommand, nil, readWriteModel.NewAlpha(m.alphaGenerator.getAndIncrement()), readWriteModel.RequestType_REQUEST_COMMAND, nil, nil, readWriteModel.RequestType_EMPTY, readWriteModel.NewRequestTermination(), cbusOptions)
		return readWriteModel.NewCBusMessageToServer(request, requestContext, cbusOptions), nil
	case *calGetstatusField:
		calData := readWriteModel.NewCALDataGetStatus(field.parameter, field.count, readWriteModel.CALCommandTypeContainer_CALCommandGetStatus, nil, requestContext)
		//TODO: we need support for bridged commands
		command := readWriteModel.NewCBusPointToPointCommandDirect(field.unitAddress, 0x0000, calData, cbusOptions)
		header := readWriteModel.NewCBusHeader(readWriteModel.PriorityClass_Class4, false, 0, readWriteModel.DestinationAddressType_PointToPoint)
		cbusCommand := readWriteModel.NewCBusCommandPointToPoint(command, header, cbusOptions)
		request := readWriteModel.NewRequestCommand(cbusCommand, nil, readWriteModel.NewAlpha(m.alphaGenerator.getAndIncrement()), readWriteModel.RequestType_REQUEST_COMMAND, nil, nil, readWriteModel.RequestType_EMPTY, readWriteModel.NewRequestTermination(), cbusOptions)
		return readWriteModel.NewCBusMessageToServer(request, requestContext, cbusOptions), nil
	default:
		return nil, errors.Errorf("Unmapped type %T", field)
	}
}
