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
type BACnetContextTagEnumerated struct {
	*BACnetContextTag
	Data []int8
}

// The corresponding interface
type IBACnetContextTagEnumerated interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagEnumerated) DataType() BACnetDataType {
	return BACnetDataType_ENUMERATED
}

func (m *BACnetContextTagEnumerated) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, isNotOpeningOrClosingTag bool) {
	m.BACnetContextTag.Header = header
	m.BACnetContextTag.TagNumber = tagNumber
	m.BACnetContextTag.ActualLength = actualLength
	m.BACnetContextTag.IsNotOpeningOrClosingTag = isNotOpeningOrClosingTag
}

func NewBACnetContextTagEnumerated(data []int8, header *BACnetTagHeader, tagNumber uint8, actualLength uint32, isNotOpeningOrClosingTag bool) *BACnetContextTag {
	child := &BACnetContextTagEnumerated{
		Data:             data,
		BACnetContextTag: NewBACnetContextTag(header, tagNumber, actualLength, isNotOpeningOrClosingTag),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagEnumerated(structType interface{}) *BACnetContextTagEnumerated {
	castFunc := func(typ interface{}) *BACnetContextTagEnumerated {
		if casted, ok := typ.(BACnetContextTagEnumerated); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetContextTagEnumerated); ok {
			return casted
		}
		if casted, ok := typ.(BACnetContextTag); ok {
			return CastBACnetContextTagEnumerated(casted.Child)
		}
		if casted, ok := typ.(*BACnetContextTag); ok {
			return CastBACnetContextTagEnumerated(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetContextTagEnumerated) GetTypeName() string {
	return "BACnetContextTagEnumerated"
}

func (m *BACnetContextTagEnumerated) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetContextTagEnumerated) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *BACnetContextTagEnumerated) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetContextTagEnumeratedParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagEnumerated"); pullErr != nil {
		return nil, pullErr
	}

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]int8, 0)
	{
		_dataLength := actualLength
		_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
		for readBuffer.GetPos() < _dataEndPos {
			_item, _err := readBuffer.ReadInt8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagEnumerated"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagEnumerated{
		Data:             data,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEnumerated"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEnumerated"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
