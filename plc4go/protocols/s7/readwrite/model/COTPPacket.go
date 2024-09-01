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

// COTPPacket is the corresponding interface of COTPPacket
type COTPPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTpduCode returns TpduCode (discriminator field)
	GetTpduCode() uint8
	// GetParameters returns Parameters (property field)
	GetParameters() []COTPParameter
	// GetPayload returns Payload (property field)
	GetPayload() S7Message
}

// COTPPacketExactly can be used when we want exactly this type and not a type which fulfills COTPPacket.
// This is useful for switch cases.
type COTPPacketExactly interface {
	COTPPacket
	isCOTPPacket() bool
}

// _COTPPacket is the data-structure of this message
type _COTPPacket struct {
	_COTPPacketChildRequirements
	Parameters []COTPParameter
	Payload    S7Message

	// Arguments.
	CotpLen uint16
}

type _COTPPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetTpduCode() uint8
}

type COTPPacketParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type COTPPacketChild interface {
	utils.Serializable
	InitializeParent(parent COTPPacket, parameters []COTPParameter, payload S7Message)
	GetParent() *COTPPacket

	GetTypeName() string
	COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacket) GetParameters() []COTPParameter {
	return m.Parameters
}

func (m *_COTPPacket) GetPayload() S7Message {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacket factory function for _COTPPacket
func NewCOTPPacket(parameters []COTPParameter, payload S7Message, cotpLen uint16) *_COTPPacket {
	return &_COTPPacket{Parameters: parameters, Payload: payload, CotpLen: cotpLen}
}

// Deprecated: use the interface for direct cast
func CastCOTPPacket(structType any) COTPPacket {
	if casted, ok := structType.(COTPPacket); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacket); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacket) GetTypeName() string {
	return "COTPPacket"
}

func (m *_COTPPacket) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8
	// Discriminator Field (tpduCode)
	lengthInBits += 8

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += m.Payload.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_COTPPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPPacketParse(ctx context.Context, theBytes []byte, cotpLen uint16) (COTPPacket, error) {
	return COTPPacketParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cotpLen)
}

func COTPPacketParseWithBufferProducer[T COTPPacket](cotpLen uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		buffer, err := COTPPacketParseWithBuffer(ctx, readBuffer, cotpLen)
		if err != nil {
			var zero T
			return zero, err
		}
		return buffer.(T), err
	}
}

func COTPPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos

	headerLength, err := ReadImplicitField[uint8](ctx, "headerLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'headerLength' field"))
	}
	_ = headerLength

	tpduCode, err := ReadDiscriminatorField[uint8](ctx, "tpduCode", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tpduCode' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type COTPPacketChildSerializeRequirement interface {
		COTPPacket
		InitializeParent(COTPPacket, []COTPParameter, S7Message)
		GetParent() COTPPacket
	}
	var _childTemp any
	var _child COTPPacketChildSerializeRequirement
	var typeSwitchError error
	switch {
	case tpduCode == 0xF0: // COTPPacketData
		_childTemp, typeSwitchError = COTPPacketDataParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0xE0: // COTPPacketConnectionRequest
		_childTemp, typeSwitchError = COTPPacketConnectionRequestParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0xD0: // COTPPacketConnectionResponse
		_childTemp, typeSwitchError = COTPPacketConnectionResponseParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0x80: // COTPPacketDisconnectRequest
		_childTemp, typeSwitchError = COTPPacketDisconnectRequestParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0xC0: // COTPPacketDisconnectResponse
		_childTemp, typeSwitchError = COTPPacketDisconnectResponseParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0x70: // COTPPacketTpduError
		_childTemp, typeSwitchError = COTPPacketTpduErrorParseWithBuffer(ctx, readBuffer, cotpLen)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [tpduCode=%v]", tpduCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of COTPPacket")
	}
	_child = _childTemp.(COTPPacketChildSerializeRequirement)

	parameters, err := ReadLengthArrayField[COTPParameter](ctx, "parameters", ReadComplex[COTPParameter](COTPParameterParseWithBufferProducer[COTPParameter]((uint8)(uint8((uint8(headerLength)+uint8(uint8(1))))-uint8((positionAware.GetPos()-startPos)))), readBuffer), int(int32((int32(headerLength)+int32(int32(1))))-int32((positionAware.GetPos()-startPos))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}

	_payload, err := ReadOptionalField[S7Message](ctx, "payload", ReadComplex[S7Message](S7MessageParseWithBuffer, readBuffer), bool((positionAware.GetPos()-startPos) < (cotpLen)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	var payload S7Message
	if _payload != nil {
		payload = *_payload
	}

	if closeErr := readBuffer.CloseContext("COTPPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacket")
	}

	// Finish initializing
	_child.InitializeParent(_child, parameters, payload)
	return _child, nil
}

func (pm *_COTPPacket) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("COTPPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for COTPPacket")
	}

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8((uint8((utils.InlineIf((bool((m.GetPayload()) != (nil))), func() any { return uint8((m.GetPayload()).GetLengthInBytes(ctx)) }, func() any { return uint8(uint8(0)) }).(uint8))) + uint8(uint8(1)))))
	_headerLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("headerLength", 8, uint8((headerLength)))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	if err := WriteDiscriminatorField(ctx, "tpduCode", m.GetTpduCode(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'tpduCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteComplexTypeArrayField(ctx, "parameters", m.GetParameters(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'parameters' field")
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload S7Message = nil
	if m.GetPayload() != nil {
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		payload = m.GetPayload()
		_payloadErr := writeBuffer.WriteSerializable(ctx, payload)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	if popErr := writeBuffer.PopContext("COTPPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for COTPPacket")
	}
	return nil
}

////
// Arguments Getter

func (m *_COTPPacket) GetCotpLen() uint16 {
	return m.CotpLen
}

//
////

func (m *_COTPPacket) isCOTPPacket() bool {
	return true
}

func (m *_COTPPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
