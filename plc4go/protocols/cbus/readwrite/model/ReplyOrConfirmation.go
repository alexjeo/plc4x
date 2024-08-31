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

// ReplyOrConfirmation is the corresponding interface of ReplyOrConfirmation
type ReplyOrConfirmation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() byte
	// GetIsAlpha returns IsAlpha (virtual field)
	GetIsAlpha() bool
}

// ReplyOrConfirmationExactly can be used when we want exactly this type and not a type which fulfills ReplyOrConfirmation.
// This is useful for switch cases.
type ReplyOrConfirmationExactly interface {
	ReplyOrConfirmation
	isReplyOrConfirmation() bool
}

// _ReplyOrConfirmation is the data-structure of this message
type _ReplyOrConfirmation struct {
	_ReplyOrConfirmationChildRequirements
	PeekedByte byte

	// Arguments.
	CBusOptions    CBusOptions
	RequestContext RequestContext
}

type _ReplyOrConfirmationChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetIsAlpha() bool
	GetPeekedByte() byte
}

type ReplyOrConfirmationParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ReplyOrConfirmation, serializeChildFunction func() error) error
	GetTypeName() string
}

type ReplyOrConfirmationChild interface {
	utils.Serializable
	InitializeParent(parent ReplyOrConfirmation, peekedByte byte)
	GetParent() *ReplyOrConfirmation

	GetTypeName() string
	ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmation) GetPeekedByte() byte {
	return m.PeekedByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ReplyOrConfirmation) GetIsAlpha() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((bool((m.GetPeekedByte()) >= (0x67)))) && bool((bool((m.GetPeekedByte()) <= (0x7A)))))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyOrConfirmation factory function for _ReplyOrConfirmation
func NewReplyOrConfirmation(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmation {
	return &_ReplyOrConfirmation{PeekedByte: peekedByte, CBusOptions: cBusOptions, RequestContext: requestContext}
}

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmation(structType any) ReplyOrConfirmation {
	if casted, ok := structType.(ReplyOrConfirmation); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmation); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmation) GetTypeName() string {
	return "ReplyOrConfirmation"
}

func (m *_ReplyOrConfirmation) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_ReplyOrConfirmation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReplyOrConfirmationParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (ReplyOrConfirmation, error) {
	return ReplyOrConfirmationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func ReplyOrConfirmationParseWithBufferProducer[T ReplyOrConfirmation](cBusOptions CBusOptions, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := ReplyOrConfirmationParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func ReplyOrConfirmationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ReplyOrConfirmation, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReplyOrConfirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedByte, err := ReadPeekField[byte](ctx, "peekedByte", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedByte' field"))
	}

	isAlpha, err := ReadVirtualField[bool](ctx, "isAlpha", (*bool)(nil), bool((bool((peekedByte) >= (0x67)))) && bool((bool((peekedByte) <= (0x7A)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAlpha' field"))
	}
	_ = isAlpha

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ReplyOrConfirmationChildSerializeRequirement interface {
		ReplyOrConfirmation
		InitializeParent(ReplyOrConfirmation, byte)
		GetParent() ReplyOrConfirmation
	}
	var _childTemp any
	var _child ReplyOrConfirmationChildSerializeRequirement
	var typeSwitchError error
	switch {
	case isAlpha == bool(false) && peekedByte == 0x21: // ServerErrorReply
		_childTemp, typeSwitchError = ServerErrorReplyParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	case isAlpha == bool(true): // ReplyOrConfirmationConfirmation
		_childTemp, typeSwitchError = ReplyOrConfirmationConfirmationParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	case isAlpha == bool(false): // ReplyOrConfirmationReply
		_childTemp, typeSwitchError = ReplyOrConfirmationReplyParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [isAlpha=%v, peekedByte=%v]", isAlpha, peekedByte)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ReplyOrConfirmation")
	}
	_child = _childTemp.(ReplyOrConfirmationChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmation")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedByte)
	return _child, nil
}

func (pm *_ReplyOrConfirmation) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ReplyOrConfirmation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ReplyOrConfirmation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmation")
	}
	// Virtual field
	isAlpha := m.GetIsAlpha()
	_ = isAlpha
	if _isAlphaErr := writeBuffer.WriteVirtual(ctx, "isAlpha", m.GetIsAlpha()); _isAlphaErr != nil {
		return errors.Wrap(_isAlphaErr, "Error serializing 'isAlpha' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ReplyOrConfirmation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ReplyOrConfirmation")
	}
	return nil
}

////
// Arguments Getter

func (m *_ReplyOrConfirmation) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_ReplyOrConfirmation) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_ReplyOrConfirmation) isReplyOrConfirmation() bool {
	return true
}

func (m *_ReplyOrConfirmation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
