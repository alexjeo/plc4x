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

// BACnetConstructedDataChangesPending is the corresponding interface of BACnetConstructedDataChangesPending
type BACnetConstructedDataChangesPending interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetChangesPending returns ChangesPending (property field)
	GetChangesPending() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataChangesPending is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataChangesPending()
	// CreateBuilder creates a BACnetConstructedDataChangesPendingBuilder
	CreateBACnetConstructedDataChangesPendingBuilder() BACnetConstructedDataChangesPendingBuilder
}

// _BACnetConstructedDataChangesPending is the data-structure of this message
type _BACnetConstructedDataChangesPending struct {
	BACnetConstructedDataContract
	ChangesPending BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataChangesPending = (*_BACnetConstructedDataChangesPending)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataChangesPending)(nil)

// NewBACnetConstructedDataChangesPending factory function for _BACnetConstructedDataChangesPending
func NewBACnetConstructedDataChangesPending(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, changesPending BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChangesPending {
	if changesPending == nil {
		panic("changesPending of type BACnetApplicationTagBoolean for BACnetConstructedDataChangesPending must not be nil")
	}
	_result := &_BACnetConstructedDataChangesPending{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ChangesPending:                changesPending,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataChangesPendingBuilder is a builder for BACnetConstructedDataChangesPending
type BACnetConstructedDataChangesPendingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(changesPending BACnetApplicationTagBoolean) BACnetConstructedDataChangesPendingBuilder
	// WithChangesPending adds ChangesPending (property field)
	WithChangesPending(BACnetApplicationTagBoolean) BACnetConstructedDataChangesPendingBuilder
	// WithChangesPendingBuilder adds ChangesPending (property field) which is build by the builder
	WithChangesPendingBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataChangesPendingBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataChangesPending or returns an error if something is wrong
	Build() (BACnetConstructedDataChangesPending, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataChangesPending
}

// NewBACnetConstructedDataChangesPendingBuilder() creates a BACnetConstructedDataChangesPendingBuilder
func NewBACnetConstructedDataChangesPendingBuilder() BACnetConstructedDataChangesPendingBuilder {
	return &_BACnetConstructedDataChangesPendingBuilder{_BACnetConstructedDataChangesPending: new(_BACnetConstructedDataChangesPending)}
}

type _BACnetConstructedDataChangesPendingBuilder struct {
	*_BACnetConstructedDataChangesPending

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataChangesPendingBuilder) = (*_BACnetConstructedDataChangesPendingBuilder)(nil)

func (b *_BACnetConstructedDataChangesPendingBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataChangesPending
}

func (b *_BACnetConstructedDataChangesPendingBuilder) WithMandatoryFields(changesPending BACnetApplicationTagBoolean) BACnetConstructedDataChangesPendingBuilder {
	return b.WithChangesPending(changesPending)
}

func (b *_BACnetConstructedDataChangesPendingBuilder) WithChangesPending(changesPending BACnetApplicationTagBoolean) BACnetConstructedDataChangesPendingBuilder {
	b.ChangesPending = changesPending
	return b
}

func (b *_BACnetConstructedDataChangesPendingBuilder) WithChangesPendingBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataChangesPendingBuilder {
	builder := builderSupplier(b.ChangesPending.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.ChangesPending, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataChangesPendingBuilder) Build() (BACnetConstructedDataChangesPending, error) {
	if b.ChangesPending == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'changesPending' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataChangesPending.deepCopy(), nil
}

func (b *_BACnetConstructedDataChangesPendingBuilder) MustBuild() BACnetConstructedDataChangesPending {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataChangesPendingBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataChangesPendingBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataChangesPendingBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataChangesPendingBuilder().(*_BACnetConstructedDataChangesPendingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataChangesPendingBuilder creates a BACnetConstructedDataChangesPendingBuilder
func (b *_BACnetConstructedDataChangesPending) CreateBACnetConstructedDataChangesPendingBuilder() BACnetConstructedDataChangesPendingBuilder {
	if b == nil {
		return NewBACnetConstructedDataChangesPendingBuilder()
	}
	return &_BACnetConstructedDataChangesPendingBuilder{_BACnetConstructedDataChangesPending: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChangesPending) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChangesPending) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANGES_PENDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChangesPending) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChangesPending) GetChangesPending() BACnetApplicationTagBoolean {
	return m.ChangesPending
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChangesPending) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetChangesPending())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChangesPending(structType any) BACnetConstructedDataChangesPending {
	if casted, ok := structType.(BACnetConstructedDataChangesPending); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChangesPending); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChangesPending) GetTypeName() string {
	return "BACnetConstructedDataChangesPending"
}

func (m *_BACnetConstructedDataChangesPending) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (changesPending)
	lengthInBits += m.ChangesPending.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChangesPending) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataChangesPending) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataChangesPending BACnetConstructedDataChangesPending, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChangesPending"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChangesPending")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	changesPending, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "changesPending", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'changesPending' field"))
	}
	m.ChangesPending = changesPending

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), changesPending)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChangesPending"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChangesPending")
	}

	return m, nil
}

func (m *_BACnetConstructedDataChangesPending) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChangesPending) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChangesPending"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChangesPending")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "changesPending", m.GetChangesPending(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'changesPending' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChangesPending"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChangesPending")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChangesPending) IsBACnetConstructedDataChangesPending() {}

func (m *_BACnetConstructedDataChangesPending) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataChangesPending) deepCopy() *_BACnetConstructedDataChangesPending {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataChangesPendingCopy := &_BACnetConstructedDataChangesPending{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.ChangesPending),
	}
	_BACnetConstructedDataChangesPendingCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataChangesPendingCopy
}

func (m *_BACnetConstructedDataChangesPending) String() string {
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
