/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAck is the data-structure of this message
type BACnetServiceAck struct {

	// Arguments.
	ServiceAckLength uint16
	Child            IBACnetServiceAckChild
}

// IBACnetServiceAck is the corresponding interface of BACnetServiceAck
type IBACnetServiceAck interface {
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetConfirmedServiceChoice
	// GetServiceAckPayloadLength returns ServiceAckPayloadLength (virtual field)
	GetServiceAckPayloadLength() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetServiceAckParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetServiceAck, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetServiceAckChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetServiceAck)
	GetParent() *BACnetServiceAck

	GetTypeName() string
	IBACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetServiceAck) GetServiceAckPayloadLength() uint16 {
	return uint16(utils.InlineIf(bool(bool((m.ServiceAckLength) > (0))), func() interface{} { return uint16(uint16(uint16(m.ServiceAckLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAck factory function for BACnetServiceAck
func NewBACnetServiceAck(serviceAckLength uint16) *BACnetServiceAck {
	return &BACnetServiceAck{ServiceAckLength: serviceAckLength}
}

func CastBACnetServiceAck(structType interface{}) *BACnetServiceAck {
	if casted, ok := structType.(BACnetServiceAck); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetServiceAckChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetServiceAck) GetTypeName() string {
	return "BACnetServiceAck"
}

func (m *BACnetServiceAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetServiceAck) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetServiceAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (*BACnetServiceAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	if pullErr := readBuffer.PullContext("serviceChoice"); pullErr != nil {
		return nil, pullErr
	}
	serviceChoice_temp, _serviceChoiceErr := BACnetConfirmedServiceChoiceParse(readBuffer)
	var serviceChoice BACnetConfirmedServiceChoice = serviceChoice_temp
	if closeErr := readBuffer.CloseContext("serviceChoice"); closeErr != nil {
		return nil, closeErr
	}
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Virtual field
	_serviceAckPayloadLength := utils.InlineIf(bool(bool((serviceAckLength) > (0))), func() interface{} { return uint16(uint16(uint16(serviceAckLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16)
	serviceAckPayloadLength := uint16(_serviceAckPayloadLength)
	_ = serviceAckPayloadLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetServiceAckChild interface {
		InitializeParent(*BACnetServiceAck)
		GetParent() *BACnetServiceAck
	}
	var _child BACnetServiceAckChild
	var typeSwitchError error
	switch {
	case serviceChoice == BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY: // BACnetServiceAckGetAlarmSummary
		_child, typeSwitchError = BACnetServiceAckGetAlarmSummaryParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY: // BACnetServiceAckGetEnrollmentSummary
		_child, typeSwitchError = BACnetServiceAckGetEnrollmentSummaryParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION: // BACnetServiceAckGetEventInformation
		_child, typeSwitchError = BACnetServiceAckGetEventInformationParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_READ_FILE: // BACnetServiceAckAtomicReadFile
		_child, typeSwitchError = BACnetServiceAckAtomicReadFileParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_ATOMIC_WRITE_FILE: // BACnetServiceAckAtomicWriteFile
		_child, typeSwitchError = BACnetServiceAckAtomicWriteFileParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT: // BACnetServiceAckCreateObject
		_child, typeSwitchError = BACnetServiceAckCreateObjectParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY: // BACnetServiceAckReadProperty
		_child, typeSwitchError = BACnetServiceAckReadPropertyParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_MULTIPLE: // BACnetServiceAckReadPropertyMultiple
		_child, typeSwitchError = BACnetServiceAckReadPropertyMultipleParse(readBuffer, serviceAckLength, serviceAckPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_RANGE: // BACnetServiceAckReadRange
		_child, typeSwitchError = BACnetServiceAckReadRangeParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER: // BACnetServiceAckConfirmedPrivateTransfer
		_child, typeSwitchError = BACnetServiceAckConfirmedPrivateTransferParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_VT_OPEN: // BACnetServiceAckVTOpen
		_child, typeSwitchError = BACnetServiceAckVTOpenParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_VT_DATA: // BACnetServiceAckVTData
		_child, typeSwitchError = BACnetServiceAckVTDataParse(readBuffer, serviceAckLength)
	case serviceChoice == BACnetConfirmedServiceChoice_AUTHENTICATE: // BACnetServiceAckAuthenticate
		_child, typeSwitchError = BACnetServiceAckAuthenticateParse(readBuffer, serviceAckLength, serviceAckPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_REQUEST_KEY: // BACnetServiceAckRequestKey
		_child, typeSwitchError = BACnetServiceAckRequestKeyParse(readBuffer, serviceAckLength, serviceAckPayloadLength)
	case serviceChoice == BACnetConfirmedServiceChoice_READ_PROPERTY_CONDITIONAL: // BACnetServiceAckReadPropertyConditional
		_child, typeSwitchError = BACnetServiceAckReadPropertyConditionalParse(readBuffer, serviceAckLength, serviceAckPayloadLength)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAck"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *BACnetServiceAck) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetServiceAck) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetServiceAck, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetServiceAck"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := BACnetConfirmedServiceChoice(child.GetServiceChoice())
	if pushErr := writeBuffer.PushContext("serviceChoice"); pushErr != nil {
		return pushErr
	}
	_serviceChoiceErr := serviceChoice.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("serviceChoice"); popErr != nil {
		return popErr
	}

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}
	// Virtual field
	if _serviceAckPayloadLengthErr := writeBuffer.WriteVirtual("serviceAckPayloadLength", m.GetServiceAckPayloadLength()); _serviceAckPayloadLengthErr != nil {
		return errors.Wrap(_serviceAckPayloadLengthErr, "Error serializing 'serviceAckPayloadLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetServiceAck"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetServiceAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
