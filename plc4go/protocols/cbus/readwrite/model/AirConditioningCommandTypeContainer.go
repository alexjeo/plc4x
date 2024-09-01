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

// AirConditioningCommandTypeContainer is an enum
type AirConditioningCommandTypeContainer uint8

type IAirConditioningCommandTypeContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumBytes() uint8
	CommandType() AirConditioningCommandType
}

const (
	AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOff            AirConditioningCommandTypeContainer = 0x01
	AirConditioningCommandTypeContainer_AirConditioningCommandZoneHvacPlantStatus        AirConditioningCommandTypeContainer = 0x05
	AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidityPlantStatus    AirConditioningCommandTypeContainer = 0x0D
	AirConditioningCommandTypeContainer_AirConditioningCommandZoneTemperature            AirConditioningCommandTypeContainer = 0x15
	AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidity               AirConditioningCommandTypeContainer = 0x1D
	AirConditioningCommandTypeContainer_AirConditioningCommandRefresh                    AirConditioningCommandTypeContainer = 0x21
	AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHvacMode            AirConditioningCommandTypeContainer = 0x2F
	AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHvacLevel          AirConditioningCommandTypeContainer = 0x36
	AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHumidityMode        AirConditioningCommandTypeContainer = 0x47
	AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHumidityLevel      AirConditioningCommandTypeContainer = 0x4E
	AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacUpperGuardLimit     AirConditioningCommandTypeContainer = 0x55
	AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacLowerGuardLimit     AirConditioningCommandTypeContainer = 0x5D
	AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacSetbackLimit        AirConditioningCommandTypeContainer = 0x65
	AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityUpperGuardLimit AirConditioningCommandTypeContainer = 0x6D
	AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityLowerGuardLimit AirConditioningCommandTypeContainer = 0x75
	AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOn             AirConditioningCommandTypeContainer = 0x79
	AirConditioningCommandTypeContainer_AirConditioningCommandSetHumiditySetbackLimit    AirConditioningCommandTypeContainer = 0x7D
	AirConditioningCommandTypeContainer_AirConditioningCommandHvacScheduleEntry          AirConditioningCommandTypeContainer = 0x89
	AirConditioningCommandTypeContainer_AirConditioningCommandHumidityScheduleEntry      AirConditioningCommandTypeContainer = 0xA9
)

var AirConditioningCommandTypeContainerValues []AirConditioningCommandTypeContainer

func init() {
	_ = errors.New
	AirConditioningCommandTypeContainerValues = []AirConditioningCommandTypeContainer{
		AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOff,
		AirConditioningCommandTypeContainer_AirConditioningCommandZoneHvacPlantStatus,
		AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidityPlantStatus,
		AirConditioningCommandTypeContainer_AirConditioningCommandZoneTemperature,
		AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidity,
		AirConditioningCommandTypeContainer_AirConditioningCommandRefresh,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHvacMode,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHvacLevel,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHumidityMode,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHumidityLevel,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacUpperGuardLimit,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacLowerGuardLimit,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacSetbackLimit,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityUpperGuardLimit,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityLowerGuardLimit,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOn,
		AirConditioningCommandTypeContainer_AirConditioningCommandSetHumiditySetbackLimit,
		AirConditioningCommandTypeContainer_AirConditioningCommandHvacScheduleEntry,
		AirConditioningCommandTypeContainer_AirConditioningCommandHumidityScheduleEntry,
	}
}

func (e AirConditioningCommandTypeContainer) NumBytes() uint8 {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x05:
		{ /* '0x05' */
			return 5
		}
	case 0x0D:
		{ /* '0x0D' */
			return 5
		}
	case 0x15:
		{ /* '0x15' */
			return 5
		}
	case 0x1D:
		{ /* '0x1D' */
			return 5
		}
	case 0x21:
		{ /* '0x21' */
			return 1
		}
	case 0x2F:
		{ /* '0x2F' */
			return 7
		}
	case 0x36:
		{ /* '0x36' */
			return 6
		}
	case 0x47:
		{ /* '0x47' */
			return 7
		}
	case 0x4E:
		{ /* '0x4E' */
			return 6
		}
	case 0x55:
		{ /* '0x55' */
			return 5
		}
	case 0x5D:
		{ /* '0x5D' */
			return 5
		}
	case 0x65:
		{ /* '0x65' */
			return 5
		}
	case 0x6D:
		{ /* '0x6D' */
			return 5
		}
	case 0x75:
		{ /* '0x75' */
			return 5
		}
	case 0x79:
		{ /* '0x79' */
			return 1
		}
	case 0x7D:
		{ /* '0x7D' */
			return 5
		}
	case 0x89:
		{ /* '0x89' */
			return 9
		}
	case 0xA9:
		{ /* '0xA9' */
			return 9
		}
	default:
		{
			return 0
		}
	}
}

func AirConditioningCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (enum AirConditioningCommandTypeContainer, ok bool) {
	for _, sizeValue := range AirConditioningCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, true
		}
	}
	return 0, false
}

func (e AirConditioningCommandTypeContainer) CommandType() AirConditioningCommandType {
	switch e {
	case 0x01:
		{ /* '0x01' */
			return AirConditioningCommandType_SET_ZONE_GROUP_OFF
		}
	case 0x05:
		{ /* '0x05' */
			return AirConditioningCommandType_ZONE_HVAC_PLANT_STATUS
		}
	case 0x0D:
		{ /* '0x0D' */
			return AirConditioningCommandType_ZONE_HUMIDITY_PLANT_STATUS
		}
	case 0x15:
		{ /* '0x15' */
			return AirConditioningCommandType_ZONE_TEMPERATURE
		}
	case 0x1D:
		{ /* '0x1D' */
			return AirConditioningCommandType_ZONE_HUMIDITY
		}
	case 0x21:
		{ /* '0x21' */
			return AirConditioningCommandType_REFRESH
		}
	case 0x2F:
		{ /* '0x2F' */
			return AirConditioningCommandType_SET_ZONE_HVAC_MODE
		}
	case 0x36:
		{ /* '0x36' */
			return AirConditioningCommandType_SET_PLANT_HVAC_LEVEL
		}
	case 0x47:
		{ /* '0x47' */
			return AirConditioningCommandType_SET_ZONE_HUMIDITY_MODE
		}
	case 0x4E:
		{ /* '0x4E' */
			return AirConditioningCommandType_SET_PLANT_HUMIDITY_LEVEL
		}
	case 0x55:
		{ /* '0x55' */
			return AirConditioningCommandType_SET_HVAC_UPPER_GUARD_LIMIT
		}
	case 0x5D:
		{ /* '0x5D' */
			return AirConditioningCommandType_SET_HVAC_LOWER_GUARD_LIMIT
		}
	case 0x65:
		{ /* '0x65' */
			return AirConditioningCommandType_SET_HVAC_SETBACK_LIMIT
		}
	case 0x6D:
		{ /* '0x6D' */
			return AirConditioningCommandType_SET_HUMIDITY_UPPER_GUARD_LIMIT
		}
	case 0x75:
		{ /* '0x75' */
			return AirConditioningCommandType_SET_HUMIDITY_LOWER_GUARD_LIMIT
		}
	case 0x79:
		{ /* '0x79' */
			return AirConditioningCommandType_SET_ZONE_GROUP_ON
		}
	case 0x7D:
		{ /* '0x7D' */
			return AirConditioningCommandType_SET_HUMIDITY_SETBACK_LIMIT
		}
	case 0x89:
		{ /* '0x89' */
			return AirConditioningCommandType_HVAC_SCHEDULE_ENTRY
		}
	case 0xA9:
		{ /* '0xA9' */
			return AirConditioningCommandType_HUMIDITY_SCHEDULE_ENTRY
		}
	default:
		{
			return 0
		}
	}
}

func AirConditioningCommandTypeContainerFirstEnumForFieldCommandType(value AirConditioningCommandType) (enum AirConditioningCommandTypeContainer, ok bool) {
	for _, sizeValue := range AirConditioningCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, true
		}
	}
	return 0, false
}
func AirConditioningCommandTypeContainerByValue(value uint8) (enum AirConditioningCommandTypeContainer, ok bool) {
	switch value {
	case 0x01:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOff, true
	case 0x05:
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneHvacPlantStatus, true
	case 0x0D:
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidityPlantStatus, true
	case 0x15:
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneTemperature, true
	case 0x1D:
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidity, true
	case 0x21:
		return AirConditioningCommandTypeContainer_AirConditioningCommandRefresh, true
	case 0x2F:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHvacMode, true
	case 0x36:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHvacLevel, true
	case 0x47:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHumidityMode, true
	case 0x4E:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHumidityLevel, true
	case 0x55:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacUpperGuardLimit, true
	case 0x5D:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacLowerGuardLimit, true
	case 0x65:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacSetbackLimit, true
	case 0x6D:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityUpperGuardLimit, true
	case 0x75:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityLowerGuardLimit, true
	case 0x79:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOn, true
	case 0x7D:
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHumiditySetbackLimit, true
	case 0x89:
		return AirConditioningCommandTypeContainer_AirConditioningCommandHvacScheduleEntry, true
	case 0xA9:
		return AirConditioningCommandTypeContainer_AirConditioningCommandHumidityScheduleEntry, true
	}
	return 0, false
}

func AirConditioningCommandTypeContainerByName(value string) (enum AirConditioningCommandTypeContainer, ok bool) {
	switch value {
	case "AirConditioningCommandSetZoneGroupOff":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOff, true
	case "AirConditioningCommandZoneHvacPlantStatus":
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneHvacPlantStatus, true
	case "AirConditioningCommandZoneHumidityPlantStatus":
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidityPlantStatus, true
	case "AirConditioningCommandZoneTemperature":
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneTemperature, true
	case "AirConditioningCommandZoneHumidity":
		return AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidity, true
	case "AirConditioningCommandRefresh":
		return AirConditioningCommandTypeContainer_AirConditioningCommandRefresh, true
	case "AirConditioningCommandSetZoneHvacMode":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHvacMode, true
	case "AirConditioningCommandSetPlantHvacLevel":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHvacLevel, true
	case "AirConditioningCommandSetZoneHumidityMode":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHumidityMode, true
	case "AirConditioningCommandSetPlantHumidityLevel":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHumidityLevel, true
	case "AirConditioningCommandSetHvacUpperGuardLimit":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacUpperGuardLimit, true
	case "AirConditioningCommandSetHvacLowerGuardLimit":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacLowerGuardLimit, true
	case "AirConditioningCommandSetHvacSetbackLimit":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacSetbackLimit, true
	case "AirConditioningCommandSetHumidityUpperGuardLimit":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityUpperGuardLimit, true
	case "AirConditioningCommandSetHumidityLowerGuardLimit":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityLowerGuardLimit, true
	case "AirConditioningCommandSetZoneGroupOn":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOn, true
	case "AirConditioningCommandSetHumiditySetbackLimit":
		return AirConditioningCommandTypeContainer_AirConditioningCommandSetHumiditySetbackLimit, true
	case "AirConditioningCommandHvacScheduleEntry":
		return AirConditioningCommandTypeContainer_AirConditioningCommandHvacScheduleEntry, true
	case "AirConditioningCommandHumidityScheduleEntry":
		return AirConditioningCommandTypeContainer_AirConditioningCommandHumidityScheduleEntry, true
	}
	return 0, false
}

func AirConditioningCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range AirConditioningCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAirConditioningCommandTypeContainer(structType any) AirConditioningCommandTypeContainer {
	castFunc := func(typ any) AirConditioningCommandTypeContainer {
		if sAirConditioningCommandTypeContainer, ok := typ.(AirConditioningCommandTypeContainer); ok {
			return sAirConditioningCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m AirConditioningCommandTypeContainer) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m AirConditioningCommandTypeContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningCommandTypeContainerParse(ctx context.Context, theBytes []byte) (AirConditioningCommandTypeContainer, error) {
	return AirConditioningCommandTypeContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningCommandTypeContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningCommandTypeContainer, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("AirConditioningCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AirConditioningCommandTypeContainer")
	}
	if enum, ok := AirConditioningCommandTypeContainerByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AirConditioningCommandTypeContainer")
		return AirConditioningCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e AirConditioningCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AirConditioningCommandTypeContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("AirConditioningCommandTypeContainer", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e AirConditioningCommandTypeContainer) GetValue() uint8 {
	return uint8(e)
}

func (e AirConditioningCommandTypeContainer) AirConditioningCommandTypeContainerGetNumBytes() uint8 {
	return e.NumBytes()
}
func (e AirConditioningCommandTypeContainer) AirConditioningCommandTypeContainerGetCommandType() AirConditioningCommandType {
	return e.CommandType()
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AirConditioningCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOff:
		return "AirConditioningCommandSetZoneGroupOff"
	case AirConditioningCommandTypeContainer_AirConditioningCommandZoneHvacPlantStatus:
		return "AirConditioningCommandZoneHvacPlantStatus"
	case AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidityPlantStatus:
		return "AirConditioningCommandZoneHumidityPlantStatus"
	case AirConditioningCommandTypeContainer_AirConditioningCommandZoneTemperature:
		return "AirConditioningCommandZoneTemperature"
	case AirConditioningCommandTypeContainer_AirConditioningCommandZoneHumidity:
		return "AirConditioningCommandZoneHumidity"
	case AirConditioningCommandTypeContainer_AirConditioningCommandRefresh:
		return "AirConditioningCommandRefresh"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHvacMode:
		return "AirConditioningCommandSetZoneHvacMode"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHvacLevel:
		return "AirConditioningCommandSetPlantHvacLevel"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneHumidityMode:
		return "AirConditioningCommandSetZoneHumidityMode"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetPlantHumidityLevel:
		return "AirConditioningCommandSetPlantHumidityLevel"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacUpperGuardLimit:
		return "AirConditioningCommandSetHvacUpperGuardLimit"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacLowerGuardLimit:
		return "AirConditioningCommandSetHvacLowerGuardLimit"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetHvacSetbackLimit:
		return "AirConditioningCommandSetHvacSetbackLimit"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityUpperGuardLimit:
		return "AirConditioningCommandSetHumidityUpperGuardLimit"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetHumidityLowerGuardLimit:
		return "AirConditioningCommandSetHumidityLowerGuardLimit"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetZoneGroupOn:
		return "AirConditioningCommandSetZoneGroupOn"
	case AirConditioningCommandTypeContainer_AirConditioningCommandSetHumiditySetbackLimit:
		return "AirConditioningCommandSetHumiditySetbackLimit"
	case AirConditioningCommandTypeContainer_AirConditioningCommandHvacScheduleEntry:
		return "AirConditioningCommandHvacScheduleEntry"
	case AirConditioningCommandTypeContainer_AirConditioningCommandHumidityScheduleEntry:
		return "AirConditioningCommandHumidityScheduleEntry"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e AirConditioningCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
