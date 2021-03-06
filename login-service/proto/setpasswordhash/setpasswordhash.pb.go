// Code generated by protoc-gen-go.
// source: github.com/HailoOSS/login-service/proto/setpasswordhash/setpasswordhash.proto
// DO NOT EDIT!

/*
Package com_HailoOSS_service_login_setpasswordhash is a generated protocol buffer package.

It is generated from these files:
	github.com/HailoOSS/login-service/proto/setpasswordhash/setpasswordhash.proto

It has these top-level messages:
	Request
	Response
*/
package com_HailoOSS_service_login_setpasswordhash

import proto "github.com/HailoOSS/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	// defines the application this data logically belongs to
	Application *string `protobuf:"bytes,2,req,name=application" json:"application,omitempty"`
	// Who we are changing the password for
	Uid *string `protobuf:"bytes,3,req,name=uid" json:"uid,omitempty"`
	// The new password
	PasswordHash     *string `protobuf:"bytes,4,req,name=passwordHash" json:"passwordHash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetApplication() string {
	if m != nil && m.Application != nil {
		return *m.Application
	}
	return ""
}

func (m *Request) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *Request) GetPasswordHash() string {
	if m != nil && m.PasswordHash != nil {
		return *m.PasswordHash
	}
	return ""
}

// Response is empty if the call was successful
type Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func init() {
}
