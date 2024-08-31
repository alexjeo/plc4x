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

// BACnetConstructedDataFaultHighLimit is the corresponding interface of BACnetConstructedDataFaultHighLimit
type BACnetConstructedDataFaultHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataFaultHighLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFaultHighLimit.
// This is useful for switch cases.
type BACnetConstructedDataFaultHighLimitExactly interface {
	BACnetConstructedDataFaultHighLimit
	isBACnetConstructedDataFaultHighLimit() bool
}

// _BACnetConstructedDataFaultHighLimit is the data-structure of this message
type _BACnetConstructedDataFaultHighLimit struct {
	*_BACnetConstructedData
	FaultHighLimit BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFaultHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFaultHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFaultHighLimit) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFaultHighLimit) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFaultHighLimit) GetFaultHighLimit() BACnetApplicationTagUnsignedInteger {
	return m.FaultHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFaultHighLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFaultHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFaultHighLimit factory function for _BACnetConstructedDataFaultHighLimit
func NewBACnetConstructedDataFaultHighLimit(faultHighLimit BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFaultHighLimit {
	_result := &_BACnetConstructedDataFaultHighLimit{
		FaultHighLimit:         faultHighLimit,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFaultHighLimit(structType any) BACnetConstructedDataFaultHighLimit {
	if casted, ok := structType.(BACnetConstructedDataFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFaultHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataFaultHighLimit"
}

func (m *_BACnetConstructedDataFaultHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFaultHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataFaultHighLimitParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFaultHighLimit, error) {
	return BACnetConstructedDataFaultHighLimitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataFaultHighLimitParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataFaultHighLimit, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataFaultHighLimit, error) {
		return BACnetConstructedDataFaultHighLimitParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataFaultHighLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFaultHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultHighLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "faultHighLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultHighLimit' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), faultHighLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFaultHighLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFaultHighLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultHighLimit: faultHighLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFaultHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFaultHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFaultHighLimit")
		}

		// Simple Field (faultHighLimit)
		if pushErr := writeBuffer.PushContext("faultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for faultHighLimit")
		}
		_faultHighLimitErr := writeBuffer.WriteSerializable(ctx, m.GetFaultHighLimit())
		if popErr := writeBuffer.PopContext("faultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for faultHighLimit")
		}
		if _faultHighLimitErr != nil {
			return errors.Wrap(_faultHighLimitErr, "Error serializing 'faultHighLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFaultHighLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFaultHighLimit) isBACnetConstructedDataFaultHighLimit() bool {
	return true
}

func (m *_BACnetConstructedDataFaultHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
