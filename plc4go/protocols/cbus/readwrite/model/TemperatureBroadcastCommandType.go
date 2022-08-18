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

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// TemperatureBroadcastCommandType is an enum
type TemperatureBroadcastCommandType uint8

type ITemperatureBroadcastCommandType interface {
	NumberOfArguments() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	TemperatureBroadcastCommandType_BROADCAST_EVENT TemperatureBroadcastCommandType = 0x00
)

var TemperatureBroadcastCommandTypeValues []TemperatureBroadcastCommandType

func init() {
	_ = errors.New
	TemperatureBroadcastCommandTypeValues = []TemperatureBroadcastCommandType{
		TemperatureBroadcastCommandType_BROADCAST_EVENT,
	}
}

func (e TemperatureBroadcastCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 2
		}
	default:
		{
			return 0
		}
	}
}

func TemperatureBroadcastCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (TemperatureBroadcastCommandType, error) {
	for _, sizeValue := range TemperatureBroadcastCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func TemperatureBroadcastCommandTypeByValue(value uint8) (enum TemperatureBroadcastCommandType, ok bool) {
	switch value {
	case 0x00:
		return TemperatureBroadcastCommandType_BROADCAST_EVENT, true
	}
	return 0, false
}

func TemperatureBroadcastCommandTypeByName(value string) (enum TemperatureBroadcastCommandType, ok bool) {
	switch value {
	case "BROADCAST_EVENT":
		return TemperatureBroadcastCommandType_BROADCAST_EVENT, true
	}
	return 0, false
}

func TemperatureBroadcastCommandTypeKnows(value uint8) bool {
	for _, typeValue := range TemperatureBroadcastCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTemperatureBroadcastCommandType(structType interface{}) TemperatureBroadcastCommandType {
	castFunc := func(typ interface{}) TemperatureBroadcastCommandType {
		if sTemperatureBroadcastCommandType, ok := typ.(TemperatureBroadcastCommandType); ok {
			return sTemperatureBroadcastCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m TemperatureBroadcastCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m TemperatureBroadcastCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TemperatureBroadcastCommandTypeParse(readBuffer utils.ReadBuffer) (TemperatureBroadcastCommandType, error) {
	val, err := readBuffer.ReadUint8("TemperatureBroadcastCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TemperatureBroadcastCommandType")
	}
	if enum, ok := TemperatureBroadcastCommandTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return TemperatureBroadcastCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e TemperatureBroadcastCommandType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("TemperatureBroadcastCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TemperatureBroadcastCommandType) PLC4XEnumName() string {
	switch e {
	case TemperatureBroadcastCommandType_BROADCAST_EVENT:
		return "BROADCAST_EVENT"
	}
	return ""
}

func (e TemperatureBroadcastCommandType) String() string {
	return e.PLC4XEnumName()
}
