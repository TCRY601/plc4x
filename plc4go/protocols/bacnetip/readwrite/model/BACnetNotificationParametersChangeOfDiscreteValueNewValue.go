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

// BACnetNotificationParametersChangeOfDiscreteValueNewValue is the data-structure of this message
type BACnetNotificationParametersChangeOfDiscreteValueNewValue struct {
	OpeningTag      *BACnetOpeningTag
	PeekedTagHeader *BACnetTagHeader
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber uint8
	Child     IBACnetNotificationParametersChangeOfDiscreteValueNewValueChild
}

// IBACnetNotificationParametersChangeOfDiscreteValueNewValue is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValue
type IBACnetNotificationParametersChangeOfDiscreteValueNewValue interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetNotificationParametersChangeOfDiscreteValueNewValueParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetNotificationParametersChangeOfDiscreteValueNewValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetNotificationParametersChangeOfDiscreteValueNewValueChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag)
	GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue

	GetTypeName() string
	IBACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetPeekedIsContextTag() bool {
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValue factory function for BACnetNotificationParametersChangeOfDiscreteValueNewValue
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return &BACnetNotificationParametersChangeOfDiscreteValueNewValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetNotificationParametersChangeOfDiscreteValueNewValue(structType interface{}) *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetNotificationParametersChangeOfDiscreteValueNewValueChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValue"
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfDiscreteValueNewValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Virtual field
	_peekedIsContextTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))
	peekedIsContextTag := bool(_peekedIsContextTag)
	_ = peekedIsContextTag

	// Validation
	if !(bool(bool(!(peekedIsContextTag))) || bool(bool(bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, utils.ParseValidationError{"unexpected opening or closing tag"}
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetNotificationParametersChangeOfDiscreteValueNewValueChild interface {
		InitializeParent(*BACnetNotificationParametersChangeOfDiscreteValueNewValue, *BACnetOpeningTag, *BACnetTagHeader, *BACnetClosingTag)
		GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue
	}
	var _child BACnetNotificationParametersChangeOfDiscreteValueNewValueChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsigned
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueUnsignedParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueInteger
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueIntegerParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumerated
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueEnumeratedParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterString
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueCharacterStringParse(readBuffer, tagNumber)
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetStringParse(readBuffer, tagNumber)
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateParse(readBuffer, tagNumber)
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTime
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetTimeParse(readBuffer, tagNumber)
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifierParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetime
		_child, typeSwitchError = BACnetNotificationParametersChangeOfDiscreteValueNewValueDatetimeParse(readBuffer, tagNumber)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), openingTag, peekedTagHeader, closingTag)
	return _child.GetParent(), nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetNotificationParametersChangeOfDiscreteValueNewValue, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual("peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValue"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
