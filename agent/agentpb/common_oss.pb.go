// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common_oss.proto

package agentpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EnterpriseMeta struct {
}

func (m *EnterpriseMeta) Reset()         { *m = EnterpriseMeta{} }
func (m *EnterpriseMeta) String() string { return proto.CompactTextString(m) }
func (*EnterpriseMeta) ProtoMessage()    {}
func (*EnterpriseMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcf35e841fcc50ea, []int{0}
}
func (m *EnterpriseMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnterpriseMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnterpriseMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnterpriseMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterpriseMeta.Merge(m, src)
}
func (m *EnterpriseMeta) XXX_Size() int {
	return m.Size()
}
func (m *EnterpriseMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterpriseMeta.DiscardUnknown(m)
}

var xxx_messageInfo_EnterpriseMeta proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EnterpriseMeta)(nil), "agentpb.EnterpriseMeta")
}

func init() { proto.RegisterFile("common_oss.proto", fileDescriptor_bcf35e841fcc50ea) }

var fileDescriptor_bcf35e841fcc50ea = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0xcf, 0x2f, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4c,
	0x4f, 0xcd, 0x2b, 0x29, 0x48, 0x92, 0x92, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b,
	0x27, 0x95, 0xa6, 0xe9, 0xa7, 0x94, 0x16, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x41, 0x14, 0x4a, 0x89,
	0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99, 0xfa, 0x20, 0x16, 0x44, 0x54, 0x49, 0x80, 0x8b, 0xcf, 0x35,
	0xaf, 0x24, 0xb5, 0xa8, 0xa0, 0x28, 0xb3, 0x38, 0xd5, 0x37, 0xb5, 0x24, 0xd1, 0x49, 0xe1, 0xc4,
	0x43, 0x39, 0x86, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x19, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0x89, 0x0d, 0xac, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xbf, 0x18, 0x0f, 0x8d,
	0x00, 0x00, 0x00,
}

func (m *EnterpriseMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnterpriseMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintCommonOss(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EnterpriseMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovCommonOss(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommonOss(x uint64) (n int) {
	return sovCommonOss(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnterpriseMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommonOss
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
			return fmt.Errorf("proto: EnterpriseMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnterpriseMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommonOss(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommonOss
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommonOss
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
func skipCommonOss(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommonOss
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
					return 0, ErrIntOverflowCommonOss
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
					return 0, ErrIntOverflowCommonOss
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
				return 0, ErrInvalidLengthCommonOss
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCommonOss
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommonOss
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
				next, err := skipCommonOss(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCommonOss
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
	ErrInvalidLengthCommonOss = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommonOss   = fmt.Errorf("proto: integer overflow")
)
