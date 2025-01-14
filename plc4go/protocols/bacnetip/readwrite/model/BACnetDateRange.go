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

// BACnetDateRange is the corresponding interface of BACnetDateRange
type BACnetDateRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetStartDate returns StartDate (property field)
	GetStartDate() BACnetApplicationTagDate
	// GetEndDate returns EndDate (property field)
	GetEndDate() BACnetApplicationTagDate
	// IsBACnetDateRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDateRange()
	// CreateBuilder creates a BACnetDateRangeBuilder
	CreateBACnetDateRangeBuilder() BACnetDateRangeBuilder
}

// _BACnetDateRange is the data-structure of this message
type _BACnetDateRange struct {
	StartDate BACnetApplicationTagDate
	EndDate   BACnetApplicationTagDate
}

var _ BACnetDateRange = (*_BACnetDateRange)(nil)

// NewBACnetDateRange factory function for _BACnetDateRange
func NewBACnetDateRange(startDate BACnetApplicationTagDate, endDate BACnetApplicationTagDate) *_BACnetDateRange {
	if startDate == nil {
		panic("startDate of type BACnetApplicationTagDate for BACnetDateRange must not be nil")
	}
	if endDate == nil {
		panic("endDate of type BACnetApplicationTagDate for BACnetDateRange must not be nil")
	}
	return &_BACnetDateRange{StartDate: startDate, EndDate: endDate}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetDateRangeBuilder is a builder for BACnetDateRange
type BACnetDateRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startDate BACnetApplicationTagDate, endDate BACnetApplicationTagDate) BACnetDateRangeBuilder
	// WithStartDate adds StartDate (property field)
	WithStartDate(BACnetApplicationTagDate) BACnetDateRangeBuilder
	// WithStartDateBuilder adds StartDate (property field) which is build by the builder
	WithStartDateBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetDateRangeBuilder
	// WithEndDate adds EndDate (property field)
	WithEndDate(BACnetApplicationTagDate) BACnetDateRangeBuilder
	// WithEndDateBuilder adds EndDate (property field) which is build by the builder
	WithEndDateBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetDateRangeBuilder
	// Build builds the BACnetDateRange or returns an error if something is wrong
	Build() (BACnetDateRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetDateRange
}

// NewBACnetDateRangeBuilder() creates a BACnetDateRangeBuilder
func NewBACnetDateRangeBuilder() BACnetDateRangeBuilder {
	return &_BACnetDateRangeBuilder{_BACnetDateRange: new(_BACnetDateRange)}
}

type _BACnetDateRangeBuilder struct {
	*_BACnetDateRange

	err *utils.MultiError
}

var _ (BACnetDateRangeBuilder) = (*_BACnetDateRangeBuilder)(nil)

func (b *_BACnetDateRangeBuilder) WithMandatoryFields(startDate BACnetApplicationTagDate, endDate BACnetApplicationTagDate) BACnetDateRangeBuilder {
	return b.WithStartDate(startDate).WithEndDate(endDate)
}

func (b *_BACnetDateRangeBuilder) WithStartDate(startDate BACnetApplicationTagDate) BACnetDateRangeBuilder {
	b.StartDate = startDate
	return b
}

func (b *_BACnetDateRangeBuilder) WithStartDateBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetDateRangeBuilder {
	builder := builderSupplier(b.StartDate.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.StartDate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetDateRangeBuilder) WithEndDate(endDate BACnetApplicationTagDate) BACnetDateRangeBuilder {
	b.EndDate = endDate
	return b
}

func (b *_BACnetDateRangeBuilder) WithEndDateBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetDateRangeBuilder {
	builder := builderSupplier(b.EndDate.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.EndDate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetDateRangeBuilder) Build() (BACnetDateRange, error) {
	if b.StartDate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'startDate' not set"))
	}
	if b.EndDate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'endDate' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetDateRange.deepCopy(), nil
}

func (b *_BACnetDateRangeBuilder) MustBuild() BACnetDateRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetDateRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetDateRangeBuilder().(*_BACnetDateRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetDateRangeBuilder creates a BACnetDateRangeBuilder
func (b *_BACnetDateRange) CreateBACnetDateRangeBuilder() BACnetDateRangeBuilder {
	if b == nil {
		return NewBACnetDateRangeBuilder()
	}
	return &_BACnetDateRangeBuilder{_BACnetDateRange: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateRange) GetStartDate() BACnetApplicationTagDate {
	return m.StartDate
}

func (m *_BACnetDateRange) GetEndDate() BACnetApplicationTagDate {
	return m.EndDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetDateRange(structType any) BACnetDateRange {
	if casted, ok := structType.(BACnetDateRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateRange) GetTypeName() string {
	return "BACnetDateRange"
}

func (m *_BACnetDateRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (startDate)
	lengthInBits += m.StartDate.GetLengthInBits(ctx)

	// Simple field (endDate)
	lengthInBits += m.EndDate.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateRangeParse(ctx context.Context, theBytes []byte) (BACnetDateRange, error) {
	return BACnetDateRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDateRangeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
		return BACnetDateRangeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDateRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDateRange, error) {
	v, err := (&_BACnetDateRange{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetDateRange) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetDateRange BACnetDateRange, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "startDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startDate' field"))
	}
	m.StartDate = startDate

	endDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "endDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endDate' field"))
	}
	m.EndDate = endDate

	if closeErr := readBuffer.CloseContext("BACnetDateRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateRange")
	}

	return m, nil
}

func (m *_BACnetDateRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDateRange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateRange")
	}

	if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "startDate", m.GetStartDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'startDate' field")
	}

	if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "endDate", m.GetEndDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'endDate' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateRange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateRange")
	}
	return nil
}

func (m *_BACnetDateRange) IsBACnetDateRange() {}

func (m *_BACnetDateRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetDateRange) deepCopy() *_BACnetDateRange {
	if m == nil {
		return nil
	}
	_BACnetDateRangeCopy := &_BACnetDateRange{
		utils.DeepCopy[BACnetApplicationTagDate](m.StartDate),
		utils.DeepCopy[BACnetApplicationTagDate](m.EndDate),
	}
	return _BACnetDateRangeCopy
}

func (m *_BACnetDateRange) String() string {
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
