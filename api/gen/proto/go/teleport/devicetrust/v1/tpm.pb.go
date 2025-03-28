// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: teleport/devicetrust/v1/tpm.proto

package devicetrustv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Encapsulates the value of a PCR at a point at time.
// See https://pkg.go.dev/github.com/google/go-attestation/attest#PCR
type TPMPCR struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the PCR index in the PCR bank
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// the digest currently held in the PCR
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	// the hash algorithm used to produce the digest in this PCR bank. This value
	// is the underlying value of the Go crypto.Hash type.
	DigestAlg     uint64 `protobuf:"varint,3,opt,name=digest_alg,json=digestAlg,proto3" json:"digest_alg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TPMPCR) Reset() {
	*x = TPMPCR{}
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TPMPCR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPMPCR) ProtoMessage() {}

func (x *TPMPCR) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPMPCR.ProtoReflect.Descriptor instead.
func (*TPMPCR) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_tpm_proto_rawDescGZIP(), []int{0}
}

func (x *TPMPCR) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *TPMPCR) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *TPMPCR) GetDigestAlg() uint64 {
	if x != nil {
		return x.DigestAlg
	}
	return 0
}

// Encapsulates the result of a quote operation against the TPM over a PCR
// using an attestation key.
// See https://pkg.go.dev/github.com/google/go-attestation/attest#Quote
type TPMQuote struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Quote         []byte                 `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	Signature     []byte                 `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TPMQuote) Reset() {
	*x = TPMQuote{}
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TPMQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPMQuote) ProtoMessage() {}

func (x *TPMQuote) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPMQuote.ProtoReflect.Descriptor instead.
func (*TPMQuote) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_tpm_proto_rawDescGZIP(), []int{1}
}

func (x *TPMQuote) GetQuote() []byte {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *TPMQuote) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// The quotes, PCRs and event log from a TPM that attest to the booted state
// of the machine.
// See https://pkg.go.dev/github.com/google/go-attestation/attest#PlatformParameters
// Excludes TPMVersion and Public since these are already known values.
type TPMPlatformParameters struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Quotes        []*TPMQuote            `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
	Pcrs          []*TPMPCR              `protobuf:"bytes,2,rep,name=pcrs,proto3" json:"pcrs,omitempty"`
	EventLog      []byte                 `protobuf:"bytes,3,opt,name=event_log,json=eventLog,proto3" json:"event_log,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TPMPlatformParameters) Reset() {
	*x = TPMPlatformParameters{}
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TPMPlatformParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPMPlatformParameters) ProtoMessage() {}

func (x *TPMPlatformParameters) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPMPlatformParameters.ProtoReflect.Descriptor instead.
func (*TPMPlatformParameters) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_tpm_proto_rawDescGZIP(), []int{2}
}

func (x *TPMPlatformParameters) GetQuotes() []*TPMQuote {
	if x != nil {
		return x.Quotes
	}
	return nil
}

func (x *TPMPlatformParameters) GetPcrs() []*TPMPCR {
	if x != nil {
		return x.Pcrs
	}
	return nil
}

func (x *TPMPlatformParameters) GetEventLog() []byte {
	if x != nil {
		return x.EventLog
	}
	return nil
}

// Holds the record of a TPM platform attestation, including the platform
// parameters sent by the device and the nonce the server generated. This allows
// a historical platform attestation to be revalidated and allows us to compare
// the incoming state of a device (e.g during authentication) against the
// historical state in order to detect potentially malicious actions.
type TPMPlatformAttestation struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Nonce              []byte                 `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	PlatformParameters *TPMPlatformParameters `protobuf:"bytes,2,opt,name=platform_parameters,json=platformParameters,proto3" json:"platform_parameters,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *TPMPlatformAttestation) Reset() {
	*x = TPMPlatformAttestation{}
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TPMPlatformAttestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPMPlatformAttestation) ProtoMessage() {}

func (x *TPMPlatformAttestation) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_tpm_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPMPlatformAttestation.ProtoReflect.Descriptor instead.
func (*TPMPlatformAttestation) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_tpm_proto_rawDescGZIP(), []int{3}
}

func (x *TPMPlatformAttestation) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *TPMPlatformAttestation) GetPlatformParameters() *TPMPlatformParameters {
	if x != nil {
		return x.PlatformParameters
	}
	return nil
}

var File_teleport_devicetrust_v1_tpm_proto protoreflect.FileDescriptor

var file_teleport_devicetrust_v1_tpm_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x70, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x55, 0x0a, 0x06,
	0x54, 0x50, 0x4d, 0x50, 0x43, 0x52, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x6c, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x41, 0x6c, 0x67, 0x22, 0x3e, 0x0a, 0x08, 0x54, 0x50, 0x4d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x15, 0x54, 0x50, 0x4d, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a,
	0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x50, 0x4d, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x04, 0x70, 0x63, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x50, 0x4d, 0x50, 0x43, 0x52, 0x52, 0x04, 0x70, 0x63, 0x72, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x54,
	0x50, 0x4d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x13, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x50, 0x4d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x12, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42, 0x5a, 0x5a, 0x58,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teleport_devicetrust_v1_tpm_proto_rawDescOnce sync.Once
	file_teleport_devicetrust_v1_tpm_proto_rawDescData []byte
)

func file_teleport_devicetrust_v1_tpm_proto_rawDescGZIP() []byte {
	file_teleport_devicetrust_v1_tpm_proto_rawDescOnce.Do(func() {
		file_teleport_devicetrust_v1_tpm_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_devicetrust_v1_tpm_proto_rawDesc), len(file_teleport_devicetrust_v1_tpm_proto_rawDesc)))
	})
	return file_teleport_devicetrust_v1_tpm_proto_rawDescData
}

var file_teleport_devicetrust_v1_tpm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_devicetrust_v1_tpm_proto_goTypes = []any{
	(*TPMPCR)(nil),                 // 0: teleport.devicetrust.v1.TPMPCR
	(*TPMQuote)(nil),               // 1: teleport.devicetrust.v1.TPMQuote
	(*TPMPlatformParameters)(nil),  // 2: teleport.devicetrust.v1.TPMPlatformParameters
	(*TPMPlatformAttestation)(nil), // 3: teleport.devicetrust.v1.TPMPlatformAttestation
}
var file_teleport_devicetrust_v1_tpm_proto_depIdxs = []int32{
	1, // 0: teleport.devicetrust.v1.TPMPlatformParameters.quotes:type_name -> teleport.devicetrust.v1.TPMQuote
	0, // 1: teleport.devicetrust.v1.TPMPlatformParameters.pcrs:type_name -> teleport.devicetrust.v1.TPMPCR
	2, // 2: teleport.devicetrust.v1.TPMPlatformAttestation.platform_parameters:type_name -> teleport.devicetrust.v1.TPMPlatformParameters
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_teleport_devicetrust_v1_tpm_proto_init() }
func file_teleport_devicetrust_v1_tpm_proto_init() {
	if File_teleport_devicetrust_v1_tpm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_devicetrust_v1_tpm_proto_rawDesc), len(file_teleport_devicetrust_v1_tpm_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_devicetrust_v1_tpm_proto_goTypes,
		DependencyIndexes: file_teleport_devicetrust_v1_tpm_proto_depIdxs,
		MessageInfos:      file_teleport_devicetrust_v1_tpm_proto_msgTypes,
	}.Build()
	File_teleport_devicetrust_v1_tpm_proto = out.File
	file_teleport_devicetrust_v1_tpm_proto_goTypes = nil
	file_teleport_devicetrust_v1_tpm_proto_depIdxs = nil
}
