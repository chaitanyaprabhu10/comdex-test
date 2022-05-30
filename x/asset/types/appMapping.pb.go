// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/appMapping.proto

package types

import (
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

type AppMapping struct {
	Id               uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	ShortName        string              `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty" yaml:"short_name"`
	MintGenesisToken []*MintGenesisToken `protobuf:"bytes,4,rep,name=mint_genesis_token,json=mintGenesisToken,proto3" json:"mint_genesis_token,omitempty" yaml:"mint_genesis_token"`
}

func (m *AppMapping) Reset()         { *m = AppMapping{} }
func (m *AppMapping) String() string { return proto.CompactTextString(m) }
func (*AppMapping) ProtoMessage()    {}
func (*AppMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfd2a3ad9e2bc5f, []int{0}
}
func (m *AppMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppMapping.Merge(m, src)
}
func (m *AppMapping) XXX_Size() int {
	return m.Size()
}
func (m *AppMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AppMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AppMapping proto.InternalMessageInfo

type MintGenesisToken struct {
	AssetId       uint64                                  `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	GenesisSupply *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=genesis_supply,json=genesisSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"genesis_supply,omitempty"`
	IsgovToken    bool                                    `protobuf:"varint,3,opt,name=isgovToken,proto3" json:"isgovToken,omitempty" yaml:"isgovToken"`
	Recipient     string                                  `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
}

func (m *MintGenesisToken) Reset()         { *m = MintGenesisToken{} }
func (m *MintGenesisToken) String() string { return proto.CompactTextString(m) }
func (*MintGenesisToken) ProtoMessage()    {}
func (*MintGenesisToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfd2a3ad9e2bc5f, []int{1}
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

func init() {
	proto.RegisterType((*AppMapping)(nil), "comdex.asset.v1beta1.AppMapping")
	proto.RegisterType((*MintGenesisToken)(nil), "comdex.asset.v1beta1.MintGenesisToken")
}

func init() {
	proto.RegisterFile("comdex/asset/v1beta1/appMapping.proto", fileDescriptor_2cfd2a3ad9e2bc5f)
}

var fileDescriptor_2cfd2a3ad9e2bc5f = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0x93, 0x6b, 0x04, 0x3d, 0x57, 0xb4, 0x87, 0x39, 0xa4, 0x80, 0x84, 0x73, 0x32, 0xa2,
	0x3a, 0x21, 0x35, 0x56, 0x0b, 0x2c, 0x6c, 0x64, 0x41, 0x1d, 0x8a, 0x54, 0xc3, 0xc4, 0x12, 0xe5,
	0x12, 0x37, 0xb5, 0x7a, 0xb1, 0xad, 0xb3, 0x5b, 0x71, 0x1b, 0x1f, 0x81, 0x8f, 0xc1, 0x47, 0xe9,
	0xd8, 0x11, 0x31, 0x44, 0x90, 0x5b, 0x99, 0xf2, 0x09, 0x50, 0x9c, 0xa4, 0x57, 0x4a, 0xa7, 0xd8,
	0x79, 0x7e, 0xef, 0xf3, 0xfe, 0xf1, 0x0b, 0x5e, 0xa4, 0xb2, 0xc8, 0xd8, 0x17, 0x92, 0x68, 0xcd,
	0x0c, 0xb9, 0xd8, 0x9f, 0x31, 0x93, 0xec, 0x93, 0x44, 0xa9, 0xa3, 0x44, 0x29, 0x2e, 0xf2, 0x50,
	0x2d, 0xa4, 0x91, 0x70, 0xdc, 0x62, 0xa1, 0xc5, 0xc2, 0x0e, 0x7b, 0x3a, 0xce, 0x65, 0x2e, 0x2d,
	0x40, 0x9a, 0x53, 0xcb, 0xe2, 0x3f, 0x2e, 0x00, 0xef, 0xae, 0x0d, 0xe0, 0x36, 0x18, 0xf0, 0xcc,
	0x77, 0x27, 0xee, 0xd4, 0xa3, 0x03, 0x9e, 0xc1, 0xe7, 0xc0, 0x13, 0x49, 0xc1, 0xfc, 0xc1, 0xc4,
	0x9d, 0x0e, 0xa3, 0x9d, 0xba, 0x0c, 0xb6, 0x96, 0x49, 0x31, 0x7f, 0x8b, 0x9b, 0xbf, 0x98, 0x5a,
	0x11, 0xbe, 0x06, 0x40, 0x9f, 0xca, 0x85, 0x89, 0x2d, 0xba, 0x61, 0xd1, 0xc7, 0x75, 0x19, 0x3c,
	0x6c, 0xd1, 0xb5, 0x86, 0xe9, 0xd0, 0x5e, 0x3e, 0x34, 0x51, 0x1a, 0xc0, 0x82, 0x0b, 0x13, 0xe7,
	0x4c, 0x30, 0xcd, 0x75, 0x6c, 0xe4, 0x19, 0x13, 0xbe, 0x37, 0xd9, 0x98, 0x6e, 0x1d, 0xec, 0x86,
	0x77, 0xb5, 0x10, 0x1e, 0x71, 0x61, 0xde, 0xb7, 0xf8, 0xa7, 0x86, 0x8e, 0x9e, 0xd5, 0x65, 0xf0,
	0xa4, 0xcd, 0xf2, 0xbf, 0x17, 0xa6, 0xa3, 0xe2, 0x56, 0x00, 0xfe, 0x3a, 0x00, 0xa3, 0xdb, 0x2e,
	0x30, 0x04, 0x9b, 0x36, 0x4f, 0xdc, 0xb7, 0x1e, 0x3d, 0xaa, 0xcb, 0x60, 0xa7, 0xf5, 0xed, 0x15,
	0x4c, 0xef, 0xdb, 0xe3, 0x61, 0x06, 0x8f, 0xc1, 0x76, 0x9f, 0x48, 0x9f, 0x2b, 0x35, 0x5f, 0x76,
	0xe3, 0x79, 0xf9, 0xb3, 0x0c, 0x76, 0x73, 0x6e, 0x4e, 0xcf, 0x67, 0x4d, 0xfd, 0x24, 0x95, 0xba,
	0x90, 0xba, 0xfb, 0xec, 0xe9, 0xec, 0x8c, 0x98, 0xa5, 0x62, 0x3a, 0x3c, 0x14, 0x86, 0x3e, 0xe8,
	0x1c, 0x3e, 0x5a, 0x03, 0xf8, 0x06, 0x00, 0xae, 0x73, 0x79, 0x61, 0x0b, 0xb2, 0x23, 0xdc, 0xbc,
	0x39, 0xc2, 0xb5, 0x86, 0xe9, 0x0d, 0x10, 0x1e, 0x80, 0xe1, 0x82, 0xa5, 0x5c, 0x71, 0x26, 0x8c,
	0xef, 0xd9, 0x22, 0xc6, 0x75, 0x19, 0x8c, 0xda, 0xa8, 0x6b, 0x09, 0xd3, 0x35, 0x16, 0x1d, 0x5f,
	0xfe, 0x46, 0xce, 0xf7, 0x0a, 0x39, 0x97, 0x15, 0x72, 0xaf, 0x2a, 0xe4, 0xfe, 0xaa, 0x90, 0xfb,
	0x6d, 0x85, 0x9c, 0xab, 0x15, 0x72, 0x7e, 0xac, 0x90, 0xf3, 0x99, 0xfc, 0xd3, 0x43, 0xf3, 0x0e,
	0x7b, 0xf2, 0xe4, 0x84, 0xa7, 0x3c, 0x99, 0x77, 0x77, 0xd2, 0xef, 0xa0, 0x6d, 0x68, 0x76, 0xcf,
	0xee, 0xd2, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x29, 0x45, 0x29, 0xa0, 0x02, 0x00,
	0x00,
}

func (m *AppMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintGenesisToken) > 0 {
		for iNdEx := len(m.MintGenesisToken) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintGenesisToken[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAppMapping(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ShortName) > 0 {
		i -= len(m.ShortName)
		copy(dAtA[i:], m.ShortName)
		i = encodeVarintAppMapping(dAtA, i, uint64(len(m.ShortName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAppMapping(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAppMapping(dAtA, i, uint64(m.Id))
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
		i = encodeVarintAppMapping(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x22
	}
	if m.IsgovToken {
		i--
		if m.IsgovToken {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.GenesisSupply != nil {
		{
			size := m.GenesisSupply.Size()
			i -= size
			if _, err := m.GenesisSupply.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintAppMapping(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.AssetId != 0 {
		i = encodeVarintAppMapping(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAppMapping(dAtA []byte, offset int, v uint64) int {
	offset -= sovAppMapping(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAppMapping(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAppMapping(uint64(l))
	}
	l = len(m.ShortName)
	if l > 0 {
		n += 1 + l + sovAppMapping(uint64(l))
	}
	if len(m.MintGenesisToken) > 0 {
		for _, e := range m.MintGenesisToken {
			l = e.Size()
			n += 1 + l + sovAppMapping(uint64(l))
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
		n += 1 + sovAppMapping(uint64(m.AssetId))
	}
	if m.GenesisSupply != nil {
		l = m.GenesisSupply.Size()
		n += 1 + l + sovAppMapping(uint64(l))
	}
	if m.IsgovToken {
		n += 2
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovAppMapping(uint64(l))
	}
	return n
}

func sovAppMapping(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAppMapping(x uint64) (n int) {
	return sovAppMapping(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppMapping
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
			return fmt.Errorf("proto: AppMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
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
					return ErrIntOverflowAppMapping
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
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
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
					return ErrIntOverflowAppMapping
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
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintGenesisToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
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
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintGenesisToken = append(m.MintGenesisToken, &MintGenesisToken{})
			if err := m.MintGenesisToken[len(m.MintGenesisToken)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppMapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAppMapping
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
				return ErrIntOverflowAppMapping
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
					return ErrIntOverflowAppMapping
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
					return ErrIntOverflowAppMapping
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
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.GenesisSupply = &v
			if err := m.GenesisSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsgovToken", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
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
			m.IsgovToken = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppMapping
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
				return ErrInvalidLengthAppMapping
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAppMapping
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppMapping(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAppMapping
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
func skipAppMapping(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAppMapping
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
					return 0, ErrIntOverflowAppMapping
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
					return 0, ErrIntOverflowAppMapping
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
				return 0, ErrInvalidLengthAppMapping
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAppMapping
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAppMapping
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAppMapping        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAppMapping          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAppMapping = fmt.Errorf("proto: unexpected end of group")
)
