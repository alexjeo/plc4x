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

// RequestContext is the corresponding interface of RequestContext
type RequestContext interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetSendIdentifyRequestBefore returns SendIdentifyRequestBefore (property field)
	GetSendIdentifyRequestBefore() bool
	// IsRequestContext is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequestContext()
	// CreateBuilder creates a RequestContextBuilder
	CreateRequestContextBuilder() RequestContextBuilder
}

// _RequestContext is the data-structure of this message
type _RequestContext struct {
	SendIdentifyRequestBefore bool
}

var _ RequestContext = (*_RequestContext)(nil)

// NewRequestContext factory function for _RequestContext
func NewRequestContext(sendIdentifyRequestBefore bool) *_RequestContext {
	return &_RequestContext{SendIdentifyRequestBefore: sendIdentifyRequestBefore}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RequestContextBuilder is a builder for RequestContext
type RequestContextBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(sendIdentifyRequestBefore bool) RequestContextBuilder
	// WithSendIdentifyRequestBefore adds SendIdentifyRequestBefore (property field)
	WithSendIdentifyRequestBefore(bool) RequestContextBuilder
	// Build builds the RequestContext or returns an error if something is wrong
	Build() (RequestContext, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RequestContext
}

// NewRequestContextBuilder() creates a RequestContextBuilder
func NewRequestContextBuilder() RequestContextBuilder {
	return &_RequestContextBuilder{_RequestContext: new(_RequestContext)}
}

type _RequestContextBuilder struct {
	*_RequestContext

	err *utils.MultiError
}

var _ (RequestContextBuilder) = (*_RequestContextBuilder)(nil)

func (b *_RequestContextBuilder) WithMandatoryFields(sendIdentifyRequestBefore bool) RequestContextBuilder {
	return b.WithSendIdentifyRequestBefore(sendIdentifyRequestBefore)
}

func (b *_RequestContextBuilder) WithSendIdentifyRequestBefore(sendIdentifyRequestBefore bool) RequestContextBuilder {
	b.SendIdentifyRequestBefore = sendIdentifyRequestBefore
	return b
}

func (b *_RequestContextBuilder) Build() (RequestContext, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RequestContext.deepCopy(), nil
}

func (b *_RequestContextBuilder) MustBuild() RequestContext {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RequestContextBuilder) DeepCopy() any {
	_copy := b.CreateRequestContextBuilder().(*_RequestContextBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRequestContextBuilder creates a RequestContextBuilder
func (b *_RequestContext) CreateRequestContextBuilder() RequestContextBuilder {
	if b == nil {
		return NewRequestContextBuilder()
	}
	return &_RequestContextBuilder{_RequestContext: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestContext) GetSendIdentifyRequestBefore() bool {
	return m.SendIdentifyRequestBefore
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRequestContext(structType any) RequestContext {
	if casted, ok := structType.(RequestContext); ok {
		return casted
	}
	if casted, ok := structType.(*RequestContext); ok {
		return *casted
	}
	return nil
}

func (m *_RequestContext) GetTypeName() string {
	return "RequestContext"
}

func (m *_RequestContext) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (sendIdentifyRequestBefore)
	lengthInBits += 1

	return lengthInBits
}

func (m *_RequestContext) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestContextParse(ctx context.Context, theBytes []byte) (RequestContext, error) {
	return RequestContextParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RequestContextParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (RequestContext, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (RequestContext, error) {
		return RequestContextParseWithBuffer(ctx, readBuffer)
	}
}

func RequestContextParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RequestContext, error) {
	v, err := (&_RequestContext{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_RequestContext) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__requestContext RequestContext, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestContext"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestContext")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sendIdentifyRequestBefore, err := ReadSimpleField(ctx, "sendIdentifyRequestBefore", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sendIdentifyRequestBefore' field"))
	}
	m.SendIdentifyRequestBefore = sendIdentifyRequestBefore

	if closeErr := readBuffer.CloseContext("RequestContext"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestContext")
	}

	return m, nil
}

func (m *_RequestContext) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestContext) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("RequestContext"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RequestContext")
	}

	if err := WriteSimpleField[bool](ctx, "sendIdentifyRequestBefore", m.GetSendIdentifyRequestBefore(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'sendIdentifyRequestBefore' field")
	}

	if popErr := writeBuffer.PopContext("RequestContext"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RequestContext")
	}
	return nil
}

func (m *_RequestContext) IsRequestContext() {}

func (m *_RequestContext) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RequestContext) deepCopy() *_RequestContext {
	if m == nil {
		return nil
	}
	_RequestContextCopy := &_RequestContext{
		m.SendIdentifyRequestBefore,
	}
	return _RequestContextCopy
}

func (m *_RequestContext) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
