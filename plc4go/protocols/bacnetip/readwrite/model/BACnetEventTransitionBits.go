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

// BACnetEventTransitionBits is an enum
type BACnetEventTransitionBits uint8

type IBACnetEventTransitionBits interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetEventTransitionBits_TO_OFFNORMAL BACnetEventTransitionBits = 0
	BACnetEventTransitionBits_TO_FAULT     BACnetEventTransitionBits = 1
	BACnetEventTransitionBits_TO_NORMAL    BACnetEventTransitionBits = 2
)

var BACnetEventTransitionBitsValues []BACnetEventTransitionBits

func init() {
	_ = errors.New
	BACnetEventTransitionBitsValues = []BACnetEventTransitionBits{
		BACnetEventTransitionBits_TO_OFFNORMAL,
		BACnetEventTransitionBits_TO_FAULT,
		BACnetEventTransitionBits_TO_NORMAL,
	}
}

func BACnetEventTransitionBitsByValue(value uint8) BACnetEventTransitionBits {
	switch value {
	case 0:
		return BACnetEventTransitionBits_TO_OFFNORMAL
	case 1:
		return BACnetEventTransitionBits_TO_FAULT
	case 2:
		return BACnetEventTransitionBits_TO_NORMAL
	}
	return 0
}

func BACnetEventTransitionBitsByName(value string) BACnetEventTransitionBits {
	switch value {
	case "TO_OFFNORMAL":
		return BACnetEventTransitionBits_TO_OFFNORMAL
	case "TO_FAULT":
		return BACnetEventTransitionBits_TO_FAULT
	case "TO_NORMAL":
		return BACnetEventTransitionBits_TO_NORMAL
	}
	return 0
}

func BACnetEventTransitionBitsKnows(value uint8) bool {
	for _, typeValue := range BACnetEventTransitionBitsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEventTransitionBits(structType interface{}) BACnetEventTransitionBits {
	castFunc := func(typ interface{}) BACnetEventTransitionBits {
		if sBACnetEventTransitionBits, ok := typ.(BACnetEventTransitionBits); ok {
			return sBACnetEventTransitionBits
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEventTransitionBits) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetEventTransitionBits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventTransitionBitsParse(readBuffer utils.ReadBuffer) (BACnetEventTransitionBits, error) {
	val, err := readBuffer.ReadUint8("BACnetEventTransitionBits", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetEventTransitionBitsByValue(val), nil
}

func (e BACnetEventTransitionBits) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetEventTransitionBits", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetEventTransitionBits) name() string {
	switch e {
	case BACnetEventTransitionBits_TO_OFFNORMAL:
		return "TO_OFFNORMAL"
	case BACnetEventTransitionBits_TO_FAULT:
		return "TO_FAULT"
	case BACnetEventTransitionBits_TO_NORMAL:
		return "TO_NORMAL"
	}
	return ""
}

func (e BACnetEventTransitionBits) String() string {
	return e.name()
}
