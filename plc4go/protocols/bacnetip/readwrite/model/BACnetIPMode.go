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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetIPMode is an enum
type BACnetIPMode uint8

type IBACnetIPMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetIPMode_NORMAL  BACnetIPMode = 0
	BACnetIPMode_FOREIGN BACnetIPMode = 1
	BACnetIPMode_BBMD    BACnetIPMode = 2
)

var BACnetIPModeValues []BACnetIPMode

func init() {
	_ = errors.New
	BACnetIPModeValues = []BACnetIPMode{
		BACnetIPMode_NORMAL,
		BACnetIPMode_FOREIGN,
		BACnetIPMode_BBMD,
	}
}

func BACnetIPModeByValue(value uint8) (enum BACnetIPMode, ok bool) {
	switch value {
	case 0:
		return BACnetIPMode_NORMAL, true
	case 1:
		return BACnetIPMode_FOREIGN, true
	case 2:
		return BACnetIPMode_BBMD, true
	}
	return 0, false
}

func BACnetIPModeByName(value string) (enum BACnetIPMode, ok bool) {
	switch value {
	case "NORMAL":
		return BACnetIPMode_NORMAL, true
	case "FOREIGN":
		return BACnetIPMode_FOREIGN, true
	case "BBMD":
		return BACnetIPMode_BBMD, true
	}
	return 0, false
}

func BACnetIPModeKnows(value uint8) bool {
	for _, typeValue := range BACnetIPModeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetIPMode(structType any) BACnetIPMode {
	castFunc := func(typ any) BACnetIPMode {
		if sBACnetIPMode, ok := typ.(BACnetIPMode); ok {
			return sBACnetIPMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetIPMode) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetIPMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetIPModeParse(ctx context.Context, theBytes []byte) (BACnetIPMode, error) {
	return BACnetIPModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetIPModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetIPMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("BACnetIPMode", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetIPMode")
	}
	if enum, ok := BACnetIPModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetIPMode")
		return BACnetIPMode(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetIPMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetIPMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("BACnetIPMode", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetIPMode) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetIPMode) PLC4XEnumName() string {
	switch e {
	case BACnetIPMode_NORMAL:
		return "NORMAL"
	case BACnetIPMode_FOREIGN:
		return "FOREIGN"
	case BACnetIPMode_BBMD:
		return "BBMD"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetIPMode) String() string {
	return e.PLC4XEnumName()
}
