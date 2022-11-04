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

// BACnetConstructedDataTotalRecordCount is the corresponding interface of BACnetConstructedDataTotalRecordCount
type BACnetConstructedDataTotalRecordCount interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTotalRecordCount returns TotalRecordCount (property field)
	GetTotalRecordCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataTotalRecordCountExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTotalRecordCount.
// This is useful for switch cases.
type BACnetConstructedDataTotalRecordCountExactly interface {
	BACnetConstructedDataTotalRecordCount
	isBACnetConstructedDataTotalRecordCount() bool
}

// _BACnetConstructedDataTotalRecordCount is the data-structure of this message
type _BACnetConstructedDataTotalRecordCount struct {
	*_BACnetConstructedData
	TotalRecordCount BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTotalRecordCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TOTAL_RECORD_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTotalRecordCount) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetTotalRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.TotalRecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTotalRecordCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetTotalRecordCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTotalRecordCount factory function for _BACnetConstructedDataTotalRecordCount
func NewBACnetConstructedDataTotalRecordCount(totalRecordCount BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTotalRecordCount {
	_result := &_BACnetConstructedDataTotalRecordCount{
		TotalRecordCount:       totalRecordCount,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTotalRecordCount(structType interface{}) BACnetConstructedDataTotalRecordCount {
	if casted, ok := structType.(BACnetConstructedDataTotalRecordCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTotalRecordCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTotalRecordCount) GetTypeName() string {
	return "BACnetConstructedDataTotalRecordCount"
}

func (m *_BACnetConstructedDataTotalRecordCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTotalRecordCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (totalRecordCount)
	lengthInBits += m.TotalRecordCount.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTotalRecordCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTotalRecordCountParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTotalRecordCount, error) {
	return BACnetConstructedDataTotalRecordCountParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataTotalRecordCountParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTotalRecordCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTotalRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTotalRecordCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (totalRecordCount)
	if pullErr := readBuffer.PullContext("totalRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for totalRecordCount")
	}
	_totalRecordCount, _totalRecordCountErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _totalRecordCountErr != nil {
		return nil, errors.Wrap(_totalRecordCountErr, "Error parsing 'totalRecordCount' field of BACnetConstructedDataTotalRecordCount")
	}
	totalRecordCount := _totalRecordCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("totalRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for totalRecordCount")
	}

	// Virtual field
	_actualValue := totalRecordCount
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTotalRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTotalRecordCount")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTotalRecordCount{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TotalRecordCount: totalRecordCount,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTotalRecordCount) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTotalRecordCount) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTotalRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTotalRecordCount")
		}

		// Simple Field (totalRecordCount)
		if pushErr := writeBuffer.PushContext("totalRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for totalRecordCount")
		}
		_totalRecordCountErr := writeBuffer.WriteSerializable(m.GetTotalRecordCount())
		if popErr := writeBuffer.PopContext("totalRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for totalRecordCount")
		}
		if _totalRecordCountErr != nil {
			return errors.Wrap(_totalRecordCountErr, "Error serializing 'totalRecordCount' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTotalRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTotalRecordCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTotalRecordCount) isBACnetConstructedDataTotalRecordCount() bool {
	return true
}

func (m *_BACnetConstructedDataTotalRecordCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}