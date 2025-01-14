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

// Services is the corresponding interface of Services
type Services interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOffsets returns Offsets (property field)
	GetOffsets() []uint16
	// GetServices returns Services (property field)
	GetServices() []CipService
	// IsServices is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServices()
	// CreateBuilder creates a ServicesBuilder
	CreateServicesBuilder() ServicesBuilder
}

// _Services is the data-structure of this message
type _Services struct {
	Offsets  []uint16
	Services []CipService

	// Arguments.
	ServicesLen uint16
}

var _ Services = (*_Services)(nil)

// NewServices factory function for _Services
func NewServices(offsets []uint16, services []CipService, servicesLen uint16) *_Services {
	return &_Services{Offsets: offsets, Services: services, ServicesLen: servicesLen}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ServicesBuilder is a builder for Services
type ServicesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(offsets []uint16, services []CipService) ServicesBuilder
	// WithOffsets adds Offsets (property field)
	WithOffsets(...uint16) ServicesBuilder
	// WithServices adds Services (property field)
	WithServices(...CipService) ServicesBuilder
	// WithArgServicesLen sets a parser argument
	WithArgServicesLen(uint16) ServicesBuilder
	// Build builds the Services or returns an error if something is wrong
	Build() (Services, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() Services
}

// NewServicesBuilder() creates a ServicesBuilder
func NewServicesBuilder() ServicesBuilder {
	return &_ServicesBuilder{_Services: new(_Services)}
}

type _ServicesBuilder struct {
	*_Services

	err *utils.MultiError
}

var _ (ServicesBuilder) = (*_ServicesBuilder)(nil)

func (b *_ServicesBuilder) WithMandatoryFields(offsets []uint16, services []CipService) ServicesBuilder {
	return b.WithOffsets(offsets...).WithServices(services...)
}

func (b *_ServicesBuilder) WithOffsets(offsets ...uint16) ServicesBuilder {
	b.Offsets = offsets
	return b
}

func (b *_ServicesBuilder) WithServices(services ...CipService) ServicesBuilder {
	b.Services = services
	return b
}

func (b *_ServicesBuilder) WithArgServicesLen(servicesLen uint16) ServicesBuilder {
	b.ServicesLen = servicesLen
	return b
}

func (b *_ServicesBuilder) Build() (Services, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._Services.deepCopy(), nil
}

func (b *_ServicesBuilder) MustBuild() Services {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ServicesBuilder) DeepCopy() any {
	_copy := b.CreateServicesBuilder().(*_ServicesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateServicesBuilder creates a ServicesBuilder
func (b *_Services) CreateServicesBuilder() ServicesBuilder {
	if b == nil {
		return NewServicesBuilder()
	}
	return &_ServicesBuilder{_Services: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Services) GetOffsets() []uint16 {
	return m.Offsets
}

func (m *_Services) GetServices() []CipService {
	return m.Services
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastServices(structType any) Services {
	if casted, ok := structType.(Services); ok {
		return casted
	}
	if casted, ok := structType.(*Services); ok {
		return *casted
	}
	return nil
}

func (m *_Services) GetTypeName() string {
	return "Services"
}

func (m *_Services) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (serviceNb)
	lengthInBits += 16

	// Array field
	if len(m.Offsets) > 0 {
		lengthInBits += 16 * uint16(len(m.Offsets))
	}

	// Array field
	if len(m.Services) > 0 {
		for _curItem, element := range m.Services {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Services), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_Services) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServicesParse(ctx context.Context, theBytes []byte, servicesLen uint16) (Services, error) {
	return ServicesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), servicesLen)
}

func ServicesParseWithBufferProducer(servicesLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (Services, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (Services, error) {
		return ServicesParseWithBuffer(ctx, readBuffer, servicesLen)
	}
}

func ServicesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, servicesLen uint16) (Services, error) {
	v, err := (&_Services{ServicesLen: servicesLen}).parse(ctx, readBuffer, servicesLen)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_Services) parse(ctx context.Context, readBuffer utils.ReadBuffer, servicesLen uint16) (__services Services, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Services"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Services")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceNb, err := ReadImplicitField[uint16](ctx, "serviceNb", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceNb' field"))
	}
	_ = serviceNb

	offsets, err := ReadCountArrayField[uint16](ctx, "offsets", ReadUnsignedShort(readBuffer, uint8(16)), uint64(serviceNb))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offsets' field"))
	}
	m.Offsets = offsets

	services, err := ReadCountArrayField[CipService](ctx, "services", ReadComplex[CipService](CipServiceParseWithBufferProducer[CipService]((bool)(bool(false)), (uint16)(uint16(servicesLen)/uint16(serviceNb))), readBuffer), uint64(serviceNb))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'services' field"))
	}
	m.Services = services

	if closeErr := readBuffer.CloseContext("Services"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Services")
	}

	return m, nil
}

func (m *_Services) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Services) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Services"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Services")
	}
	serviceNb := uint16(uint16(len(m.GetOffsets())))
	if err := WriteImplicitField(ctx, "serviceNb", serviceNb, WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'serviceNb' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "offsets", m.GetOffsets(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'offsets' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "services", m.GetServices(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'services' field")
	}

	if popErr := writeBuffer.PopContext("Services"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Services")
	}
	return nil
}

////
// Arguments Getter

func (m *_Services) GetServicesLen() uint16 {
	return m.ServicesLen
}

//
////

func (m *_Services) IsServices() {}

func (m *_Services) DeepCopy() any {
	return m.deepCopy()
}

func (m *_Services) deepCopy() *_Services {
	if m == nil {
		return nil
	}
	_ServicesCopy := &_Services{
		utils.DeepCopySlice[uint16, uint16](m.Offsets),
		utils.DeepCopySlice[CipService, CipService](m.Services),
		m.ServicesLen,
	}
	return _ServicesCopy
}

func (m *_Services) String() string {
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
