#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from __future__ import annotations

from asyncio import Transport
from dataclasses import dataclass
from typing import Any, Optional


@dataclass
class TCPTransport(Transport):
    """
    Wrapper for the TCP Python Transport. Instead of using the generic asyncio.create_connection.
    Using a custom transport allows for it to be modified if required.
    """

    loop: str
    _protocol: str
    _read_transport: Transport
    _write_transport: Transport

    @property
    def protocol(self):
        return self._protocol

    @protocol.setter
    def protocol(self, protocol: str):
        self._protocol = protocol

    def is_reading(self):
        return self._read_transport

    def pause_reading(self):
        self._read_transport.pause_reading()

    def resume_reading(self):
        self._read_transport.pause_reading()

    def abort(self):
        self._write_transport.abort()

    def can_write_eof(self):
        return self._write_transport.can_write_eof()

    def get_write_buffer_size(self):
        return self._write_transport.get_write_buffer_size()

    # def get_write_buffer_limits(self):
    #    return self._write_transport.get_write_buffer_limits()

    def set_write_buffer_limits(
        self, high: Optional[int] = 65535, low: Optional[int] = 1023
    ) -> None:
        self._write_transport.set_write_buffer_limits(high, low)

    def write(self, data: Any) -> None:
        self._write_transport.write(data)

    def writelines(self, list_of_data: list[Any]) -> None:
        self._write_transport.writelines(list_of_data)

    def write_eof(self):
        self._write_transport.write_eof()
