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
type BACnetContextTagEmpty struct {
	*BACnetContextTag
}

// The corresponding interface
type IBACnetContextTagEmpty interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagEmpty) DataType() BACnetDataType {
	return 0
}

func (m *BACnetContextTagEmpty) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, isNotOpeningOrClosingTag bool) {
	m.BACnetContextTag.Header = header
	m.BACnetContextTag.TagNumber = tagNumber
	m.BACnetContextTag.ActualLength = actualLength
	m.BACnetContextTag.IsNotOpeningOrClosingTag = isNotOpeningOrClosingTag
}

func NewBACnetContextTagEmpty(header *BACnetTagHeader, tagNumber uint8, actualLength uint32, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagEmpty{
		BACnetContextTag: NewBACnetContextTag(header, tagNumber, actualLength, isNotOpeningOrClosingTag),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagEmpty(structType interface{}) *BACnetContextTagEmpty {
	castFunc := func(typ interface{}) *BACnetContextTagEmpty {
		if casted, ok := typ.(BACnetContextTagEmpty); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagEmpty); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagEmpty(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagEmpty(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagEmpty) GetTypeName() string {
	return "BACnetContextTagEmpty"
}

func (m *BACnetContextTagEmpty) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagEmpty) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetContextTagEmpty) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagEmptyParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagEmpty"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagEmpty"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagEmpty{
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagEmpty) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEmpty"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEmpty"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagEmpty) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
