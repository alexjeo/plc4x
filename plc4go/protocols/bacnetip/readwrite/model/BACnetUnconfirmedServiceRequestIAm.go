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

// BACnetUnconfirmedServiceRequestIAm is the corresponding interface of BACnetUnconfirmedServiceRequestIAm
type BACnetUnconfirmedServiceRequestIAm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier
	// GetMaximumApduLengthAcceptedLength returns MaximumApduLengthAcceptedLength (property field)
	GetMaximumApduLengthAcceptedLength() BACnetApplicationTagUnsignedInteger
	// GetSegmentationSupported returns SegmentationSupported (property field)
	GetSegmentationSupported() BACnetSegmentationTagged
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
}

// BACnetUnconfirmedServiceRequestIAmExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestIAm.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestIAmExactly interface {
	BACnetUnconfirmedServiceRequestIAm
	isBACnetUnconfirmedServiceRequestIAm() bool
}

// _BACnetUnconfirmedServiceRequestIAm is the data-structure of this message
type _BACnetUnconfirmedServiceRequestIAm struct {
	*_BACnetUnconfirmedServiceRequest
	DeviceIdentifier                BACnetApplicationTagObjectIdentifier
	MaximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger
	SegmentationSupported           BACnetSegmentationTagged
	VendorId                        BACnetVendorIdTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIAm) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_I_AM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestIAm) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIAm) GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetMaximumApduLengthAcceptedLength() BACnetApplicationTagUnsignedInteger {
	return m.MaximumApduLengthAcceptedLength
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetSegmentationSupported() BACnetSegmentationTagged {
	return m.SegmentationSupported
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestIAm factory function for _BACnetUnconfirmedServiceRequestIAm
func NewBACnetUnconfirmedServiceRequestIAm(deviceIdentifier BACnetApplicationTagObjectIdentifier, maximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger, segmentationSupported BACnetSegmentationTagged, vendorId BACnetVendorIdTagged, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestIAm {
	_result := &_BACnetUnconfirmedServiceRequestIAm{
		DeviceIdentifier:                 deviceIdentifier,
		MaximumApduLengthAcceptedLength:  maximumApduLengthAcceptedLength,
		SegmentationSupported:            segmentationSupported,
		VendorId:                         vendorId,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestIAm(structType any) BACnetUnconfirmedServiceRequestIAm {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestIAm); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestIAm); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIAm"
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (maximumApduLengthAcceptedLength)
	lengthInBits += m.MaximumApduLengthAcceptedLength.GetLengthInBits(ctx)

	// Simple field (segmentationSupported)
	lengthInBits += m.SegmentationSupported.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestIAmParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestIAm, error) {
	return BACnetUnconfirmedServiceRequestIAmParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestIAmParseWithBufferProducer(serviceRequestLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestIAm, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceRequestIAm, error) {
		return BACnetUnconfirmedServiceRequestIAmParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetUnconfirmedServiceRequestIAmParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestIAm, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIAm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestIAm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}

	maximumApduLengthAcceptedLength, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maximumApduLengthAcceptedLength", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maximumApduLengthAcceptedLength' field"))
	}

	segmentationSupported, err := ReadSimpleField[BACnetSegmentationTagged](ctx, "segmentationSupported", ReadComplex[BACnetSegmentationTagged](BACnetSegmentationTaggedParseWithBufferProducer((uint8)(uint8(9)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentationSupported' field"))
	}

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIAm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestIAm")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestIAm{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		DeviceIdentifier:                deviceIdentifier,
		MaximumApduLengthAcceptedLength: maximumApduLengthAcceptedLength,
		SegmentationSupported:           segmentationSupported,
		VendorId:                        vendorId,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestIAm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestIAm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIAm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestIAm")
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		_deviceIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetDeviceIdentifier())
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}

		// Simple Field (maximumApduLengthAcceptedLength)
		if pushErr := writeBuffer.PushContext("maximumApduLengthAcceptedLength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maximumApduLengthAcceptedLength")
		}
		_maximumApduLengthAcceptedLengthErr := writeBuffer.WriteSerializable(ctx, m.GetMaximumApduLengthAcceptedLength())
		if popErr := writeBuffer.PopContext("maximumApduLengthAcceptedLength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maximumApduLengthAcceptedLength")
		}
		if _maximumApduLengthAcceptedLengthErr != nil {
			return errors.Wrap(_maximumApduLengthAcceptedLengthErr, "Error serializing 'maximumApduLengthAcceptedLength' field")
		}

		// Simple Field (segmentationSupported)
		if pushErr := writeBuffer.PushContext("segmentationSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for segmentationSupported")
		}
		_segmentationSupportedErr := writeBuffer.WriteSerializable(ctx, m.GetSegmentationSupported())
		if popErr := writeBuffer.PopContext("segmentationSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for segmentationSupported")
		}
		if _segmentationSupportedErr != nil {
			return errors.Wrap(_segmentationSupportedErr, "Error serializing 'segmentationSupported' field")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(ctx, m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIAm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestIAm")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestIAm) isBACnetUnconfirmedServiceRequestIAm() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestIAm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
