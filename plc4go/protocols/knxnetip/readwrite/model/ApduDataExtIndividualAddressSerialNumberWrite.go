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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtIndividualAddressSerialNumberWrite is the corresponding interface of ApduDataExtIndividualAddressSerialNumberWrite
type ApduDataExtIndividualAddressSerialNumberWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtIndividualAddressSerialNumberWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtIndividualAddressSerialNumberWrite.
// This is useful for switch cases.
type ApduDataExtIndividualAddressSerialNumberWriteExactly interface {
	ApduDataExtIndividualAddressSerialNumberWrite
	isApduDataExtIndividualAddressSerialNumberWrite() bool
}

// _ApduDataExtIndividualAddressSerialNumberWrite is the data-structure of this message
type _ApduDataExtIndividualAddressSerialNumberWrite struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) GetExtApciType() uint8 {
	return 0x1E
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtIndividualAddressSerialNumberWrite factory function for _ApduDataExtIndividualAddressSerialNumberWrite
func NewApduDataExtIndividualAddressSerialNumberWrite(length uint8) *_ApduDataExtIndividualAddressSerialNumberWrite {
	_result := &_ApduDataExtIndividualAddressSerialNumberWrite{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtIndividualAddressSerialNumberWrite(structType interface{}) ApduDataExtIndividualAddressSerialNumberWrite {
	if casted, ok := structType.(ApduDataExtIndividualAddressSerialNumberWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtIndividualAddressSerialNumberWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberWrite"
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtIndividualAddressSerialNumberWriteParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtIndividualAddressSerialNumberWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtIndividualAddressSerialNumberWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtIndividualAddressSerialNumberWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtIndividualAddressSerialNumberWrite{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtIndividualAddressSerialNumberWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtIndividualAddressSerialNumberWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) isApduDataExtIndividualAddressSerialNumberWrite() bool {
	return true
}

func (m *_ApduDataExtIndividualAddressSerialNumberWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}