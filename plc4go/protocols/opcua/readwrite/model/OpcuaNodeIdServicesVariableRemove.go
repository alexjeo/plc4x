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

// OpcuaNodeIdServicesVariableRemove is an enum
type OpcuaNodeIdServicesVariableRemove int32

type IOpcuaNodeIdServicesVariableRemove interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableRemove_RemoveCertificateMethodType_InputArguments         OpcuaNodeIdServicesVariableRemove = 12521
	OpcuaNodeIdServicesVariableRemove_RemoveConnectionMethodType_InputArguments          OpcuaNodeIdServicesVariableRemove = 14184
	OpcuaNodeIdServicesVariableRemove_RemovePublishedDataSetMethodType_InputArguments    OpcuaNodeIdServicesVariableRemove = 14508
	OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupMethodType_InputArguments       OpcuaNodeIdServicesVariableRemove = 15470
	OpcuaNodeIdServicesVariableRemove_RemoveExtensionFieldMethodType_InputArguments      OpcuaNodeIdServicesVariableRemove = 15500
	OpcuaNodeIdServicesVariableRemove_RemoveIdentityMethodType_InputArguments            OpcuaNodeIdServicesVariableRemove = 15639
	OpcuaNodeIdServicesVariableRemove_RemoveRoleMethodType_InputArguments                OpcuaNodeIdServicesVariableRemove = 16006
	OpcuaNodeIdServicesVariableRemove_RemoveApplicationMethodType_InputArguments         OpcuaNodeIdServicesVariableRemove = 16187
	OpcuaNodeIdServicesVariableRemove_RemoveEndpointMethodType_InputArguments            OpcuaNodeIdServicesVariableRemove = 16191
	OpcuaNodeIdServicesVariableRemove_RemoveDataSetFolderMethodType_InputArguments       OpcuaNodeIdServicesVariableRemove = 17201
	OpcuaNodeIdServicesVariableRemove_RemoveSubscribedDataSetMethodType_InputArguments   OpcuaNodeIdServicesVariableRemove = 23825
	OpcuaNodeIdServicesVariableRemove_RemoveUserMethodType_InputArguments                OpcuaNodeIdServicesVariableRemove = 24287
	OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupFolderMethodType_InputArguments OpcuaNodeIdServicesVariableRemove = 25292
	OpcuaNodeIdServicesVariableRemove_RemovePushTargetMethodType_InputArguments          OpcuaNodeIdServicesVariableRemove = 25380
	OpcuaNodeIdServicesVariableRemove_RemovePushTargetFolderMethodType_InputArguments    OpcuaNodeIdServicesVariableRemove = 25385
)

var OpcuaNodeIdServicesVariableRemoveValues []OpcuaNodeIdServicesVariableRemove

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableRemoveValues = []OpcuaNodeIdServicesVariableRemove{
		OpcuaNodeIdServicesVariableRemove_RemoveCertificateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveConnectionMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemovePublishedDataSetMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveExtensionFieldMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveIdentityMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveRoleMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveApplicationMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveEndpointMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveDataSetFolderMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveSubscribedDataSetMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveUserMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupFolderMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemovePushTargetMethodType_InputArguments,
		OpcuaNodeIdServicesVariableRemove_RemovePushTargetFolderMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableRemoveByValue(value int32) (enum OpcuaNodeIdServicesVariableRemove, ok bool) {
	switch value {
	case 12521:
		return OpcuaNodeIdServicesVariableRemove_RemoveCertificateMethodType_InputArguments, true
	case 14184:
		return OpcuaNodeIdServicesVariableRemove_RemoveConnectionMethodType_InputArguments, true
	case 14508:
		return OpcuaNodeIdServicesVariableRemove_RemovePublishedDataSetMethodType_InputArguments, true
	case 15470:
		return OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupMethodType_InputArguments, true
	case 15500:
		return OpcuaNodeIdServicesVariableRemove_RemoveExtensionFieldMethodType_InputArguments, true
	case 15639:
		return OpcuaNodeIdServicesVariableRemove_RemoveIdentityMethodType_InputArguments, true
	case 16006:
		return OpcuaNodeIdServicesVariableRemove_RemoveRoleMethodType_InputArguments, true
	case 16187:
		return OpcuaNodeIdServicesVariableRemove_RemoveApplicationMethodType_InputArguments, true
	case 16191:
		return OpcuaNodeIdServicesVariableRemove_RemoveEndpointMethodType_InputArguments, true
	case 17201:
		return OpcuaNodeIdServicesVariableRemove_RemoveDataSetFolderMethodType_InputArguments, true
	case 23825:
		return OpcuaNodeIdServicesVariableRemove_RemoveSubscribedDataSetMethodType_InputArguments, true
	case 24287:
		return OpcuaNodeIdServicesVariableRemove_RemoveUserMethodType_InputArguments, true
	case 25292:
		return OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupFolderMethodType_InputArguments, true
	case 25380:
		return OpcuaNodeIdServicesVariableRemove_RemovePushTargetMethodType_InputArguments, true
	case 25385:
		return OpcuaNodeIdServicesVariableRemove_RemovePushTargetFolderMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRemoveByName(value string) (enum OpcuaNodeIdServicesVariableRemove, ok bool) {
	switch value {
	case "RemoveCertificateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveCertificateMethodType_InputArguments, true
	case "RemoveConnectionMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveConnectionMethodType_InputArguments, true
	case "RemovePublishedDataSetMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemovePublishedDataSetMethodType_InputArguments, true
	case "RemoveSecurityGroupMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupMethodType_InputArguments, true
	case "RemoveExtensionFieldMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveExtensionFieldMethodType_InputArguments, true
	case "RemoveIdentityMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveIdentityMethodType_InputArguments, true
	case "RemoveRoleMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveRoleMethodType_InputArguments, true
	case "RemoveApplicationMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveApplicationMethodType_InputArguments, true
	case "RemoveEndpointMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveEndpointMethodType_InputArguments, true
	case "RemoveDataSetFolderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveDataSetFolderMethodType_InputArguments, true
	case "RemoveSubscribedDataSetMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveSubscribedDataSetMethodType_InputArguments, true
	case "RemoveUserMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveUserMethodType_InputArguments, true
	case "RemoveSecurityGroupFolderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupFolderMethodType_InputArguments, true
	case "RemovePushTargetMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemovePushTargetMethodType_InputArguments, true
	case "RemovePushTargetFolderMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableRemove_RemovePushTargetFolderMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableRemoveKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableRemoveValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableRemove(structType any) OpcuaNodeIdServicesVariableRemove {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableRemove {
		if sOpcuaNodeIdServicesVariableRemove, ok := typ.(OpcuaNodeIdServicesVariableRemove); ok {
			return sOpcuaNodeIdServicesVariableRemove
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableRemove) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableRemove) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableRemoveParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableRemove, error) {
	return OpcuaNodeIdServicesVariableRemoveParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableRemoveParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableRemove, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableRemove", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableRemove")
	}
	if enum, ok := OpcuaNodeIdServicesVariableRemoveByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableRemove")
		return OpcuaNodeIdServicesVariableRemove(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableRemove) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableRemove) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableRemove", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableRemove) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableRemove) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableRemove_RemoveCertificateMethodType_InputArguments:
		return "RemoveCertificateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveConnectionMethodType_InputArguments:
		return "RemoveConnectionMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemovePublishedDataSetMethodType_InputArguments:
		return "RemovePublishedDataSetMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupMethodType_InputArguments:
		return "RemoveSecurityGroupMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveExtensionFieldMethodType_InputArguments:
		return "RemoveExtensionFieldMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveIdentityMethodType_InputArguments:
		return "RemoveIdentityMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveRoleMethodType_InputArguments:
		return "RemoveRoleMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveApplicationMethodType_InputArguments:
		return "RemoveApplicationMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveEndpointMethodType_InputArguments:
		return "RemoveEndpointMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveDataSetFolderMethodType_InputArguments:
		return "RemoveDataSetFolderMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveSubscribedDataSetMethodType_InputArguments:
		return "RemoveSubscribedDataSetMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveUserMethodType_InputArguments:
		return "RemoveUserMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemoveSecurityGroupFolderMethodType_InputArguments:
		return "RemoveSecurityGroupFolderMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemovePushTargetMethodType_InputArguments:
		return "RemovePushTargetMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableRemove_RemovePushTargetFolderMethodType_InputArguments:
		return "RemovePushTargetFolderMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableRemove) String() string {
	return e.PLC4XEnumName()
}
