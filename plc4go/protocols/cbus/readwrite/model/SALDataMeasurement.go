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

// SALDataMeasurement is the corresponding interface of SALDataMeasurement
type SALDataMeasurement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetMeasurementData returns MeasurementData (property field)
	GetMeasurementData() MeasurementData
	// IsSALDataMeasurement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataMeasurement()
	// CreateBuilder creates a SALDataMeasurementBuilder
	CreateSALDataMeasurementBuilder() SALDataMeasurementBuilder
}

// _SALDataMeasurement is the data-structure of this message
type _SALDataMeasurement struct {
	SALDataContract
	MeasurementData MeasurementData
}

var _ SALDataMeasurement = (*_SALDataMeasurement)(nil)
var _ SALDataRequirements = (*_SALDataMeasurement)(nil)

// NewSALDataMeasurement factory function for _SALDataMeasurement
func NewSALDataMeasurement(salData SALData, measurementData MeasurementData) *_SALDataMeasurement {
	if measurementData == nil {
		panic("measurementData of type MeasurementData for SALDataMeasurement must not be nil")
	}
	_result := &_SALDataMeasurement{
		SALDataContract: NewSALData(salData),
		MeasurementData: measurementData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataMeasurementBuilder is a builder for SALDataMeasurement
type SALDataMeasurementBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(measurementData MeasurementData) SALDataMeasurementBuilder
	// WithMeasurementData adds MeasurementData (property field)
	WithMeasurementData(MeasurementData) SALDataMeasurementBuilder
	// WithMeasurementDataBuilder adds MeasurementData (property field) which is build by the builder
	WithMeasurementDataBuilder(func(MeasurementDataBuilder) MeasurementDataBuilder) SALDataMeasurementBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SALDataBuilder
	// Build builds the SALDataMeasurement or returns an error if something is wrong
	Build() (SALDataMeasurement, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataMeasurement
}

// NewSALDataMeasurementBuilder() creates a SALDataMeasurementBuilder
func NewSALDataMeasurementBuilder() SALDataMeasurementBuilder {
	return &_SALDataMeasurementBuilder{_SALDataMeasurement: new(_SALDataMeasurement)}
}

type _SALDataMeasurementBuilder struct {
	*_SALDataMeasurement

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataMeasurementBuilder) = (*_SALDataMeasurementBuilder)(nil)

func (b *_SALDataMeasurementBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
	contract.(*_SALData)._SubType = b._SALDataMeasurement
}

func (b *_SALDataMeasurementBuilder) WithMandatoryFields(measurementData MeasurementData) SALDataMeasurementBuilder {
	return b.WithMeasurementData(measurementData)
}

func (b *_SALDataMeasurementBuilder) WithMeasurementData(measurementData MeasurementData) SALDataMeasurementBuilder {
	b.MeasurementData = measurementData
	return b
}

func (b *_SALDataMeasurementBuilder) WithMeasurementDataBuilder(builderSupplier func(MeasurementDataBuilder) MeasurementDataBuilder) SALDataMeasurementBuilder {
	builder := builderSupplier(b.MeasurementData.CreateMeasurementDataBuilder())
	var err error
	b.MeasurementData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "MeasurementDataBuilder failed"))
	}
	return b
}

func (b *_SALDataMeasurementBuilder) Build() (SALDataMeasurement, error) {
	if b.MeasurementData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'measurementData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataMeasurement.deepCopy(), nil
}

func (b *_SALDataMeasurementBuilder) MustBuild() SALDataMeasurement {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SALDataMeasurementBuilder) Done() SALDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSALDataBuilder().(*_SALDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SALDataMeasurementBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataMeasurementBuilder) DeepCopy() any {
	_copy := b.CreateSALDataMeasurementBuilder().(*_SALDataMeasurementBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataMeasurementBuilder creates a SALDataMeasurementBuilder
func (b *_SALDataMeasurement) CreateSALDataMeasurementBuilder() SALDataMeasurementBuilder {
	if b == nil {
		return NewSALDataMeasurementBuilder()
	}
	return &_SALDataMeasurementBuilder{_SALDataMeasurement: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataMeasurement) GetApplicationId() ApplicationId {
	return ApplicationId_MEASUREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataMeasurement) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataMeasurement) GetMeasurementData() MeasurementData {
	return m.MeasurementData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataMeasurement(structType any) SALDataMeasurement {
	if casted, ok := structType.(SALDataMeasurement); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataMeasurement); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataMeasurement) GetTypeName() string {
	return "SALDataMeasurement"
}

func (m *_SALDataMeasurement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (measurementData)
	lengthInBits += m.MeasurementData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataMeasurement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataMeasurement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataMeasurement SALDataMeasurement, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataMeasurement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataMeasurement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	measurementData, err := ReadSimpleField[MeasurementData](ctx, "measurementData", ReadComplex[MeasurementData](MeasurementDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'measurementData' field"))
	}
	m.MeasurementData = measurementData

	if closeErr := readBuffer.CloseContext("SALDataMeasurement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataMeasurement")
	}

	return m, nil
}

func (m *_SALDataMeasurement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataMeasurement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataMeasurement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataMeasurement")
		}

		if err := WriteSimpleField[MeasurementData](ctx, "measurementData", m.GetMeasurementData(), WriteComplex[MeasurementData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'measurementData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataMeasurement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataMeasurement")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataMeasurement) IsSALDataMeasurement() {}

func (m *_SALDataMeasurement) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataMeasurement) deepCopy() *_SALDataMeasurement {
	if m == nil {
		return nil
	}
	_SALDataMeasurementCopy := &_SALDataMeasurement{
		m.SALDataContract.(*_SALData).deepCopy(),
		utils.DeepCopy[MeasurementData](m.MeasurementData),
	}
	_SALDataMeasurementCopy.SALDataContract.(*_SALData)._SubType = m
	return _SALDataMeasurementCopy
}

func (m *_SALDataMeasurement) String() string {
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
