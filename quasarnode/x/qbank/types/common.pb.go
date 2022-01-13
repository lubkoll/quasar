// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qbank/common.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QCoins struct {
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *QCoins) Reset()         { *m = QCoins{} }
func (m *QCoins) String() string { return proto.CompactTextString(m) }
func (*QCoins) ProtoMessage()    {}
func (*QCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_910d3f484eab02ad, []int{0}
}
func (m *QCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QCoins.Merge(m, src)
}
func (m *QCoins) XXX_Size() int {
	return m.Size()
}
func (m *QCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_QCoins.DiscardUnknown(m)
}

var xxx_messageInfo_QCoins proto.InternalMessageInfo

func (m *QCoins) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

type QDenoms struct {
	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *QDenoms) Reset()         { *m = QDenoms{} }
func (m *QDenoms) String() string { return proto.CompactTextString(m) }
func (*QDenoms) ProtoMessage()    {}
func (*QDenoms) Descriptor() ([]byte, []int) {
	return fileDescriptor_910d3f484eab02ad, []int{1}
}
func (m *QDenoms) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QDenoms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QDenoms.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QDenoms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QDenoms.Merge(m, src)
}
func (m *QDenoms) XXX_Size() int {
	return m.Size()
}
func (m *QDenoms) XXX_DiscardUnknown() {
	xxx_messageInfo_QDenoms.DiscardUnknown(m)
}

var xxx_messageInfo_QDenoms proto.InternalMessageInfo

func (m *QDenoms) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func init() {
	proto.RegisterType((*QCoins)(nil), "abag.quasarnode.qbank.QCoins")
	proto.RegisterType((*QDenoms)(nil), "abag.quasarnode.qbank.QDenoms")
}

func init() { proto.RegisterFile("qbank/common.proto", fileDescriptor_910d3f484eab02ad) }

var fileDescriptor_910d3f484eab02ad = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x14, 0x84, 0x13, 0x01, 0x41, 0x98, 0x5d, 0x04, 0x12, 0x54, 0xc8, 0x45, 0xdd, 0xd0, 0x0d, 0x7e,
	0x14, 0x6e, 0x10, 0xb8, 0x40, 0xbb, 0x64, 0x67, 0x27, 0x56, 0x88, 0x22, 0xfb, 0xb5, 0x7d, 0x29,
	0x82, 0x5b, 0x70, 0x0e, 0x4e, 0xd2, 0x65, 0x97, 0xac, 0x00, 0x25, 0x17, 0x41, 0xfe, 0x91, 0xe8,
	0xca, 0xa3, 0xa7, 0x99, 0x6f, 0xe4, 0x61, 0xf9, 0x4a, 0x49, 0xdb, 0x42, 0x89, 0xc6, 0xa0, 0x15,
	0xcb, 0x35, 0x76, 0x98, 0x9f, 0x4b, 0x25, 0x6b, 0xb1, 0xda, 0x48, 0x92, 0x6b, 0x8b, 0x95, 0x16,
	0xde, 0x33, 0xe2, 0x25, 0x92, 0x41, 0x02, 0x25, 0x49, 0xc3, 0xeb, 0x4c, 0xe9, 0x4e, 0xce, 0xa0,
	0xc4, 0x26, 0xc6, 0x46, 0x67, 0x35, 0xd6, 0xe8, 0x25, 0x38, 0x15, 0xae, 0x93, 0x96, 0x65, 0xf3,
	0x47, 0x6c, 0x2c, 0xe5, 0x92, 0x1d, 0x39, 0x37, 0x5d, 0xa4, 0xd7, 0x07, 0xd3, 0xd3, 0xfb, 0x4b,
	0x11, 0x78, 0xc2, 0xf1, 0x44, 0xe4, 0x09, 0x67, 0x2d, 0xee, 0xb6, 0xdf, 0xe3, 0xe4, 0xf3, 0x67,
	0x3c, 0xad, 0x9b, 0xee, 0x65, 0xa3, 0x44, 0x89, 0x06, 0x62, 0x79, 0x78, 0x6e, 0xa9, 0x6a, 0xa1,
	0x7b, 0x5f, 0x6a, 0xf2, 0x01, 0x5a, 0x04, 0xf2, 0xe4, 0x86, 0x1d, 0xcf, 0x9f, 0xb4, 0x45, 0x43,
	0xf9, 0x15, 0xcb, 0x2a, 0xaf, 0x7c, 0xdd, 0x49, 0x71, 0xe8, 0x98, 0x8b, 0x78, 0x2b, 0x8a, 0x6d,
	0xcf, 0xd3, 0x5d, 0xcf, 0xd3, 0xdf, 0x9e, 0xa7, 0x1f, 0x03, 0x4f, 0x76, 0x03, 0x4f, 0xbe, 0x06,
	0x9e, 0x3c, 0xef, 0x77, 0xba, 0x1d, 0xe0, 0x7f, 0x07, 0x78, 0x83, 0xb0, 0x96, 0x6f, 0x56, 0x99,
	0xff, 0xe0, 0xc3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xce, 0x82, 0x25, 0x43, 0x01, 0x00,
	0x00,
}

func (m *QCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QDenoms) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QDenoms) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QDenoms) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func (m *QDenoms) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QCoins) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: QCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *QDenoms) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: QDenoms: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QDenoms: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
