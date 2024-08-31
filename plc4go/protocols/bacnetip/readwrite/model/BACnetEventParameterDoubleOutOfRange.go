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

// BACnetEventParameterDoubleOutOfRange is the corresponding interface of BACnetEventParameterDoubleOutOfRange
type BACnetEventParameterDoubleOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagDouble
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagDouble
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagDouble
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterDoubleOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterDoubleOutOfRange.
// This is useful for switch cases.
type BACnetEventParameterDoubleOutOfRangeExactly interface {
	BACnetEventParameterDoubleOutOfRange
	isBACnetEventParameterDoubleOutOfRange() bool
}

// _BACnetEventParameterDoubleOutOfRange is the data-structure of this message
type _BACnetEventParameterDoubleOutOfRange struct {
	*_BACnetEventParameter
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagDouble
	HighLimit  BACnetContextTagDouble
	Deadband   BACnetContextTagDouble
	ClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterDoubleOutOfRange) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterDoubleOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetLowLimit() BACnetContextTagDouble {
	return m.LowLimit
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetHighLimit() BACnetContextTagDouble {
	return m.HighLimit
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetDeadband() BACnetContextTagDouble {
	return m.Deadband
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterDoubleOutOfRange factory function for _BACnetEventParameterDoubleOutOfRange
func NewBACnetEventParameterDoubleOutOfRange(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagDouble, highLimit BACnetContextTagDouble, deadband BACnetContextTagDouble, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterDoubleOutOfRange {
	_result := &_BACnetEventParameterDoubleOutOfRange{
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		LowLimit:              lowLimit,
		HighLimit:             highLimit,
		Deadband:              deadband,
		ClosingTag:            closingTag,
		_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterDoubleOutOfRange(structType any) BACnetEventParameterDoubleOutOfRange {
	if casted, ok := structType.(BACnetEventParameterDoubleOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterDoubleOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetTypeName() string {
	return "BACnetEventParameterDoubleOutOfRange"
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterDoubleOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterDoubleOutOfRangeParse(ctx context.Context, theBytes []byte) (BACnetEventParameterDoubleOutOfRange, error) {
	return BACnetEventParameterDoubleOutOfRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterDoubleOutOfRangeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterDoubleOutOfRange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterDoubleOutOfRange, error) {
		return BACnetEventParameterDoubleOutOfRangeParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetEventParameterDoubleOutOfRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterDoubleOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventParameterDoubleOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterDoubleOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(14))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}

	lowLimit, err := ReadSimpleField[BACnetContextTagDouble](ctx, "lowLimit", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}

	highLimit, err := ReadSimpleField[BACnetContextTagDouble](ctx, "highLimit", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}

	deadband, err := ReadSimpleField[BACnetContextTagDouble](ctx, "deadband", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(14))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterDoubleOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterDoubleOutOfRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterDoubleOutOfRange{
		_BACnetEventParameter: &_BACnetEventParameter{},
		OpeningTag:            openingTag,
		TimeDelay:             timeDelay,
		LowLimit:              lowLimit,
		HighLimit:             highLimit,
		Deadband:              deadband,
		ClosingTag:            closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterDoubleOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterDoubleOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterDoubleOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterDoubleOutOfRange")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(ctx, m.GetTimeDelay())
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (lowLimit)
		if pushErr := writeBuffer.PushContext("lowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowLimit")
		}
		_lowLimitErr := writeBuffer.WriteSerializable(ctx, m.GetLowLimit())
		if popErr := writeBuffer.PopContext("lowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowLimit")
		}
		if _lowLimitErr != nil {
			return errors.Wrap(_lowLimitErr, "Error serializing 'lowLimit' field")
		}

		// Simple Field (highLimit)
		if pushErr := writeBuffer.PushContext("highLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highLimit")
		}
		_highLimitErr := writeBuffer.WriteSerializable(ctx, m.GetHighLimit())
		if popErr := writeBuffer.PopContext("highLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highLimit")
		}
		if _highLimitErr != nil {
			return errors.Wrap(_highLimitErr, "Error serializing 'highLimit' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deadband")
		}
		_deadbandErr := writeBuffer.WriteSerializable(ctx, m.GetDeadband())
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deadband")
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterDoubleOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterDoubleOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterDoubleOutOfRange) isBACnetEventParameterDoubleOutOfRange() bool {
	return true
}

func (m *_BACnetEventParameterDoubleOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
