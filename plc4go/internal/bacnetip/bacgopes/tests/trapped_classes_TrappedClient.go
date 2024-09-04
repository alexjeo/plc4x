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

package tests

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// TrappedClientContract provides a set of functions which can be overwritten by a sub struct
type TrappedClientContract interface {
	utils.Serializable
	Request(bacgopes.Args, bacgopes.KWArgs) error
	Confirmation(args bacgopes.Args, kwargs bacgopes.KWArgs) error
}

// TrappedClient  An instance of this class sits at the top of a stack.
//
//go:generate plc4xGenerator -type=TrappedClient -prefix=trapped_classes_
type TrappedClient struct {
	TrappedClientContract `ignore:"true"`
	bacgopes.Client

	requestSent          bacgopes.PDU
	confirmationReceived bacgopes.PDU

	log zerolog.Logger
}

func NewTrappedClient(localLog zerolog.Logger, opts ...func(*TrappedClient)) (*TrappedClient, error) {
	t := &TrappedClient{
		log: localLog,
	}
	t.TrappedClientContract = t
	for _, opt := range opts {
		opt(t)
	}
	var err error
	t.Client, err = bacgopes.NewClient(localLog, t)
	if err != nil {
		return nil, errors.Wrap(err, "error building client")
	}
	return t, nil
}

func WithTrappedClientContract(trappedClientContract TrappedClientContract) func(*TrappedClient) {
	return func(t *TrappedClient) {
		t.TrappedClientContract = trappedClientContract
	}
}

func (t *TrappedClient) GetRequestSent() bacgopes.PDU {
	return t.requestSent
}

func (t *TrappedClient) GetConfirmationReceived() bacgopes.PDU {
	return t.confirmationReceived
}

func (t *TrappedClient) Request(args bacgopes.Args, kwargs bacgopes.KWArgs) error {
	t.log.Debug().Stringer("args", args).Stringer("kwargs", kwargs).Msg("Request")
	// a reference for checking
	t.requestSent = args.Get0PDU()

	// continue with regular processing
	return t.Client.Request(args, kwargs)
}

func (t *TrappedClient) Confirmation(args bacgopes.Args, kwargs bacgopes.KWArgs) error {
	t.log.Debug().Stringer("args", args).Stringer("kwargs", kwargs).Msg("Confirmation")
	// a reference for checking
	t.confirmationReceived = args.Get0PDU()
	return nil
}
