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

package bacnetip

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/pkg/errors"

	"github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
)

type Real struct {
	*Atomic[float32]
	*CommonMath
}

func NewReal(arg Arg) (*Real, error) {
	b := &Real{}
	b.Atomic = NewAtomic[float32](b)

	if arg == nil {
		return b, nil
	}
	switch arg := arg.(type) {
	case *tag:
		err := b.Decode(arg)
		if err != nil {
			return nil, errors.Wrap(err, "error decoding")
		}
		return b, nil
	case float32:
		b.value = arg
	case float64:
		b.IsValid(arg)
		b.value = float32(arg)
	case int:
		b.value = float32(arg)
	case *Real:
		b.value = arg.value
	default:
		return nil, errors.Errorf("invalid constructor datatype: %T", arg)
	}

	return b, nil
}

func (d *Real) Encode(tag Tag) {
	var _b = make([]byte, 4)
	binary.BigEndian.PutUint32(_b, math.Float32bits(d.value))
	tag.setAppData(uint(model.BACnetDataType_REAL), _b)
}

func (d *Real) Decode(tag Tag) error {
	if tag.GetTagClass() != model.TagClass_APPLICATION_TAGS || tag.GetTagNumber() != uint(model.BACnetDataType_REAL) {
		return errors.New("Real application tag required")
	}
	if len(tag.GetTagData()) != 4 {
		return errors.New("invalid tag length")
	}

	// extract the data
	d.value = math.Float32frombits(binary.BigEndian.Uint32(tag.GetTagData()))
	return nil
}

func (d *Real) IsValid(arg any) bool {
	switch arg := arg.(type) {
	case float32:
		return true
	case float64:
		if arg > math.MaxFloat32 || -arg > math.MaxFloat32 {
			return false
		}
		return true
	default:
		return false
	}
}

func (d *Real) String() string {
	return fmt.Sprintf("Real(%g)", d.value)
}