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

package io

import (
	"context"

	"github.com/pkg/errors"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

type DataReaderSimpleByteArray struct {
	*DataReaderSimpleBase[[]byte]
}

var _ DataReader[[]byte] = (*DataReaderSimpleByteArray)(nil)

func NewDataReaderSimpleByteArray(readBuffer utils.ReadBuffer, bitLength uint8) *DataReaderSimpleByteArray {
	return &DataReaderSimpleByteArray{
		DataReaderSimpleBase: NewDataReaderSimpleBase[[]byte](readBuffer, uint(bitLength)),
	}
}

func (d *DataReaderSimpleByteArray) Read(ctx context.Context, logicalName string, readerArgs ...utils.WithReaderArgs) ([]byte, error) {
	if d.bitLength%8 != 0 {
		return nil, errors.New("bitLength must be a multiple of 8")
	}
	return d.readBuffer.ReadByteArray(logicalName, int(d.bitLength/8), readerArgs...)
}