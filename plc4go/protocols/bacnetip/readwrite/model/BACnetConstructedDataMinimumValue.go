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

// BACnetConstructedDataMinimumValue is the corresponding interface of BACnetConstructedDataMinimumValue
type BACnetConstructedDataMinimumValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinimumValue returns MinimumValue (property field)
	GetMinimumValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataMinimumValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMinimumValue.
// This is useful for switch cases.
type BACnetConstructedDataMinimumValueExactly interface {
	BACnetConstructedDataMinimumValue
	isBACnetConstructedDataMinimumValue() bool
}

// _BACnetConstructedDataMinimumValue is the data-structure of this message
type _BACnetConstructedDataMinimumValue struct {
	*_BACnetConstructedData
	MinimumValue BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMinimumValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValue) GetMinimumValue() BACnetApplicationTagReal {
	return m.MinimumValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMinimumValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumValue factory function for _BACnetConstructedDataMinimumValue
func NewBACnetConstructedDataMinimumValue(minimumValue BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumValue {
	_result := &_BACnetConstructedDataMinimumValue{
		MinimumValue:           minimumValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumValue(structType any) BACnetConstructedDataMinimumValue {
	if casted, ok := structType.(BACnetConstructedDataMinimumValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumValue) GetTypeName() string {
	return "BACnetConstructedDataMinimumValue"
}

func (m *_BACnetConstructedDataMinimumValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (minimumValue)
	lengthInBits += m.MinimumValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataMinimumValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMinimumValue, error) {
	return BACnetConstructedDataMinimumValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataMinimumValueParseWithBufferProducer(tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataMinimumValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConstructedDataMinimumValue, error) {
		return BACnetConstructedDataMinimumValueParseWithBuffer(ctx, readBuffer, tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
	}
}

func BACnetConstructedDataMinimumValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMinimumValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minimumValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "minimumValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minimumValue' field"))
	}

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), minimumValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMinimumValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MinimumValue: minimumValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMinimumValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumValue")
		}

		// Simple Field (minimumValue)
		if pushErr := writeBuffer.PushContext("minimumValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumValue")
		}
		_minimumValueErr := writeBuffer.WriteSerializable(ctx, m.GetMinimumValue())
		if popErr := writeBuffer.PopContext("minimumValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumValue")
		}
		if _minimumValueErr != nil {
			return errors.Wrap(_minimumValueErr, "Error serializing 'minimumValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumValue) isBACnetConstructedDataMinimumValue() bool {
	return true
}

func (m *_BACnetConstructedDataMinimumValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
