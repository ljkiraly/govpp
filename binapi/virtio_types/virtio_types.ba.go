// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.10.0-dev
//  VPP:              23.10-rc0~170-g6f1548434
// source: core/virtio_types.api.json

// Package virtio_types contains generated bindings for API file virtio_types.api.
//
// Contents:
// -  2 enums
package virtio_types

import (
	"strconv"

	api "go.fd.io/govpp/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "virtio_types"
	APIVersion = "1.0.0"
	VersionCrc = 0x7a70a44e
)

// VirtioNetFeaturesFirst32 defines enum 'virtio_net_features_first_32'.
type VirtioNetFeaturesFirst32 uint32

const (
	VIRTIO_NET_F_API_CSUM              VirtioNetFeaturesFirst32 = 1
	VIRTIO_NET_F_API_GUEST_CSUM        VirtioNetFeaturesFirst32 = 2
	VIRTIO_NET_F_API_GUEST_TSO4        VirtioNetFeaturesFirst32 = 128
	VIRTIO_NET_F_API_GUEST_TSO6        VirtioNetFeaturesFirst32 = 256
	VIRTIO_NET_F_API_GUEST_UFO         VirtioNetFeaturesFirst32 = 1024
	VIRTIO_NET_F_API_HOST_TSO4         VirtioNetFeaturesFirst32 = 2048
	VIRTIO_NET_F_API_HOST_TSO6         VirtioNetFeaturesFirst32 = 4096
	VIRTIO_NET_F_API_HOST_UFO          VirtioNetFeaturesFirst32 = 16384
	VIRTIO_NET_F_API_MRG_RXBUF         VirtioNetFeaturesFirst32 = 32768
	VIRTIO_NET_F_API_CTRL_VQ           VirtioNetFeaturesFirst32 = 131072
	VIRTIO_NET_F_API_GUEST_ANNOUNCE    VirtioNetFeaturesFirst32 = 2097152
	VIRTIO_NET_F_API_MQ                VirtioNetFeaturesFirst32 = 4194304
	VHOST_F_API_LOG_ALL                VirtioNetFeaturesFirst32 = 67108864
	VIRTIO_F_API_ANY_LAYOUT            VirtioNetFeaturesFirst32 = 134217728
	VIRTIO_F_API_INDIRECT_DESC         VirtioNetFeaturesFirst32 = 268435456
	VHOST_USER_F_API_PROTOCOL_FEATURES VirtioNetFeaturesFirst32 = 1073741824
)

var (
	VirtioNetFeaturesFirst32_name = map[uint32]string{
		1:          "VIRTIO_NET_F_API_CSUM",
		2:          "VIRTIO_NET_F_API_GUEST_CSUM",
		128:        "VIRTIO_NET_F_API_GUEST_TSO4",
		256:        "VIRTIO_NET_F_API_GUEST_TSO6",
		1024:       "VIRTIO_NET_F_API_GUEST_UFO",
		2048:       "VIRTIO_NET_F_API_HOST_TSO4",
		4096:       "VIRTIO_NET_F_API_HOST_TSO6",
		16384:      "VIRTIO_NET_F_API_HOST_UFO",
		32768:      "VIRTIO_NET_F_API_MRG_RXBUF",
		131072:     "VIRTIO_NET_F_API_CTRL_VQ",
		2097152:    "VIRTIO_NET_F_API_GUEST_ANNOUNCE",
		4194304:    "VIRTIO_NET_F_API_MQ",
		67108864:   "VHOST_F_API_LOG_ALL",
		134217728:  "VIRTIO_F_API_ANY_LAYOUT",
		268435456:  "VIRTIO_F_API_INDIRECT_DESC",
		1073741824: "VHOST_USER_F_API_PROTOCOL_FEATURES",
	}
	VirtioNetFeaturesFirst32_value = map[string]uint32{
		"VIRTIO_NET_F_API_CSUM":              1,
		"VIRTIO_NET_F_API_GUEST_CSUM":        2,
		"VIRTIO_NET_F_API_GUEST_TSO4":        128,
		"VIRTIO_NET_F_API_GUEST_TSO6":        256,
		"VIRTIO_NET_F_API_GUEST_UFO":         1024,
		"VIRTIO_NET_F_API_HOST_TSO4":         2048,
		"VIRTIO_NET_F_API_HOST_TSO6":         4096,
		"VIRTIO_NET_F_API_HOST_UFO":          16384,
		"VIRTIO_NET_F_API_MRG_RXBUF":         32768,
		"VIRTIO_NET_F_API_CTRL_VQ":           131072,
		"VIRTIO_NET_F_API_GUEST_ANNOUNCE":    2097152,
		"VIRTIO_NET_F_API_MQ":                4194304,
		"VHOST_F_API_LOG_ALL":                67108864,
		"VIRTIO_F_API_ANY_LAYOUT":            134217728,
		"VIRTIO_F_API_INDIRECT_DESC":         268435456,
		"VHOST_USER_F_API_PROTOCOL_FEATURES": 1073741824,
	}
)

func (x VirtioNetFeaturesFirst32) String() string {
	s, ok := VirtioNetFeaturesFirst32_name[uint32(x)]
	if ok {
		return s
	}
	return "VirtioNetFeaturesFirst32(" + strconv.Itoa(int(x)) + ")"
}

// VirtioNetFeaturesLast32 defines enum 'virtio_net_features_last_32'.
type VirtioNetFeaturesLast32 uint32

const (
	VIRTIO_F_API_VERSION_1 VirtioNetFeaturesLast32 = 1
)

var (
	VirtioNetFeaturesLast32_name = map[uint32]string{
		1: "VIRTIO_F_API_VERSION_1",
	}
	VirtioNetFeaturesLast32_value = map[string]uint32{
		"VIRTIO_F_API_VERSION_1": 1,
	}
)

func (x VirtioNetFeaturesLast32) String() string {
	s, ok := VirtioNetFeaturesLast32_name[uint32(x)]
	if ok {
		return s
	}
	return "VirtioNetFeaturesLast32(" + strconv.Itoa(int(x)) + ")"
}
