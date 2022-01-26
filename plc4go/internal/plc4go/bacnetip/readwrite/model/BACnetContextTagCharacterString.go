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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetContextTagCharacterString struct {
	*BACnetContextTag
	Encoding          BACnetCharacterEncoding
	Value             string
	ActualLengthInBit uint16
}

// The corresponding interface
type IBACnetContextTagCharacterString interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagCharacterString) DataType() BACnetDataType {
	return BACnetDataType_CHARACTER_STRING
}

func (m *BACnetContextTagCharacterString) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, isNotOpeningOrClosingTag bool) {
	m.BACnetContextTag.Header = header
	m.BACnetContextTag.TagNumber = tagNumber
	m.BACnetContextTag.ActualLength = actualLength
	m.BACnetContextTag.IsNotOpeningOrClosingTag = isNotOpeningOrClosingTag
}

func NewBACnetContextTagCharacterString(encoding BACnetCharacterEncoding, value string, actualLengthInBit uint16, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagCharacterString{
		Encoding:          encoding,
		Value:             value,
		ActualLengthInBit: actualLengthInBit,
		BACnetContextTag:  NewBACnetContextTag(header, tagNumber, actualLength, isNotOpeningOrClosingTag),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagCharacterString(structType interface{}) *BACnetContextTagCharacterString {
	castFunc := func(typ interface{}) *BACnetContextTagCharacterString {
		if casted, ok := typ.(BACnetContextTagCharacterString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagCharacterString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagCharacterString(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagCharacterString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagCharacterString) GetTypeName() string {
	return "BACnetContextTagCharacterString"
}

func (m *BACnetContextTagCharacterString) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagCharacterString) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (encoding)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.ActualLengthInBit)

	return lengthInBits
}

func (m *BACnetContextTagCharacterString) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagCharacterStringParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagCharacterString"); pullErr != nil {
		return nil, pullErr
	}

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Simple Field (encoding)
	if pullErr := readBuffer.PullContext("encoding"); pullErr != nil {
		return nil, pullErr
	}
	_encoding, _encodingErr := BACnetCharacterEncodingParse(readBuffer)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field")
	}
	encoding := _encoding
	if closeErr := readBuffer.CloseContext("encoding"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_actualLengthInBit := uint16(uint16(actualLength)*uint16(uint16(8))) - uint16(uint16(8))
	actualLengthInBit := uint16(_actualLengthInBit)

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadString("value", uint32(actualLengthInBit))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetContextTagCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagCharacterString{
		Encoding:          encoding,
		Value:             value,
		ActualLengthInBit: actualLengthInBit,
		BACnetContextTag:  &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagCharacterString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (encoding)
		if pushErr := writeBuffer.PushContext("encoding"); pushErr != nil {
			return pushErr
		}
		_encodingErr := m.Encoding.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("encoding"); popErr != nil {
			return popErr
		}
		if _encodingErr != nil {
			return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
		}
		// Virtual field
		if _actualLengthInBitErr := writeBuffer.WriteVirtual("actualLengthInBit", m.ActualLengthInBit); _actualLengthInBitErr != nil {
			return errors.Wrap(_actualLengthInBitErr, "Error serializing 'actualLengthInBit' field")
		}

		// Simple Field (value)
		value := string(m.Value)
		_valueErr := writeBuffer.WriteString("value", uint32(m.ActualLengthInBit), "UTF-8", (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
