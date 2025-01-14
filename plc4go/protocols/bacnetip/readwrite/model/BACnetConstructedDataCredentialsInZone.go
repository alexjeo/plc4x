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

// BACnetConstructedDataCredentialsInZone is the corresponding interface of BACnetConstructedDataCredentialsInZone
type BACnetConstructedDataCredentialsInZone interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetCredentialsInZone returns CredentialsInZone (property field)
	GetCredentialsInZone() []BACnetDeviceObjectReference
	// IsBACnetConstructedDataCredentialsInZone is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCredentialsInZone()
	// CreateBuilder creates a BACnetConstructedDataCredentialsInZoneBuilder
	CreateBACnetConstructedDataCredentialsInZoneBuilder() BACnetConstructedDataCredentialsInZoneBuilder
}

// _BACnetConstructedDataCredentialsInZone is the data-structure of this message
type _BACnetConstructedDataCredentialsInZone struct {
	BACnetConstructedDataContract
	CredentialsInZone []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataCredentialsInZone = (*_BACnetConstructedDataCredentialsInZone)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCredentialsInZone)(nil)

// NewBACnetConstructedDataCredentialsInZone factory function for _BACnetConstructedDataCredentialsInZone
func NewBACnetConstructedDataCredentialsInZone(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, credentialsInZone []BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialsInZone {
	_result := &_BACnetConstructedDataCredentialsInZone{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CredentialsInZone:             credentialsInZone,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCredentialsInZoneBuilder is a builder for BACnetConstructedDataCredentialsInZone
type BACnetConstructedDataCredentialsInZoneBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(credentialsInZone []BACnetDeviceObjectReference) BACnetConstructedDataCredentialsInZoneBuilder
	// WithCredentialsInZone adds CredentialsInZone (property field)
	WithCredentialsInZone(...BACnetDeviceObjectReference) BACnetConstructedDataCredentialsInZoneBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataCredentialsInZone or returns an error if something is wrong
	Build() (BACnetConstructedDataCredentialsInZone, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCredentialsInZone
}

// NewBACnetConstructedDataCredentialsInZoneBuilder() creates a BACnetConstructedDataCredentialsInZoneBuilder
func NewBACnetConstructedDataCredentialsInZoneBuilder() BACnetConstructedDataCredentialsInZoneBuilder {
	return &_BACnetConstructedDataCredentialsInZoneBuilder{_BACnetConstructedDataCredentialsInZone: new(_BACnetConstructedDataCredentialsInZone)}
}

type _BACnetConstructedDataCredentialsInZoneBuilder struct {
	*_BACnetConstructedDataCredentialsInZone

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataCredentialsInZoneBuilder) = (*_BACnetConstructedDataCredentialsInZoneBuilder)(nil)

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataCredentialsInZone
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) WithMandatoryFields(credentialsInZone []BACnetDeviceObjectReference) BACnetConstructedDataCredentialsInZoneBuilder {
	return b.WithCredentialsInZone(credentialsInZone...)
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) WithCredentialsInZone(credentialsInZone ...BACnetDeviceObjectReference) BACnetConstructedDataCredentialsInZoneBuilder {
	b.CredentialsInZone = credentialsInZone
	return b
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) Build() (BACnetConstructedDataCredentialsInZone, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataCredentialsInZone.deepCopy(), nil
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) MustBuild() BACnetConstructedDataCredentialsInZone {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataCredentialsInZoneBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataCredentialsInZoneBuilder().(*_BACnetConstructedDataCredentialsInZoneBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataCredentialsInZoneBuilder creates a BACnetConstructedDataCredentialsInZoneBuilder
func (b *_BACnetConstructedDataCredentialsInZone) CreateBACnetConstructedDataCredentialsInZoneBuilder() BACnetConstructedDataCredentialsInZoneBuilder {
	if b == nil {
		return NewBACnetConstructedDataCredentialsInZoneBuilder()
	}
	return &_BACnetConstructedDataCredentialsInZoneBuilder{_BACnetConstructedDataCredentialsInZone: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialsInZone) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIALS_IN_ZONE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialsInZone) GetCredentialsInZone() []BACnetDeviceObjectReference {
	return m.CredentialsInZone
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialsInZone(structType any) BACnetConstructedDataCredentialsInZone {
	if casted, ok := structType.(BACnetConstructedDataCredentialsInZone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialsInZone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialsInZone) GetTypeName() string {
	return "BACnetConstructedDataCredentialsInZone"
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.CredentialsInZone) > 0 {
		for _, element := range m.CredentialsInZone {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialsInZone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCredentialsInZone) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCredentialsInZone BACnetConstructedDataCredentialsInZone, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialsInZone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialsInZone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	credentialsInZone, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "credentialsInZone", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'credentialsInZone' field"))
	}
	m.CredentialsInZone = credentialsInZone

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialsInZone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialsInZone")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCredentialsInZone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCredentialsInZone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialsInZone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialsInZone")
		}

		if err := WriteComplexTypeArrayField(ctx, "credentialsInZone", m.GetCredentialsInZone(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'credentialsInZone' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialsInZone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialsInZone")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialsInZone) IsBACnetConstructedDataCredentialsInZone() {}

func (m *_BACnetConstructedDataCredentialsInZone) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCredentialsInZone) deepCopy() *_BACnetConstructedDataCredentialsInZone {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCredentialsInZoneCopy := &_BACnetConstructedDataCredentialsInZone{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetDeviceObjectReference, BACnetDeviceObjectReference](m.CredentialsInZone),
	}
	_BACnetConstructedDataCredentialsInZoneCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCredentialsInZoneCopy
}

func (m *_BACnetConstructedDataCredentialsInZone) String() string {
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
