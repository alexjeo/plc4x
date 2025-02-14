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

// OpcuaMessageRequest is the corresponding interface of OpcuaMessageRequest
type OpcuaMessageRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MessagePDU
	// GetSecurityHeader returns SecurityHeader (property field)
	GetSecurityHeader() SecurityHeader
	// GetMessage returns Message (property field)
	GetMessage() Payload
	// IsOpcuaMessageRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaMessageRequest()
	// CreateBuilder creates a OpcuaMessageRequestBuilder
	CreateOpcuaMessageRequestBuilder() OpcuaMessageRequestBuilder
}

// _OpcuaMessageRequest is the data-structure of this message
type _OpcuaMessageRequest struct {
	MessagePDUContract
	SecurityHeader SecurityHeader
	Message        Payload

	// Arguments.
	TotalLength uint32
}

var _ OpcuaMessageRequest = (*_OpcuaMessageRequest)(nil)
var _ MessagePDURequirements = (*_OpcuaMessageRequest)(nil)

// NewOpcuaMessageRequest factory function for _OpcuaMessageRequest
func NewOpcuaMessageRequest(chunk ChunkType, securityHeader SecurityHeader, message Payload, totalLength uint32, binary bool) *_OpcuaMessageRequest {
	if securityHeader == nil {
		panic("securityHeader of type SecurityHeader for OpcuaMessageRequest must not be nil")
	}
	if message == nil {
		panic("message of type Payload for OpcuaMessageRequest must not be nil")
	}
	_result := &_OpcuaMessageRequest{
		MessagePDUContract: NewMessagePDU(chunk, binary),
		SecurityHeader:     securityHeader,
		Message:            message,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaMessageRequestBuilder is a builder for OpcuaMessageRequest
type OpcuaMessageRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(securityHeader SecurityHeader, message Payload) OpcuaMessageRequestBuilder
	// WithSecurityHeader adds SecurityHeader (property field)
	WithSecurityHeader(SecurityHeader) OpcuaMessageRequestBuilder
	// WithSecurityHeaderBuilder adds SecurityHeader (property field) which is build by the builder
	WithSecurityHeaderBuilder(func(SecurityHeaderBuilder) SecurityHeaderBuilder) OpcuaMessageRequestBuilder
	// WithMessage adds Message (property field)
	WithMessage(Payload) OpcuaMessageRequestBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(PayloadBuilder) PayloadBuilder) OpcuaMessageRequestBuilder
	// WithArgTotalLength sets a parser argument
	WithArgTotalLength(uint32) OpcuaMessageRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MessagePDUBuilder
	// Build builds the OpcuaMessageRequest or returns an error if something is wrong
	Build() (OpcuaMessageRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaMessageRequest
}

// NewOpcuaMessageRequestBuilder() creates a OpcuaMessageRequestBuilder
func NewOpcuaMessageRequestBuilder() OpcuaMessageRequestBuilder {
	return &_OpcuaMessageRequestBuilder{_OpcuaMessageRequest: new(_OpcuaMessageRequest)}
}

type _OpcuaMessageRequestBuilder struct {
	*_OpcuaMessageRequest

	parentBuilder *_MessagePDUBuilder

	err *utils.MultiError
}

var _ (OpcuaMessageRequestBuilder) = (*_OpcuaMessageRequestBuilder)(nil)

func (b *_OpcuaMessageRequestBuilder) setParent(contract MessagePDUContract) {
	b.MessagePDUContract = contract
	contract.(*_MessagePDU)._SubType = b._OpcuaMessageRequest
}

func (b *_OpcuaMessageRequestBuilder) WithMandatoryFields(securityHeader SecurityHeader, message Payload) OpcuaMessageRequestBuilder {
	return b.WithSecurityHeader(securityHeader).WithMessage(message)
}

func (b *_OpcuaMessageRequestBuilder) WithSecurityHeader(securityHeader SecurityHeader) OpcuaMessageRequestBuilder {
	b.SecurityHeader = securityHeader
	return b
}

func (b *_OpcuaMessageRequestBuilder) WithSecurityHeaderBuilder(builderSupplier func(SecurityHeaderBuilder) SecurityHeaderBuilder) OpcuaMessageRequestBuilder {
	builder := builderSupplier(b.SecurityHeader.CreateSecurityHeaderBuilder())
	var err error
	b.SecurityHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SecurityHeaderBuilder failed"))
	}
	return b
}

func (b *_OpcuaMessageRequestBuilder) WithMessage(message Payload) OpcuaMessageRequestBuilder {
	b.Message = message
	return b
}

func (b *_OpcuaMessageRequestBuilder) WithMessageBuilder(builderSupplier func(PayloadBuilder) PayloadBuilder) OpcuaMessageRequestBuilder {
	builder := builderSupplier(b.Message.CreatePayloadBuilder())
	var err error
	b.Message, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PayloadBuilder failed"))
	}
	return b
}

func (b *_OpcuaMessageRequestBuilder) WithArgTotalLength(totalLength uint32) OpcuaMessageRequestBuilder {
	b.TotalLength = totalLength
	return b
}

func (b *_OpcuaMessageRequestBuilder) Build() (OpcuaMessageRequest, error) {
	if b.SecurityHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'securityHeader' not set"))
	}
	if b.Message == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpcuaMessageRequest.deepCopy(), nil
}

func (b *_OpcuaMessageRequestBuilder) MustBuild() OpcuaMessageRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_OpcuaMessageRequestBuilder) Done() MessagePDUBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMessagePDUBuilder().(*_MessagePDUBuilder)
	}
	return b.parentBuilder
}

func (b *_OpcuaMessageRequestBuilder) buildForMessagePDU() (MessagePDU, error) {
	return b.Build()
}

func (b *_OpcuaMessageRequestBuilder) DeepCopy() any {
	_copy := b.CreateOpcuaMessageRequestBuilder().(*_OpcuaMessageRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpcuaMessageRequestBuilder creates a OpcuaMessageRequestBuilder
func (b *_OpcuaMessageRequest) CreateOpcuaMessageRequestBuilder() OpcuaMessageRequestBuilder {
	if b == nil {
		return NewOpcuaMessageRequestBuilder()
	}
	return &_OpcuaMessageRequestBuilder{_OpcuaMessageRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaMessageRequest) GetMessageType() string {
	return "MSG"
}

func (m *_OpcuaMessageRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaMessageRequest) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaMessageRequest) GetSecurityHeader() SecurityHeader {
	return m.SecurityHeader
}

func (m *_OpcuaMessageRequest) GetMessage() Payload {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaMessageRequest(structType any) OpcuaMessageRequest {
	if casted, ok := structType.(OpcuaMessageRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaMessageRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaMessageRequest) GetTypeName() string {
	return "OpcuaMessageRequest"
}

func (m *_OpcuaMessageRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (securityHeader)
	lengthInBits += m.SecurityHeader.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaMessageRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaMessageRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, totalLength uint32, response bool, binary bool) (__opcuaMessageRequest OpcuaMessageRequest, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaMessageRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaMessageRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	securityHeader, err := ReadSimpleField[SecurityHeader](ctx, "securityHeader", ReadComplex[SecurityHeader](SecurityHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityHeader' field"))
	}
	m.SecurityHeader = securityHeader

	message, err := ReadSimpleField[Payload](ctx, "message", ReadComplex[Payload](PayloadParseWithBufferProducer[Payload]((bool)(binary), (uint32)(uint32(uint32(totalLength)-uint32(securityHeader.GetLengthInBytes(ctx)))-uint32(uint32(16)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaMessageRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaMessageRequest")
	}

	return m, nil
}

func (m *_OpcuaMessageRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaMessageRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaMessageRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaMessageRequest")
		}

		if err := WriteSimpleField[SecurityHeader](ctx, "securityHeader", m.GetSecurityHeader(), WriteComplex[SecurityHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'securityHeader' field")
		}

		if err := WriteSimpleField[Payload](ctx, "message", m.GetMessage(), WriteComplex[Payload](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaMessageRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaMessageRequest")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_OpcuaMessageRequest) GetTotalLength() uint32 {
	return m.TotalLength
}

//
////

func (m *_OpcuaMessageRequest) IsOpcuaMessageRequest() {}

func (m *_OpcuaMessageRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaMessageRequest) deepCopy() *_OpcuaMessageRequest {
	if m == nil {
		return nil
	}
	_OpcuaMessageRequestCopy := &_OpcuaMessageRequest{
		m.MessagePDUContract.(*_MessagePDU).deepCopy(),
		utils.DeepCopy[SecurityHeader](m.SecurityHeader),
		utils.DeepCopy[Payload](m.Message),
		m.TotalLength,
	}
	_OpcuaMessageRequestCopy.MessagePDUContract.(*_MessagePDU)._SubType = m
	return _OpcuaMessageRequestCopy
}

func (m *_OpcuaMessageRequest) String() string {
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
