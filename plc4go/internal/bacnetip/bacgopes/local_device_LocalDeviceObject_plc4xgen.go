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

// Code generated by "plc4xGenerator -type=LocalDeviceObject -prefix=local_device_"; DO NOT EDIT.

package bacgopes

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *LocalDeviceObject) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *LocalDeviceObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("LocalDeviceObject"); err != nil {
		return err
	}
	if d.NumberOfAPDURetries != nil {
		{
			_value := fmt.Sprintf("%v", d.NumberOfAPDURetries)

			if err := writeBuffer.WriteString("numberOfAPDURetries", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.APDUTimeout != nil {
		{
			_value := fmt.Sprintf("%v", d.APDUTimeout)

			if err := writeBuffer.WriteString("aPDUTimeout", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.SegmentationSupported != nil {
		if err := writeBuffer.PushContext("segmentationSupported"); err != nil {
			return err
		}
		if err := d.SegmentationSupported.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
			return err
		}
		if err := writeBuffer.PopContext("segmentationSupported"); err != nil {
			return err
		}
	}
	if d.APDUSegmentTimeout != nil {
		{
			_value := fmt.Sprintf("%v", d.APDUSegmentTimeout)

			if err := writeBuffer.WriteString("aPDUSegmentTimeout", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if d.MaxSegmentsAccepted != nil {
		if err := writeBuffer.PushContext("maxSegmentsAccepted"); err != nil {
			return err
		}
		if err := d.MaxSegmentsAccepted.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
			return err
		}
		if err := writeBuffer.PopContext("maxSegmentsAccepted"); err != nil {
			return err
		}
	}
	if d.MaximumApduLengthAccepted != nil {
		if err := writeBuffer.PushContext("maximumApduLengthAccepted"); err != nil {
			return err
		}
		if err := d.MaximumApduLengthAccepted.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
			return err
		}
		if err := writeBuffer.PopContext("maximumApduLengthAccepted"); err != nil {
			return err
		}
	}

	if err := writeBuffer.WriteString("objectName", uint32(len(d.ObjectName)*8), d.ObjectName); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("objectIdentifier", uint32(len(d.ObjectIdentifier)*8), d.ObjectIdentifier); err != nil {
		return err
	}

	if err := writeBuffer.WriteUint16("vendorIdentifier", 16, d.VendorIdentifier); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("objectList", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.ObjectList {
		if err := writeBuffer.WriteString("", uint32(len(elem)*8), elem); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("objectList", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("LocalDeviceObject"); err != nil {
		return err
	}
	return nil
}

func (d *LocalDeviceObject) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
