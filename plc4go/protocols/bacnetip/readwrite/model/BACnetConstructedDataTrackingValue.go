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

// BACnetConstructedDataTrackingValue is the corresponding interface of BACnetConstructedDataTrackingValue
type BACnetConstructedDataTrackingValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTrackingValue returns TrackingValue (property field)
	GetTrackingValue() BACnetLifeSafetyStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLifeSafetyStateTagged
	// IsBACnetConstructedDataTrackingValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTrackingValue()
	// CreateBuilder creates a BACnetConstructedDataTrackingValueBuilder
	CreateBACnetConstructedDataTrackingValueBuilder() BACnetConstructedDataTrackingValueBuilder
}

// _BACnetConstructedDataTrackingValue is the data-structure of this message
type _BACnetConstructedDataTrackingValue struct {
	BACnetConstructedDataContract
	TrackingValue BACnetLifeSafetyStateTagged
}

var _ BACnetConstructedDataTrackingValue = (*_BACnetConstructedDataTrackingValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTrackingValue)(nil)

// NewBACnetConstructedDataTrackingValue factory function for _BACnetConstructedDataTrackingValue
func NewBACnetConstructedDataTrackingValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, trackingValue BACnetLifeSafetyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrackingValue {
	if trackingValue == nil {
		panic("trackingValue of type BACnetLifeSafetyStateTagged for BACnetConstructedDataTrackingValue must not be nil")
	}
	_result := &_BACnetConstructedDataTrackingValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TrackingValue:                 trackingValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTrackingValueBuilder is a builder for BACnetConstructedDataTrackingValue
type BACnetConstructedDataTrackingValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(trackingValue BACnetLifeSafetyStateTagged) BACnetConstructedDataTrackingValueBuilder
	// WithTrackingValue adds TrackingValue (property field)
	WithTrackingValue(BACnetLifeSafetyStateTagged) BACnetConstructedDataTrackingValueBuilder
	// WithTrackingValueBuilder adds TrackingValue (property field) which is build by the builder
	WithTrackingValueBuilder(func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetConstructedDataTrackingValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTrackingValue or returns an error if something is wrong
	Build() (BACnetConstructedDataTrackingValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTrackingValue
}

// NewBACnetConstructedDataTrackingValueBuilder() creates a BACnetConstructedDataTrackingValueBuilder
func NewBACnetConstructedDataTrackingValueBuilder() BACnetConstructedDataTrackingValueBuilder {
	return &_BACnetConstructedDataTrackingValueBuilder{_BACnetConstructedDataTrackingValue: new(_BACnetConstructedDataTrackingValue)}
}

type _BACnetConstructedDataTrackingValueBuilder struct {
	*_BACnetConstructedDataTrackingValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTrackingValueBuilder) = (*_BACnetConstructedDataTrackingValueBuilder)(nil)

func (b *_BACnetConstructedDataTrackingValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataTrackingValue
}

func (b *_BACnetConstructedDataTrackingValueBuilder) WithMandatoryFields(trackingValue BACnetLifeSafetyStateTagged) BACnetConstructedDataTrackingValueBuilder {
	return b.WithTrackingValue(trackingValue)
}

func (b *_BACnetConstructedDataTrackingValueBuilder) WithTrackingValue(trackingValue BACnetLifeSafetyStateTagged) BACnetConstructedDataTrackingValueBuilder {
	b.TrackingValue = trackingValue
	return b
}

func (b *_BACnetConstructedDataTrackingValueBuilder) WithTrackingValueBuilder(builderSupplier func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetConstructedDataTrackingValueBuilder {
	builder := builderSupplier(b.TrackingValue.CreateBACnetLifeSafetyStateTaggedBuilder())
	var err error
	b.TrackingValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTrackingValueBuilder) Build() (BACnetConstructedDataTrackingValue, error) {
	if b.TrackingValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'trackingValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTrackingValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataTrackingValueBuilder) MustBuild() BACnetConstructedDataTrackingValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTrackingValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTrackingValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTrackingValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTrackingValueBuilder().(*_BACnetConstructedDataTrackingValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTrackingValueBuilder creates a BACnetConstructedDataTrackingValueBuilder
func (b *_BACnetConstructedDataTrackingValue) CreateBACnetConstructedDataTrackingValueBuilder() BACnetConstructedDataTrackingValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataTrackingValueBuilder()
	}
	return &_BACnetConstructedDataTrackingValueBuilder{_BACnetConstructedDataTrackingValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrackingValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTrackingValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRACKING_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrackingValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrackingValue) GetTrackingValue() BACnetLifeSafetyStateTagged {
	return m.TrackingValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTrackingValue) GetActualValue() BACnetLifeSafetyStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLifeSafetyStateTagged(m.GetTrackingValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrackingValue(structType any) BACnetConstructedDataTrackingValue {
	if casted, ok := structType.(BACnetConstructedDataTrackingValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrackingValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrackingValue) GetTypeName() string {
	return "BACnetConstructedDataTrackingValue"
}

func (m *_BACnetConstructedDataTrackingValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (trackingValue)
	lengthInBits += m.TrackingValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTrackingValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTrackingValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTrackingValue BACnetConstructedDataTrackingValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrackingValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrackingValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	trackingValue, err := ReadSimpleField[BACnetLifeSafetyStateTagged](ctx, "trackingValue", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'trackingValue' field"))
	}
	m.TrackingValue = trackingValue

	actualValue, err := ReadVirtualField[BACnetLifeSafetyStateTagged](ctx, "actualValue", (*BACnetLifeSafetyStateTagged)(nil), trackingValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrackingValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrackingValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTrackingValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrackingValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrackingValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrackingValue")
		}

		if err := WriteSimpleField[BACnetLifeSafetyStateTagged](ctx, "trackingValue", m.GetTrackingValue(), WriteComplex[BACnetLifeSafetyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'trackingValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrackingValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrackingValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrackingValue) IsBACnetConstructedDataTrackingValue() {}

func (m *_BACnetConstructedDataTrackingValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTrackingValue) deepCopy() *_BACnetConstructedDataTrackingValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTrackingValueCopy := &_BACnetConstructedDataTrackingValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetLifeSafetyStateTagged](m.TrackingValue),
	}
	_BACnetConstructedDataTrackingValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTrackingValueCopy
}

func (m *_BACnetConstructedDataTrackingValue) String() string {
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
