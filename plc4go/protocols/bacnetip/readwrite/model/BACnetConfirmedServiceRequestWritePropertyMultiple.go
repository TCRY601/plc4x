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

// BACnetConfirmedServiceRequestWritePropertyMultiple is the data-structure of this message
type BACnetConfirmedServiceRequestWritePropertyMultiple struct {
	*BACnetConfirmedServiceRequest
	Data []*BACnetWriteAccessSpecification

	// Arguments.
	ServiceRequestLength        uint16
	ServiceRequestPayloadLength uint16
}

// IBACnetConfirmedServiceRequestWritePropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestWritePropertyMultiple
type IBACnetConfirmedServiceRequestWritePropertyMultiple interface {
	IBACnetConfirmedServiceRequest
	// GetData returns Data (property field)
	GetData() []*BACnetWriteAccessSpecification
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetData() []*BACnetWriteAccessSpecification {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestWritePropertyMultiple factory function for BACnetConfirmedServiceRequestWritePropertyMultiple
func NewBACnetConfirmedServiceRequestWritePropertyMultiple(data []*BACnetWriteAccessSpecification, serviceRequestLength uint16, serviceRequestPayloadLength uint16) *BACnetConfirmedServiceRequestWritePropertyMultiple {
	_result := &BACnetConfirmedServiceRequestWritePropertyMultiple{
		Data:                          data,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestWritePropertyMultiple(structType interface{}) *BACnetConfirmedServiceRequestWritePropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWritePropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestWritePropertyMultiple(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestWritePropertyMultiple(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWritePropertyMultiple"
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestWritePropertyMultipleParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16, serviceRequestPayloadLength uint16) (*BACnetConfirmedServiceRequestWritePropertyMultiple, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]*BACnetWriteAccessSpecification, 0)
	{
		_dataLength := serviceRequestPayloadLength
		_dataEndPos := positionAware.GetPos() + uint16(_dataLength)
		for positionAware.GetPos() < _dataEndPos {
			_item, _err := BACnetWriteAccessSpecificationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, CastBACnetWriteAccessSpecification(_item))
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestWritePropertyMultiple{
		Data:                          data,
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWritePropertyMultiple"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestWritePropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
