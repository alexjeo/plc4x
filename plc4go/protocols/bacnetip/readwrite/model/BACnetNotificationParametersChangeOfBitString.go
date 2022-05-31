/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersChangeOfBitString is the data-structure of this message
type BACnetNotificationParametersChangeOfBitString struct {
	*BACnetNotificationParameters
	InnerOpeningTag   *BACnetOpeningTag
	ChangeOfBitString *BACnetContextTagBitString
	StatusFlags       *BACnetStatusFlagsTagged
	InnerClosingTag   *BACnetClosingTag

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
}

// IBACnetNotificationParametersChangeOfBitString is the corresponding interface of BACnetNotificationParametersChangeOfBitString
type IBACnetNotificationParametersChangeOfBitString interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetChangeOfBitString returns ChangeOfBitString (property field)
	GetChangeOfBitString() *BACnetContextTagBitString
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() *BACnetStatusFlagsTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetNotificationParametersChangeOfBitString) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfBitString) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfBitString) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersChangeOfBitString) GetChangeOfBitString() *BACnetContextTagBitString {
	return m.ChangeOfBitString
}

func (m *BACnetNotificationParametersChangeOfBitString) GetStatusFlags() *BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *BACnetNotificationParametersChangeOfBitString) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfBitString factory function for BACnetNotificationParametersChangeOfBitString
func NewBACnetNotificationParametersChangeOfBitString(innerOpeningTag *BACnetOpeningTag, changeOfBitString *BACnetContextTagBitString, statusFlags *BACnetStatusFlagsTagged, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParametersChangeOfBitString {
	_result := &BACnetNotificationParametersChangeOfBitString{
		InnerOpeningTag:              innerOpeningTag,
		ChangeOfBitString:            changeOfBitString,
		StatusFlags:                  statusFlags,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectType),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfBitString(structType interface{}) *BACnetNotificationParametersChangeOfBitString {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfBitString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfBitString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfBitString(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfBitString(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfBitString) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfBitString"
}

func (m *BACnetNotificationParametersChangeOfBitString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfBitString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (changeOfBitString)
	lengthInBits += m.ChangeOfBitString.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfBitString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfBitStringParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfBitString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (changeOfBitString)
	if pullErr := readBuffer.PullContext("changeOfBitString"); pullErr != nil {
		return nil, pullErr
	}
	_changeOfBitString, _changeOfBitStringErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _changeOfBitStringErr != nil {
		return nil, errors.Wrap(_changeOfBitStringErr, "Error parsing 'changeOfBitString' field")
	}
	changeOfBitString := CastBACnetContextTagBitString(_changeOfBitString)
	if closeErr := readBuffer.CloseContext("changeOfBitString"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, pullErr
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlagsTagged(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfBitString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfBitString{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ChangeOfBitString:            CastBACnetContextTagBitString(changeOfBitString),
		StatusFlags:                  CastBACnetStatusFlagsTagged(statusFlags),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfBitString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfBitString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (changeOfBitString)
		if pushErr := writeBuffer.PushContext("changeOfBitString"); pushErr != nil {
			return pushErr
		}
		_changeOfBitStringErr := m.ChangeOfBitString.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changeOfBitString"); popErr != nil {
			return popErr
		}
		if _changeOfBitStringErr != nil {
			return errors.Wrap(_changeOfBitStringErr, "Error serializing 'changeOfBitString' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return pushErr
		}
		_statusFlagsErr := m.StatusFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return popErr
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfBitString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
