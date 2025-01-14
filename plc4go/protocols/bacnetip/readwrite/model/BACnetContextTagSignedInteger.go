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

// BACnetContextTagSignedInteger is the corresponding interface of BACnetContextTagSignedInteger
type BACnetContextTagSignedInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
	// IsBACnetContextTagSignedInteger is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetContextTagSignedInteger()
	// CreateBuilder creates a BACnetContextTagSignedIntegerBuilder
	CreateBACnetContextTagSignedIntegerBuilder() BACnetContextTagSignedIntegerBuilder
}

// _BACnetContextTagSignedInteger is the data-structure of this message
type _BACnetContextTagSignedInteger struct {
	BACnetContextTagContract
	Payload BACnetTagPayloadSignedInteger
}

var _ BACnetContextTagSignedInteger = (*_BACnetContextTagSignedInteger)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagSignedInteger)(nil)

// NewBACnetContextTagSignedInteger factory function for _BACnetContextTagSignedInteger
func NewBACnetContextTagSignedInteger(header BACnetTagHeader, payload BACnetTagPayloadSignedInteger, tagNumberArgument uint8) *_BACnetContextTagSignedInteger {
	if payload == nil {
		panic("payload of type BACnetTagPayloadSignedInteger for BACnetContextTagSignedInteger must not be nil")
	}
	_result := &_BACnetContextTagSignedInteger{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetContextTagSignedIntegerBuilder is a builder for BACnetContextTagSignedInteger
type BACnetContextTagSignedIntegerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadSignedInteger) BACnetContextTagSignedIntegerBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadSignedInteger) BACnetContextTagSignedIntegerBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadSignedIntegerBuilder) BACnetTagPayloadSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetContextTagBuilder
	// Build builds the BACnetContextTagSignedInteger or returns an error if something is wrong
	Build() (BACnetContextTagSignedInteger, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetContextTagSignedInteger
}

// NewBACnetContextTagSignedIntegerBuilder() creates a BACnetContextTagSignedIntegerBuilder
func NewBACnetContextTagSignedIntegerBuilder() BACnetContextTagSignedIntegerBuilder {
	return &_BACnetContextTagSignedIntegerBuilder{_BACnetContextTagSignedInteger: new(_BACnetContextTagSignedInteger)}
}

type _BACnetContextTagSignedIntegerBuilder struct {
	*_BACnetContextTagSignedInteger

	parentBuilder *_BACnetContextTagBuilder

	err *utils.MultiError
}

var _ (BACnetContextTagSignedIntegerBuilder) = (*_BACnetContextTagSignedIntegerBuilder)(nil)

func (b *_BACnetContextTagSignedIntegerBuilder) setParent(contract BACnetContextTagContract) {
	b.BACnetContextTagContract = contract
	contract.(*_BACnetContextTag)._SubType = b._BACnetContextTagSignedInteger
}

func (b *_BACnetContextTagSignedIntegerBuilder) WithMandatoryFields(payload BACnetTagPayloadSignedInteger) BACnetContextTagSignedIntegerBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetContextTagSignedIntegerBuilder) WithPayload(payload BACnetTagPayloadSignedInteger) BACnetContextTagSignedIntegerBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetContextTagSignedIntegerBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadSignedIntegerBuilder) BACnetTagPayloadSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadSignedIntegerBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetContextTagSignedIntegerBuilder) Build() (BACnetContextTagSignedInteger, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetContextTagSignedInteger.deepCopy(), nil
}

func (b *_BACnetContextTagSignedIntegerBuilder) MustBuild() BACnetContextTagSignedInteger {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetContextTagSignedIntegerBuilder) Done() BACnetContextTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetContextTagBuilder().(*_BACnetContextTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetContextTagSignedIntegerBuilder) buildForBACnetContextTag() (BACnetContextTag, error) {
	return b.Build()
}

func (b *_BACnetContextTagSignedIntegerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetContextTagSignedIntegerBuilder().(*_BACnetContextTagSignedIntegerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetContextTagSignedIntegerBuilder creates a BACnetContextTagSignedIntegerBuilder
func (b *_BACnetContextTagSignedInteger) CreateBACnetContextTagSignedIntegerBuilder() BACnetContextTagSignedIntegerBuilder {
	if b == nil {
		return NewBACnetContextTagSignedIntegerBuilder()
	}
	return &_BACnetContextTagSignedIntegerBuilder{_BACnetContextTagSignedInteger: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagSignedInteger) GetDataType() BACnetDataType {
	return BACnetDataType_SIGNED_INTEGER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagSignedInteger) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagSignedInteger) GetPayload() BACnetTagPayloadSignedInteger {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagSignedInteger) GetActualValue() uint64 {
	ctx := context.Background()
	_ = ctx
	return uint64(m.GetPayload().GetActualValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetContextTagSignedInteger(structType any) BACnetContextTagSignedInteger {
	if casted, ok := structType.(BACnetContextTagSignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagSignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagSignedInteger) GetTypeName() string {
	return "BACnetContextTagSignedInteger"
}

func (m *_BACnetContextTagSignedInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagSignedInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagSignedInteger) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagSignedInteger BACnetContextTagSignedInteger, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagSignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagSignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadSignedInteger](ctx, "payload", ReadComplex[BACnetTagPayloadSignedInteger](BACnetTagPayloadSignedIntegerParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[uint64](ctx, "actualValue", (*uint64)(nil), payload.GetActualValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagSignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagSignedInteger")
	}

	return m, nil
}

func (m *_BACnetContextTagSignedInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagSignedInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagSignedInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagSignedInteger")
		}

		if err := WriteSimpleField[BACnetTagPayloadSignedInteger](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagSignedInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagSignedInteger")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagSignedInteger) IsBACnetContextTagSignedInteger() {}

func (m *_BACnetContextTagSignedInteger) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetContextTagSignedInteger) deepCopy() *_BACnetContextTagSignedInteger {
	if m == nil {
		return nil
	}
	_BACnetContextTagSignedIntegerCopy := &_BACnetContextTagSignedInteger{
		m.BACnetContextTagContract.(*_BACnetContextTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadSignedInteger](m.Payload),
	}
	_BACnetContextTagSignedIntegerCopy.BACnetContextTagContract.(*_BACnetContextTag)._SubType = m
	return _BACnetContextTagSignedIntegerCopy
}

func (m *_BACnetContextTagSignedInteger) String() string {
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
