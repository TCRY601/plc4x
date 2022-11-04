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

// BACnetConstructedDataCOVPeriod is the corresponding interface of BACnetConstructedDataCOVPeriod
type BACnetConstructedDataCOVPeriod interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovPeriod returns CovPeriod (property field)
	GetCovPeriod() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataCOVPeriodExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCOVPeriod.
// This is useful for switch cases.
type BACnetConstructedDataCOVPeriodExactly interface {
	BACnetConstructedDataCOVPeriod
	isBACnetConstructedDataCOVPeriod() bool
}

// _BACnetConstructedDataCOVPeriod is the data-structure of this message
type _BACnetConstructedDataCOVPeriod struct {
	*_BACnetConstructedData
	CovPeriod BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCOVPeriod) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCOVPeriod) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_PERIOD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCOVPeriod) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCOVPeriod) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCOVPeriod) GetCovPeriod() BACnetApplicationTagUnsignedInteger {
	return m.CovPeriod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCOVPeriod) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovPeriod())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCOVPeriod factory function for _BACnetConstructedDataCOVPeriod
func NewBACnetConstructedDataCOVPeriod(covPeriod BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCOVPeriod {
	_result := &_BACnetConstructedDataCOVPeriod{
		CovPeriod:              covPeriod,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCOVPeriod(structType interface{}) BACnetConstructedDataCOVPeriod {
	if casted, ok := structType.(BACnetConstructedDataCOVPeriod); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCOVPeriod); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCOVPeriod) GetTypeName() string {
	return "BACnetConstructedDataCOVPeriod"
}

func (m *_BACnetConstructedDataCOVPeriod) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCOVPeriod) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (covPeriod)
	lengthInBits += m.CovPeriod.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCOVPeriod) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCOVPeriodParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCOVPeriod, error) {
	return BACnetConstructedDataCOVPeriodParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument) // TODO: get endianness from mspec
}

func BACnetConstructedDataCOVPeriodParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCOVPeriod, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCOVPeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCOVPeriod")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (covPeriod)
	if pullErr := readBuffer.PullContext("covPeriod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for covPeriod")
	}
	_covPeriod, _covPeriodErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _covPeriodErr != nil {
		return nil, errors.Wrap(_covPeriodErr, "Error parsing 'covPeriod' field of BACnetConstructedDataCOVPeriod")
	}
	covPeriod := _covPeriod.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("covPeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covPeriod")
	}

	// Virtual field
	_actualValue := covPeriod
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCOVPeriod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCOVPeriod")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCOVPeriod{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CovPeriod: covPeriod,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCOVPeriod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCOVPeriod) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCOVPeriod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCOVPeriod")
		}

		// Simple Field (covPeriod)
		if pushErr := writeBuffer.PushContext("covPeriod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covPeriod")
		}
		_covPeriodErr := writeBuffer.WriteSerializable(m.GetCovPeriod())
		if popErr := writeBuffer.PopContext("covPeriod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covPeriod")
		}
		if _covPeriodErr != nil {
			return errors.Wrap(_covPeriodErr, "Error serializing 'covPeriod' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCOVPeriod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCOVPeriod")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCOVPeriod) isBACnetConstructedDataCOVPeriod() bool {
	return true
}

func (m *_BACnetConstructedDataCOVPeriod) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}