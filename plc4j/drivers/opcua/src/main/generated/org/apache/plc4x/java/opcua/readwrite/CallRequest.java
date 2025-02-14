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

public class CallRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 712;
  }

  // Properties.
  protected final RequestHeader requestHeader;
  protected final List<CallMethodRequest> methodsToCall;

  public CallRequest(RequestHeader requestHeader, List<CallMethodRequest> methodsToCall) {
    super();
    this.requestHeader = requestHeader;
    this.methodsToCall = methodsToCall;
  }

  public RequestHeader getRequestHeader() {
    return requestHeader;
  }

  public List<CallMethodRequest> getMethodsToCall() {
    return methodsToCall;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CallRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Implicit Field (noOfMethodsToCall) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfMethodsToCall =
        (int) ((((getMethodsToCall()) == (null)) ? -(1) : COUNT(getMethodsToCall())));
    writeImplicitField("noOfMethodsToCall", noOfMethodsToCall, writeSignedInt(writeBuffer, 32));

    // Array Field (methodsToCall)
    writeComplexTypeArrayField("methodsToCall", methodsToCall, writeBuffer);

    writeBuffer.popContext("CallRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CallRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Implicit Field (noOfMethodsToCall)
    lengthInBits += 32;

    // Array field
    if (methodsToCall != null) {
      int i = 0;
      for (CallMethodRequest element : methodsToCall) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= methodsToCall.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("CallRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    RequestHeader requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () ->
                    (RequestHeader) ExtensionObjectDefinition.staticParse(readBuffer, (int) (391)),
                readBuffer));

    int noOfMethodsToCall = readImplicitField("noOfMethodsToCall", readSignedInt(readBuffer, 32));

    List<CallMethodRequest> methodsToCall =
        readCountArrayField(
            "methodsToCall",
            readComplex(
                () ->
                    (CallMethodRequest)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (706)),
                readBuffer),
            noOfMethodsToCall);

    readBuffer.closeContext("CallRequest");
    // Create the instance
    return new CallRequestBuilderImpl(requestHeader, methodsToCall);
  }

  public static class CallRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final RequestHeader requestHeader;
    private final List<CallMethodRequest> methodsToCall;

    public CallRequestBuilderImpl(
        RequestHeader requestHeader, List<CallMethodRequest> methodsToCall) {
      this.requestHeader = requestHeader;
      this.methodsToCall = methodsToCall;
    }

    public CallRequest build() {
      CallRequest callRequest = new CallRequest(requestHeader, methodsToCall);
      return callRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CallRequest)) {
      return false;
    }
    CallRequest that = (CallRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getMethodsToCall() == that.getMethodsToCall())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getRequestHeader(), getMethodsToCall());
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
