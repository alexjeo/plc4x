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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckReadRange is the data-structure of this message
type BACnetServiceAckReadRange struct {
	*BACnetServiceAck
	ObjectIdentifier    *BACnetContextTagObjectIdentifier
	PropertyIdentifier  *BACnetPropertyIdentifierTagged
	PropertyArrayIndex  *BACnetContextTagUnsignedInteger
	ResultFlags         *BACnetResultFlagsTagged
	ItemCount           *BACnetContextTagUnsignedInteger
	ItemData            *BACnetConstructedData
	FirstSequenceNumber *BACnetContextTagUnsignedInteger

	// Arguments.
	ServiceAckLength uint16
}

// IBACnetServiceAckReadRange is the corresponding interface of BACnetServiceAckReadRange
type IBACnetServiceAckReadRange interface {
	IBACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() *BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() *BACnetContextTagUnsignedInteger
	// GetResultFlags returns ResultFlags (property field)
	GetResultFlags() *BACnetResultFlagsTagged
	// GetItemCount returns ItemCount (property field)
	GetItemCount() *BACnetContextTagUnsignedInteger
	// GetItemData returns ItemData (property field)
	GetItemData() *BACnetConstructedData
	// GetFirstSequenceNumber returns FirstSequenceNumber (property field)
	GetFirstSequenceNumber() *BACnetContextTagUnsignedInteger
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

func (m *BACnetServiceAckReadRange) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_RANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckReadRange) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckReadRange) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetServiceAckReadRange) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetServiceAckReadRange) GetPropertyIdentifier() *BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *BACnetServiceAckReadRange) GetPropertyArrayIndex() *BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *BACnetServiceAckReadRange) GetResultFlags() *BACnetResultFlagsTagged {
	return m.ResultFlags
}

func (m *BACnetServiceAckReadRange) GetItemCount() *BACnetContextTagUnsignedInteger {
	return m.ItemCount
}

func (m *BACnetServiceAckReadRange) GetItemData() *BACnetConstructedData {
	return m.ItemData
}

func (m *BACnetServiceAckReadRange) GetFirstSequenceNumber() *BACnetContextTagUnsignedInteger {
	return m.FirstSequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckReadRange factory function for BACnetServiceAckReadRange
func NewBACnetServiceAckReadRange(objectIdentifier *BACnetContextTagObjectIdentifier, propertyIdentifier *BACnetPropertyIdentifierTagged, propertyArrayIndex *BACnetContextTagUnsignedInteger, resultFlags *BACnetResultFlagsTagged, itemCount *BACnetContextTagUnsignedInteger, itemData *BACnetConstructedData, firstSequenceNumber *BACnetContextTagUnsignedInteger, serviceAckLength uint16) *BACnetServiceAckReadRange {
	_result := &BACnetServiceAckReadRange{
		ObjectIdentifier:    objectIdentifier,
		PropertyIdentifier:  propertyIdentifier,
		PropertyArrayIndex:  propertyArrayIndex,
		ResultFlags:         resultFlags,
		ItemCount:           itemCount,
		ItemData:            itemData,
		FirstSequenceNumber: firstSequenceNumber,
		BACnetServiceAck:    NewBACnetServiceAck(serviceAckLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckReadRange(structType interface{}) *BACnetServiceAckReadRange {
	if casted, ok := structType.(BACnetServiceAckReadRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckReadRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckReadRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckReadRange(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckReadRange) GetTypeName() string {
	return "BACnetServiceAckReadRange"
}

func (m *BACnetServiceAckReadRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckReadRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += (*m.PropertyArrayIndex).GetLengthInBits()
	}

	// Simple field (resultFlags)
	lengthInBits += m.ResultFlags.GetLengthInBits()

	// Simple field (itemCount)
	lengthInBits += m.ItemCount.GetLengthInBits()

	// Optional Field (itemData)
	if m.ItemData != nil {
		lengthInBits += (*m.ItemData).GetLengthInBits()
	}

	// Optional Field (firstSequenceNumber)
	if m.FirstSequenceNumber != nil {
		lengthInBits += (*m.FirstSequenceNumber).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetServiceAckReadRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckReadRangeParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (*BACnetServiceAckReadRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadRange"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetPropertyIdentifierTagged(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if a given expression evaluates to false)
	var propertyArrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyArrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyArrayIndex' field")
		default:
			propertyArrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("propertyArrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (resultFlags)
	if pullErr := readBuffer.PullContext("resultFlags"); pullErr != nil {
		return nil, pullErr
	}
	_resultFlags, _resultFlagsErr := BACnetResultFlagsTaggedParse(readBuffer, uint8(uint8(3)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _resultFlagsErr != nil {
		return nil, errors.Wrap(_resultFlagsErr, "Error parsing 'resultFlags' field")
	}
	resultFlags := CastBACnetResultFlagsTagged(_resultFlags)
	if closeErr := readBuffer.CloseContext("resultFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (itemCount)
	if pullErr := readBuffer.PullContext("itemCount"); pullErr != nil {
		return nil, pullErr
	}
	_itemCount, _itemCountErr := BACnetContextTagParse(readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _itemCountErr != nil {
		return nil, errors.Wrap(_itemCountErr, "Error parsing 'itemCount' field")
	}
	itemCount := CastBACnetContextTagUnsignedInteger(_itemCount)
	if closeErr := readBuffer.CloseContext("itemCount"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (itemData) (Can be skipped, if a given expression evaluates to false)
	var itemData *BACnetConstructedData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("itemData"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(5), objectIdentifier.GetObjectType(), propertyIdentifier.GetValue())
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'itemData' field")
		default:
			itemData = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("itemData"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (firstSequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var firstSequenceNumber *BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("firstSequenceNumber"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(6), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'firstSequenceNumber' field")
		default:
			firstSequenceNumber = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("firstSequenceNumber"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadRange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckReadRange{
		ObjectIdentifier:    CastBACnetContextTagObjectIdentifier(objectIdentifier),
		PropertyIdentifier:  CastBACnetPropertyIdentifierTagged(propertyIdentifier),
		PropertyArrayIndex:  CastBACnetContextTagUnsignedInteger(propertyArrayIndex),
		ResultFlags:         CastBACnetResultFlagsTagged(resultFlags),
		ItemCount:           CastBACnetContextTagUnsignedInteger(itemCount),
		ItemData:            CastBACnetConstructedData(itemData),
		FirstSequenceNumber: CastBACnetContextTagUnsignedInteger(firstSequenceNumber),
		BACnetServiceAck:    &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckReadRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadRange"); pushErr != nil {
			return pushErr
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return pushErr
		}
		_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return popErr
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		// Simple Field (propertyIdentifier)
		if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
			return pushErr
		}
		_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
			return popErr
		}
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
		var propertyArrayIndex *BACnetContextTagUnsignedInteger = nil
		if m.PropertyArrayIndex != nil {
			if pushErr := writeBuffer.PushContext("propertyArrayIndex"); pushErr != nil {
				return pushErr
			}
			propertyArrayIndex = m.PropertyArrayIndex
			_propertyArrayIndexErr := propertyArrayIndex.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("propertyArrayIndex"); popErr != nil {
				return popErr
			}
			if _propertyArrayIndexErr != nil {
				return errors.Wrap(_propertyArrayIndexErr, "Error serializing 'propertyArrayIndex' field")
			}
		}

		// Simple Field (resultFlags)
		if pushErr := writeBuffer.PushContext("resultFlags"); pushErr != nil {
			return pushErr
		}
		_resultFlagsErr := m.ResultFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("resultFlags"); popErr != nil {
			return popErr
		}
		if _resultFlagsErr != nil {
			return errors.Wrap(_resultFlagsErr, "Error serializing 'resultFlags' field")
		}

		// Simple Field (itemCount)
		if pushErr := writeBuffer.PushContext("itemCount"); pushErr != nil {
			return pushErr
		}
		_itemCountErr := m.ItemCount.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("itemCount"); popErr != nil {
			return popErr
		}
		if _itemCountErr != nil {
			return errors.Wrap(_itemCountErr, "Error serializing 'itemCount' field")
		}

		// Optional Field (itemData) (Can be skipped, if the value is null)
		var itemData *BACnetConstructedData = nil
		if m.ItemData != nil {
			if pushErr := writeBuffer.PushContext("itemData"); pushErr != nil {
				return pushErr
			}
			itemData = m.ItemData
			_itemDataErr := itemData.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("itemData"); popErr != nil {
				return popErr
			}
			if _itemDataErr != nil {
				return errors.Wrap(_itemDataErr, "Error serializing 'itemData' field")
			}
		}

		// Optional Field (firstSequenceNumber) (Can be skipped, if the value is null)
		var firstSequenceNumber *BACnetContextTagUnsignedInteger = nil
		if m.FirstSequenceNumber != nil {
			if pushErr := writeBuffer.PushContext("firstSequenceNumber"); pushErr != nil {
				return pushErr
			}
			firstSequenceNumber = m.FirstSequenceNumber
			_firstSequenceNumberErr := firstSequenceNumber.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("firstSequenceNumber"); popErr != nil {
				return popErr
			}
			if _firstSequenceNumberErr != nil {
				return errors.Wrap(_firstSequenceNumberErr, "Error serializing 'firstSequenceNumber' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadRange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
