// Code generated by protoc-gen-go.
// source: util.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	util.proto

It has these top-level messages:
	MovieMessage
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type MovieMessage struct {
	TypeName    string `protobuf:"bytes,1,opt,name=type_name" json:"type_name,omitempty"`
	MessageData []byte `protobuf:"bytes,2,opt,name=message_data,proto3" json:"message_data,omitempty"`
}

func (m *MovieMessage) Reset()         { *m = MovieMessage{} }
func (m *MovieMessage) String() string { return proto1.CompactTextString(m) }
func (*MovieMessage) ProtoMessage()    {}

func init() {
}