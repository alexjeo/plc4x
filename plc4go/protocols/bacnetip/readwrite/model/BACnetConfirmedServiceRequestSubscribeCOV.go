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

// BACnetConfirmedServiceRequestSubscribeCOV is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOV
type BACnetConfirmedServiceRequestSubscribeCOV interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetIssueConfirmed returns IssueConfirmed (property field)
	GetIssueConfirmed() BACnetContextTagBoolean
	// GetLifetimeInSeconds returns LifetimeInSeconds (property field)
	GetLifetimeInSeconds() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestSubscribeCOVExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestSubscribeCOV.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestSubscribeCOVExactly interface {
	BACnetConfirmedServiceRequestSubscribeCOV
	isBACnetConfirmedServiceRequestSubscribeCOV() bool
}

// _BACnetConfirmedServiceRequestSubscribeCOV is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOV struct {
	*_BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	MonitoredObjectIdentifier   BACnetContextTagObjectIdentifier
	IssueConfirmed              BACnetContextTagBoolean
	LifetimeInSeconds           BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetIssueConfirmed() BACnetContextTagBoolean {
	return m.IssueConfirmed
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetLifetimeInSeconds() BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOV factory function for _BACnetConfirmedServiceRequestSubscribeCOV
func NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, issueConfirmed BACnetContextTagBoolean, lifetimeInSeconds BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestSubscribeCOV {
	_result := &_BACnetConfirmedServiceRequestSubscribeCOV{
		SubscriberProcessIdentifier:    subscriberProcessIdentifier,
		MonitoredObjectIdentifier:      monitoredObjectIdentifier,
		IssueConfirmed:                 issueConfirmed,
		LifetimeInSeconds:              lifetimeInSeconds,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOV(structType any) BACnetConfirmedServiceRequestSubscribeCOV {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOV); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOV); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOV"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Optional Field (issueConfirmed)
	if m.IssueConfirmed != nil {
		lengthInBits += m.IssueConfirmed.GetLengthInBits(ctx)
	}

	// Optional Field (lifetimeInSeconds)
	if m.LifetimeInSeconds != nil {
		lengthInBits += m.LifetimeInSeconds.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestSubscribeCOV, error) {
	return BACnetConfirmedServiceRequestSubscribeCOVParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestSubscribeCOVParseWithBufferProducer(serviceRequestLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestSubscribeCOV, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestSubscribeCOV, error) {
		return BACnetConfirmedServiceRequestSubscribeCOVParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	}
}

func BACnetConfirmedServiceRequestSubscribeCOVParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestSubscribeCOV, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOV"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOV")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}

	_issueConfirmed, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmed", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmed' field"))
	}
	var issueConfirmed BACnetContextTagBoolean
	if _issueConfirmed != nil {
		issueConfirmed = *_issueConfirmed
	}

	_lifetimeInSeconds, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetimeInSeconds", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetimeInSeconds' field"))
	}
	var lifetimeInSeconds BACnetContextTagUnsignedInteger
	if _lifetimeInSeconds != nil {
		lifetimeInSeconds = *_lifetimeInSeconds
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOV"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOV")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestSubscribeCOV{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		SubscriberProcessIdentifier: subscriberProcessIdentifier,
		MonitoredObjectIdentifier:   monitoredObjectIdentifier,
		IssueConfirmed:              issueConfirmed,
		LifetimeInSeconds:           lifetimeInSeconds,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOV"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOV")
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subscriberProcessIdentifier")
		}
		_subscriberProcessIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetSubscriberProcessIdentifier())
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subscriberProcessIdentifier")
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
		}
		_monitoredObjectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetMonitoredObjectIdentifier())
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Optional Field (issueConfirmed) (Can be skipped, if the value is null)
		var issueConfirmed BACnetContextTagBoolean = nil
		if m.GetIssueConfirmed() != nil {
			if pushErr := writeBuffer.PushContext("issueConfirmed"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for issueConfirmed")
			}
			issueConfirmed = m.GetIssueConfirmed()
			_issueConfirmedErr := writeBuffer.WriteSerializable(ctx, issueConfirmed)
			if popErr := writeBuffer.PopContext("issueConfirmed"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for issueConfirmed")
			}
			if _issueConfirmedErr != nil {
				return errors.Wrap(_issueConfirmedErr, "Error serializing 'issueConfirmed' field")
			}
		}

		// Optional Field (lifetimeInSeconds) (Can be skipped, if the value is null)
		var lifetimeInSeconds BACnetContextTagUnsignedInteger = nil
		if m.GetLifetimeInSeconds() != nil {
			if pushErr := writeBuffer.PushContext("lifetimeInSeconds"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for lifetimeInSeconds")
			}
			lifetimeInSeconds = m.GetLifetimeInSeconds()
			_lifetimeInSecondsErr := writeBuffer.WriteSerializable(ctx, lifetimeInSeconds)
			if popErr := writeBuffer.PopContext("lifetimeInSeconds"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for lifetimeInSeconds")
			}
			if _lifetimeInSecondsErr != nil {
				return errors.Wrap(_lifetimeInSecondsErr, "Error serializing 'lifetimeInSeconds' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOV"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOV")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) isBACnetConfirmedServiceRequestSubscribeCOV() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOV) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
