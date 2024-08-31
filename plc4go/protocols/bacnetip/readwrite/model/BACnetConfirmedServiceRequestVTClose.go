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

// BACnetConfirmedServiceRequestVTClose is the corresponding interface of BACnetConfirmedServiceRequestVTClose
type BACnetConfirmedServiceRequestVTClose interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetListOfRemoteVtSessionIdentifiers returns ListOfRemoteVtSessionIdentifiers (property field)
	GetListOfRemoteVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger
}

// BACnetConfirmedServiceRequestVTCloseExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestVTClose.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestVTCloseExactly interface {
	BACnetConfirmedServiceRequestVTClose
	isBACnetConfirmedServiceRequestVTClose() bool
}

// _BACnetConfirmedServiceRequestVTClose is the data-structure of this message
type _BACnetConfirmedServiceRequestVTClose struct {
	*_BACnetConfirmedServiceRequest
	ListOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger

	// Arguments.
	ServiceRequestPayloadLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetListOfRemoteVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger {
	return m.ListOfRemoteVtSessionIdentifiers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestVTClose factory function for _BACnetConfirmedServiceRequestVTClose
func NewBACnetConfirmedServiceRequestVTClose(listOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestVTClose {
	_result := &_BACnetConfirmedServiceRequestVTClose{
		ListOfRemoteVtSessionIdentifiers: listOfRemoteVtSessionIdentifiers,
		_BACnetConfirmedServiceRequest:   NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTClose(structType any) BACnetConfirmedServiceRequestVTClose {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTClose); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTClose); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTClose"
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.ListOfRemoteVtSessionIdentifiers) > 0 {
		for _, element := range m.ListOfRemoteVtSessionIdentifiers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestVTCloseParse(ctx context.Context, theBytes []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (BACnetConfirmedServiceRequestVTClose, error) {
	return BACnetConfirmedServiceRequestVTCloseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestPayloadLength, serviceRequestLength)
}

func BACnetConfirmedServiceRequestVTCloseParseWithBufferProducer(serviceRequestPayloadLength uint32, serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestVTClose, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestVTClose, error) {
		return BACnetConfirmedServiceRequestVTCloseParseWithBuffer(ctx, readBuffer, serviceRequestPayloadLength, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestVTCloseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (BACnetConfirmedServiceRequestVTClose, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTClose"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTClose")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	listOfRemoteVtSessionIdentifiers, err := ReadLengthArrayField[BACnetApplicationTagUnsignedInteger](ctx, "listOfRemoteVtSessionIdentifiers", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfRemoteVtSessionIdentifiers' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTClose"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTClose")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestVTClose{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ListOfRemoteVtSessionIdentifiers: listOfRemoteVtSessionIdentifiers,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTClose"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTClose")
		}

		// Array Field (listOfRemoteVtSessionIdentifiers)
		if pushErr := writeBuffer.PushContext("listOfRemoteVtSessionIdentifiers", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfRemoteVtSessionIdentifiers")
		}
		for _curItem, _element := range m.GetListOfRemoteVtSessionIdentifiers() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetListOfRemoteVtSessionIdentifiers()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfRemoteVtSessionIdentifiers' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfRemoteVtSessionIdentifiers", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfRemoteVtSessionIdentifiers")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTClose"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTClose")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestVTClose) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestVTClose) isBACnetConfirmedServiceRequestVTClose() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestVTClose) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
