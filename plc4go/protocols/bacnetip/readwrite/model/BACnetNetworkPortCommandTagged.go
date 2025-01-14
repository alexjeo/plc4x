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

// BACnetNetworkPortCommandTagged is the corresponding interface of BACnetNetworkPortCommandTagged
type BACnetNetworkPortCommandTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNetworkPortCommand
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetNetworkPortCommandTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNetworkPortCommandTagged()
	// CreateBuilder creates a BACnetNetworkPortCommandTaggedBuilder
	CreateBACnetNetworkPortCommandTaggedBuilder() BACnetNetworkPortCommandTaggedBuilder
}

// _BACnetNetworkPortCommandTagged is the data-structure of this message
type _BACnetNetworkPortCommandTagged struct {
	Header           BACnetTagHeader
	Value            BACnetNetworkPortCommand
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetNetworkPortCommandTagged = (*_BACnetNetworkPortCommandTagged)(nil)

// NewBACnetNetworkPortCommandTagged factory function for _BACnetNetworkPortCommandTagged
func NewBACnetNetworkPortCommandTagged(header BACnetTagHeader, value BACnetNetworkPortCommand, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetNetworkPortCommandTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetNetworkPortCommandTagged must not be nil")
	}
	return &_BACnetNetworkPortCommandTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNetworkPortCommandTaggedBuilder is a builder for BACnetNetworkPortCommandTagged
type BACnetNetworkPortCommandTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetNetworkPortCommand, proprietaryValue uint32) BACnetNetworkPortCommandTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetNetworkPortCommandTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNetworkPortCommandTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetNetworkPortCommand) BACnetNetworkPortCommandTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetNetworkPortCommandTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetNetworkPortCommandTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetNetworkPortCommandTaggedBuilder
	// Build builds the BACnetNetworkPortCommandTagged or returns an error if something is wrong
	Build() (BACnetNetworkPortCommandTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNetworkPortCommandTagged
}

// NewBACnetNetworkPortCommandTaggedBuilder() creates a BACnetNetworkPortCommandTaggedBuilder
func NewBACnetNetworkPortCommandTaggedBuilder() BACnetNetworkPortCommandTaggedBuilder {
	return &_BACnetNetworkPortCommandTaggedBuilder{_BACnetNetworkPortCommandTagged: new(_BACnetNetworkPortCommandTagged)}
}

type _BACnetNetworkPortCommandTaggedBuilder struct {
	*_BACnetNetworkPortCommandTagged

	err *utils.MultiError
}

var _ (BACnetNetworkPortCommandTaggedBuilder) = (*_BACnetNetworkPortCommandTaggedBuilder)(nil)

func (b *_BACnetNetworkPortCommandTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetNetworkPortCommand, proprietaryValue uint32) BACnetNetworkPortCommandTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetNetworkPortCommandTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetNetworkPortCommandTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) WithValue(value BACnetNetworkPortCommand) BACnetNetworkPortCommandTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetNetworkPortCommandTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetNetworkPortCommandTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetNetworkPortCommandTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetNetworkPortCommandTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) Build() (BACnetNetworkPortCommandTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNetworkPortCommandTagged.deepCopy(), nil
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) MustBuild() BACnetNetworkPortCommandTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNetworkPortCommandTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNetworkPortCommandTaggedBuilder().(*_BACnetNetworkPortCommandTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNetworkPortCommandTaggedBuilder creates a BACnetNetworkPortCommandTaggedBuilder
func (b *_BACnetNetworkPortCommandTagged) CreateBACnetNetworkPortCommandTaggedBuilder() BACnetNetworkPortCommandTaggedBuilder {
	if b == nil {
		return NewBACnetNetworkPortCommandTaggedBuilder()
	}
	return &_BACnetNetworkPortCommandTaggedBuilder{_BACnetNetworkPortCommandTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNetworkPortCommandTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNetworkPortCommandTagged) GetValue() BACnetNetworkPortCommand {
	return m.Value
}

func (m *_BACnetNetworkPortCommandTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetNetworkPortCommandTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNetworkPortCommandTagged(structType any) BACnetNetworkPortCommandTagged {
	if casted, ok := structType.(BACnetNetworkPortCommandTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNetworkPortCommandTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNetworkPortCommandTagged) GetTypeName() string {
	return "BACnetNetworkPortCommandTagged"
}

func (m *_BACnetNetworkPortCommandTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetNetworkPortCommandTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNetworkPortCommandTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetNetworkPortCommandTagged, error) {
	return BACnetNetworkPortCommandTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetNetworkPortCommandTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkPortCommandTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetNetworkPortCommandTagged, error) {
		return BACnetNetworkPortCommandTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetNetworkPortCommandTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNetworkPortCommandTagged, error) {
	v, err := (&_BACnetNetworkPortCommandTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetNetworkPortCommandTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetNetworkPortCommandTagged BACnetNetworkPortCommandTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNetworkPortCommandTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNetworkPortCommandTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetNetworkPortCommand](ctx, "value", readBuffer, EnsureType[BACnetNetworkPortCommand](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetNetworkPortCommand_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetNetworkPortCommandTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNetworkPortCommandTagged")
	}

	return m, nil
}

func (m *_BACnetNetworkPortCommandTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNetworkPortCommandTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetNetworkPortCommandTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNetworkPortCommandTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetNetworkPortCommand](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNetworkPortCommandTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNetworkPortCommandTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetNetworkPortCommandTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetNetworkPortCommandTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetNetworkPortCommandTagged) IsBACnetNetworkPortCommandTagged() {}

func (m *_BACnetNetworkPortCommandTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNetworkPortCommandTagged) deepCopy() *_BACnetNetworkPortCommandTagged {
	if m == nil {
		return nil
	}
	_BACnetNetworkPortCommandTaggedCopy := &_BACnetNetworkPortCommandTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetNetworkPortCommandTaggedCopy
}

func (m *_BACnetNetworkPortCommandTagged) String() string {
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
