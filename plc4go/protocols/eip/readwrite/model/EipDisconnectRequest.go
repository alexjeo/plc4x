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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// EipDisconnectRequest is the corresponding interface of EipDisconnectRequest
type EipDisconnectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// IsEipDisconnectRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEipDisconnectRequest()
	// CreateBuilder creates a EipDisconnectRequestBuilder
	CreateEipDisconnectRequestBuilder() EipDisconnectRequestBuilder
}

// _EipDisconnectRequest is the data-structure of this message
type _EipDisconnectRequest struct {
	EipPacketContract
}

var _ EipDisconnectRequest = (*_EipDisconnectRequest)(nil)
var _ EipPacketRequirements = (*_EipDisconnectRequest)(nil)

// NewEipDisconnectRequest factory function for _EipDisconnectRequest
func NewEipDisconnectRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipDisconnectRequest {
	_result := &_EipDisconnectRequest{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EipDisconnectRequestBuilder is a builder for EipDisconnectRequest
type EipDisconnectRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() EipDisconnectRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() EipPacketBuilder
	// Build builds the EipDisconnectRequest or returns an error if something is wrong
	Build() (EipDisconnectRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EipDisconnectRequest
}

// NewEipDisconnectRequestBuilder() creates a EipDisconnectRequestBuilder
func NewEipDisconnectRequestBuilder() EipDisconnectRequestBuilder {
	return &_EipDisconnectRequestBuilder{_EipDisconnectRequest: new(_EipDisconnectRequest)}
}

type _EipDisconnectRequestBuilder struct {
	*_EipDisconnectRequest

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (EipDisconnectRequestBuilder) = (*_EipDisconnectRequestBuilder)(nil)

func (b *_EipDisconnectRequestBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
	contract.(*_EipPacket)._SubType = b._EipDisconnectRequest
}

func (b *_EipDisconnectRequestBuilder) WithMandatoryFields() EipDisconnectRequestBuilder {
	return b
}

func (b *_EipDisconnectRequestBuilder) Build() (EipDisconnectRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EipDisconnectRequest.deepCopy(), nil
}

func (b *_EipDisconnectRequestBuilder) MustBuild() EipDisconnectRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EipDisconnectRequestBuilder) Done() EipPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewEipPacketBuilder().(*_EipPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_EipDisconnectRequestBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_EipDisconnectRequestBuilder) DeepCopy() any {
	_copy := b.CreateEipDisconnectRequestBuilder().(*_EipDisconnectRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEipDisconnectRequestBuilder creates a EipDisconnectRequestBuilder
func (b *_EipDisconnectRequest) CreateEipDisconnectRequestBuilder() EipDisconnectRequestBuilder {
	if b == nil {
		return NewEipDisconnectRequestBuilder()
	}
	return &_EipDisconnectRequestBuilder{_EipDisconnectRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipDisconnectRequest) GetCommand() uint16 {
	return 0x0066
}

func (m *_EipDisconnectRequest) GetResponse() bool {
	return false
}

func (m *_EipDisconnectRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipDisconnectRequest) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// Deprecated: use the interface for direct cast
func CastEipDisconnectRequest(structType any) EipDisconnectRequest {
	if casted, ok := structType.(EipDisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*EipDisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_EipDisconnectRequest) GetTypeName() string {
	return "EipDisconnectRequest"
}

func (m *_EipDisconnectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_EipDisconnectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EipDisconnectRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__eipDisconnectRequest EipDisconnectRequest, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipDisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipDisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EipDisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipDisconnectRequest")
	}

	return m, nil
}

func (m *_EipDisconnectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipDisconnectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipDisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipDisconnectRequest")
		}

		if popErr := writeBuffer.PopContext("EipDisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipDisconnectRequest")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipDisconnectRequest) IsEipDisconnectRequest() {}

func (m *_EipDisconnectRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EipDisconnectRequest) deepCopy() *_EipDisconnectRequest {
	if m == nil {
		return nil
	}
	_EipDisconnectRequestCopy := &_EipDisconnectRequest{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
	}
	_EipDisconnectRequestCopy.EipPacketContract.(*_EipPacket)._SubType = m
	return _EipDisconnectRequestCopy
}

func (m *_EipDisconnectRequest) String() string {
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
