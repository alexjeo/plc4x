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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type NLMEstablishConnectionToNetwork struct {
	*NLM
	DestinationNetworkAddress uint16
	TerminationTime           uint8
}

// The corresponding interface
type INLMEstablishConnectionToNetwork interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *NLMEstablishConnectionToNetwork) MessageType() uint8 {
	return 0x08
}

func (m *NLMEstablishConnectionToNetwork) InitializeParent(parent *NLM, vendorId *uint16) {
	m.VendorId = vendorId
}

func NewNLMEstablishConnectionToNetwork(destinationNetworkAddress uint16, terminationTime uint8, vendorId *uint16) *NLM {
	child := &NLMEstablishConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		TerminationTime:           terminationTime,
		NLM:                       NewNLM(vendorId),
	}
	child.Child = child
	return child.NLM
}

func CastNLMEstablishConnectionToNetwork(structType interface{}) *NLMEstablishConnectionToNetwork {
	castFunc := func(typ interface{}) *NLMEstablishConnectionToNetwork {
		if casted, ok := typ.(NLMEstablishConnectionToNetwork); ok {
			return &casted
		}
		if casted, ok := typ.(*NLMEstablishConnectionToNetwork); ok {
			return casted
		}
		if casted, ok := typ.(NLM); ok {
			return CastNLMEstablishConnectionToNetwork(casted.Child)
		}
		if casted, ok := typ.(*NLM); ok {
			return CastNLMEstablishConnectionToNetwork(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *NLMEstablishConnectionToNetwork) GetTypeName() string {
	return "NLMEstablishConnectionToNetwork"
}

func (m *NLMEstablishConnectionToNetwork) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *NLMEstablishConnectionToNetwork) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (terminationTime)
	lengthInBits += 8

	return lengthInBits
}

func (m *NLMEstablishConnectionToNetwork) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func NLMEstablishConnectionToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLM, error) {
	if pullErr := readBuffer.PullContext("NLMEstablishConnectionToNetwork"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (terminationTime)
	_terminationTime, _terminationTimeErr := readBuffer.ReadUint8("terminationTime", 8)
	if _terminationTimeErr != nil {
		return nil, errors.Wrap(_terminationTimeErr, "Error parsing 'terminationTime' field")
	}
	terminationTime := _terminationTime

	if closeErr := readBuffer.CloseContext("NLMEstablishConnectionToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NLMEstablishConnectionToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		TerminationTime:           terminationTime,
		NLM:                       &NLM{},
	}
	_child.NLM.Child = _child
	return _child.NLM, nil
}

func (m *NLMEstablishConnectionToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMEstablishConnectionToNetwork"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.DestinationNetworkAddress)
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		// Simple Field (terminationTime)
		terminationTime := uint8(m.TerminationTime)
		_terminationTimeErr := writeBuffer.WriteUint8("terminationTime", 8, (terminationTime))
		if _terminationTimeErr != nil {
			return errors.Wrap(_terminationTimeErr, "Error serializing 'terminationTime' field")
		}

		if popErr := writeBuffer.PopContext("NLMEstablishConnectionToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMEstablishConnectionToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}