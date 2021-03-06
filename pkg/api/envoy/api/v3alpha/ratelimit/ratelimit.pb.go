// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v3alpha/ratelimit/ratelimit.proto

package envoy_api_v3alpha_ratelimit

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
// configuration supplies a default limit for the *remote_address* key. If there is a desire to
// raise the limit for 10.0.0.1 or block it entirely it can be specified directly in the
// configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"]
//
// What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
// configured that way in the service).
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
// Like (1) we can raise/block specific IP addresses if we want with an override configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"]
//
// What it does: Limits all traffic for an authenticated client "foo"
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//
// What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
type RateLimitDescriptor struct {
	// Descriptor entries.
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa87ac9676b2311, []int{0}
}
func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(m, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	// Descriptor key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Descriptor value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa87ac9676b2311, []int{0, 0}
}
func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(m, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return m.Size()
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimitDescriptor)(nil), "envoy.api.v3alpha.ratelimit.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "envoy.api.v3alpha.ratelimit.RateLimitDescriptor.Entry")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/ratelimit/ratelimit.proto", fileDescriptor_3fa87ac9676b2311)
}

var fileDescriptor_3fa87ac9676b2311 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x2f, 0x4a,
	0x2c, 0x49, 0xcd, 0xc9, 0xcc, 0xcd, 0x2c, 0x41, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xa4, 0xc1, 0x8a, 0xf5, 0x12, 0x0b, 0x32, 0xf5, 0xa0, 0x8a, 0xf5, 0xe0, 0x4a, 0xa4, 0xc4, 0xcb,
	0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x2e, 0xa5, 0x5d, 0x8c, 0x5c,
	0xc2, 0x41, 0x89, 0x25, 0xa9, 0x3e, 0x20, 0x65, 0x2e, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25,
	0xf9, 0x45, 0x42, 0x31, 0x5c, 0xec, 0xa9, 0x79, 0x25, 0x45, 0x99, 0xa9, 0xc5, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x66, 0x7a, 0x78, 0xcc, 0xd7, 0xc3, 0x62, 0x84, 0x9e, 0x6b, 0x5e, 0x49,
	0x51, 0xa5, 0x13, 0xd7, 0xae, 0x97, 0x07, 0x98, 0x59, 0x27, 0x31, 0x32, 0x71, 0x30, 0x06, 0xc1,
	0x8c, 0x94, 0x72, 0xe5, 0x62, 0x05, 0xcb, 0x0a, 0x49, 0x73, 0x31, 0x67, 0xa7, 0x56, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x71, 0x82, 0x94, 0xb2, 0x14, 0x31, 0x29, 0x30, 0x06, 0x81, 0x44,
	0x85, 0xe4, 0xb9, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xd0, 0xa5, 0x21, 0xe2, 0x4e,
	0x9e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x23, 0x97, 0x66,
	0x66, 0x3e, 0xc4, 0x8d, 0x05, 0x45, 0xf9, 0x15, 0x95, 0xf8, 0x9c, 0xeb, 0xc4, 0x17, 0x04, 0x63,
	0x06, 0x80, 0x42, 0x21, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x1c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc0, 0x64, 0x6d, 0xd0, 0x73, 0x01, 0x00, 0x00,
}

func (m *RateLimitDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRatelimit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RateLimitDescriptor_Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimitDescriptor_Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimitDescriptor_Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintRatelimit(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRatelimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovRatelimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RateLimitDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovRatelimit(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RateLimitDescriptor_Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRatelimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRatelimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRatelimit(x uint64) (n int) {
	return sovRatelimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimitDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
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
			return fmt.Errorf("proto: RateLimitDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimitDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
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
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &RateLimitDescriptor_Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RateLimitDescriptor_Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRatelimit
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRatelimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRatelimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRatelimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRatelimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRatelimit
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRatelimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRatelimit
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
					return 0, ErrIntOverflowRatelimit
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
					return 0, ErrIntOverflowRatelimit
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
				return 0, ErrInvalidLengthRatelimit
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRatelimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRatelimit
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
				next, err := skipRatelimit(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRatelimit
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
	ErrInvalidLengthRatelimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRatelimit   = fmt.Errorf("proto: integer overflow")
)
