// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/tokenmint/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//Will become governance proposal- will trigger token minting & sending
type MsgMintNewTokensRequest struct {
	From         string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	AppMappingId uint64 `protobuf:"varint,2,opt,name=app_mapping_id,json=appMappingId,proto3" json:"app_mapping_id,omitempty" yaml:"app_mapping_id"`
	AssetId      uint64 `protobuf:"varint,3,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
}

func (m *MsgMintNewTokensRequest) Reset()         { *m = MsgMintNewTokensRequest{} }
func (m *MsgMintNewTokensRequest) String() string { return proto.CompactTextString(m) }
func (*MsgMintNewTokensRequest) ProtoMessage()    {}
func (*MsgMintNewTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_877e0eff92180c18, []int{0}
}
func (m *MsgMintNewTokensRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintNewTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintNewTokensRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintNewTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintNewTokensRequest.Merge(m, src)
}
func (m *MsgMintNewTokensRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintNewTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintNewTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintNewTokensRequest proto.InternalMessageInfo

type MsgMintNewTokensResponse struct {
}

func (m *MsgMintNewTokensResponse) Reset()         { *m = MsgMintNewTokensResponse{} }
func (m *MsgMintNewTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintNewTokensResponse) ProtoMessage()    {}
func (*MsgMintNewTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_877e0eff92180c18, []int{1}
}
func (m *MsgMintNewTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintNewTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintNewTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintNewTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintNewTokensResponse.Merge(m, src)
}
func (m *MsgMintNewTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintNewTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintNewTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintNewTokensResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMintNewTokensRequest)(nil), "comdex.tokenmint.v1beta1.MsgMintNewTokensRequest")
	proto.RegisterType((*MsgMintNewTokensResponse)(nil), "comdex.tokenmint.v1beta1.MsgMintNewTokensResponse")
}

func init() { proto.RegisterFile("comdex/tokenmint/v1beta1/tx.proto", fileDescriptor_877e0eff92180c18) }

var fileDescriptor_877e0eff92180c18 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xc2, 0x40,
	0x18, 0xc7, 0x7b, 0x42, 0x44, 0x4f, 0x22, 0xa6, 0xd1, 0xd8, 0x30, 0x5c, 0xb1, 0x2e, 0x38, 0xd8,
	0x06, 0x98, 0x74, 0x83, 0xcd, 0xa1, 0x9a, 0x34, 0xba, 0xb8, 0x90, 0x03, 0x8e, 0x7a, 0x91, 0xf6,
	0x4e, 0xee, 0x50, 0xd8, 0x7c, 0x04, 0x1f, 0xc3, 0x47, 0x21, 0x4e, 0x8c, 0x4e, 0x8d, 0xb6, 0x6f,
	0xc0, 0x13, 0x98, 0xf6, 0x1a, 0x83, 0x1a, 0x06, 0xb7, 0xbb, 0xef, 0xfb, 0xfd, 0x7f, 0xb9, 0xfb,
	0x3e, 0x78, 0xd4, 0x67, 0xc1, 0x80, 0x4c, 0x1d, 0xc9, 0xee, 0x49, 0x18, 0xd0, 0x50, 0x3a, 0x8f,
	0x8d, 0x1e, 0x91, 0xb8, 0xe1, 0xc8, 0xa9, 0xcd, 0xc7, 0x4c, 0x32, 0xdd, 0x50, 0x88, 0xfd, 0x8d,
	0xd8, 0x39, 0x52, 0xdd, 0xf7, 0x99, 0xcf, 0x32, 0xc8, 0x49, 0x4f, 0x8a, 0xb7, 0xde, 0x00, 0x3c,
	0x74, 0x85, 0xef, 0xd2, 0x50, 0x5e, 0x92, 0xa7, 0xeb, 0x34, 0x25, 0x3c, 0xf2, 0x30, 0x21, 0x42,
	0xea, 0xc7, 0xb0, 0x38, 0x1c, 0xb3, 0xc0, 0x00, 0x35, 0x50, 0xdf, 0xee, 0x54, 0x96, 0x91, 0xb9,
	0x33, 0xc3, 0xc1, 0xe8, 0xdc, 0x4a, 0xab, 0x96, 0x97, 0x35, 0xf5, 0x2b, 0xb8, 0x8b, 0x39, 0xef,
	0x06, 0x98, 0x73, 0x1a, 0xfa, 0x5d, 0x3a, 0x30, 0x36, 0x6a, 0xa0, 0x5e, 0xec, 0x9c, 0xc4, 0x91,
	0x59, 0x6e, 0x73, 0xee, 0xaa, 0xc6, 0xc5, 0x60, 0x19, 0x99, 0x07, 0x2a, 0xfe, 0x93, 0xb7, 0xbc,
	0x32, 0x5e, 0xc1, 0xf4, 0x33, 0xb8, 0x85, 0x85, 0x20, 0x32, 0x55, 0x15, 0x32, 0x15, 0x8a, 0x23,
	0xb3, 0xd4, 0x4e, 0x6b, 0x99, 0xa5, 0x92, 0x5b, 0x72, 0xc8, 0xf2, 0x4a, 0x58, 0xf5, 0xac, 0x2a,
	0x34, 0xfe, 0xfe, 0x45, 0x70, 0x16, 0x0a, 0xd2, 0x7c, 0x06, 0xb0, 0xe0, 0x0a, 0x5f, 0x9f, 0xc1,
	0xbd, 0xdf, 0x8c, 0xde, 0xb0, 0xd7, 0x4d, 0xcd, 0x5e, 0x33, 0x9b, 0x6a, 0xf3, 0x3f, 0x11, 0xf5,
	0x84, 0xce, 0xcd, 0xfc, 0x13, 0x69, 0xaf, 0x31, 0xd2, 0xe6, 0x31, 0x02, 0x8b, 0x18, 0x81, 0x8f,
	0x18, 0x81, 0x97, 0x04, 0x69, 0x8b, 0x04, 0x69, 0xef, 0x09, 0xd2, 0x6e, 0x5b, 0x3e, 0x95, 0x77,
	0x93, 0x5e, 0xea, 0x76, 0x94, 0xff, 0x94, 0x0d, 0x87, 0xb4, 0x4f, 0xf1, 0x28, 0xbf, 0x3b, 0xab,
	0xdb, 0x97, 0x33, 0x4e, 0x44, 0x6f, 0x33, 0xdb, 0x64, 0xeb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xf2, 0xcb, 0xd3, 0x1e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	MsgMintNewTokens(ctx context.Context, in *MsgMintNewTokensRequest, opts ...grpc.CallOption) (*MsgMintNewTokensResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MsgMintNewTokens(ctx context.Context, in *MsgMintNewTokensRequest, opts ...grpc.CallOption) (*MsgMintNewTokensResponse, error) {
	out := new(MsgMintNewTokensResponse)
	err := c.cc.Invoke(ctx, "/comdex.tokenmint.v1beta1.Msg/MsgMintNewTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MsgMintNewTokens(context.Context, *MsgMintNewTokensRequest) (*MsgMintNewTokensResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MsgMintNewTokens(ctx context.Context, req *MsgMintNewTokensRequest) (*MsgMintNewTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgMintNewTokens not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MsgMintNewTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintNewTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgMintNewTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.tokenmint.v1beta1.Msg/MsgMintNewTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgMintNewTokens(ctx, req.(*MsgMintNewTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.tokenmint.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgMintNewTokens",
			Handler:    _Msg_MsgMintNewTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/tokenmint/v1beta1/tx.proto",
}

func (m *MsgMintNewTokensRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintNewTokensRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintNewTokensRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x18
	}
	if m.AppMappingId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AppMappingId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintNewTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintNewTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintNewTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMintNewTokensRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.AppMappingId != 0 {
		n += 1 + sovTx(uint64(m.AppMappingId))
	}
	if m.AssetId != 0 {
		n += 1 + sovTx(uint64(m.AssetId))
	}
	return n
}

func (m *MsgMintNewTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMintNewTokensRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintNewTokensRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintNewTokensRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppMappingId", wireType)
			}
			m.AppMappingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppMappingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMintNewTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintNewTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintNewTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
