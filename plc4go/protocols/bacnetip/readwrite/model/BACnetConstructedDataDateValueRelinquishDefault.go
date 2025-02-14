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

// BACnetConstructedDataDateValueRelinquishDefault is the corresponding interface of BACnetConstructedDataDateValueRelinquishDefault
type BACnetConstructedDataDateValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
	// IsBACnetConstructedDataDateValueRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDateValueRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataDateValueRelinquishDefaultBuilder
	CreateBACnetConstructedDataDateValueRelinquishDefaultBuilder() BACnetConstructedDataDateValueRelinquishDefaultBuilder
}

// _BACnetConstructedDataDateValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataDateValueRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetApplicationTagDate
}

var _ BACnetConstructedDataDateValueRelinquishDefault = (*_BACnetConstructedDataDateValueRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDateValueRelinquishDefault)(nil)

// NewBACnetConstructedDataDateValueRelinquishDefault factory function for _BACnetConstructedDataDateValueRelinquishDefault
func NewBACnetConstructedDataDateValueRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetApplicationTagDate, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateValueRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetApplicationTagDate for BACnetConstructedDataDateValueRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataDateValueRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDateValueRelinquishDefaultBuilder is a builder for BACnetConstructedDataDateValueRelinquishDefault
type BACnetConstructedDataDateValueRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetApplicationTagDate) BACnetConstructedDataDateValueRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetApplicationTagDate) BACnetConstructedDataDateValueRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetConstructedDataDateValueRelinquishDefaultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDateValueRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataDateValueRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDateValueRelinquishDefault
}

// NewBACnetConstructedDataDateValueRelinquishDefaultBuilder() creates a BACnetConstructedDataDateValueRelinquishDefaultBuilder
func NewBACnetConstructedDataDateValueRelinquishDefaultBuilder() BACnetConstructedDataDateValueRelinquishDefaultBuilder {
	return &_BACnetConstructedDataDateValueRelinquishDefaultBuilder{_BACnetConstructedDataDateValueRelinquishDefault: new(_BACnetConstructedDataDateValueRelinquishDefault)}
}

type _BACnetConstructedDataDateValueRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataDateValueRelinquishDefault

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDateValueRelinquishDefaultBuilder) = (*_BACnetConstructedDataDateValueRelinquishDefaultBuilder)(nil)

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDateValueRelinquishDefault
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetApplicationTagDate) BACnetConstructedDataDateValueRelinquishDefaultBuilder {
	return b.WithRelinquishDefault(relinquishDefault)
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetApplicationTagDate) BACnetConstructedDataDateValueRelinquishDefaultBuilder {
	b.RelinquishDefault = relinquishDefault
	return b
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetConstructedDataDateValueRelinquishDefaultBuilder {
	builder := builderSupplier(b.RelinquishDefault.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.RelinquishDefault, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) Build() (BACnetConstructedDataDateValueRelinquishDefault, error) {
	if b.RelinquishDefault == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDateValueRelinquishDefault.deepCopy(), nil
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataDateValueRelinquishDefault {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDateValueRelinquishDefaultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDateValueRelinquishDefaultBuilder().(*_BACnetConstructedDataDateValueRelinquishDefaultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDateValueRelinquishDefaultBuilder creates a BACnetConstructedDataDateValueRelinquishDefaultBuilder
func (b *_BACnetConstructedDataDateValueRelinquishDefault) CreateBACnetConstructedDataDateValueRelinquishDefaultBuilder() BACnetConstructedDataDateValueRelinquishDefaultBuilder {
	if b == nil {
		return NewBACnetConstructedDataDateValueRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataDateValueRelinquishDefaultBuilder{_BACnetConstructedDataDateValueRelinquishDefault: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATE_VALUE
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagDate {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetActualValue() BACnetApplicationTagDate {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDate(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateValueRelinquishDefault(structType any) BACnetConstructedDataDateValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataDateValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataDateValueRelinquishDefault"
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDateValueRelinquishDefault BACnetConstructedDataDateValueRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "relinquishDefault", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetApplicationTagDate](ctx, "actualValue", (*BACnetApplicationTagDate)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateValueRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateValueRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateValueRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) IsBACnetConstructedDataDateValueRelinquishDefault() {
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) deepCopy() *_BACnetConstructedDataDateValueRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDateValueRelinquishDefaultCopy := &_BACnetConstructedDataDateValueRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDate](m.RelinquishDefault),
	}
	_BACnetConstructedDataDateValueRelinquishDefaultCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDateValueRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) String() string {
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
