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

// BACnetServicesSupportedTagged is the corresponding interface of BACnetServicesSupportedTagged
type BACnetServicesSupportedTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetWriteGroup returns WriteGroup (virtual field)
	GetWriteGroup() bool
	// GetSubscribeCovPropertyMultiple returns SubscribeCovPropertyMultiple (virtual field)
	GetSubscribeCovPropertyMultiple() bool
	// GetConfirmedCovNotificationMultiple returns ConfirmedCovNotificationMultiple (virtual field)
	GetConfirmedCovNotificationMultiple() bool
	// GetUnconfirmedCovNotificationMultiple returns UnconfirmedCovNotificationMultiple (virtual field)
	GetUnconfirmedCovNotificationMultiple() bool
	// GetWhoIs returns WhoIs (virtual field)
	GetWhoIs() bool
	// GetReadRange returns ReadRange (virtual field)
	GetReadRange() bool
	// GetUtcTimeSynchronization returns UtcTimeSynchronization (virtual field)
	GetUtcTimeSynchronization() bool
	// GetLifeSafetyOperation returns LifeSafetyOperation (virtual field)
	GetLifeSafetyOperation() bool
	// GetSubscribeCovProperty returns SubscribeCovProperty (virtual field)
	GetSubscribeCovProperty() bool
	// GetGetEventInformation returns GetEventInformation (virtual field)
	GetGetEventInformation() bool
	// IsBACnetServicesSupportedTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetServicesSupportedTagged()
	// CreateBuilder creates a BACnetServicesSupportedTaggedBuilder
	CreateBACnetServicesSupportedTaggedBuilder() BACnetServicesSupportedTaggedBuilder
}

// _BACnetServicesSupportedTagged is the data-structure of this message
type _BACnetServicesSupportedTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetServicesSupportedTagged = (*_BACnetServicesSupportedTagged)(nil)

// NewBACnetServicesSupportedTagged factory function for _BACnetServicesSupportedTagged
func NewBACnetServicesSupportedTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetServicesSupportedTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetServicesSupportedTagged must not be nil")
	}
	if payload == nil {
		panic("payload of type BACnetTagPayloadBitString for BACnetServicesSupportedTagged must not be nil")
	}
	return &_BACnetServicesSupportedTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetServicesSupportedTaggedBuilder is a builder for BACnetServicesSupportedTagged
type BACnetServicesSupportedTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetServicesSupportedTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetServicesSupportedTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetServicesSupportedTaggedBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadBitString) BACnetServicesSupportedTaggedBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetServicesSupportedTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetServicesSupportedTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetServicesSupportedTaggedBuilder
	// Build builds the BACnetServicesSupportedTagged or returns an error if something is wrong
	Build() (BACnetServicesSupportedTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetServicesSupportedTagged
}

// NewBACnetServicesSupportedTaggedBuilder() creates a BACnetServicesSupportedTaggedBuilder
func NewBACnetServicesSupportedTaggedBuilder() BACnetServicesSupportedTaggedBuilder {
	return &_BACnetServicesSupportedTaggedBuilder{_BACnetServicesSupportedTagged: new(_BACnetServicesSupportedTagged)}
}

type _BACnetServicesSupportedTaggedBuilder struct {
	*_BACnetServicesSupportedTagged

	err *utils.MultiError
}

var _ (BACnetServicesSupportedTaggedBuilder) = (*_BACnetServicesSupportedTaggedBuilder)(nil)

func (b *_BACnetServicesSupportedTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetServicesSupportedTaggedBuilder {
	return b.WithHeader(header).WithPayload(payload)
}

func (b *_BACnetServicesSupportedTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetServicesSupportedTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetServicesSupportedTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetServicesSupportedTaggedBuilder {
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

func (b *_BACnetServicesSupportedTaggedBuilder) WithPayload(payload BACnetTagPayloadBitString) BACnetServicesSupportedTaggedBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetServicesSupportedTaggedBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetServicesSupportedTaggedBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadBitStringBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadBitStringBuilder failed"))
	}
	return b
}

func (b *_BACnetServicesSupportedTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetServicesSupportedTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetServicesSupportedTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetServicesSupportedTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetServicesSupportedTaggedBuilder) Build() (BACnetServicesSupportedTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetServicesSupportedTagged.deepCopy(), nil
}

func (b *_BACnetServicesSupportedTaggedBuilder) MustBuild() BACnetServicesSupportedTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetServicesSupportedTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetServicesSupportedTaggedBuilder().(*_BACnetServicesSupportedTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetServicesSupportedTaggedBuilder creates a BACnetServicesSupportedTaggedBuilder
func (b *_BACnetServicesSupportedTagged) CreateBACnetServicesSupportedTaggedBuilder() BACnetServicesSupportedTaggedBuilder {
	if b == nil {
		return NewBACnetServicesSupportedTaggedBuilder()
	}
	return &_BACnetServicesSupportedTaggedBuilder{_BACnetServicesSupportedTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServicesSupportedTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetServicesSupportedTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetServicesSupportedTagged) GetWriteGroup() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetSubscribeCovPropertyMultiple() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetConfirmedCovNotificationMultiple() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetUnconfirmedCovNotificationMultiple() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (3))), func() any { return bool(m.GetPayload().GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetWhoIs() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (4))), func() any { return bool(m.GetPayload().GetData()[4]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetReadRange() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (5))), func() any { return bool(m.GetPayload().GetData()[5]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetUtcTimeSynchronization() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (6))), func() any { return bool(m.GetPayload().GetData()[6]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetLifeSafetyOperation() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (7))), func() any { return bool(m.GetPayload().GetData()[7]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetSubscribeCovProperty() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (8))), func() any { return bool(m.GetPayload().GetData()[8]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetServicesSupportedTagged) GetGetEventInformation() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (9))), func() any { return bool(m.GetPayload().GetData()[9]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetServicesSupportedTagged(structType any) BACnetServicesSupportedTagged {
	if casted, ok := structType.(BACnetServicesSupportedTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServicesSupportedTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServicesSupportedTagged) GetTypeName() string {
	return "BACnetServicesSupportedTagged"
}

func (m *_BACnetServicesSupportedTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetServicesSupportedTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServicesSupportedTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetServicesSupportedTagged, error) {
	return BACnetServicesSupportedTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetServicesSupportedTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServicesSupportedTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetServicesSupportedTagged, error) {
		return BACnetServicesSupportedTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetServicesSupportedTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetServicesSupportedTagged, error) {
	v, err := (&_BACnetServicesSupportedTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetServicesSupportedTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetServicesSupportedTagged BACnetServicesSupportedTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServicesSupportedTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServicesSupportedTagged")
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

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	writeGroup, err := ReadVirtualField[bool](ctx, "writeGroup", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeGroup' field"))
	}
	_ = writeGroup

	subscribeCovPropertyMultiple, err := ReadVirtualField[bool](ctx, "subscribeCovPropertyMultiple", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscribeCovPropertyMultiple' field"))
	}
	_ = subscribeCovPropertyMultiple

	confirmedCovNotificationMultiple, err := ReadVirtualField[bool](ctx, "confirmedCovNotificationMultiple", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'confirmedCovNotificationMultiple' field"))
	}
	_ = confirmedCovNotificationMultiple

	unconfirmedCovNotificationMultiple, err := ReadVirtualField[bool](ctx, "unconfirmedCovNotificationMultiple", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (3))), func() any { return bool(payload.GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unconfirmedCovNotificationMultiple' field"))
	}
	_ = unconfirmedCovNotificationMultiple

	whoIs, err := ReadVirtualField[bool](ctx, "whoIs", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (4))), func() any { return bool(payload.GetData()[4]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'whoIs' field"))
	}
	_ = whoIs

	readRange, err := ReadVirtualField[bool](ctx, "readRange", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (5))), func() any { return bool(payload.GetData()[5]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readRange' field"))
	}
	_ = readRange

	utcTimeSynchronization, err := ReadVirtualField[bool](ctx, "utcTimeSynchronization", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (6))), func() any { return bool(payload.GetData()[6]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'utcTimeSynchronization' field"))
	}
	_ = utcTimeSynchronization

	lifeSafetyOperation, err := ReadVirtualField[bool](ctx, "lifeSafetyOperation", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (7))), func() any { return bool(payload.GetData()[7]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifeSafetyOperation' field"))
	}
	_ = lifeSafetyOperation

	subscribeCovProperty, err := ReadVirtualField[bool](ctx, "subscribeCovProperty", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (8))), func() any { return bool(payload.GetData()[8]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscribeCovProperty' field"))
	}
	_ = subscribeCovProperty

	getEventInformation, err := ReadVirtualField[bool](ctx, "getEventInformation", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (9))), func() any { return bool(payload.GetData()[9]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'getEventInformation' field"))
	}
	_ = getEventInformation

	if closeErr := readBuffer.CloseContext("BACnetServicesSupportedTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServicesSupportedTagged")
	}

	return m, nil
}

func (m *_BACnetServicesSupportedTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServicesSupportedTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetServicesSupportedTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetServicesSupportedTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	writeGroup := m.GetWriteGroup()
	_ = writeGroup
	if _writeGroupErr := writeBuffer.WriteVirtual(ctx, "writeGroup", m.GetWriteGroup()); _writeGroupErr != nil {
		return errors.Wrap(_writeGroupErr, "Error serializing 'writeGroup' field")
	}
	// Virtual field
	subscribeCovPropertyMultiple := m.GetSubscribeCovPropertyMultiple()
	_ = subscribeCovPropertyMultiple
	if _subscribeCovPropertyMultipleErr := writeBuffer.WriteVirtual(ctx, "subscribeCovPropertyMultiple", m.GetSubscribeCovPropertyMultiple()); _subscribeCovPropertyMultipleErr != nil {
		return errors.Wrap(_subscribeCovPropertyMultipleErr, "Error serializing 'subscribeCovPropertyMultiple' field")
	}
	// Virtual field
	confirmedCovNotificationMultiple := m.GetConfirmedCovNotificationMultiple()
	_ = confirmedCovNotificationMultiple
	if _confirmedCovNotificationMultipleErr := writeBuffer.WriteVirtual(ctx, "confirmedCovNotificationMultiple", m.GetConfirmedCovNotificationMultiple()); _confirmedCovNotificationMultipleErr != nil {
		return errors.Wrap(_confirmedCovNotificationMultipleErr, "Error serializing 'confirmedCovNotificationMultiple' field")
	}
	// Virtual field
	unconfirmedCovNotificationMultiple := m.GetUnconfirmedCovNotificationMultiple()
	_ = unconfirmedCovNotificationMultiple
	if _unconfirmedCovNotificationMultipleErr := writeBuffer.WriteVirtual(ctx, "unconfirmedCovNotificationMultiple", m.GetUnconfirmedCovNotificationMultiple()); _unconfirmedCovNotificationMultipleErr != nil {
		return errors.Wrap(_unconfirmedCovNotificationMultipleErr, "Error serializing 'unconfirmedCovNotificationMultiple' field")
	}
	// Virtual field
	whoIs := m.GetWhoIs()
	_ = whoIs
	if _whoIsErr := writeBuffer.WriteVirtual(ctx, "whoIs", m.GetWhoIs()); _whoIsErr != nil {
		return errors.Wrap(_whoIsErr, "Error serializing 'whoIs' field")
	}
	// Virtual field
	readRange := m.GetReadRange()
	_ = readRange
	if _readRangeErr := writeBuffer.WriteVirtual(ctx, "readRange", m.GetReadRange()); _readRangeErr != nil {
		return errors.Wrap(_readRangeErr, "Error serializing 'readRange' field")
	}
	// Virtual field
	utcTimeSynchronization := m.GetUtcTimeSynchronization()
	_ = utcTimeSynchronization
	if _utcTimeSynchronizationErr := writeBuffer.WriteVirtual(ctx, "utcTimeSynchronization", m.GetUtcTimeSynchronization()); _utcTimeSynchronizationErr != nil {
		return errors.Wrap(_utcTimeSynchronizationErr, "Error serializing 'utcTimeSynchronization' field")
	}
	// Virtual field
	lifeSafetyOperation := m.GetLifeSafetyOperation()
	_ = lifeSafetyOperation
	if _lifeSafetyOperationErr := writeBuffer.WriteVirtual(ctx, "lifeSafetyOperation", m.GetLifeSafetyOperation()); _lifeSafetyOperationErr != nil {
		return errors.Wrap(_lifeSafetyOperationErr, "Error serializing 'lifeSafetyOperation' field")
	}
	// Virtual field
	subscribeCovProperty := m.GetSubscribeCovProperty()
	_ = subscribeCovProperty
	if _subscribeCovPropertyErr := writeBuffer.WriteVirtual(ctx, "subscribeCovProperty", m.GetSubscribeCovProperty()); _subscribeCovPropertyErr != nil {
		return errors.Wrap(_subscribeCovPropertyErr, "Error serializing 'subscribeCovProperty' field")
	}
	// Virtual field
	getEventInformation := m.GetGetEventInformation()
	_ = getEventInformation
	if _getEventInformationErr := writeBuffer.WriteVirtual(ctx, "getEventInformation", m.GetGetEventInformation()); _getEventInformationErr != nil {
		return errors.Wrap(_getEventInformationErr, "Error serializing 'getEventInformation' field")
	}

	if popErr := writeBuffer.PopContext("BACnetServicesSupportedTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetServicesSupportedTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetServicesSupportedTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetServicesSupportedTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetServicesSupportedTagged) IsBACnetServicesSupportedTagged() {}

func (m *_BACnetServicesSupportedTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetServicesSupportedTagged) deepCopy() *_BACnetServicesSupportedTagged {
	if m == nil {
		return nil
	}
	_BACnetServicesSupportedTaggedCopy := &_BACnetServicesSupportedTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		utils.DeepCopy[BACnetTagPayloadBitString](m.Payload),
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetServicesSupportedTaggedCopy
}

func (m *_BACnetServicesSupportedTagged) String() string {
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
