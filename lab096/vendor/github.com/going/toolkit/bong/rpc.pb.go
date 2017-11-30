// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package bong

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Id               *uint64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Method           *string `protobuf:"bytes,2,req,name=method" json:"method,omitempty"`
	Body             []byte  `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Request) Reset()         { *this = Request{} }
func (this *Request) String() string { return proto.CompactTextString(this) }
func (*Request) ProtoMessage()       {}

func (this *Request) GetId() uint64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Request) GetMethod() string {
	if this != nil && this.Method != nil {
		return *this.Method
	}
	return ""
}

func (this *Request) GetBody() []byte {
	if this != nil {
		return this.Body
	}
	return nil
}

type Response struct {
	Id               *uint64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Error            *string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Body             []byte  `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Response) Reset()         { *this = Response{} }
func (this *Response) String() string { return proto.CompactTextString(this) }
func (*Response) ProtoMessage()       {}

func (this *Response) GetId() uint64 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Response) GetError() string {
	if this != nil && this.Error != nil {
		return *this.Error
	}
	return ""
}

func (this *Response) GetBody() []byte {
	if this != nil {
		return this.Body
	}
	return nil
}

func init() {
}
