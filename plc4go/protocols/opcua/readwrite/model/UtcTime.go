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

// UtcTime is the corresponding interface of UtcTime
type UtcTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsUtcTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUtcTime()
	// CreateBuilder creates a UtcTimeBuilder
	CreateUtcTimeBuilder() UtcTimeBuilder
}

// _UtcTime is the data-structure of this message
type _UtcTime struct {
}

var _ UtcTime = (*_UtcTime)(nil)

// NewUtcTime factory function for _UtcTime
func NewUtcTime() *_UtcTime {
	return &_UtcTime{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UtcTimeBuilder is a builder for UtcTime
type UtcTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() UtcTimeBuilder
	// Build builds the UtcTime or returns an error if something is wrong
	Build() (UtcTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UtcTime
}

// NewUtcTimeBuilder() creates a UtcTimeBuilder
func NewUtcTimeBuilder() UtcTimeBuilder {
	return &_UtcTimeBuilder{_UtcTime: new(_UtcTime)}
}

type _UtcTimeBuilder struct {
	*_UtcTime

	err *utils.MultiError
}

var _ (UtcTimeBuilder) = (*_UtcTimeBuilder)(nil)

func (b *_UtcTimeBuilder) WithMandatoryFields() UtcTimeBuilder {
	return b
}

func (b *_UtcTimeBuilder) Build() (UtcTime, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UtcTime.deepCopy(), nil
}

func (b *_UtcTimeBuilder) MustBuild() UtcTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_UtcTimeBuilder) DeepCopy() any {
	_copy := b.CreateUtcTimeBuilder().(*_UtcTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUtcTimeBuilder creates a UtcTimeBuilder
func (b *_UtcTime) CreateUtcTimeBuilder() UtcTimeBuilder {
	if b == nil {
		return NewUtcTimeBuilder()
	}
	return &_UtcTimeBuilder{_UtcTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUtcTime(structType any) UtcTime {
	if casted, ok := structType.(UtcTime); ok {
		return casted
	}
	if casted, ok := structType.(*UtcTime); ok {
		return *casted
	}
	return nil
}

func (m *_UtcTime) GetTypeName() string {
	return "UtcTime"
}

func (m *_UtcTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_UtcTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UtcTimeParse(ctx context.Context, theBytes []byte) (UtcTime, error) {
	return UtcTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UtcTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (UtcTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (UtcTime, error) {
		return UtcTimeParseWithBuffer(ctx, readBuffer)
	}
}

func UtcTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UtcTime, error) {
	v, err := (&_UtcTime{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_UtcTime) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__utcTime UtcTime, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UtcTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UtcTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("UtcTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UtcTime")
	}

	return m, nil
}

func (m *_UtcTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UtcTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("UtcTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for UtcTime")
	}

	if popErr := writeBuffer.PopContext("UtcTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for UtcTime")
	}
	return nil
}

func (m *_UtcTime) IsUtcTime() {}

func (m *_UtcTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UtcTime) deepCopy() *_UtcTime {
	if m == nil {
		return nil
	}
	_UtcTimeCopy := &_UtcTime{}
	return _UtcTimeCopy
}

func (m *_UtcTime) String() string {
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
