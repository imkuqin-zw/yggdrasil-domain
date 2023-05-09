// Code generated by protoc-gen-yggdrasil-domain. DO NOT EDIT.

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the yggdrasil package it is being compiled against.
var _ = proto.Message(nil)
var _ = fmt.GoStringer(nil)

func (e *AccountBalanceChangeEvent) ID() string {
	return fmt.Sprintf("%d:%d", e.BusinessType, e.BusinessUniqueNo)
}

func (*AccountBalanceChangeEvent) Topic() string {
	return "AccountBalanceChange"
}

func (e *AccountBalanceChangeEvent) Content() []byte {
	content, _ := proto.Marshal(e)
	return content
}

func (e *AccountBalanceChangeEvent) Unmarshal(content []byte) error {
	return proto.Unmarshal(content, e)
}