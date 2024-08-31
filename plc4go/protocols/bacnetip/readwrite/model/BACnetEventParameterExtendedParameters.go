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

// BACnetEventParameterExtendedParameters is the corresponding interface of BACnetEventParameterExtendedParameters
type BACnetEventParameterExtendedParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
	// GetTimeValue returns TimeValue (property field)
	GetTimeValue() BACnetApplicationTagTime
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetReference returns Reference (property field)
	GetReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetIsOpeningTag returns IsOpeningTag (virtual field)
	GetIsOpeningTag() bool
	// GetIsClosingTag returns IsClosingTag (virtual field)
	GetIsClosingTag() bool
}

// BACnetEventParameterExtendedParametersExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterExtendedParameters.
// This is useful for switch cases.
type BACnetEventParameterExtendedParametersExactly interface {
	BACnetEventParameterExtendedParameters
	isBACnetEventParameterExtendedParameters() bool
}

// _BACnetEventParameterExtendedParameters is the data-structure of this message
type _BACnetEventParameterExtendedParameters struct {
	OpeningTag           BACnetOpeningTag
	PeekedTagHeader      BACnetTagHeader
	NullValue            BACnetApplicationTagNull
	RealValue            BACnetApplicationTagReal
	UnsignedValue        BACnetApplicationTagUnsignedInteger
	BooleanValue         BACnetApplicationTagBoolean
	IntegerValue         BACnetApplicationTagSignedInteger
	DoubleValue          BACnetApplicationTagDouble
	OctetStringValue     BACnetApplicationTagOctetString
	CharacterStringValue BACnetApplicationTagCharacterString
	BitStringValue       BACnetApplicationTagBitString
	EnumeratedValue      BACnetApplicationTagEnumerated
	DateValue            BACnetApplicationTagDate
	TimeValue            BACnetApplicationTagTime
	ObjectIdentifier     BACnetApplicationTagObjectIdentifier
	Reference            BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterExtendedParameters) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterExtendedParameters) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetEventParameterExtendedParameters) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

func (m *_BACnetEventParameterExtendedParameters) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

func (m *_BACnetEventParameterExtendedParameters) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

func (m *_BACnetEventParameterExtendedParameters) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

func (m *_BACnetEventParameterExtendedParameters) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

func (m *_BACnetEventParameterExtendedParameters) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

func (m *_BACnetEventParameterExtendedParameters) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

func (m *_BACnetEventParameterExtendedParameters) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

func (m *_BACnetEventParameterExtendedParameters) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

func (m *_BACnetEventParameterExtendedParameters) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

func (m *_BACnetEventParameterExtendedParameters) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

func (m *_BACnetEventParameterExtendedParameters) GetTimeValue() BACnetApplicationTagTime {
	return m.TimeValue
}

func (m *_BACnetEventParameterExtendedParameters) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetEventParameterExtendedParameters) GetReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.Reference
}

func (m *_BACnetEventParameterExtendedParameters) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetEventParameterExtendedParameters) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	nullValue := m.NullValue
	_ = nullValue
	realValue := m.RealValue
	_ = realValue
	unsignedValue := m.UnsignedValue
	_ = unsignedValue
	booleanValue := m.BooleanValue
	_ = booleanValue
	integerValue := m.IntegerValue
	_ = integerValue
	doubleValue := m.DoubleValue
	_ = doubleValue
	octetStringValue := m.OctetStringValue
	_ = octetStringValue
	characterStringValue := m.CharacterStringValue
	_ = characterStringValue
	bitStringValue := m.BitStringValue
	_ = bitStringValue
	enumeratedValue := m.EnumeratedValue
	_ = enumeratedValue
	dateValue := m.DateValue
	_ = dateValue
	timeValue := m.TimeValue
	_ = timeValue
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	reference := m.Reference
	_ = reference
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetEventParameterExtendedParameters) GetIsOpeningTag() bool {
	ctx := context.Background()
	_ = ctx
	nullValue := m.NullValue
	_ = nullValue
	realValue := m.RealValue
	_ = realValue
	unsignedValue := m.UnsignedValue
	_ = unsignedValue
	booleanValue := m.BooleanValue
	_ = booleanValue
	integerValue := m.IntegerValue
	_ = integerValue
	doubleValue := m.DoubleValue
	_ = doubleValue
	octetStringValue := m.OctetStringValue
	_ = octetStringValue
	characterStringValue := m.CharacterStringValue
	_ = characterStringValue
	bitStringValue := m.BitStringValue
	_ = bitStringValue
	enumeratedValue := m.EnumeratedValue
	_ = enumeratedValue
	dateValue := m.DateValue
	_ = dateValue
	timeValue := m.TimeValue
	_ = timeValue
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	reference := m.Reference
	_ = reference
	return bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x6)))
}

func (m *_BACnetEventParameterExtendedParameters) GetIsClosingTag() bool {
	ctx := context.Background()
	_ = ctx
	nullValue := m.NullValue
	_ = nullValue
	realValue := m.RealValue
	_ = realValue
	unsignedValue := m.UnsignedValue
	_ = unsignedValue
	booleanValue := m.BooleanValue
	_ = booleanValue
	integerValue := m.IntegerValue
	_ = integerValue
	doubleValue := m.DoubleValue
	_ = doubleValue
	octetStringValue := m.OctetStringValue
	_ = octetStringValue
	characterStringValue := m.CharacterStringValue
	_ = characterStringValue
	bitStringValue := m.BitStringValue
	_ = bitStringValue
	enumeratedValue := m.EnumeratedValue
	_ = enumeratedValue
	dateValue := m.DateValue
	_ = dateValue
	timeValue := m.TimeValue
	_ = timeValue
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	reference := m.Reference
	_ = reference
	return bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x7)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterExtendedParameters factory function for _BACnetEventParameterExtendedParameters
func NewBACnetEventParameterExtendedParameters(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull, realValue BACnetApplicationTagReal, unsignedValue BACnetApplicationTagUnsignedInteger, booleanValue BACnetApplicationTagBoolean, integerValue BACnetApplicationTagSignedInteger, doubleValue BACnetApplicationTagDouble, octetStringValue BACnetApplicationTagOctetString, characterStringValue BACnetApplicationTagCharacterString, bitStringValue BACnetApplicationTagBitString, enumeratedValue BACnetApplicationTagEnumerated, dateValue BACnetApplicationTagDate, timeValue BACnetApplicationTagTime, objectIdentifier BACnetApplicationTagObjectIdentifier, reference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterExtendedParameters {
	return &_BACnetEventParameterExtendedParameters{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, NullValue: nullValue, RealValue: realValue, UnsignedValue: unsignedValue, BooleanValue: booleanValue, IntegerValue: integerValue, DoubleValue: doubleValue, OctetStringValue: octetStringValue, CharacterStringValue: characterStringValue, BitStringValue: bitStringValue, EnumeratedValue: enumeratedValue, DateValue: dateValue, TimeValue: timeValue, ObjectIdentifier: objectIdentifier, Reference: reference, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterExtendedParameters(structType any) BACnetEventParameterExtendedParameters {
	if casted, ok := structType.(BACnetEventParameterExtendedParameters); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterExtendedParameters); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterExtendedParameters) GetTypeName() string {
	return "BACnetEventParameterExtendedParameters"
}

func (m *_BACnetEventParameterExtendedParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (nullValue)
	if m.NullValue != nil {
		lengthInBits += m.NullValue.GetLengthInBits(ctx)
	}

	// Optional Field (realValue)
	if m.RealValue != nil {
		lengthInBits += m.RealValue.GetLengthInBits(ctx)
	}

	// Optional Field (unsignedValue)
	if m.UnsignedValue != nil {
		lengthInBits += m.UnsignedValue.GetLengthInBits(ctx)
	}

	// Optional Field (booleanValue)
	if m.BooleanValue != nil {
		lengthInBits += m.BooleanValue.GetLengthInBits(ctx)
	}

	// Optional Field (integerValue)
	if m.IntegerValue != nil {
		lengthInBits += m.IntegerValue.GetLengthInBits(ctx)
	}

	// Optional Field (doubleValue)
	if m.DoubleValue != nil {
		lengthInBits += m.DoubleValue.GetLengthInBits(ctx)
	}

	// Optional Field (octetStringValue)
	if m.OctetStringValue != nil {
		lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)
	}

	// Optional Field (characterStringValue)
	if m.CharacterStringValue != nil {
		lengthInBits += m.CharacterStringValue.GetLengthInBits(ctx)
	}

	// Optional Field (bitStringValue)
	if m.BitStringValue != nil {
		lengthInBits += m.BitStringValue.GetLengthInBits(ctx)
	}

	// Optional Field (enumeratedValue)
	if m.EnumeratedValue != nil {
		lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)
	}

	// Optional Field (dateValue)
	if m.DateValue != nil {
		lengthInBits += m.DateValue.GetLengthInBits(ctx)
	}

	// Optional Field (timeValue)
	if m.TimeValue != nil {
		lengthInBits += m.TimeValue.GetLengthInBits(ctx)
	}

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)
	}

	// Optional Field (reference)
	if m.Reference != nil {
		lengthInBits += m.Reference.GetLengthInBits(ctx)
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterExtendedParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterExtendedParametersParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterExtendedParameters, error) {
	return BACnetEventParameterExtendedParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterExtendedParametersParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterExtendedParameters, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterExtendedParameters, error) {
		return BACnetEventParameterExtendedParametersParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterExtendedParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterExtendedParameters, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetEventParameterExtendedParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterExtendedParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	isOpeningTag, err := ReadVirtualField[bool](ctx, "isOpeningTag", (*bool)(nil), bool((peekedTagHeader.GetLengthValueType()) == (0x6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isOpeningTag' field"))
	}
	_ = isOpeningTag

	isClosingTag, err := ReadVirtualField[bool](ctx, "isClosingTag", (*bool)(nil), bool((peekedTagHeader.GetLengthValueType()) == (0x7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isClosingTag' field"))
	}
	_ = isClosingTag

	_nullValue, err := ReadOptionalField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x0))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	var nullValue BACnetApplicationTagNull
	if _nullValue != nil {
		nullValue = *_nullValue
	}

	_realValue, err := ReadOptionalField[BACnetApplicationTagReal](ctx, "realValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x4))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'realValue' field"))
	}
	var realValue BACnetApplicationTagReal
	if _realValue != nil {
		realValue = *_realValue
	}

	_unsignedValue, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "unsignedValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x2))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unsignedValue' field"))
	}
	var unsignedValue BACnetApplicationTagUnsignedInteger
	if _unsignedValue != nil {
		unsignedValue = *_unsignedValue
	}

	_booleanValue, err := ReadOptionalField[BACnetApplicationTagBoolean](ctx, "booleanValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x1))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'booleanValue' field"))
	}
	var booleanValue BACnetApplicationTagBoolean
	if _booleanValue != nil {
		booleanValue = *_booleanValue
	}

	_integerValue, err := ReadOptionalField[BACnetApplicationTagSignedInteger](ctx, "integerValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x3))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerValue' field"))
	}
	var integerValue BACnetApplicationTagSignedInteger
	if _integerValue != nil {
		integerValue = *_integerValue
	}

	_doubleValue, err := ReadOptionalField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x5))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	var doubleValue BACnetApplicationTagDouble
	if _doubleValue != nil {
		doubleValue = *_doubleValue
	}

	_octetStringValue, err := ReadOptionalField[BACnetApplicationTagOctetString](ctx, "octetStringValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x6))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octetStringValue' field"))
	}
	var octetStringValue BACnetApplicationTagOctetString
	if _octetStringValue != nil {
		octetStringValue = *_octetStringValue
	}

	_characterStringValue, err := ReadOptionalField[BACnetApplicationTagCharacterString](ctx, "characterStringValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x7))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'characterStringValue' field"))
	}
	var characterStringValue BACnetApplicationTagCharacterString
	if _characterStringValue != nil {
		characterStringValue = *_characterStringValue
	}

	_bitStringValue, err := ReadOptionalField[BACnetApplicationTagBitString](ctx, "bitStringValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x8))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	var bitStringValue BACnetApplicationTagBitString
	if _bitStringValue != nil {
		bitStringValue = *_bitStringValue
	}

	_enumeratedValue, err := ReadOptionalField[BACnetApplicationTagEnumerated](ctx, "enumeratedValue", ReadComplex[BACnetApplicationTagEnumerated](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagEnumerated](), readBuffer), bool(bool(bool((peekedTagNumber) == (0x9))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enumeratedValue' field"))
	}
	var enumeratedValue BACnetApplicationTagEnumerated
	if _enumeratedValue != nil {
		enumeratedValue = *_enumeratedValue
	}

	_dateValue, err := ReadOptionalField[BACnetApplicationTagDate](ctx, "dateValue", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer), bool(bool(bool((peekedTagNumber) == (0xA))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateValue' field"))
	}
	var dateValue BACnetApplicationTagDate
	if _dateValue != nil {
		dateValue = *_dateValue
	}

	_timeValue, err := ReadOptionalField[BACnetApplicationTagTime](ctx, "timeValue", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer), bool(bool(bool((peekedTagNumber) == (0xB))) && bool(!(isOpeningTag))) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	var timeValue BACnetApplicationTagTime
	if _timeValue != nil {
		timeValue = *_timeValue
	}

	_objectIdentifier, err := ReadOptionalField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer), bool(bool((peekedTagNumber) == (0xC))) && bool(!(isOpeningTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	var objectIdentifier BACnetApplicationTagObjectIdentifier
	if _objectIdentifier != nil {
		objectIdentifier = *_objectIdentifier
	}

	_reference, err := ReadOptionalField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "reference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer), bool(isOpeningTag) && bool(!(isClosingTag)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reference' field"))
	}
	var reference BACnetDeviceObjectPropertyReferenceEnclosed
	if _reference != nil {
		reference = *_reference
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterExtendedParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterExtendedParameters")
	}

	// Create the instance
	return &_BACnetEventParameterExtendedParameters{
		TagNumber:            tagNumber,
		OpeningTag:           openingTag,
		PeekedTagHeader:      peekedTagHeader,
		NullValue:            nullValue,
		RealValue:            realValue,
		UnsignedValue:        unsignedValue,
		BooleanValue:         booleanValue,
		IntegerValue:         integerValue,
		DoubleValue:          doubleValue,
		OctetStringValue:     octetStringValue,
		CharacterStringValue: characterStringValue,
		BitStringValue:       bitStringValue,
		EnumeratedValue:      enumeratedValue,
		DateValue:            dateValue,
		TimeValue:            timeValue,
		ObjectIdentifier:     objectIdentifier,
		Reference:            reference,
		ClosingTag:           closingTag,
	}, nil
}

func (m *_BACnetEventParameterExtendedParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterExtendedParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterExtendedParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterExtendedParameters")
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
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	isOpeningTag := m.GetIsOpeningTag()
	_ = isOpeningTag
	if _isOpeningTagErr := writeBuffer.WriteVirtual(ctx, "isOpeningTag", m.GetIsOpeningTag()); _isOpeningTagErr != nil {
		return errors.Wrap(_isOpeningTagErr, "Error serializing 'isOpeningTag' field")
	}
	// Virtual field
	isClosingTag := m.GetIsClosingTag()
	_ = isClosingTag
	if _isClosingTagErr := writeBuffer.WriteVirtual(ctx, "isClosingTag", m.GetIsClosingTag()); _isClosingTagErr != nil {
		return errors.Wrap(_isClosingTagErr, "Error serializing 'isClosingTag' field")
	}

	// Optional Field (nullValue) (Can be skipped, if the value is null)
	var nullValue BACnetApplicationTagNull = nil
	if m.GetNullValue() != nil {
		if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nullValue")
		}
		nullValue = m.GetNullValue()
		_nullValueErr := writeBuffer.WriteSerializable(ctx, nullValue)
		if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nullValue")
		}
		if _nullValueErr != nil {
			return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
		}
	}

	// Optional Field (realValue) (Can be skipped, if the value is null)
	var realValue BACnetApplicationTagReal = nil
	if m.GetRealValue() != nil {
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for realValue")
		}
		realValue = m.GetRealValue()
		_realValueErr := writeBuffer.WriteSerializable(ctx, realValue)
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for realValue")
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}
	}

	// Optional Field (unsignedValue) (Can be skipped, if the value is null)
	var unsignedValue BACnetApplicationTagUnsignedInteger = nil
	if m.GetUnsignedValue() != nil {
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		unsignedValue = m.GetUnsignedValue()
		_unsignedValueErr := writeBuffer.WriteSerializable(ctx, unsignedValue)
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}
	}

	// Optional Field (booleanValue) (Can be skipped, if the value is null)
	var booleanValue BACnetApplicationTagBoolean = nil
	if m.GetBooleanValue() != nil {
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		booleanValue = m.GetBooleanValue()
		_booleanValueErr := writeBuffer.WriteSerializable(ctx, booleanValue)
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}
	}

	// Optional Field (integerValue) (Can be skipped, if the value is null)
	var integerValue BACnetApplicationTagSignedInteger = nil
	if m.GetIntegerValue() != nil {
		if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for integerValue")
		}
		integerValue = m.GetIntegerValue()
		_integerValueErr := writeBuffer.WriteSerializable(ctx, integerValue)
		if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for integerValue")
		}
		if _integerValueErr != nil {
			return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
		}
	}

	// Optional Field (doubleValue) (Can be skipped, if the value is null)
	var doubleValue BACnetApplicationTagDouble = nil
	if m.GetDoubleValue() != nil {
		if pushErr := writeBuffer.PushContext("doubleValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doubleValue")
		}
		doubleValue = m.GetDoubleValue()
		_doubleValueErr := writeBuffer.WriteSerializable(ctx, doubleValue)
		if popErr := writeBuffer.PopContext("doubleValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doubleValue")
		}
		if _doubleValueErr != nil {
			return errors.Wrap(_doubleValueErr, "Error serializing 'doubleValue' field")
		}
	}

	// Optional Field (octetStringValue) (Can be skipped, if the value is null)
	var octetStringValue BACnetApplicationTagOctetString = nil
	if m.GetOctetStringValue() != nil {
		if pushErr := writeBuffer.PushContext("octetStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for octetStringValue")
		}
		octetStringValue = m.GetOctetStringValue()
		_octetStringValueErr := writeBuffer.WriteSerializable(ctx, octetStringValue)
		if popErr := writeBuffer.PopContext("octetStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for octetStringValue")
		}
		if _octetStringValueErr != nil {
			return errors.Wrap(_octetStringValueErr, "Error serializing 'octetStringValue' field")
		}
	}

	// Optional Field (characterStringValue) (Can be skipped, if the value is null)
	var characterStringValue BACnetApplicationTagCharacterString = nil
	if m.GetCharacterStringValue() != nil {
		if pushErr := writeBuffer.PushContext("characterStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for characterStringValue")
		}
		characterStringValue = m.GetCharacterStringValue()
		_characterStringValueErr := writeBuffer.WriteSerializable(ctx, characterStringValue)
		if popErr := writeBuffer.PopContext("characterStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for characterStringValue")
		}
		if _characterStringValueErr != nil {
			return errors.Wrap(_characterStringValueErr, "Error serializing 'characterStringValue' field")
		}
	}

	// Optional Field (bitStringValue) (Can be skipped, if the value is null)
	var bitStringValue BACnetApplicationTagBitString = nil
	if m.GetBitStringValue() != nil {
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitStringValue")
		}
		bitStringValue = m.GetBitStringValue()
		_bitStringValueErr := writeBuffer.WriteSerializable(ctx, bitStringValue)
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitStringValue")
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}
	}

	// Optional Field (enumeratedValue) (Can be skipped, if the value is null)
	var enumeratedValue BACnetApplicationTagEnumerated = nil
	if m.GetEnumeratedValue() != nil {
		if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enumeratedValue")
		}
		enumeratedValue = m.GetEnumeratedValue()
		_enumeratedValueErr := writeBuffer.WriteSerializable(ctx, enumeratedValue)
		if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enumeratedValue")
		}
		if _enumeratedValueErr != nil {
			return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
		}
	}

	// Optional Field (dateValue) (Can be skipped, if the value is null)
	var dateValue BACnetApplicationTagDate = nil
	if m.GetDateValue() != nil {
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateValue")
		}
		dateValue = m.GetDateValue()
		_dateValueErr := writeBuffer.WriteSerializable(ctx, dateValue)
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateValue")
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}
	}

	// Optional Field (timeValue) (Can be skipped, if the value is null)
	var timeValue BACnetApplicationTagTime = nil
	if m.GetTimeValue() != nil {
		if pushErr := writeBuffer.PushContext("timeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeValue")
		}
		timeValue = m.GetTimeValue()
		_timeValueErr := writeBuffer.WriteSerializable(ctx, timeValue)
		if popErr := writeBuffer.PopContext("timeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeValue")
		}
		if _timeValueErr != nil {
			return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
		}
	}

	// Optional Field (objectIdentifier) (Can be skipped, if the value is null)
	var objectIdentifier BACnetApplicationTagObjectIdentifier = nil
	if m.GetObjectIdentifier() != nil {
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		objectIdentifier = m.GetObjectIdentifier()
		_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, objectIdentifier)
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}
	}

	// Optional Field (reference) (Can be skipped, if the value is null)
	var reference BACnetDeviceObjectPropertyReferenceEnclosed = nil
	if m.GetReference() != nil {
		if pushErr := writeBuffer.PushContext("reference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reference")
		}
		reference = m.GetReference()
		_referenceErr := writeBuffer.WriteSerializable(ctx, reference)
		if popErr := writeBuffer.PopContext("reference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reference")
		}
		if _referenceErr != nil {
			return errors.Wrap(_referenceErr, "Error serializing 'reference' field")
		}
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

	if popErr := writeBuffer.PopContext("BACnetEventParameterExtendedParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterExtendedParameters")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterExtendedParameters) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterExtendedParameters) isBACnetEventParameterExtendedParameters() bool {
	return true
}

func (m *_BACnetEventParameterExtendedParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
