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

// BACnetConstructedDataBinaryLightingOutputAll is the corresponding interface of BACnetConstructedDataBinaryLightingOutputAll
type BACnetConstructedDataBinaryLightingOutputAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataBinaryLightingOutputAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryLightingOutputAll()
	// CreateBuilder creates a BACnetConstructedDataBinaryLightingOutputAllBuilder
	CreateBACnetConstructedDataBinaryLightingOutputAllBuilder() BACnetConstructedDataBinaryLightingOutputAllBuilder
}

// _BACnetConstructedDataBinaryLightingOutputAll is the data-structure of this message
type _BACnetConstructedDataBinaryLightingOutputAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataBinaryLightingOutputAll = (*_BACnetConstructedDataBinaryLightingOutputAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryLightingOutputAll)(nil)

// NewBACnetConstructedDataBinaryLightingOutputAll factory function for _BACnetConstructedDataBinaryLightingOutputAll
func NewBACnetConstructedDataBinaryLightingOutputAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryLightingOutputAll {
	_result := &_BACnetConstructedDataBinaryLightingOutputAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBinaryLightingOutputAllBuilder is a builder for BACnetConstructedDataBinaryLightingOutputAll
type BACnetConstructedDataBinaryLightingOutputAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataBinaryLightingOutputAllBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBinaryLightingOutputAll or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryLightingOutputAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryLightingOutputAll
}

// NewBACnetConstructedDataBinaryLightingOutputAllBuilder() creates a BACnetConstructedDataBinaryLightingOutputAllBuilder
func NewBACnetConstructedDataBinaryLightingOutputAllBuilder() BACnetConstructedDataBinaryLightingOutputAllBuilder {
	return &_BACnetConstructedDataBinaryLightingOutputAllBuilder{_BACnetConstructedDataBinaryLightingOutputAll: new(_BACnetConstructedDataBinaryLightingOutputAll)}
}

type _BACnetConstructedDataBinaryLightingOutputAllBuilder struct {
	*_BACnetConstructedDataBinaryLightingOutputAll

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryLightingOutputAllBuilder) = (*_BACnetConstructedDataBinaryLightingOutputAllBuilder)(nil)

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataBinaryLightingOutputAll
}

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) WithMandatoryFields() BACnetConstructedDataBinaryLightingOutputAllBuilder {
	return b
}

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) Build() (BACnetConstructedDataBinaryLightingOutputAll, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryLightingOutputAll.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) MustBuild() BACnetConstructedDataBinaryLightingOutputAll {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryLightingOutputAllBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryLightingOutputAllBuilder().(*_BACnetConstructedDataBinaryLightingOutputAllBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryLightingOutputAllBuilder creates a BACnetConstructedDataBinaryLightingOutputAllBuilder
func (b *_BACnetConstructedDataBinaryLightingOutputAll) CreateBACnetConstructedDataBinaryLightingOutputAllBuilder() BACnetConstructedDataBinaryLightingOutputAllBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryLightingOutputAllBuilder()
	}
	return &_BACnetConstructedDataBinaryLightingOutputAllBuilder{_BACnetConstructedDataBinaryLightingOutputAll: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryLightingOutputAll(structType any) BACnetConstructedDataBinaryLightingOutputAll {
	if casted, ok := structType.(BACnetConstructedDataBinaryLightingOutputAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryLightingOutputAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) GetTypeName() string {
	return "BACnetConstructedDataBinaryLightingOutputAll"
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryLightingOutputAll BACnetConstructedDataBinaryLightingOutputAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryLightingOutputAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryLightingOutputAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryLightingOutputAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryLightingOutputAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryLightingOutputAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryLightingOutputAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryLightingOutputAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryLightingOutputAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) IsBACnetConstructedDataBinaryLightingOutputAll() {
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) deepCopy() *_BACnetConstructedDataBinaryLightingOutputAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryLightingOutputAllCopy := &_BACnetConstructedDataBinaryLightingOutputAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	_BACnetConstructedDataBinaryLightingOutputAllCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryLightingOutputAllCopy
}

func (m *_BACnetConstructedDataBinaryLightingOutputAll) String() string {
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
