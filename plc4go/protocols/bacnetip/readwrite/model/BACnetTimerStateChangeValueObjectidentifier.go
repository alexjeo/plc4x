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

// BACnetTimerStateChangeValueObjectidentifier is the corresponding interface of BACnetTimerStateChangeValueObjectidentifier
type BACnetTimerStateChangeValueObjectidentifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
}

// BACnetTimerStateChangeValueObjectidentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueObjectidentifier.
// This is useful for switch cases.
type BACnetTimerStateChangeValueObjectidentifierExactly interface {
	BACnetTimerStateChangeValueObjectidentifier
	isBACnetTimerStateChangeValueObjectidentifier() bool
}

// _BACnetTimerStateChangeValueObjectidentifier is the data-structure of this message
type _BACnetTimerStateChangeValueObjectidentifier struct {
	*_BACnetTimerStateChangeValue
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueObjectidentifier) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueObjectidentifier factory function for _BACnetTimerStateChangeValueObjectidentifier
func NewBACnetTimerStateChangeValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueObjectidentifier {
	_result := &_BACnetTimerStateChangeValueObjectidentifier{
		ObjectidentifierValue:        objectidentifierValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueObjectidentifier(structType any) BACnetTimerStateChangeValueObjectidentifier {
	if casted, ok := structType.(BACnetTimerStateChangeValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetTypeName() string {
	return "BACnetTimerStateChangeValueObjectidentifier"
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerStateChangeValueObjectidentifierParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueObjectidentifier, error) {
	return BACnetTimerStateChangeValueObjectidentifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetTimerStateChangeValueObjectidentifierParseWithBufferProducer(objectTypeArgument BACnetObjectType) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerStateChangeValueObjectidentifier, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerStateChangeValueObjectidentifier, error) {
		return BACnetTimerStateChangeValueObjectidentifierParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	}
}

func BACnetTimerStateChangeValueObjectidentifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectidentifierValue, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectidentifierValue", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectidentifierValue' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueObjectidentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueObjectidentifier{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		ObjectidentifierValue: objectidentifierValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueObjectidentifier")
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectidentifierValue")
		}
		_objectidentifierValueErr := writeBuffer.WriteSerializable(ctx, m.GetObjectidentifierValue())
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectidentifierValue")
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueObjectidentifier")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) isBACnetTimerStateChangeValueObjectidentifier() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
