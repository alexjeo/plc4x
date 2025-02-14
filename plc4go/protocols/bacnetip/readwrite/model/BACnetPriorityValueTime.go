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

// BACnetPriorityValueTime is the corresponding interface of BACnetPriorityValueTime
type BACnetPriorityValueTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// IsBACnetPriorityValueTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueTime()
	// CreateBuilder creates a BACnetPriorityValueTimeBuilder
	CreateBACnetPriorityValueTimeBuilder() BACnetPriorityValueTimeBuilder
}

// _BACnetPriorityValueTime is the data-structure of this message
type _BACnetPriorityValueTime struct {
	BACnetPriorityValueContract
	TimeValue BACnetApplicationTagTime
}

var _ BACnetPriorityValueTime = (*_BACnetPriorityValueTime)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueTime)(nil)

// NewBACnetPriorityValueTime factory function for _BACnetPriorityValueTime
func NewBACnetPriorityValueTime(peekedTagHeader BACnetTagHeader, timeValue BACnetApplicationTagTime, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueTime {
	if timeValue == nil {
		panic("timeValue of type BACnetApplicationTagTime for BACnetPriorityValueTime must not be nil")
	}
	_result := &_BACnetPriorityValueTime{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		TimeValue:                   timeValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueTimeBuilder is a builder for BACnetPriorityValueTime
type BACnetPriorityValueTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetPriorityValueTimeBuilder
	// WithTimeValue adds TimeValue (property field)
	WithTimeValue(BACnetApplicationTagTime) BACnetPriorityValueTimeBuilder
	// WithTimeValueBuilder adds TimeValue (property field) which is build by the builder
	WithTimeValueBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetPriorityValueTimeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPriorityValueBuilder
	// Build builds the BACnetPriorityValueTime or returns an error if something is wrong
	Build() (BACnetPriorityValueTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueTime
}

// NewBACnetPriorityValueTimeBuilder() creates a BACnetPriorityValueTimeBuilder
func NewBACnetPriorityValueTimeBuilder() BACnetPriorityValueTimeBuilder {
	return &_BACnetPriorityValueTimeBuilder{_BACnetPriorityValueTime: new(_BACnetPriorityValueTime)}
}

type _BACnetPriorityValueTimeBuilder struct {
	*_BACnetPriorityValueTime

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueTimeBuilder) = (*_BACnetPriorityValueTimeBuilder)(nil)

func (b *_BACnetPriorityValueTimeBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
	contract.(*_BACnetPriorityValue)._SubType = b._BACnetPriorityValueTime
}

func (b *_BACnetPriorityValueTimeBuilder) WithMandatoryFields(timeValue BACnetApplicationTagTime) BACnetPriorityValueTimeBuilder {
	return b.WithTimeValue(timeValue)
}

func (b *_BACnetPriorityValueTimeBuilder) WithTimeValue(timeValue BACnetApplicationTagTime) BACnetPriorityValueTimeBuilder {
	b.TimeValue = timeValue
	return b
}

func (b *_BACnetPriorityValueTimeBuilder) WithTimeValueBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetPriorityValueTimeBuilder {
	builder := builderSupplier(b.TimeValue.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.TimeValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetPriorityValueTimeBuilder) Build() (BACnetPriorityValueTime, error) {
	if b.TimeValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueTime.deepCopy(), nil
}

func (b *_BACnetPriorityValueTimeBuilder) MustBuild() BACnetPriorityValueTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPriorityValueTimeBuilder) Done() BACnetPriorityValueBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPriorityValueBuilder().(*_BACnetPriorityValueBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPriorityValueTimeBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueTimeBuilder().(*_BACnetPriorityValueTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueTimeBuilder creates a BACnetPriorityValueTimeBuilder
func (b *_BACnetPriorityValueTime) CreateBACnetPriorityValueTimeBuilder() BACnetPriorityValueTimeBuilder {
	if b == nil {
		return NewBACnetPriorityValueTimeBuilder()
	}
	return &_BACnetPriorityValueTimeBuilder{_BACnetPriorityValueTime: b.deepCopy()}
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

func (m *_BACnetPriorityValueTime) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueTime) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueTime(structType any) BACnetPriorityValueTime {
	if casted, ok := structType.(BACnetPriorityValueTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueTime) GetTypeName() string {
	return "BACnetPriorityValueTime"
}

func (m *_BACnetPriorityValueTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (timeValue)
	lengthInBits += m.TimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueTime BACnetPriorityValueTime, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeValue, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	m.TimeValue = timeValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueTime")
	}

	return m, nil
}

func (m *_BACnetPriorityValueTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "timeValue", m.GetTimeValue(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueTime")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueTime) IsBACnetPriorityValueTime() {}

func (m *_BACnetPriorityValueTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueTime) deepCopy() *_BACnetPriorityValueTime {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueTimeCopy := &_BACnetPriorityValueTime{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagTime](m.TimeValue),
	}
	_BACnetPriorityValueTimeCopy.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueTimeCopy
}

func (m *_BACnetPriorityValueTime) String() string {
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
