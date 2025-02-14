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

// BACnetConfirmedServiceRequestReadRangeRangeByTime is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeByTime
type BACnetConfirmedServiceRequestReadRangeRangeByTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceTime returns ReferenceTime (property field)
	GetReferenceTime() BACnetDateTime
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagSignedInteger
	// IsBACnetConfirmedServiceRequestReadRangeRangeByTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadRangeRangeByTime()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
	CreateBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder() BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
}

// _BACnetConfirmedServiceRequestReadRangeRangeByTime is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRangeByTime struct {
	BACnetConfirmedServiceRequestReadRangeRangeContract
	ReferenceTime BACnetDateTime
	Count         BACnetApplicationTagSignedInteger
}

var _ BACnetConfirmedServiceRequestReadRangeRangeByTime = (*_BACnetConfirmedServiceRequestReadRangeRangeByTime)(nil)
var _ BACnetConfirmedServiceRequestReadRangeRangeRequirements = (*_BACnetConfirmedServiceRequestReadRangeRangeByTime)(nil)

// NewBACnetConfirmedServiceRequestReadRangeRangeByTime factory function for _BACnetConfirmedServiceRequestReadRangeRangeByTime
func NewBACnetConfirmedServiceRequestReadRangeRangeByTime(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag, referenceTime BACnetDateTime, count BACnetApplicationTagSignedInteger) *_BACnetConfirmedServiceRequestReadRangeRangeByTime {
	if referenceTime == nil {
		panic("referenceTime of type BACnetDateTime for BACnetConfirmedServiceRequestReadRangeRangeByTime must not be nil")
	}
	if count == nil {
		panic("count of type BACnetApplicationTagSignedInteger for BACnetConfirmedServiceRequestReadRangeRangeByTime must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestReadRangeRangeByTime{
		BACnetConfirmedServiceRequestReadRangeRangeContract: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
		ReferenceTime: referenceTime,
		Count:         count,
	}
	_result.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder is a builder for BACnetConfirmedServiceRequestReadRangeRangeByTime
type BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referenceTime BACnetDateTime, count BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
	// WithReferenceTime adds ReferenceTime (property field)
	WithReferenceTime(BACnetDateTime) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
	// WithReferenceTimeBuilder adds ReferenceTime (property field) which is build by the builder
	WithReferenceTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
	// WithCount adds Count (property field)
	WithCount(BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
	// WithCountBuilder adds Count (property field) which is build by the builder
	WithCountBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestReadRangeRangeBuilder
	// Build builds the BACnetConfirmedServiceRequestReadRangeRangeByTime or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReadRangeRangeByTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReadRangeRangeByTime
}

// NewBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder() creates a BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
func NewBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder() BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	return &_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder{_BACnetConfirmedServiceRequestReadRangeRangeByTime: new(_BACnetConfirmedServiceRequestReadRangeRangeByTime)}
}

type _BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder struct {
	*_BACnetConfirmedServiceRequestReadRangeRangeByTime

	parentBuilder *_BACnetConfirmedServiceRequestReadRangeRangeBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) = (*_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) setParent(contract BACnetConfirmedServiceRequestReadRangeRangeContract) {
	b.BACnetConfirmedServiceRequestReadRangeRangeContract = contract
	contract.(*_BACnetConfirmedServiceRequestReadRangeRange)._SubType = b._BACnetConfirmedServiceRequestReadRangeRangeByTime
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) WithMandatoryFields(referenceTime BACnetDateTime, count BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	return b.WithReferenceTime(referenceTime).WithCount(count)
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) WithReferenceTime(referenceTime BACnetDateTime) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	b.ReferenceTime = referenceTime
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) WithReferenceTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	builder := builderSupplier(b.ReferenceTime.CreateBACnetDateTimeBuilder())
	var err error
	b.ReferenceTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) WithCount(count BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	b.Count = count
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) WithCountBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	builder := builderSupplier(b.Count.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.Count, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) Build() (BACnetConfirmedServiceRequestReadRangeRangeByTime, error) {
	if b.ReferenceTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceTime' not set"))
	}
	if b.Count == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'count' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestReadRangeRangeByTime.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) MustBuild() BACnetConfirmedServiceRequestReadRangeRangeByTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) Done() BACnetConfirmedServiceRequestReadRangeRangeBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestReadRangeRangeBuilder().(*_BACnetConfirmedServiceRequestReadRangeRangeBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) buildForBACnetConfirmedServiceRequestReadRangeRange() (BACnetConfirmedServiceRequestReadRangeRange, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder().(*_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder creates a BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder
func (b *_BACnetConfirmedServiceRequestReadRangeRangeByTime) CreateBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder() BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder()
	}
	return &_BACnetConfirmedServiceRequestReadRangeRangeByTimeBuilder{_BACnetConfirmedServiceRequestReadRangeRangeByTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetParent() BACnetConfirmedServiceRequestReadRangeRangeContract {
	return m.BACnetConfirmedServiceRequestReadRangeRangeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetReferenceTime() BACnetDateTime {
	return m.ReferenceTime
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetCount() BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRangeByTime(structType any) BACnetConfirmedServiceRequestReadRangeRangeByTime {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeByTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeByTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeByTime"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).getLengthInBits(ctx))

	// Simple field (referenceTime)
	lengthInBits += m.ReferenceTime.GetLengthInBits(ctx)

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestReadRangeRange) (__bACnetConfirmedServiceRequestReadRangeRangeByTime BACnetConfirmedServiceRequestReadRangeRangeByTime, err error) {
	m.BACnetConfirmedServiceRequestReadRangeRangeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeByTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceTime, err := ReadSimpleField[BACnetDateTime](ctx, "referenceTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTime' field"))
	}
	m.ReferenceTime = referenceTime

	count, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "count", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}
	m.Count = count

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeByTime")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeByTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "referenceTime", m.GetReferenceTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTime' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "count", m.GetCount(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeByTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeByTime")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) IsBACnetConfirmedServiceRequestReadRangeRangeByTime() {
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) deepCopy() *_BACnetConfirmedServiceRequestReadRangeRangeByTime {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReadRangeRangeByTimeCopy := &_BACnetConfirmedServiceRequestReadRangeRangeByTime{
		m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).deepCopy(),
		utils.DeepCopy[BACnetDateTime](m.ReferenceTime),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.Count),
	}
	_BACnetConfirmedServiceRequestReadRangeRangeByTimeCopy.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange)._SubType = m
	return _BACnetConfirmedServiceRequestReadRangeRangeByTimeCopy
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByTime) String() string {
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
