// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/platform/proto/trace/trace.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_platform_trace is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/platform/proto/trace/trace.proto

It has these top-level messages:
	Event
*/
package com_HailoOSS_kernel_platform_trace

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Event_Type int32

const (
	Event_IN              Event_Type = 1
	Event_OUT             Event_Type = 2
	Event_REQ             Event_Type = 3
	Event_REP             Event_Type = 4
	Event_ATTEMPT_TIMEOUT Event_Type = 5
	Event_START           Event_Type = 6
)

var Event_Type_name = map[int32]string{
	1: "IN",
	2: "OUT",
	3: "REQ",
	4: "REP",
	5: "ATTEMPT_TIMEOUT",
	6: "START",
}
var Event_Type_value = map[string]int32{
	"IN":              1,
	"OUT":             2,
	"REQ":             3,
	"REP":             4,
	"ATTEMPT_TIMEOUT": 5,
	"START":           6,
}

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}
func (x Event_Type) String() string {
	return proto.EnumName(Event_Type_name, int32(x))
}
func (x Event_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Event_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_Type_value, data, "Event_Type")
	if err != nil {
		return err
	}
	*x = Event_Type(value)
	return nil
}

type Event struct {
	Timestamp         *int64      `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	TraceId           *string     `protobuf:"bytes,2,req,name=traceId" json:"traceId,omitempty"`
	Type              *Event_Type `protobuf:"varint,3,req,name=type,enum=com.HailoOSS.kernel.platform.trace.Event_Type" json:"type,omitempty"`
	MessageId         *string     `protobuf:"bytes,4,opt,name=messageId" json:"messageId,omitempty"`
	ParentMessageId   *string     `protobuf:"bytes,5,opt,name=parentMessageId" json:"parentMessageId,omitempty"`
	From              *string     `protobuf:"bytes,6,opt,name=from" json:"from,omitempty"`
	FromEndpoint      *string     `protobuf:"bytes,16,opt,name=fromEndpoint" json:"fromEndpoint,omitempty"`
	To                *string     `protobuf:"bytes,7,opt,name=to" json:"to,omitempty"`
	Hostname          *string     `protobuf:"bytes,8,opt,name=hostname" json:"hostname,omitempty"`
	Az                *string     `protobuf:"bytes,9,opt,name=az" json:"az,omitempty"`
	Payload           *string     `protobuf:"bytes,10,opt,name=payload" json:"payload,omitempty"`
	ErrorCode         *string     `protobuf:"bytes,11,opt,name=errorCode" json:"errorCode,omitempty"`
	ErrorDescription  *string     `protobuf:"bytes,12,opt,name=errorDescription" json:"errorDescription,omitempty"`
	HandlerInstanceId *string     `protobuf:"bytes,13,opt,name=handlerInstanceId" json:"handlerInstanceId,omitempty"`
	Duration          *int64      `protobuf:"varint,14,opt,name=duration" json:"duration,omitempty"`
	PersistentTrace   *bool       `protobuf:"varint,15,opt,name=persistentTrace,def=0" json:"persistentTrace,omitempty"`
	XXX_unrecognized  []byte      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}

const Default_Event_PersistentTrace bool = false

func (m *Event) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Event) GetTraceId() string {
	if m != nil && m.TraceId != nil {
		return *m.TraceId
	}
	return ""
}

func (m *Event) GetType() Event_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Event_IN
}

func (m *Event) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Event) GetParentMessageId() string {
	if m != nil && m.ParentMessageId != nil {
		return *m.ParentMessageId
	}
	return ""
}

func (m *Event) GetFrom() string {
	if m != nil && m.From != nil {
		return *m.From
	}
	return ""
}

func (m *Event) GetFromEndpoint() string {
	if m != nil && m.FromEndpoint != nil {
		return *m.FromEndpoint
	}
	return ""
}

func (m *Event) GetTo() string {
	if m != nil && m.To != nil {
		return *m.To
	}
	return ""
}

func (m *Event) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Event) GetAz() string {
	if m != nil && m.Az != nil {
		return *m.Az
	}
	return ""
}

func (m *Event) GetPayload() string {
	if m != nil && m.Payload != nil {
		return *m.Payload
	}
	return ""
}

func (m *Event) GetErrorCode() string {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ""
}

func (m *Event) GetErrorDescription() string {
	if m != nil && m.ErrorDescription != nil {
		return *m.ErrorDescription
	}
	return ""
}

func (m *Event) GetHandlerInstanceId() string {
	if m != nil && m.HandlerInstanceId != nil {
		return *m.HandlerInstanceId
	}
	return ""
}

func (m *Event) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Event) GetPersistentTrace() bool {
	if m != nil && m.PersistentTrace != nil {
		return *m.PersistentTrace
	}
	return Default_Event_PersistentTrace
}

func init() {
	proto.RegisterEnum("com.HailoOSS.kernel.platform.trace.Event_Type", Event_Type_name, Event_Type_value)
}