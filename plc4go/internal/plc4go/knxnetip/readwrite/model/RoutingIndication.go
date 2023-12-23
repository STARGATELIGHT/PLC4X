/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type RoutingIndication struct {
	Parent *KnxNetIpMessage
}

// The corresponding interface
type IRoutingIndication interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *RoutingIndication) MsgType() uint16 {
	return 0x0530
}

func (m *RoutingIndication) InitializeParent(parent *KnxNetIpMessage) {
}

func NewRoutingIndication() *KnxNetIpMessage {
	child := &RoutingIndication{
		Parent: NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastRoutingIndication(structType interface{}) *RoutingIndication {
	castFunc := func(typ interface{}) *RoutingIndication {
		if casted, ok := typ.(RoutingIndication); ok {
			return &casted
		}
		if casted, ok := typ.(*RoutingIndication); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastRoutingIndication(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastRoutingIndication(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *RoutingIndication) GetTypeName() string {
	return "RoutingIndication"
}

func (m *RoutingIndication) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *RoutingIndication) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *RoutingIndication) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func RoutingIndicationParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("RoutingIndication"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("RoutingIndication"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &RoutingIndication{
		Parent: &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *RoutingIndication) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RoutingIndication"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("RoutingIndication"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *RoutingIndication) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}