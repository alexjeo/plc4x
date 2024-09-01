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
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1Command is the corresponding interface of DF1Command
type DF1Command interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandCode returns CommandCode (discriminator field)
	GetCommandCode() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetTransactionCounter returns TransactionCounter (property field)
	GetTransactionCounter() uint16
}

// DF1CommandExactly can be used when we want exactly this type and not a type which fulfills DF1Command.
// This is useful for switch cases.
type DF1CommandExactly interface {
	DF1Command
	isDF1Command() bool
}

// _DF1Command is the data-structure of this message
type _DF1Command struct {
	_DF1CommandChildRequirements
	Status             uint8
	TransactionCounter uint16
}

type _DF1CommandChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandCode() uint8
}

type DF1CommandParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1Command, serializeChildFunction func() error) error
	GetTypeName() string
}

type DF1CommandChild interface {
	utils.Serializable
	InitializeParent(parent DF1Command, status uint8, transactionCounter uint16)
	GetParent() *DF1Command

	GetTypeName() string
	DF1Command
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1Command) GetStatus() uint8 {
	return m.Status
}

func (m *_DF1Command) GetTransactionCounter() uint16 {
	return m.TransactionCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1Command factory function for _DF1Command
func NewDF1Command(status uint8, transactionCounter uint16) *_DF1Command {
	return &_DF1Command{Status: status, TransactionCounter: transactionCounter}
}

// Deprecated: use the interface for direct cast
func CastDF1Command(structType any) DF1Command {
	if casted, ok := structType.(DF1Command); ok {
		return casted
	}
	if casted, ok := structType.(*DF1Command); ok {
		return *casted
	}
	return nil
}

func (m *_DF1Command) GetTypeName() string {
	return "DF1Command"
}

func (m *_DF1Command) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandCode)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (transactionCounter)
	lengthInBits += 16

	return lengthInBits
}

func (m *_DF1Command) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1CommandParse(ctx context.Context, theBytes []byte) (DF1Command, error) {
	return DF1CommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DF1CommandParseWithBufferProducer[T DF1Command]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := DF1CommandParseWithBuffer(ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func DF1CommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DF1Command, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1Command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1Command")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	commandCode, err := ReadDiscriminatorField[uint8](ctx, "commandCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandCode' field"))
	}

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}

	transactionCounter, err := ReadSimpleField(ctx, "transactionCounter", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transactionCounter' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type DF1CommandChildSerializeRequirement interface {
		DF1Command
		InitializeParent(DF1Command, uint8, uint16)
		GetParent() DF1Command
	}
	var _childTemp any
	var _child DF1CommandChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandCode == 0x01: // DF1UnprotectedReadRequest
		_childTemp, typeSwitchError = DF1UnprotectedReadRequestParseWithBuffer(ctx, readBuffer)
	case commandCode == 0x41: // DF1UnprotectedReadResponse
		_childTemp, typeSwitchError = DF1UnprotectedReadResponseParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandCode=%v]", commandCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of DF1Command")
	}
	_child = _childTemp.(DF1CommandChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("DF1Command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1Command")
	}

	// Finish initializing
	_child.InitializeParent(_child, status, transactionCounter)
	return _child, nil
}

func (pm *_DF1Command) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DF1Command, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DF1Command"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DF1Command")
	}

	if err := WriteDiscriminatorField(ctx, "commandCode", m.GetCommandCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'commandCode' field")
	}

	// Simple Field (status)
	status := uint8(m.GetStatus())
	_statusErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("status", 8, uint8((status)))
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	// Simple Field (transactionCounter)
	transactionCounter := uint16(m.GetTransactionCounter())
	_transactionCounterErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("transactionCounter", 16, uint16((transactionCounter)))
	if _transactionCounterErr != nil {
		return errors.Wrap(_transactionCounterErr, "Error serializing 'transactionCounter' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DF1Command"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DF1Command")
	}
	return nil
}

func (m *_DF1Command) isDF1Command() bool {
	return true
}

func (m *_DF1Command) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
