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

// BACnetEventParameterChangeOfValueCivCriteriaBitmask is the corresponding interface of BACnetEventParameterChangeOfValueCivCriteriaBitmask
type BACnetEventParameterChangeOfValueCivCriteriaBitmask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameterChangeOfValueCivCriteria
	// GetBitmask returns Bitmask (property field)
	GetBitmask() BACnetContextTagBitString
}

// BACnetEventParameterChangeOfValueCivCriteriaBitmaskExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfValueCivCriteriaBitmask.
// This is useful for switch cases.
type BACnetEventParameterChangeOfValueCivCriteriaBitmaskExactly interface {
	BACnetEventParameterChangeOfValueCivCriteriaBitmask
	isBACnetEventParameterChangeOfValueCivCriteriaBitmask() bool
}

// _BACnetEventParameterChangeOfValueCivCriteriaBitmask is the data-structure of this message
type _BACnetEventParameterChangeOfValueCivCriteriaBitmask struct {
	*_BACnetEventParameterChangeOfValueCivCriteria
	Bitmask BACnetContextTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) InitializeParent(parent BACnetEventParameterChangeOfValueCivCriteria, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) GetParent() BACnetEventParameterChangeOfValueCivCriteria {
	return m._BACnetEventParameterChangeOfValueCivCriteria
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) GetBitmask() BACnetContextTagBitString {
	return m.Bitmask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfValueCivCriteriaBitmask factory function for _BACnetEventParameterChangeOfValueCivCriteriaBitmask
func NewBACnetEventParameterChangeOfValueCivCriteriaBitmask(bitmask BACnetContextTagBitString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfValueCivCriteriaBitmask {
	_result := &_BACnetEventParameterChangeOfValueCivCriteriaBitmask{
		Bitmask: bitmask,
		_BACnetEventParameterChangeOfValueCivCriteria: NewBACnetEventParameterChangeOfValueCivCriteria(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetEventParameterChangeOfValueCivCriteria._BACnetEventParameterChangeOfValueCivCriteriaChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfValueCivCriteriaBitmask(structType any) BACnetEventParameterChangeOfValueCivCriteriaBitmask {
	if casted, ok := structType.(BACnetEventParameterChangeOfValueCivCriteriaBitmask); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfValueCivCriteriaBitmask); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) GetTypeName() string {
	return "BACnetEventParameterChangeOfValueCivCriteriaBitmask"
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bitmask)
	lengthInBits += m.Bitmask.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterChangeOfValueCivCriteriaBitmaskParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterChangeOfValueCivCriteriaBitmask, error) {
	return BACnetEventParameterChangeOfValueCivCriteriaBitmaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterChangeOfValueCivCriteriaBitmaskParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfValueCivCriteriaBitmask, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfValueCivCriteriaBitmask, error) {
		return BACnetEventParameterChangeOfValueCivCriteriaBitmaskParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterChangeOfValueCivCriteriaBitmaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfValueCivCriteriaBitmask, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfValueCivCriteriaBitmask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfValueCivCriteriaBitmask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bitmask, err := ReadSimpleField[BACnetContextTagBitString](ctx, "bitmask", ReadComplex[BACnetContextTagBitString](BACnetContextTagParseWithBufferProducer[BACnetContextTagBitString]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BIT_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitmask' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfValueCivCriteriaBitmask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfValueCivCriteriaBitmask")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfValueCivCriteriaBitmask{
		_BACnetEventParameterChangeOfValueCivCriteria: &_BACnetEventParameterChangeOfValueCivCriteria{
			TagNumber: tagNumber,
		},
		Bitmask: bitmask,
	}
	_child._BACnetEventParameterChangeOfValueCivCriteria._BACnetEventParameterChangeOfValueCivCriteriaChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfValueCivCriteriaBitmask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfValueCivCriteriaBitmask")
		}

		// Simple Field (bitmask)
		if pushErr := writeBuffer.PushContext("bitmask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitmask")
		}
		_bitmaskErr := writeBuffer.WriteSerializable(ctx, m.GetBitmask())
		if popErr := writeBuffer.PopContext("bitmask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitmask")
		}
		if _bitmaskErr != nil {
			return errors.Wrap(_bitmaskErr, "Error serializing 'bitmask' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfValueCivCriteriaBitmask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfValueCivCriteriaBitmask")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) isBACnetEventParameterChangeOfValueCivCriteriaBitmask() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfValueCivCriteriaBitmask) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
