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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class SimpleTypeDescription extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15007;
  }

  // Properties.
  protected final NodeId dataTypeId;
  protected final QualifiedName name;
  protected final NodeId baseDataType;
  protected final short builtInType;

  public SimpleTypeDescription(
      NodeId dataTypeId, QualifiedName name, NodeId baseDataType, short builtInType) {
    super();
    this.dataTypeId = dataTypeId;
    this.name = name;
    this.baseDataType = baseDataType;
    this.builtInType = builtInType;
  }

  public NodeId getDataTypeId() {
    return dataTypeId;
  }

  public QualifiedName getName() {
    return name;
  }

  public NodeId getBaseDataType() {
    return baseDataType;
  }

  public short getBuiltInType() {
    return builtInType;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SimpleTypeDescription");

    // Simple Field (dataTypeId)
    writeSimpleField("dataTypeId", dataTypeId, writeComplex(writeBuffer));

    // Simple Field (name)
    writeSimpleField("name", name, writeComplex(writeBuffer));

    // Simple Field (baseDataType)
    writeSimpleField("baseDataType", baseDataType, writeComplex(writeBuffer));

    // Simple Field (builtInType)
    writeSimpleField("builtInType", builtInType, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("SimpleTypeDescription");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SimpleTypeDescription _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dataTypeId)
    lengthInBits += dataTypeId.getLengthInBits();

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    // Simple field (baseDataType)
    lengthInBits += baseDataType.getLengthInBits();

    // Simple field (builtInType)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("SimpleTypeDescription");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId dataTypeId =
        readSimpleField(
            "dataTypeId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    QualifiedName name =
        readSimpleField(
            "name", readComplex(() -> QualifiedName.staticParse(readBuffer), readBuffer));

    NodeId baseDataType =
        readSimpleField(
            "baseDataType", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    short builtInType = readSimpleField("builtInType", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("SimpleTypeDescription");
    // Create the instance
    return new SimpleTypeDescriptionBuilderImpl(dataTypeId, name, baseDataType, builtInType);
  }

  public static class SimpleTypeDescriptionBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId dataTypeId;
    private final QualifiedName name;
    private final NodeId baseDataType;
    private final short builtInType;

    public SimpleTypeDescriptionBuilderImpl(
        NodeId dataTypeId, QualifiedName name, NodeId baseDataType, short builtInType) {
      this.dataTypeId = dataTypeId;
      this.name = name;
      this.baseDataType = baseDataType;
      this.builtInType = builtInType;
    }

    public SimpleTypeDescription build() {
      SimpleTypeDescription simpleTypeDescription =
          new SimpleTypeDescription(dataTypeId, name, baseDataType, builtInType);
      return simpleTypeDescription;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SimpleTypeDescription)) {
      return false;
    }
    SimpleTypeDescription that = (SimpleTypeDescription) o;
    return (getDataTypeId() == that.getDataTypeId())
        && (getName() == that.getName())
        && (getBaseDataType() == that.getBaseDataType())
        && (getBuiltInType() == that.getBuiltInType())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getDataTypeId(), getName(), getBaseDataType(), getBuiltInType());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}