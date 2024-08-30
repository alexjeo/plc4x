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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetProtocolLevelTagged is the corresponding interface of BACnetProtocolLevelTagged
type BACnetProtocolLevelTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProtocolLevel
}

// BACnetProtocolLevelTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetProtocolLevelTagged.
// This is useful for switch cases.
type BACnetProtocolLevelTaggedExactly interface {
	BACnetProtocolLevelTagged
	isBACnetProtocolLevelTagged() bool
}

// _BACnetProtocolLevelTagged is the data-structure of this message
type _BACnetProtocolLevelTagged struct {
	Header BACnetTagHeader
	Value  BACnetProtocolLevel

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProtocolLevelTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetProtocolLevelTagged) GetValue() BACnetProtocolLevel {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProtocolLevelTagged factory function for _BACnetProtocolLevelTagged
func NewBACnetProtocolLevelTagged(header BACnetTagHeader, value BACnetProtocolLevel, tagNumber uint8, tagClass TagClass) *_BACnetProtocolLevelTagged {
	return &_BACnetProtocolLevelTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetProtocolLevelTagged(structType any) BACnetProtocolLevelTagged {
	if casted, ok := structType.(BACnetProtocolLevelTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProtocolLevelTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProtocolLevelTagged) GetTypeName() string {
	return "BACnetProtocolLevelTagged"
}

func (m *_BACnetProtocolLevelTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetProtocolLevelTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetProtocolLevelTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetProtocolLevelTagged, error) {
	return BACnetProtocolLevelTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetProtocolLevelTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetProtocolLevelTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetProtocolLevelTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProtocolLevelTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetProtocolLevelTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetProtocolLevel_PHYSICAL)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetProtocolLevelTagged")
	}
	var value BACnetProtocolLevel
	if _value != nil {
		value = _value.(BACnetProtocolLevel)
	}

	if closeErr := readBuffer.CloseContext("BACnetProtocolLevelTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProtocolLevelTagged")
	}

	// Create the instance
	return &_BACnetProtocolLevelTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetProtocolLevelTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProtocolLevelTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetProtocolLevelTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProtocolLevelTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(ctx, writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProtocolLevelTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProtocolLevelTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetProtocolLevelTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetProtocolLevelTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetProtocolLevelTagged) isBACnetProtocolLevelTagged() bool {
	return true
}

func (m *_BACnetProtocolLevelTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
