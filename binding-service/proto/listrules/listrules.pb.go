// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/binding-service/proto/listrules/listrules.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_kernel_binding_listrules is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/binding-service/proto/listrules/listrules.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_kernel_binding_listrules

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"
import com_HailoOSS_kernel_binding "github.com/HailoOSS/binding-service/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Service          *string `protobuf:"bytes,1,req,name=service" json:"service,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetService() string {
	if m != nil && m.Service != nil {
		return *m.Service
	}
	return ""
}

type Response struct {
	Rules            []*com_HailoOSS_kernel_binding.BindingRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetRules() []*com_HailoOSS_kernel_binding.BindingRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
}