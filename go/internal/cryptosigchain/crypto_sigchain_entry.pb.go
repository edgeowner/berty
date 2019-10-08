// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/crypto_sigchain_entry.proto

package cryptosigchain

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SigChainEntry_SigChainEntryType int32

const (
	SigChainEntry_SigChainEntryTypeUndefined SigChainEntry_SigChainEntryType = 0
	SigChainEntry_SigChainEntryTypeInitChain SigChainEntry_SigChainEntryType = 1
	SigChainEntry_SigChainEntryTypeAddKey    SigChainEntry_SigChainEntryType = 2
	SigChainEntry_SigChainEntryTypeRemoveKey SigChainEntry_SigChainEntryType = 3
)

var SigChainEntry_SigChainEntryType_name = map[int32]string{
	0: "SigChainEntryTypeUndefined",
	1: "SigChainEntryTypeInitChain",
	2: "SigChainEntryTypeAddKey",
	3: "SigChainEntryTypeRemoveKey",
}

var SigChainEntry_SigChainEntryType_value = map[string]int32{
	"SigChainEntryTypeUndefined": 0,
	"SigChainEntryTypeInitChain": 1,
	"SigChainEntryTypeAddKey":    2,
	"SigChainEntryTypeRemoveKey": 3,
}

func (x SigChainEntry_SigChainEntryType) String() string {
	return proto.EnumName(SigChainEntry_SigChainEntryType_name, int32(x))
}

func (SigChainEntry_SigChainEntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10e430be580c9ede, []int{0, 0}
}

type SigChainEntry struct {
	EntryHash             []byte                          `protobuf:"bytes,1,opt,name=entry_hash,json=entryHash,proto3" json:"entry_hash,omitempty"`
	EntryTypeCode         SigChainEntry_SigChainEntryType `protobuf:"varint,2,opt,name=entry_type_code,json=entryTypeCode,proto3,enum=sigchain.SigChainEntry_SigChainEntryType" json:"entry_type_code,omitempty"`
	ParentEntryHash       []byte                          `protobuf:"bytes,3,opt,name=parent_entry_hash,json=parentEntryHash,proto3" json:"parent_entry_hash,omitempty"`
	CreatedAt             time.Time                       `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	ExpiringAt            time.Time                       `protobuf:"bytes,5,opt,name=expiring_at,json=expiringAt,proto3,stdtime" json:"expiring_at"`
	SignerPublicKeyBytes  []byte                          `protobuf:"bytes,6,opt,name=signer_public_key_bytes,json=signerPublicKeyBytes,proto3" json:"signer_public_key_bytes,omitempty"`
	SubjectPublicKeyBytes []byte                          `protobuf:"bytes,7,opt,name=subject_public_key_bytes,json=subjectPublicKeyBytes,proto3" json:"subject_public_key_bytes,omitempty"`
	Signature             []byte                          `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SigChainEntry) Reset()         { *m = SigChainEntry{} }
func (m *SigChainEntry) String() string { return proto.CompactTextString(m) }
func (*SigChainEntry) ProtoMessage()    {}
func (*SigChainEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_10e430be580c9ede, []int{0}
}
func (m *SigChainEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigChainEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigChainEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigChainEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigChainEntry.Merge(m, src)
}
func (m *SigChainEntry) XXX_Size() int {
	return m.Size()
}
func (m *SigChainEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SigChainEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SigChainEntry proto.InternalMessageInfo

func (m *SigChainEntry) GetEntryHash() []byte {
	if m != nil {
		return m.EntryHash
	}
	return nil
}

func (m *SigChainEntry) GetEntryTypeCode() SigChainEntry_SigChainEntryType {
	if m != nil {
		return m.EntryTypeCode
	}
	return SigChainEntry_SigChainEntryTypeUndefined
}

func (m *SigChainEntry) GetParentEntryHash() []byte {
	if m != nil {
		return m.ParentEntryHash
	}
	return nil
}

func (m *SigChainEntry) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *SigChainEntry) GetExpiringAt() time.Time {
	if m != nil {
		return m.ExpiringAt
	}
	return time.Time{}
}

func (m *SigChainEntry) GetSignerPublicKeyBytes() []byte {
	if m != nil {
		return m.SignerPublicKeyBytes
	}
	return nil
}

func (m *SigChainEntry) GetSubjectPublicKeyBytes() []byte {
	if m != nil {
		return m.SubjectPublicKeyBytes
	}
	return nil
}

func (m *SigChainEntry) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("sigchain.SigChainEntry_SigChainEntryType", SigChainEntry_SigChainEntryType_name, SigChainEntry_SigChainEntryType_value)
	proto.RegisterType((*SigChainEntry)(nil), "sigchain.SigChainEntry")
}

func init() {
	proto.RegisterFile("internal/crypto_sigchain_entry.proto", fileDescriptor_10e430be580c9ede)
}

var fileDescriptor_10e430be580c9ede = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0xb6, 0xa5, 0x24, 0x5b, 0x4a, 0x53, 0x0b, 0x54, 0x2b, 0x80, 0x13, 0x55, 0x20, 0x05,
	0x24, 0x6c, 0xa9, 0x08, 0x71, 0x44, 0x49, 0x14, 0x09, 0xd4, 0x0b, 0x84, 0x72, 0xe1, 0x62, 0xad,
	0xed, 0xe9, 0x7a, 0x21, 0xd9, 0xb5, 0xd6, 0x63, 0xc4, 0xfe, 0x45, 0x3e, 0xab, 0xc7, 0x1e, 0x11,
	0x07, 0x40, 0xc9, 0x8f, 0x20, 0xaf, 0xeb, 0x42, 0x9b, 0x5e, 0xb8, 0xed, 0xbc, 0x37, 0xef, 0xbd,
	0xf5, 0x8c, 0x97, 0x3e, 0x16, 0x12, 0x41, 0x4b, 0x36, 0x0b, 0x13, 0x6d, 0x72, 0x54, 0x51, 0x21,
	0x78, 0x92, 0x31, 0x21, 0x23, 0x90, 0xa8, 0x4d, 0x90, 0x6b, 0x85, 0xca, 0x6d, 0x35, 0x68, 0xb7,
	0xc7, 0x95, 0xe2, 0x33, 0x08, 0x2d, 0x1e, 0x97, 0xa7, 0x21, 0x8a, 0x39, 0x14, 0xc8, 0xe6, 0x79,
	0xdd, 0xda, 0x7d, 0xce, 0x05, 0x66, 0x65, 0x1c, 0x24, 0x6a, 0x1e, 0x72, 0xc5, 0xd5, 0xdf, 0xce,
	0xaa, 0xb2, 0x85, 0x3d, 0xd5, 0xed, 0x87, 0x3f, 0xb6, 0xe8, 0xee, 0x07, 0xc1, 0xc7, 0x95, 0xf9,
	0xa4, 0x4a, 0x74, 0x1f, 0x51, 0x6a, 0xa3, 0xa3, 0x8c, 0x15, 0x99, 0x47, 0xfa, 0x64, 0x70, 0x67,
	0xda, 0xb6, 0xc8, 0x1b, 0x56, 0x64, 0xee, 0x7b, 0xba, 0x57, 0xd3, 0x68, 0x72, 0x88, 0x12, 0x95,
	0x82, 0xb7, 0xd1, 0x27, 0x83, 0xbb, 0x47, 0x4f, 0x83, 0xe6, 0x92, 0xc1, 0x15, 0xc3, 0xab, 0xd5,
	0x89, 0xc9, 0x61, 0xba, 0x0b, 0xcd, 0x71, 0xac, 0x52, 0x70, 0x9f, 0xd1, 0xfd, 0x9c, 0x69, 0x90,
	0x18, 0xfd, 0x13, 0xbc, 0x69, 0x83, 0xf7, 0x6a, 0x62, 0x72, 0x19, 0x3f, 0xa6, 0x34, 0xd1, 0xc0,
	0x10, 0xd2, 0x88, 0xa1, 0xb7, 0xd5, 0x27, 0x83, 0x9d, 0xa3, 0x6e, 0x50, 0x0f, 0x25, 0x68, 0x3e,
	0x35, 0x38, 0x69, 0x86, 0x32, 0x6a, 0x9d, 0xfd, 0xec, 0x39, 0x8b, 0x5f, 0x3d, 0x32, 0x6d, 0x5f,
	0xe8, 0x86, 0xe8, 0x4e, 0xe8, 0x0e, 0x7c, 0xcb, 0x85, 0x16, 0x92, 0x57, 0x2e, 0xb7, 0xfe, 0xc3,
	0x85, 0x36, 0xc2, 0x21, 0xba, 0x2f, 0xe9, 0x41, 0x21, 0xb8, 0x04, 0x1d, 0xe5, 0x65, 0x3c, 0x13,
	0x49, 0xf4, 0x05, 0x4c, 0x14, 0x1b, 0x84, 0xc2, 0xdb, 0xb6, 0xb7, 0xbf, 0x57, 0xd3, 0xef, 0x2c,
	0x7b, 0x0c, 0x66, 0x54, 0x71, 0xee, 0x2b, 0xea, 0x15, 0x65, 0xfc, 0x19, 0x12, 0x5c, 0xd7, 0xdd,
	0xb6, 0xba, 0xfb, 0x17, 0xfc, 0x35, 0xe1, 0x43, 0xda, 0xae, 0x0c, 0x19, 0x96, 0x1a, 0xbc, 0x56,
	0xbd, 0x98, 0x4b, 0xe0, 0x70, 0x41, 0xe8, 0xfe, 0xda, 0xa8, 0x5d, 0x9f, 0x76, 0xd7, 0xc0, 0x8f,
	0x32, 0x85, 0x53, 0x21, 0x21, 0xed, 0x38, 0x37, 0xf2, 0x6f, 0xa5, 0x40, 0x8b, 0x74, 0x88, 0xfb,
	0x80, 0x1e, 0xac, 0xf1, 0xc3, 0x34, 0x3d, 0x06, 0xd3, 0xd9, 0xb8, 0x51, 0x3c, 0x85, 0xb9, 0xfa,
	0x0a, 0x15, 0xbf, 0x39, 0x7a, 0x7d, 0xb6, 0xf4, 0xc9, 0xf9, 0xd2, 0x27, 0xbf, 0x97, 0x3e, 0x59,
	0xac, 0x7c, 0xe7, 0x7c, 0xe5, 0x3b, 0xdf, 0x57, 0xbe, 0xf3, 0xe9, 0x49, 0x0c, 0x1a, 0x4d, 0x80,
	0x90, 0x64, 0x21, 0x57, 0xe1, 0xb5, 0x47, 0xd0, 0xfc, 0x48, 0xf1, 0xb6, 0xdd, 0xc5, 0x8b, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xc4, 0x5e, 0x91, 0x26, 0x03, 0x00, 0x00,
}

func (m *SigChainEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigChainEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigChainEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SubjectPublicKeyBytes) > 0 {
		i -= len(m.SubjectPublicKeyBytes)
		copy(dAtA[i:], m.SubjectPublicKeyBytes)
		i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(len(m.SubjectPublicKeyBytes)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SignerPublicKeyBytes) > 0 {
		i -= len(m.SignerPublicKeyBytes)
		copy(dAtA[i:], m.SignerPublicKeyBytes)
		i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(len(m.SignerPublicKeyBytes)))
		i--
		dAtA[i] = 0x32
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiringAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiringAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.ParentEntryHash) > 0 {
		i -= len(m.ParentEntryHash)
		copy(dAtA[i:], m.ParentEntryHash)
		i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(len(m.ParentEntryHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EntryTypeCode != 0 {
		i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(m.EntryTypeCode))
		i--
		dAtA[i] = 0x10
	}
	if len(m.EntryHash) > 0 {
		i -= len(m.EntryHash)
		copy(dAtA[i:], m.EntryHash)
		i = encodeVarintCryptoSigchainEntry(dAtA, i, uint64(len(m.EntryHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCryptoSigchainEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovCryptoSigchainEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SigChainEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EntryHash)
	if l > 0 {
		n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	}
	if m.EntryTypeCode != 0 {
		n += 1 + sovCryptoSigchainEntry(uint64(m.EntryTypeCode))
	}
	l = len(m.ParentEntryHash)
	if l > 0 {
		n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiringAt)
	n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	l = len(m.SignerPublicKeyBytes)
	if l > 0 {
		n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	}
	l = len(m.SubjectPublicKeyBytes)
	if l > 0 {
		n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovCryptoSigchainEntry(uint64(l))
	}
	return n
}

func sovCryptoSigchainEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCryptoSigchainEntry(x uint64) (n int) {
	return sovCryptoSigchainEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SigChainEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCryptoSigchainEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SigChainEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigChainEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntryHash = append(m.EntryHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EntryHash == nil {
				m.EntryHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntryTypeCode", wireType)
			}
			m.EntryTypeCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EntryTypeCode |= SigChainEntry_SigChainEntryType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentEntryHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentEntryHash = append(m.ParentEntryHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentEntryHash == nil {
				m.ParentEntryHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiringAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiringAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerPublicKeyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerPublicKeyBytes = append(m.SignerPublicKeyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.SignerPublicKeyBytes == nil {
				m.SignerPublicKeyBytes = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectPublicKeyBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectPublicKeyBytes = append(m.SubjectPublicKeyBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.SubjectPublicKeyBytes == nil {
				m.SubjectPublicKeyBytes = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCryptoSigchainEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCryptoSigchainEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCryptoSigchainEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCryptoSigchainEntry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCryptoSigchainEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCryptoSigchainEntry
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCryptoSigchainEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCryptoSigchainEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCryptoSigchainEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCryptoSigchainEntry
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCryptoSigchainEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCryptoSigchainEntry   = fmt.Errorf("proto: integer overflow")
)
