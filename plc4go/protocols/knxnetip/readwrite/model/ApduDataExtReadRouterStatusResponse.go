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

// ApduDataExtReadRouterStatusResponse is the corresponding interface of ApduDataExtReadRouterStatusResponse
type ApduDataExtReadRouterStatusResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtReadRouterStatusResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtReadRouterStatusResponse.
// This is useful for switch cases.
type ApduDataExtReadRouterStatusResponseExactly interface {
	ApduDataExtReadRouterStatusResponse
	isApduDataExtReadRouterStatusResponse() bool
}

// _ApduDataExtReadRouterStatusResponse is the data-structure of this message
type _ApduDataExtReadRouterStatusResponse struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtReadRouterStatusResponse) GetExtApciType() uint8 {
	return 0x0E
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtReadRouterStatusResponse) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtReadRouterStatusResponse) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtReadRouterStatusResponse factory function for _ApduDataExtReadRouterStatusResponse
func NewApduDataExtReadRouterStatusResponse(length uint8) *_ApduDataExtReadRouterStatusResponse {
	_result := &_ApduDataExtReadRouterStatusResponse{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtReadRouterStatusResponse(structType interface{}) ApduDataExtReadRouterStatusResponse {
	if casted, ok := structType.(ApduDataExtReadRouterStatusResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtReadRouterStatusResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtReadRouterStatusResponse) GetTypeName() string {
	return "ApduDataExtReadRouterStatusResponse"
}

func (m *_ApduDataExtReadRouterStatusResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtReadRouterStatusResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtReadRouterStatusResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtReadRouterStatusResponseParse(theBytes []byte, length uint8) (ApduDataExtReadRouterStatusResponse, error) {
	return ApduDataExtReadRouterStatusResponseParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), length) // TODO: get endianness from mspec
}

func ApduDataExtReadRouterStatusResponseParseWithBuffer(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtReadRouterStatusResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtReadRouterStatusResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtReadRouterStatusResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtReadRouterStatusResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtReadRouterStatusResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtReadRouterStatusResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtReadRouterStatusResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtReadRouterStatusResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtReadRouterStatusResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtReadRouterStatusResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtReadRouterStatusResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtReadRouterStatusResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtReadRouterStatusResponse) isApduDataExtReadRouterStatusResponse() bool {
	return true
}

func (m *_ApduDataExtReadRouterStatusResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}