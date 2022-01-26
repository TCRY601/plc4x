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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BVLCDeleteForeignDeviceTableEntry struct {
	*BVLC
}

// The corresponding interface
type IBVLCDeleteForeignDeviceTableEntry interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCDeleteForeignDeviceTableEntry) BvlcFunction() uint8 {
	return 0x08
}

func (m *BVLCDeleteForeignDeviceTableEntry) InitializeParent(parent *BVLC, bvlcPayloadLength uint16) {
	m.BVLC.BvlcPayloadLength = bvlcPayloadLength
}

func NewBVLCDeleteForeignDeviceTableEntry(bvlcPayloadLength uint16) *BVLC {
	child := &BVLCDeleteForeignDeviceTableEntry{
		BVLC: NewBVLC(bvlcPayloadLength),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCDeleteForeignDeviceTableEntry(structType interface{}) *BVLCDeleteForeignDeviceTableEntry {
	castFunc := func(typ interface{}) *BVLCDeleteForeignDeviceTableEntry {
		if casted, ok := typ.(BVLCDeleteForeignDeviceTableEntry); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCDeleteForeignDeviceTableEntry); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCDeleteForeignDeviceTableEntry"
}

func (m *BVLCDeleteForeignDeviceTableEntry) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCDeleteForeignDeviceTableEntry) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BVLCDeleteForeignDeviceTableEntry) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCDeleteForeignDeviceTableEntryParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCDeleteForeignDeviceTableEntry"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BVLCDeleteForeignDeviceTableEntry"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCDeleteForeignDeviceTableEntry{
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCDeleteForeignDeviceTableEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDeleteForeignDeviceTableEntry"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BVLCDeleteForeignDeviceTableEntry"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCDeleteForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
