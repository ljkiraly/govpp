// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0-dev
//  VPP:              23.10-rc0~170-gd508588a1
// source: plugins/lldp.api.json

// Package lldp contains generated bindings for API file lldp.api.
//
// Contents:
// -  4 messages
package lldp

import (
	interface_types "github.com/ljkiraly/govpp/binapi/interface_types"
	ip_types "github.com/ljkiraly/govpp/binapi/ip_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "lldp"
	APIVersion = "2.0.0"
	VersionCrc = 0x8a7e7049
)

// configure global parameter for LLDP
//   - system_name - VPP system name
//   - tx_hold - multiplier for tx_interval when setting time-to-live (TTL)
//     value in the LLDP packets
//   - tx_interval - time interval, in seconds, between each LLDP frames
//
// LldpConfig defines message 'lldp_config'.
type LldpConfig struct {
	TxHold     uint32 `binapi:"u32,name=tx_hold" json:"tx_hold,omitempty"`
	TxInterval uint32 `binapi:"u32,name=tx_interval" json:"tx_interval,omitempty"`
	SystemName string `binapi:"string[],name=system_name" json:"system_name,omitempty"`
}

func (m *LldpConfig) Reset()               { *m = LldpConfig{} }
func (*LldpConfig) GetMessageName() string { return "lldp_config" }
func (*LldpConfig) GetCrcString() string   { return "c14445df" }
func (*LldpConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LldpConfig) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                     // m.TxHold
	size += 4                     // m.TxInterval
	size += 4 + len(m.SystemName) // m.SystemName
	return size
}
func (m *LldpConfig) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.TxHold)
	buf.EncodeUint32(m.TxInterval)
	buf.EncodeString(m.SystemName, 0)
	return buf.Bytes(), nil
}
func (m *LldpConfig) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.TxHold = buf.DecodeUint32()
	m.TxInterval = buf.DecodeUint32()
	m.SystemName = buf.DecodeString(0)
	return nil
}

// LldpConfigReply defines message 'lldp_config_reply'.
type LldpConfigReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LldpConfigReply) Reset()               { *m = LldpConfigReply{} }
func (*LldpConfigReply) GetMessageName() string { return "lldp_config_reply" }
func (*LldpConfigReply) GetCrcString() string   { return "e8d4e804" }
func (*LldpConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LldpConfigReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LldpConfigReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LldpConfigReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Interface set LLDP request
//   - sw_if_index - interface for which to enable/disable LLDP
//   - mgmt_ip4_addr - management ip4 address of the interface
//   - mgmt_ip6_addr - management ip6 address of the interface
//   - mgmt_oid - OID(Object Identifier) of the interface
//   - enable - if non-zero enable, else disable
//   - port_desc - local port description
//
// SwInterfaceSetLldp defines message 'sw_interface_set_lldp'.
type SwInterfaceSetLldp struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	MgmtIP4   ip_types.IP4Address            `binapi:"ip4_address,name=mgmt_ip4" json:"mgmt_ip4,omitempty"`
	MgmtIP6   ip_types.IP6Address            `binapi:"ip6_address,name=mgmt_ip6" json:"mgmt_ip6,omitempty"`
	MgmtOid   []byte                         `binapi:"u8[128],name=mgmt_oid" json:"mgmt_oid,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable,default=true" json:"enable,omitempty"`
	PortDesc  string                         `binapi:"string[],name=port_desc" json:"port_desc,omitempty"`
}

func (m *SwInterfaceSetLldp) Reset()               { *m = SwInterfaceSetLldp{} }
func (*SwInterfaceSetLldp) GetMessageName() string { return "sw_interface_set_lldp" }
func (*SwInterfaceSetLldp) GetCrcString() string   { return "57afbcd4" }
func (*SwInterfaceSetLldp) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetLldp) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                   // m.SwIfIndex
	size += 1 * 4               // m.MgmtIP4
	size += 1 * 16              // m.MgmtIP6
	size += 1 * 128             // m.MgmtOid
	size += 1                   // m.Enable
	size += 4 + len(m.PortDesc) // m.PortDesc
	return size
}
func (m *SwInterfaceSetLldp) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.MgmtIP4[:], 4)
	buf.EncodeBytes(m.MgmtIP6[:], 16)
	buf.EncodeBytes(m.MgmtOid, 128)
	buf.EncodeBool(m.Enable)
	buf.EncodeString(m.PortDesc, 0)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetLldp) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.MgmtIP4[:], buf.DecodeBytes(4))
	copy(m.MgmtIP6[:], buf.DecodeBytes(16))
	m.MgmtOid = make([]byte, 128)
	copy(m.MgmtOid, buf.DecodeBytes(len(m.MgmtOid)))
	m.Enable = buf.DecodeBool()
	m.PortDesc = buf.DecodeString(0)
	return nil
}

// SwInterfaceSetLldpReply defines message 'sw_interface_set_lldp_reply'.
type SwInterfaceSetLldpReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetLldpReply) Reset()               { *m = SwInterfaceSetLldpReply{} }
func (*SwInterfaceSetLldpReply) GetMessageName() string { return "sw_interface_set_lldp_reply" }
func (*SwInterfaceSetLldpReply) GetCrcString() string   { return "e8d4e804" }
func (*SwInterfaceSetLldpReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetLldpReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetLldpReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetLldpReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_lldp_binapi_init() }
func file_lldp_binapi_init() {
	api.RegisterMessage((*LldpConfig)(nil), "lldp_config_c14445df")
	api.RegisterMessage((*LldpConfigReply)(nil), "lldp_config_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceSetLldp)(nil), "sw_interface_set_lldp_57afbcd4")
	api.RegisterMessage((*SwInterfaceSetLldpReply)(nil), "sw_interface_set_lldp_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LldpConfig)(nil),
		(*LldpConfigReply)(nil),
		(*SwInterfaceSetLldp)(nil),
		(*SwInterfaceSetLldpReply)(nil),
	}
}
