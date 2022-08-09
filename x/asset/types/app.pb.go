// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/app.proto

package types

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type AppData struct {
	Id               uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	ShortName        string                                 `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty" yaml:"short_name"`
	MinGovDeposit    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=min_gov_deposit,json=minGovDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_gov_deposit" yaml:"min_gov_deposit"`
	GovTimeInSeconds float64                                `protobuf:"fixed64,5,opt,name=gov_time_in_seconds,json=govTimeInSeconds,proto3" json:"gov_time_in_seconds,omitempty" yaml:"gov_time_in_seconds"`
	GenesisToken     []MintGenesisToken                     `protobuf:"bytes,6,rep,name=genesis_token,json=genesisToken,proto3" json:"genesis_token" yaml:"genesis_token"`
}

func (m *AppData) Reset()         { *m = AppData{} }
func (m *AppData) String() string { return proto.CompactTextString(m) }
func (*AppData) ProtoMessage()    {}
func (*AppData) Descriptor() ([]byte, []int) {
	return fileDescriptor_1372b4734b6486fd, []int{0}
}
func (m *AppData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppData.Merge(m, src)
}
func (m *AppData) XXX_Size() int {
	return m.Size()
}
func (m *AppData) XXX_DiscardUnknown() {
	xxx_messageInfo_AppData.DiscardUnknown(m)
}

var xxx_messageInfo_AppData proto.InternalMessageInfo

type MintGenesisToken struct {
	AssetId       uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	GenesisSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=genesis_supply,json=genesisSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"genesis_supply"`
	IsGovToken    bool                                   `protobuf:"varint,3,opt,name=is_gov_token,json=isGovToken,proto3" json:"is_gov_token,omitempty" yaml:"is_gov_token"`
	Recipient     string                                 `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
}

func (m *MintGenesisToken) Reset()         { *m = MintGenesisToken{} }
func (m *MintGenesisToken) String() string { return proto.CompactTextString(m) }
func (*MintGenesisToken) ProtoMessage()    {}
func (*MintGenesisToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_1372b4734b6486fd, []int{1}
}
func (m *MintGenesisToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintGenesisToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintGenesisToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintGenesisToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintGenesisToken.Merge(m, src)
}
func (m *MintGenesisToken) XXX_Size() int {
	return m.Size()
}
func (m *MintGenesisToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MintGenesisToken.DiscardUnknown(m)
}

var xxx_messageInfo_MintGenesisToken proto.InternalMessageInfo

type AppAndGovTime struct {
	AppId            uint64  `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	GovTimeInSeconds float64 `protobuf:"fixed64,2,opt,name=gov_time_in_seconds,json=govTimeInSeconds,proto3" json:"gov_time_in_seconds,omitempty" yaml:"gov_time_in_seconds"`
}

func (m *AppAndGovTime) Reset()         { *m = AppAndGovTime{} }
func (m *AppAndGovTime) String() string { return proto.CompactTextString(m) }
func (*AppAndGovTime) ProtoMessage()    {}
func (*AppAndGovTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_1372b4734b6486fd, []int{2}
}
func (m *AppAndGovTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppAndGovTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppAndGovTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppAndGovTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppAndGovTime.Merge(m, src)
}
func (m *AppAndGovTime) XXX_Size() int {
	return m.Size()
}
func (m *AppAndGovTime) XXX_DiscardUnknown() {
	xxx_messageInfo_AppAndGovTime.DiscardUnknown(m)
}

var xxx_messageInfo_AppAndGovTime proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AppData)(nil), "comdex.asset.v1beta1.AppData")
	proto.RegisterType((*MintGenesisToken)(nil), "comdex.asset.v1beta1.MintGenesisToken")
	proto.RegisterType((*AppAndGovTime)(nil), "comdex.asset.v1beta1.AppAndGovTime")
}

func init() { proto.RegisterFile("comdex/asset/v1beta1/app.proto", fileDescriptor_1372b4734b6486fd) }

var fileDescriptor_1372b4734b6486fd = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0xd3, 0x36, 0x6d, 0xb6, 0x4d, 0x93, 0x3a, 0x01, 0xa2, 0x0a, 0xd9, 0xd1, 0x22, 0x55,
	0xb9, 0xd4, 0x56, 0x0b, 0x17, 0xb8, 0xc5, 0xaa, 0x14, 0x72, 0x28, 0x12, 0x6e, 0xb9, 0x70, 0xb1,
	0x1c, 0x7b, 0xeb, 0xae, 0x1a, 0xef, 0xae, 0xb2, 0xdb, 0x88, 0xfc, 0x41, 0x8f, 0x7c, 0x02, 0x47,
	0x3e, 0x25, 0xc7, 0x1e, 0x11, 0x07, 0x0b, 0x92, 0x3f, 0xf0, 0x17, 0x20, 0xef, 0xda, 0x24, 0x54,
	0xe5, 0x00, 0x27, 0xef, 0xce, 0x7b, 0xf3, 0xc6, 0xf3, 0x66, 0x16, 0x98, 0x21, 0x4d, 0x22, 0xf4,
	0xc9, 0x09, 0x38, 0x47, 0xc2, 0x99, 0x9e, 0x8c, 0x90, 0x08, 0x4e, 0x9c, 0x80, 0x31, 0x9b, 0x4d,
	0xa8, 0xa0, 0x46, 0x5b, 0xe1, 0xb6, 0xc4, 0xed, 0x02, 0x3f, 0x6c, 0xc7, 0x34, 0xa6, 0x92, 0xe0,
	0xe4, 0x27, 0xc5, 0x85, 0x5f, 0x36, 0xc0, 0x76, 0x9f, 0xb1, 0xb3, 0x40, 0x04, 0xc6, 0x3e, 0xa8,
	0xe0, 0xa8, 0xa3, 0x77, 0xf5, 0xde, 0xa6, 0x57, 0xc1, 0x91, 0xf1, 0x02, 0x6c, 0x92, 0x20, 0x41,
	0x9d, 0x4a, 0x57, 0xef, 0xd5, 0xdc, 0x46, 0x96, 0x5a, 0xbb, 0xb3, 0x20, 0x19, 0xbf, 0x81, 0x79,
	0x14, 0x7a, 0x12, 0x34, 0x5e, 0x01, 0xc0, 0xaf, 0xe9, 0x44, 0xf8, 0x92, 0xba, 0x21, 0xa9, 0x4f,
	0xb2, 0xd4, 0x3a, 0x50, 0xd4, 0x15, 0x06, 0xbd, 0x9a, 0xbc, 0xbc, 0xcb, 0xb3, 0x18, 0x68, 0x24,
	0x98, 0xf8, 0x31, 0x9d, 0xfa, 0x11, 0x62, 0x94, 0x63, 0xd1, 0xd9, 0x94, 0xa9, 0x6f, 0xe7, 0xa9,
	0xa5, 0x7d, 0x4f, 0xad, 0xa3, 0x18, 0x8b, 0xeb, 0xdb, 0x91, 0x1d, 0xd2, 0xc4, 0x09, 0x29, 0x4f,
	0x28, 0x2f, 0x3e, 0xc7, 0x3c, 0xba, 0x71, 0xc4, 0x8c, 0x21, 0x6e, 0x0f, 0x89, 0xc8, 0x52, 0xeb,
	0xa9, 0x2a, 0xf4, 0x40, 0x0e, 0x7a, 0xf5, 0x04, 0x93, 0x01, 0x9d, 0x9e, 0xa9, 0xbb, 0x71, 0x0e,
	0x5a, 0x39, 0x2c, 0x70, 0x82, 0x7c, 0x4c, 0x7c, 0x8e, 0x42, 0x4a, 0x22, 0xde, 0xd9, 0xea, 0xea,
	0x3d, 0xdd, 0x35, 0xb3, 0xd4, 0x3a, 0x54, 0x3a, 0x8f, 0x90, 0xa0, 0xd7, 0x8c, 0xe9, 0xf4, 0x12,
	0x27, 0x68, 0x48, 0x2e, 0x54, 0xc8, 0xc0, 0xa0, 0x1e, 0x23, 0x82, 0x38, 0xe6, 0xbe, 0xa0, 0x37,
	0x88, 0x74, 0xaa, 0xdd, 0x8d, 0xde, 0xee, 0xe9, 0x91, 0xfd, 0x98, 0xf7, 0xf6, 0x39, 0x26, 0x62,
	0xa0, 0xe8, 0x97, 0x39, 0xdb, 0x7d, 0x9e, 0xb7, 0x99, 0xa5, 0x56, 0xbb, 0x28, 0xba, 0x2e, 0x05,
	0xbd, 0xbd, 0x78, 0x8d, 0x0b, 0xef, 0x2a, 0xa0, 0xf9, 0x50, 0xc0, 0xb0, 0xc1, 0x8e, 0x2c, 0xe1,
	0x97, 0x13, 0x73, 0x5b, 0x59, 0x6a, 0x35, 0x94, 0x5c, 0x89, 0x40, 0x6f, 0x5b, 0x1e, 0x87, 0x91,
	0xf1, 0x01, 0xec, 0x97, 0x45, 0xf8, 0x2d, 0x63, 0xe3, 0x59, 0x31, 0x55, 0xfb, 0xdf, 0xfc, 0xf6,
	0xca, 0xae, 0x2f, 0xa4, 0x88, 0xf1, 0x1a, 0xec, 0x61, 0x2e, 0x7d, 0x57, 0x2e, 0xe4, 0xf3, 0xdf,
	0x71, 0x9f, 0x65, 0xa9, 0xd5, 0x52, 0xbf, 0xb2, 0x8e, 0x42, 0x0f, 0x60, 0x3e, 0xa0, 0x53, 0xd5,
	0xc1, 0x29, 0xa8, 0x4d, 0x50, 0x88, 0x19, 0x46, 0xa4, 0x1c, 0x7e, 0x3b, 0x4b, 0xad, 0xa6, 0xca,
	0xfb, 0x0d, 0x41, 0x6f, 0x45, 0x83, 0x77, 0x3a, 0xa8, 0xf7, 0x19, 0xeb, 0x93, 0x68, 0xa0, 0x06,
	0x62, 0xf4, 0x40, 0x35, 0x60, 0x6c, 0xe5, 0xc2, 0x41, 0x96, 0x5a, 0xf5, 0xc2, 0x05, 0x19, 0x87,
	0xde, 0x56, 0xc0, 0xd8, 0x30, 0xfa, 0xdb, 0x02, 0x54, 0xfe, 0x6f, 0x01, 0xdc, 0xf7, 0xf3, 0x9f,
	0xa6, 0xf6, 0x75, 0x61, 0x6a, 0xf3, 0x85, 0xa9, 0xdf, 0x2f, 0x4c, 0xfd, 0xc7, 0xc2, 0xd4, 0x3f,
	0x2f, 0x4d, 0xed, 0x7e, 0x69, 0x6a, 0xdf, 0x96, 0xa6, 0xf6, 0xd1, 0xf9, 0xc3, 0xd2, 0x7c, 0x2b,
	0x8e, 0xe9, 0xd5, 0x15, 0x0e, 0x71, 0x30, 0x2e, 0xee, 0x4e, 0xf9, 0x86, 0xa5, 0xbf, 0xa3, 0xaa,
	0x7c, 0x92, 0x2f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x24, 0x02, 0x69, 0xe0, 0x03, 0x00,
	0x00,
}

func (m *AppData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisToken) > 0 {
		for iNdEx := len(m.GenesisToken) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisToken[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.GovTimeInSeconds != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.GovTimeInSeconds))))
		i--
		dAtA[i] = 0x29
	}
	{
		size := m.MinGovDeposit.Size()
		i -= size
		if _, err := m.MinGovDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintApp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ShortName) > 0 {
		i -= len(m.ShortName)
		copy(dAtA[i:], m.ShortName)
		i = encodeVarintApp(dAtA, i, uint64(len(m.ShortName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintApp(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MintGenesisToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintGenesisToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintGenesisToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintApp(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x22
	}
	if m.IsGovToken {
		i--
		if m.IsGovToken {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.GenesisSupply.Size()
		i -= size
		if _, err := m.GenesisSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintApp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetId != 0 {
		i = encodeVarintApp(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AppAndGovTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppAndGovTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppAndGovTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GovTimeInSeconds != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.GovTimeInSeconds))))
		i--
		dAtA[i] = 0x11
	}
	if m.AppId != 0 {
		i = encodeVarintApp(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApp(dAtA []byte, offset int, v uint64) int {
	offset -= sovApp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApp(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = len(m.ShortName)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	l = m.MinGovDeposit.Size()
	n += 1 + l + sovApp(uint64(l))
	if m.GovTimeInSeconds != 0 {
		n += 9
	}
	if len(m.GenesisToken) > 0 {
		for _, e := range m.GenesisToken {
			l = e.Size()
			n += 1 + l + sovApp(uint64(l))
		}
	}
	return n
}

func (m *MintGenesisToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovApp(uint64(m.AssetId))
	}
	l = m.GenesisSupply.Size()
	n += 1 + l + sovApp(uint64(l))
	if m.IsGovToken {
		n += 2
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovApp(uint64(l))
	}
	return n
}

func (m *AppAndGovTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AppId != 0 {
		n += 1 + sovApp(uint64(m.AppId))
	}
	if m.GovTimeInSeconds != 0 {
		n += 9
	}
	return n
}

func sovApp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApp(x uint64) (n int) {
	return sovApp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
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
			return fmt.Errorf("proto: AppData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGovDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinGovDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovTimeInSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.GovTimeInSeconds = float64(math.Float64frombits(v))
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisToken = append(m.GenesisToken, MintGenesisToken{})
			if err := m.GenesisToken[len(m.GenesisToken)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApp
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
func (m *MintGenesisToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
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
			return fmt.Errorf("proto: MintGenesisToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintGenesisToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsGovToken", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsGovToken = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
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
				return ErrInvalidLengthApp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApp
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
func (m *AppAndGovTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApp
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
			return fmt.Errorf("proto: AppAndGovTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppAndGovTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovTimeInSeconds", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.GovTimeInSeconds = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipApp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApp
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
func skipApp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApp
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
					return 0, ErrIntOverflowApp
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
					return 0, ErrIntOverflowApp
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
				return 0, ErrInvalidLengthApp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApp = fmt.Errorf("proto: unexpected end of group")
)