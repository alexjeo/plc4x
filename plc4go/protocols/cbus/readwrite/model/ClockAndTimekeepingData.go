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

// ClockAndTimekeepingData is the corresponding interface of ClockAndTimekeepingData
type ClockAndTimekeepingData interface {
	ClockAndTimekeepingDataContract
	ClockAndTimekeepingDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsClockAndTimekeepingData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClockAndTimekeepingData()
	// CreateBuilder creates a ClockAndTimekeepingDataBuilder
	CreateClockAndTimekeepingDataBuilder() ClockAndTimekeepingDataBuilder
}

// ClockAndTimekeepingDataContract provides a set of functions which can be overwritten by a sub struct
type ClockAndTimekeepingDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() ClockAndTimekeepingCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() ClockAndTimekeepingCommandType
	// IsClockAndTimekeepingData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsClockAndTimekeepingData()
	// CreateBuilder creates a ClockAndTimekeepingDataBuilder
	CreateClockAndTimekeepingDataBuilder() ClockAndTimekeepingDataBuilder
}

// ClockAndTimekeepingDataRequirements provides a set of functions which need to be implemented by a sub struct
type ClockAndTimekeepingDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetArgument returns Argument (discriminator field)
	GetArgument() byte
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() ClockAndTimekeepingCommandType
}

// _ClockAndTimekeepingData is the data-structure of this message
type _ClockAndTimekeepingData struct {
	_SubType interface {
		ClockAndTimekeepingDataContract
		ClockAndTimekeepingDataRequirements
	}
	CommandTypeContainer ClockAndTimekeepingCommandTypeContainer
	Argument             byte
}

var _ ClockAndTimekeepingDataContract = (*_ClockAndTimekeepingData)(nil)

// NewClockAndTimekeepingData factory function for _ClockAndTimekeepingData
func NewClockAndTimekeepingData(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) *_ClockAndTimekeepingData {
	return &_ClockAndTimekeepingData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ClockAndTimekeepingDataBuilder is a builder for ClockAndTimekeepingData
type ClockAndTimekeepingDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) ClockAndTimekeepingDataBuilder
	// WithCommandTypeContainer adds CommandTypeContainer (property field)
	WithCommandTypeContainer(ClockAndTimekeepingCommandTypeContainer) ClockAndTimekeepingDataBuilder
	// WithArgument adds Argument (property field)
	WithArgument(byte) ClockAndTimekeepingDataBuilder
	// AsClockAndTimekeepingDataUpdateTime converts this build to a subType of ClockAndTimekeepingData. It is always possible to return to current builder using Done()
	AsClockAndTimekeepingDataUpdateTime() ClockAndTimekeepingDataUpdateTimeBuilder
	// AsClockAndTimekeepingDataUpdateDate converts this build to a subType of ClockAndTimekeepingData. It is always possible to return to current builder using Done()
	AsClockAndTimekeepingDataUpdateDate() ClockAndTimekeepingDataUpdateDateBuilder
	// AsClockAndTimekeepingDataRequestRefresh converts this build to a subType of ClockAndTimekeepingData. It is always possible to return to current builder using Done()
	AsClockAndTimekeepingDataRequestRefresh() ClockAndTimekeepingDataRequestRefreshBuilder
	// Build builds the ClockAndTimekeepingData or returns an error if something is wrong
	PartialBuild() (ClockAndTimekeepingDataContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() ClockAndTimekeepingDataContract
	// Build builds the ClockAndTimekeepingData or returns an error if something is wrong
	Build() (ClockAndTimekeepingData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ClockAndTimekeepingData
}

// NewClockAndTimekeepingDataBuilder() creates a ClockAndTimekeepingDataBuilder
func NewClockAndTimekeepingDataBuilder() ClockAndTimekeepingDataBuilder {
	return &_ClockAndTimekeepingDataBuilder{_ClockAndTimekeepingData: new(_ClockAndTimekeepingData)}
}

type _ClockAndTimekeepingDataChildBuilder interface {
	utils.Copyable
	setParent(ClockAndTimekeepingDataContract)
	buildForClockAndTimekeepingData() (ClockAndTimekeepingData, error)
}

type _ClockAndTimekeepingDataBuilder struct {
	*_ClockAndTimekeepingData

	childBuilder _ClockAndTimekeepingDataChildBuilder

	err *utils.MultiError
}

var _ (ClockAndTimekeepingDataBuilder) = (*_ClockAndTimekeepingDataBuilder)(nil)

func (b *_ClockAndTimekeepingDataBuilder) WithMandatoryFields(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) ClockAndTimekeepingDataBuilder {
	return b.WithCommandTypeContainer(commandTypeContainer).WithArgument(argument)
}

func (b *_ClockAndTimekeepingDataBuilder) WithCommandTypeContainer(commandTypeContainer ClockAndTimekeepingCommandTypeContainer) ClockAndTimekeepingDataBuilder {
	b.CommandTypeContainer = commandTypeContainer
	return b
}

func (b *_ClockAndTimekeepingDataBuilder) WithArgument(argument byte) ClockAndTimekeepingDataBuilder {
	b.Argument = argument
	return b
}

func (b *_ClockAndTimekeepingDataBuilder) PartialBuild() (ClockAndTimekeepingDataContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ClockAndTimekeepingData.deepCopy(), nil
}

func (b *_ClockAndTimekeepingDataBuilder) PartialMustBuild() ClockAndTimekeepingDataContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ClockAndTimekeepingDataBuilder) AsClockAndTimekeepingDataUpdateTime() ClockAndTimekeepingDataUpdateTimeBuilder {
	if cb, ok := b.childBuilder.(ClockAndTimekeepingDataUpdateTimeBuilder); ok {
		return cb
	}
	cb := NewClockAndTimekeepingDataUpdateTimeBuilder().(*_ClockAndTimekeepingDataUpdateTimeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ClockAndTimekeepingDataBuilder) AsClockAndTimekeepingDataUpdateDate() ClockAndTimekeepingDataUpdateDateBuilder {
	if cb, ok := b.childBuilder.(ClockAndTimekeepingDataUpdateDateBuilder); ok {
		return cb
	}
	cb := NewClockAndTimekeepingDataUpdateDateBuilder().(*_ClockAndTimekeepingDataUpdateDateBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ClockAndTimekeepingDataBuilder) AsClockAndTimekeepingDataRequestRefresh() ClockAndTimekeepingDataRequestRefreshBuilder {
	if cb, ok := b.childBuilder.(ClockAndTimekeepingDataRequestRefreshBuilder); ok {
		return cb
	}
	cb := NewClockAndTimekeepingDataRequestRefreshBuilder().(*_ClockAndTimekeepingDataRequestRefreshBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_ClockAndTimekeepingDataBuilder) Build() (ClockAndTimekeepingData, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForClockAndTimekeepingData()
}

func (b *_ClockAndTimekeepingDataBuilder) MustBuild() ClockAndTimekeepingData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ClockAndTimekeepingDataBuilder) DeepCopy() any {
	_copy := b.CreateClockAndTimekeepingDataBuilder().(*_ClockAndTimekeepingDataBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_ClockAndTimekeepingDataChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateClockAndTimekeepingDataBuilder creates a ClockAndTimekeepingDataBuilder
func (b *_ClockAndTimekeepingData) CreateClockAndTimekeepingDataBuilder() ClockAndTimekeepingDataBuilder {
	if b == nil {
		return NewClockAndTimekeepingDataBuilder()
	}
	return &_ClockAndTimekeepingDataBuilder{_ClockAndTimekeepingData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClockAndTimekeepingData) GetCommandTypeContainer() ClockAndTimekeepingCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_ClockAndTimekeepingData) GetArgument() byte {
	return m.Argument
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_ClockAndTimekeepingData) GetCommandType() ClockAndTimekeepingCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastClockAndTimekeepingCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingData(structType any) ClockAndTimekeepingData {
	if casted, ok := structType.(ClockAndTimekeepingData); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingData); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingData) GetTypeName() string {
	return "ClockAndTimekeepingData"
}

func (m *_ClockAndTimekeepingData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClockAndTimekeepingData) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_ClockAndTimekeepingData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func ClockAndTimekeepingDataParse[T ClockAndTimekeepingData](ctx context.Context, theBytes []byte) (T, error) {
	return ClockAndTimekeepingDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClockAndTimekeepingDataParseWithBufferProducer[T ClockAndTimekeepingData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := ClockAndTimekeepingDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func ClockAndTimekeepingDataParseWithBuffer[T ClockAndTimekeepingData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_ClockAndTimekeepingData{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_ClockAndTimekeepingData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__clockAndTimekeepingData ClockAndTimekeepingData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsClockAndTimekeepingCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[ClockAndTimekeepingCommandTypeContainer](ctx, "commandTypeContainer", "ClockAndTimekeepingCommandTypeContainer", ReadEnum(ClockAndTimekeepingCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[ClockAndTimekeepingCommandType](ctx, "commandType", (*ClockAndTimekeepingCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	argument, err := ReadSimpleField(ctx, "argument", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'argument' field"))
	}
	m.Argument = argument

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child ClockAndTimekeepingData
	switch {
	case commandType == ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE && argument == 0x01: // ClockAndTimekeepingDataUpdateTime
		if _child, err = new(_ClockAndTimekeepingDataUpdateTime).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ClockAndTimekeepingDataUpdateTime for type-switch of ClockAndTimekeepingData")
		}
	case commandType == ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE && argument == 0x02: // ClockAndTimekeepingDataUpdateDate
		if _child, err = new(_ClockAndTimekeepingDataUpdateDate).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ClockAndTimekeepingDataUpdateDate for type-switch of ClockAndTimekeepingData")
		}
	case commandType == ClockAndTimekeepingCommandType_REQUEST_REFRESH && argument == 0x03: // ClockAndTimekeepingDataRequestRefresh
		if _child, err = new(_ClockAndTimekeepingDataRequestRefresh).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ClockAndTimekeepingDataRequestRefresh for type-switch of ClockAndTimekeepingData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingData")
	}

	return _child, nil
}

func (pm *_ClockAndTimekeepingData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ClockAndTimekeepingData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ClockAndTimekeepingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingData")
	}

	if err := WriteSimpleEnumField[ClockAndTimekeepingCommandTypeContainer](ctx, "commandTypeContainer", "ClockAndTimekeepingCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[ClockAndTimekeepingCommandTypeContainer, uint8](ClockAndTimekeepingCommandTypeContainer.GetValue, ClockAndTimekeepingCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "argument", m.GetArgument(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ClockAndTimekeepingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingData")
	}
	return nil
}

func (m *_ClockAndTimekeepingData) IsClockAndTimekeepingData() {}

func (m *_ClockAndTimekeepingData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ClockAndTimekeepingData) deepCopy() *_ClockAndTimekeepingData {
	if m == nil {
		return nil
	}
	_ClockAndTimekeepingDataCopy := &_ClockAndTimekeepingData{
		nil, // will be set by child
		m.CommandTypeContainer,
		m.Argument,
	}
	return _ClockAndTimekeepingDataCopy
}
