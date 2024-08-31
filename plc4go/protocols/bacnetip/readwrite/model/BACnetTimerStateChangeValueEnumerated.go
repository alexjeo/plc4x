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

// BACnetTimerStateChangeValueEnumerated is the corresponding interface of BACnetTimerStateChangeValueEnumerated
type BACnetTimerStateChangeValueEnumerated interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
}

// BACnetTimerStateChangeValueEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueEnumerated.
// This is useful for switch cases.
type BACnetTimerStateChangeValueEnumeratedExactly interface {
	BACnetTimerStateChangeValueEnumerated
	isBACnetTimerStateChangeValueEnumerated() bool
}

// _BACnetTimerStateChangeValueEnumerated is the data-structure of this message
type _BACnetTimerStateChangeValueEnumerated struct {
	*_BACnetTimerStateChangeValue
	EnumeratedValue BACnetApplicationTagEnumerated
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueEnumerated) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueEnumerated) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueEnumerated) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueEnumerated factory function for _BACnetTimerStateChangeValueEnumerated
func NewBACnetTimerStateChangeValueEnumerated(enumeratedValue BACnetApplicationTagEnumerated, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueEnumerated {
	_result := &_BACnetTimerStateChangeValueEnumerated{
		EnumeratedValue:              enumeratedValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueEnumerated(structType any) BACnetTimerStateChangeValueEnumerated {
	if casted, ok := structType.(BACnetTimerStateChangeValueEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueEnumerated) GetTypeName() string {
	return "BACnetTimerStateChangeValueEnumerated"
}

func (m *_BACnetTimerStateChangeValueEnumerated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueEnumerated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerStateChangeValueEnumeratedParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueEnumerated, error) {
	return BACnetTimerStateChangeValueEnumeratedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetTimerStateChangeValueEnumeratedParseWithBufferProducer(objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerStateChangeValueEnumerated, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerStateChangeValueEnumerated, error) {
		return BACnetTimerStateChangeValueEnumeratedParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	}
}

func BACnetTimerStateChangeValueEnumeratedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	enumeratedValue, err := ReadSimpleField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", ReadComplex[BACnetApplicationTagEnumerated](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagEnumerated](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumeratedValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueEnumerated")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueEnumerated{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		EnumeratedValue: enumeratedValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueEnumerated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueEnumerated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueEnumerated")
		}

		// Simple Field (enumeratedValue)
		if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enumeratedValue")
		}
		_enumeratedValueErr := writeBuffer.WriteSerializable(ctx, m.GetEnumeratedValue())
		if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enumeratedValue")
		}
		if _enumeratedValueErr != nil {
			return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueEnumerated")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueEnumerated) isBACnetTimerStateChangeValueEnumerated() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
