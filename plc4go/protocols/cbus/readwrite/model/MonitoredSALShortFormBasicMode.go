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

// MonitoredSALShortFormBasicMode is the corresponding interface of MonitoredSALShortFormBasicMode
type MonitoredSALShortFormBasicMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MonitoredSAL
	// GetCounts returns Counts (property field)
	GetCounts() byte
	// GetBridgeCount returns BridgeCount (property field)
	GetBridgeCount() *uint8
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() *uint8
	// GetNoCounts returns NoCounts (property field)
	GetNoCounts() *byte
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
}

// MonitoredSALShortFormBasicModeExactly can be used when we want exactly this type and not a type which fulfills MonitoredSALShortFormBasicMode.
// This is useful for switch cases.
type MonitoredSALShortFormBasicModeExactly interface {
	MonitoredSALShortFormBasicMode
	isMonitoredSALShortFormBasicMode() bool
}

// _MonitoredSALShortFormBasicMode is the data-structure of this message
type _MonitoredSALShortFormBasicMode struct {
	*_MonitoredSAL
	Counts        byte
	BridgeCount   *uint8
	NetworkNumber *uint8
	NoCounts      *byte
	Application   ApplicationIdContainer
	SalData       SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALShortFormBasicMode) InitializeParent(parent MonitoredSAL, salType byte) {
	m.SalType = salType
}

func (m *_MonitoredSALShortFormBasicMode) GetParent() MonitoredSAL {
	return m._MonitoredSAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALShortFormBasicMode) GetCounts() byte {
	return m.Counts
}

func (m *_MonitoredSALShortFormBasicMode) GetBridgeCount() *uint8 {
	return m.BridgeCount
}

func (m *_MonitoredSALShortFormBasicMode) GetNetworkNumber() *uint8 {
	return m.NetworkNumber
}

func (m *_MonitoredSALShortFormBasicMode) GetNoCounts() *byte {
	return m.NoCounts
}

func (m *_MonitoredSALShortFormBasicMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_MonitoredSALShortFormBasicMode) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredSALShortFormBasicMode factory function for _MonitoredSALShortFormBasicMode
func NewMonitoredSALShortFormBasicMode(counts byte, bridgeCount *uint8, networkNumber *uint8, noCounts *byte, application ApplicationIdContainer, salData SALData, salType byte, cBusOptions CBusOptions) *_MonitoredSALShortFormBasicMode {
	_result := &_MonitoredSALShortFormBasicMode{
		Counts:        counts,
		BridgeCount:   bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts:      noCounts,
		Application:   application,
		SalData:       salData,
		_MonitoredSAL: NewMonitoredSAL(salType, cBusOptions),
	}
	_result._MonitoredSAL._MonitoredSALChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredSALShortFormBasicMode(structType any) MonitoredSALShortFormBasicMode {
	if casted, ok := structType.(MonitoredSALShortFormBasicMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALShortFormBasicMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALShortFormBasicMode) GetTypeName() string {
	return "MonitoredSALShortFormBasicMode"
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (bridgeCount)
	if m.BridgeCount != nil {
		lengthInBits += 8
	}

	// Optional Field (networkNumber)
	if m.NetworkNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (noCounts)
	if m.NoCounts != nil {
		lengthInBits += 8
	}

	// Simple field (application)
	lengthInBits += 8

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredSALShortFormBasicModeParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (MonitoredSALShortFormBasicMode, error) {
	return MonitoredSALShortFormBasicModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func MonitoredSALShortFormBasicModeParseWithBufferProducer(cBusOptions CBusOptions) func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredSALShortFormBasicMode, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (MonitoredSALShortFormBasicMode, error) {
		return MonitoredSALShortFormBasicModeParseWithBuffer(ctx, readBuffer, cBusOptions)
	}
}

func MonitoredSALShortFormBasicModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (MonitoredSALShortFormBasicMode, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredSALShortFormBasicMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALShortFormBasicMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	counts, err := ReadPeekField[byte](ctx, "counts", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'counts' field"))
	}

	bridgeCount, err := ReadOptionalField[uint8](ctx, "bridgeCount", ReadUnsignedByte(readBuffer, uint8(8)), bool((counts) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeCount' field"))
	}

	networkNumber, err := ReadOptionalField[uint8](ctx, "networkNumber", ReadUnsignedByte(readBuffer, uint8(8)), bool((counts) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumber' field"))
	}

	noCounts, err := ReadOptionalField[byte](ctx, "noCounts", ReadByte(readBuffer, 8), bool((counts) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noCounts' field"))
	}

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}

	_salData, err := ReadOptionalField[SALData](ctx, "salData", ReadComplex[SALData](SALDataParseWithBufferProducer[SALData]((ApplicationId)(application.ApplicationId())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'salData' field"))
	}
	var salData SALData
	if _salData != nil {
		salData = *_salData
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALShortFormBasicMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALShortFormBasicMode")
	}

	// Create a partially initialized instance
	_child := &_MonitoredSALShortFormBasicMode{
		_MonitoredSAL: &_MonitoredSAL{
			CBusOptions: cBusOptions,
		},
		Counts:        counts,
		BridgeCount:   bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts:      noCounts,
		Application:   application,
		SalData:       salData,
	}
	_child._MonitoredSAL._MonitoredSALChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredSALShortFormBasicMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALShortFormBasicMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALShortFormBasicMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALShortFormBasicMode")
		}

		// Optional Field (bridgeCount) (Can be skipped, if the value is null)
		var bridgeCount *uint8 = nil
		if m.GetBridgeCount() != nil {
			bridgeCount = m.GetBridgeCount()
			_bridgeCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("bridgeCount", 8, uint8(*(bridgeCount)))
			if _bridgeCountErr != nil {
				return errors.Wrap(_bridgeCountErr, "Error serializing 'bridgeCount' field")
			}
		}

		// Optional Field (networkNumber) (Can be skipped, if the value is null)
		var networkNumber *uint8 = nil
		if m.GetNetworkNumber() != nil {
			networkNumber = m.GetNetworkNumber()
			_networkNumberErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("networkNumber", 8, uint8(*(networkNumber)))
			if _networkNumberErr != nil {
				return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
			}
		}

		// Optional Field (noCounts) (Can be skipped, if the value is null)
		var noCounts *byte = nil
		if m.GetNoCounts() != nil {
			noCounts = m.GetNoCounts()
			_noCountsErr := /*TODO: migrate me*/ writeBuffer.WriteByte("noCounts", *(noCounts))
			if _noCountsErr != nil {
				return errors.Wrap(_noCountsErr, "Error serializing 'noCounts' field")
			}
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Optional Field (salData) (Can be skipped, if the value is null)
		var salData SALData = nil
		if m.GetSalData() != nil {
			if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for salData")
			}
			salData = m.GetSalData()
			_salDataErr := writeBuffer.WriteSerializable(ctx, salData)
			if popErr := writeBuffer.PopContext("salData"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for salData")
			}
			if _salDataErr != nil {
				return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
			}
		}

		if popErr := writeBuffer.PopContext("MonitoredSALShortFormBasicMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALShortFormBasicMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALShortFormBasicMode) isMonitoredSALShortFormBasicMode() bool {
	return true
}

func (m *_MonitoredSALShortFormBasicMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
