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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataObjectType is the corresponding interface of BACnetConstructedDataObjectType
type BACnetConstructedDataObjectType interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetObjectType returns ObjectType (property field)
	GetObjectType() BACnetObjectTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectTypeTagged
}

// BACnetConstructedDataObjectTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataObjectType.
// This is useful for switch cases.
type BACnetConstructedDataObjectTypeExactly interface {
	BACnetConstructedDataObjectType
	isBACnetConstructedDataObjectType() bool
}

// _BACnetConstructedDataObjectType is the data-structure of this message
type _BACnetConstructedDataObjectType struct {
	*_BACnetConstructedData
	ObjectType BACnetObjectTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataObjectType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataObjectType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataObjectType) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataObjectType) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataObjectType) GetObjectType() BACnetObjectTypeTagged {
	return m.ObjectType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataObjectType) GetActualValue() BACnetObjectTypeTagged {
	return CastBACnetObjectTypeTagged(m.GetObjectType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataObjectType factory function for _BACnetConstructedDataObjectType
func NewBACnetConstructedDataObjectType(objectType BACnetObjectTypeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataObjectType {
	_result := &_BACnetConstructedDataObjectType{
		ObjectType:             objectType,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataObjectType(structType interface{}) BACnetConstructedDataObjectType {
	if casted, ok := structType.(BACnetConstructedDataObjectType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataObjectType) GetTypeName() string {
	return "BACnetConstructedDataObjectType"
}

func (m *_BACnetConstructedDataObjectType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataObjectType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectType)
	lengthInBits += m.ObjectType.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataObjectType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataObjectTypeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectType, error) {
	return BACnetConstructedDataObjectTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataObjectTypeParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectType)
	if pullErr := readBuffer.PullContext("objectType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectType")
	}
	_objectType, _objectTypeErr := BACnetObjectTypeTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _objectTypeErr != nil {
		return nil, errors.Wrap(_objectTypeErr, "Error parsing 'objectType' field of BACnetConstructedDataObjectType")
	}
	objectType := _objectType.(BACnetObjectTypeTagged)
	if closeErr := readBuffer.CloseContext("objectType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectType")
	}

	// Virtual field
	_actualValue := objectType
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectType")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataObjectType{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ObjectType: objectType,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataObjectType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataObjectType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectType")
		}

		// Simple Field (objectType)
		if pushErr := writeBuffer.PushContext("objectType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectType")
		}
		_objectTypeErr := writeBuffer.WriteSerializable(m.GetObjectType())
		if popErr := writeBuffer.PopContext("objectType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectType")
		}
		if _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectType")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataObjectType) isBACnetConstructedDataObjectType() bool {
	return true
}

func (m *_BACnetConstructedDataObjectType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}