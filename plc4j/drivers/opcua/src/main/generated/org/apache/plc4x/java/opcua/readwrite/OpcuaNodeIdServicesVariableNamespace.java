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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableNamespace {
  NamespaceMetadataType_NamespaceUri((int) 11617L),
  NamespaceMetadataType_NamespaceVersion((int) 11618L),
  NamespaceMetadataType_NamespacePublicationDate((int) 11619L),
  NamespaceMetadataType_IsNamespaceSubset((int) 11620L),
  NamespaceMetadataType_StaticNodeIdTypes((int) 11621L),
  NamespaceMetadataType_StaticNumericNodeIdRange((int) 11622L),
  NamespaceMetadataType_StaticStringNodeIdPattern((int) 11623L),
  NamespaceMetadataType_NamespaceFile_Size((int) 11625L),
  NamespaceMetadataType_NamespaceFile_OpenCount((int) 11628L),
  NamespaceMetadataType_NamespaceFile_Open_InputArguments((int) 11630L),
  NamespaceMetadataType_NamespaceFile_Open_OutputArguments((int) 11631L),
  NamespaceMetadataType_NamespaceFile_Close_InputArguments((int) 11633L),
  NamespaceMetadataType_NamespaceFile_Read_InputArguments((int) 11635L),
  NamespaceMetadataType_NamespaceFile_Read_OutputArguments((int) 11636L),
  NamespaceMetadataType_NamespaceFile_Write_InputArguments((int) 11638L),
  NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments((int) 11640L),
  NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments((int) 11641L),
  NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments((int) 11643L),
  NamespaceMetadataType_NamespaceFile_Writable((int) 12690L),
  NamespaceMetadataType_NamespaceFile_UserWritable((int) 12691L),
  NamespaceMetadataType_NamespaceFile_MimeType((int) 13399L),
  NamespaceMetadataType_DefaultRolePermissions((int) 16137L),
  NamespaceMetadataType_DefaultUserRolePermissions((int) 16138L),
  NamespaceMetadataType_DefaultAccessRestrictions((int) 16139L),
  NamespaceMetadataType_NamespaceFile_MaxByteStringLength((int) 24246L),
  NamespaceMetadataType_NamespaceFile_LastModifiedTime((int) 25202L),
  NamespaceMetadataType_ConfigurationVersion((int) 25267L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableNamespace> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableNamespace value :
        OpcuaNodeIdServicesVariableNamespace.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableNamespace(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableNamespace enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}