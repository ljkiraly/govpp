// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0-dev
//  VPP:              23.10-rc0~170-gd508588a1
// source: plugins/trace.api.json

// Package trace contains generated bindings for API file trace.api.
//
// Contents:
// -  6 messages
package trace

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "trace"
	APIVersion = "1.0.0"
	VersionCrc = 0x397cbf90
)

// iOAM6 Trace - Set the iOAM6 trace profile
//   - trace_type - Type of trace requested
//   - num_elts - Number of trace elements to be inserted
//   - node_id - Trace Node ID
//   - trace_tsp- Timestamp resolution
//   - app_data - Application specific opaque
//
// TraceProfileAdd defines message 'trace_profile_add'.
type TraceProfileAdd struct {
	TraceType uint8  `binapi:"u8,name=trace_type" json:"trace_type,omitempty"`
	NumElts   uint8  `binapi:"u8,name=num_elts" json:"num_elts,omitempty"`
	TraceTsp  uint8  `binapi:"u8,name=trace_tsp" json:"trace_tsp,omitempty"`
	NodeID    uint32 `binapi:"u32,name=node_id" json:"node_id,omitempty"`
	AppData   uint32 `binapi:"u32,name=app_data" json:"app_data,omitempty"`
}

func (m *TraceProfileAdd) Reset()               { *m = TraceProfileAdd{} }
func (*TraceProfileAdd) GetMessageName() string { return "trace_profile_add" }
func (*TraceProfileAdd) GetCrcString() string   { return "de08aa6d" }
func (*TraceProfileAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceProfileAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.TraceType
	size += 1 // m.NumElts
	size += 1 // m.TraceTsp
	size += 4 // m.NodeID
	size += 4 // m.AppData
	return size
}
func (m *TraceProfileAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.TraceType)
	buf.EncodeUint8(m.NumElts)
	buf.EncodeUint8(m.TraceTsp)
	buf.EncodeUint32(m.NodeID)
	buf.EncodeUint32(m.AppData)
	return buf.Bytes(), nil
}
func (m *TraceProfileAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TraceType = buf.DecodeUint8()
	m.NumElts = buf.DecodeUint8()
	m.TraceTsp = buf.DecodeUint8()
	m.NodeID = buf.DecodeUint32()
	m.AppData = buf.DecodeUint32()
	return nil
}

// TraceProfileAddReply defines message 'trace_profile_add_reply'.
type TraceProfileAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceProfileAddReply) Reset()               { *m = TraceProfileAddReply{} }
func (*TraceProfileAddReply) GetMessageName() string { return "trace_profile_add_reply" }
func (*TraceProfileAddReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceProfileAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceProfileAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceProfileAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceProfileAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Delete trace Profile
// TraceProfileDel defines message 'trace_profile_del'.
type TraceProfileDel struct{}

func (m *TraceProfileDel) Reset()               { *m = TraceProfileDel{} }
func (*TraceProfileDel) GetMessageName() string { return "trace_profile_del" }
func (*TraceProfileDel) GetCrcString() string   { return "51077d14" }
func (*TraceProfileDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceProfileDel) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TraceProfileDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TraceProfileDel) Unmarshal(b []byte) error {
	return nil
}

// TraceProfileDelReply defines message 'trace_profile_del_reply'.
type TraceProfileDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceProfileDelReply) Reset()               { *m = TraceProfileDelReply{} }
func (*TraceProfileDelReply) GetMessageName() string { return "trace_profile_del_reply" }
func (*TraceProfileDelReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceProfileDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceProfileDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceProfileDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceProfileDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Show trace Profile
// TraceProfileShowConfig defines message 'trace_profile_show_config'.
type TraceProfileShowConfig struct{}

func (m *TraceProfileShowConfig) Reset()               { *m = TraceProfileShowConfig{} }
func (*TraceProfileShowConfig) GetMessageName() string { return "trace_profile_show_config" }
func (*TraceProfileShowConfig) GetCrcString() string   { return "51077d14" }
func (*TraceProfileShowConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceProfileShowConfig) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TraceProfileShowConfig) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TraceProfileShowConfig) Unmarshal(b []byte) error {
	return nil
}

// Show trace config response
//   - retval - return value for request
//   - trace_type - Type of trace requested
//   - num_elts - Number of trace elements to be inserted
//   - node_id - Trace Node ID
//   - trace_tsp- Timestamp resolution
//   - app_data - Application specific opaque
//
// TraceProfileShowConfigReply defines message 'trace_profile_show_config_reply'.
type TraceProfileShowConfigReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	TraceType uint8  `binapi:"u8,name=trace_type" json:"trace_type,omitempty"`
	NumElts   uint8  `binapi:"u8,name=num_elts" json:"num_elts,omitempty"`
	TraceTsp  uint8  `binapi:"u8,name=trace_tsp" json:"trace_tsp,omitempty"`
	NodeID    uint32 `binapi:"u32,name=node_id" json:"node_id,omitempty"`
	AppData   uint32 `binapi:"u32,name=app_data" json:"app_data,omitempty"`
}

func (m *TraceProfileShowConfigReply) Reset()               { *m = TraceProfileShowConfigReply{} }
func (*TraceProfileShowConfigReply) GetMessageName() string { return "trace_profile_show_config_reply" }
func (*TraceProfileShowConfigReply) GetCrcString() string   { return "0f1d374c" }
func (*TraceProfileShowConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceProfileShowConfigReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 1 // m.TraceType
	size += 1 // m.NumElts
	size += 1 // m.TraceTsp
	size += 4 // m.NodeID
	size += 4 // m.AppData
	return size
}
func (m *TraceProfileShowConfigReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint8(m.TraceType)
	buf.EncodeUint8(m.NumElts)
	buf.EncodeUint8(m.TraceTsp)
	buf.EncodeUint32(m.NodeID)
	buf.EncodeUint32(m.AppData)
	return buf.Bytes(), nil
}
func (m *TraceProfileShowConfigReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.TraceType = buf.DecodeUint8()
	m.NumElts = buf.DecodeUint8()
	m.TraceTsp = buf.DecodeUint8()
	m.NodeID = buf.DecodeUint32()
	m.AppData = buf.DecodeUint32()
	return nil
}

func init() { file_trace_binapi_init() }
func file_trace_binapi_init() {
	api.RegisterMessage((*TraceProfileAdd)(nil), "trace_profile_add_de08aa6d")
	api.RegisterMessage((*TraceProfileAddReply)(nil), "trace_profile_add_reply_e8d4e804")
	api.RegisterMessage((*TraceProfileDel)(nil), "trace_profile_del_51077d14")
	api.RegisterMessage((*TraceProfileDelReply)(nil), "trace_profile_del_reply_e8d4e804")
	api.RegisterMessage((*TraceProfileShowConfig)(nil), "trace_profile_show_config_51077d14")
	api.RegisterMessage((*TraceProfileShowConfigReply)(nil), "trace_profile_show_config_reply_0f1d374c")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TraceProfileAdd)(nil),
		(*TraceProfileAddReply)(nil),
		(*TraceProfileDel)(nil),
		(*TraceProfileDelReply)(nil),
		(*TraceProfileShowConfig)(nil),
		(*TraceProfileShowConfigReply)(nil),
	}
}
