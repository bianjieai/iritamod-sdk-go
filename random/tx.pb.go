// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: random/tx.proto

package random

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_irisnet_core_sdk_go_types "github.com/irisnet/core-sdk-go/types"
	types "github.com/irisnet/core-sdk-go/types"
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

// MsgRequestRandom defines an sdk.Msg type that supports requesting a random number
type MsgRequestRandom struct {
	BlockInterval uint64                                     `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty" yaml:"block_interval"`
	Consumer      string                                     `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	Oracle        bool                                       `protobuf:"varint,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
	ServiceFeeCap github_com_irisnet_core_sdk_go_types.Coins `protobuf:"bytes,4,rep,name=service_fee_cap,json=serviceFeeCap,proto3,castrepeated=github.com/irisnet/core-sdk-go/types.Coins" json:"service_fee_cap" yaml:"service_fee_cap"`
}

func (m *MsgRequestRandom) Reset()         { *m = MsgRequestRandom{} }
func (m *MsgRequestRandom) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRandom) ProtoMessage()    {}
func (*MsgRequestRandom) Descriptor() ([]byte, []int) {
	return fileDescriptor_8734007206ce5490, []int{0}
}
func (m *MsgRequestRandom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRandom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRandom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRandom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRandom.Merge(m, src)
}
func (m *MsgRequestRandom) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRandom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRandom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRandom proto.InternalMessageInfo

func (m *MsgRequestRandom) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

func (m *MsgRequestRandom) GetConsumer() string {
	if m != nil {
		return m.Consumer
	}
	return ""
}

func (m *MsgRequestRandom) GetOracle() bool {
	if m != nil {
		return m.Oracle
	}
	return false
}

func (m *MsgRequestRandom) GetServiceFeeCap() github_com_irisnet_core_sdk_go_types.Coins {
	if m != nil {
		return m.ServiceFeeCap
	}
	return nil
}

// MsgRequestRandomResponse defines the Msg/RequestRandom response type.
type MsgRequestRandomResponse struct {
}

func (m *MsgRequestRandomResponse) Reset()         { *m = MsgRequestRandomResponse{} }
func (m *MsgRequestRandomResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRandomResponse) ProtoMessage()    {}
func (*MsgRequestRandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8734007206ce5490, []int{1}
}
func (m *MsgRequestRandomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRandomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRandomResponse.Merge(m, src)
}
func (m *MsgRequestRandomResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRandomResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestRandom)(nil), "irismod.random.MsgRequestRandom")
	proto.RegisterType((*MsgRequestRandomResponse)(nil), "irismod.random.MsgRequestRandomResponse")
}

func init() { proto.RegisterFile("random/tx.proto", fileDescriptor_8734007206ce5490) }

var fileDescriptor_8734007206ce5490 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0xe3, 0x6e, 0x55, 0x15, 0xa3, 0x6d, 0x51, 0x04, 0x55, 0x9a, 0x43, 0x12, 0xe5, 0x14,
	0x21, 0xd5, 0x56, 0xcb, 0x8d, 0x13, 0x4a, 0x25, 0x24, 0x24, 0x7a, 0xc9, 0x11, 0x0e, 0x2b, 0xc7,
	0x3b, 0x04, 0xd3, 0xc4, 0x13, 0x6c, 0xef, 0x8a, 0x3e, 0x05, 0x9c, 0x79, 0x04, 0x9e, 0xa4, 0xc7,
	0x1e, 0x39, 0x2d, 0x68, 0xf7, 0x0d, 0xfa, 0x04, 0x28, 0x7f, 0x8a, 0xd8, 0xbd, 0x70, 0xf3, 0xf8,
	0x9b, 0xf9, 0xec, 0xdf, 0x37, 0xf4, 0xd8, 0x08, 0x3d, 0xc7, 0x86, 0xbb, 0x2f, 0xac, 0x35, 0xe8,
	0xd0, 0x3f, 0x52, 0x46, 0xd9, 0x06, 0xe7, 0x6c, 0x10, 0xc2, 0xa7, 0x15, 0x56, 0xd8, 0x4b, 0xbc,
	0x3b, 0x0d, 0x5d, 0x61, 0x24, 0xd1, 0x36, 0x68, 0x79, 0x29, 0x2c, 0xf0, 0xe5, 0x79, 0x09, 0x4e,
	0x9c, 0x73, 0x89, 0x4a, 0x0f, 0x7a, 0xfa, 0x7d, 0x8f, 0x3e, 0xb9, 0xb2, 0x55, 0x01, 0x9f, 0x17,
	0x60, 0x5d, 0xd1, 0x5b, 0xf9, 0xaf, 0xe8, 0x51, 0x59, 0xa3, 0xbc, 0x9e, 0x29, 0xed, 0xc0, 0x2c,
	0x45, 0x1d, 0x90, 0x84, 0x64, 0xfb, 0xf9, 0xe9, 0xfd, 0x2a, 0x7e, 0x76, 0x23, 0x9a, 0xfa, 0x65,
	0xba, 0xad, 0xa7, 0xc5, 0xb4, 0xbf, 0x78, 0x33, 0xd6, 0x7e, 0x48, 0x0f, 0x25, 0x6a, 0xbb, 0x68,
	0xc0, 0x04, 0x7b, 0x09, 0xc9, 0x1e, 0x15, 0x7f, 0x6b, 0xff, 0x84, 0x1e, 0xa0, 0x11, 0xb2, 0x86,
	0x60, 0x92, 0x90, 0xec, 0xb0, 0x18, 0x2b, 0xff, 0x2b, 0xa1, 0xc7, 0x16, 0xcc, 0x52, 0x49, 0x98,
	0x7d, 0x00, 0x98, 0x49, 0xd1, 0x06, 0xfb, 0xc9, 0x24, 0x7b, 0x7c, 0x71, 0xca, 0x06, 0x0a, 0xd6,
	0x51, 0xb0, 0x91, 0x82, 0x5d, 0xa2, 0xd2, 0xf9, 0xdb, 0xdb, 0x55, 0xec, 0xdd, 0xaf, 0xe2, 0x93,
	0xe1, 0x5b, 0x3b, 0xf3, 0xe9, 0x8f, 0x5f, 0xf1, 0xf3, 0x4a, 0xb9, 0x8f, 0x8b, 0x92, 0x49, 0x6c,
	0x78, 0x17, 0x99, 0x06, 0xc7, 0x25, 0x1a, 0x38, 0xb3, 0xf3, 0xeb, 0xb3, 0x0a, 0xb9, 0xbb, 0x69,
	0xc1, 0xf6, 0x66, 0xb6, 0x98, 0x8e, 0xf3, 0xaf, 0x01, 0x2e, 0x45, 0x9b, 0x86, 0x34, 0xd8, 0xcd,
	0xa6, 0x00, 0xdb, 0xa2, 0xb6, 0x70, 0x51, 0xd2, 0xc9, 0x95, 0xad, 0xfc, 0xf7, 0x74, 0xba, 0x9d,
	0x5d, 0xc2, 0xb6, 0xf7, 0xc2, 0x76, 0x1d, 0xc2, 0xec, 0x7f, 0x1d, 0x0f, 0x6f, 0xe4, 0xf9, 0xed,
	0x3a, 0x22, 0x77, 0xeb, 0x88, 0xfc, 0x5e, 0x47, 0xe4, 0xdb, 0x26, 0xf2, 0xee, 0x36, 0x91, 0xf7,
	0x73, 0x13, 0x79, 0xef, 0xb2, 0x7f, 0xa0, 0x4a, 0x25, 0xf4, 0x27, 0x05, 0x42, 0x75, 0x78, 0x4e,
	0x3c, 0x70, 0x0d, 0xe6, 0xe5, 0x41, 0xbf, 0xe7, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x6b, 0x8e, 0x2a, 0x40, 0x02, 0x00, 0x00,
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
	// RequestRandom defines a method for requesting a new random number.
	RequestRandom(ctx context.Context, in *MsgRequestRandom, opts ...grpc.CallOption) (*MsgRequestRandomResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestRandom(ctx context.Context, in *MsgRequestRandom, opts ...grpc.CallOption) (*MsgRequestRandomResponse, error) {
	out := new(MsgRequestRandomResponse)
	err := c.cc.Invoke(ctx, "/irismod.random.Msg/RequestRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// RequestRandom defines a method for requesting a new random number.
	RequestRandom(context.Context, *MsgRequestRandom) (*MsgRequestRandomResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestRandom(ctx context.Context, req *MsgRequestRandom) (*MsgRequestRandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRandom not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestRandom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.random.Msg/RequestRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestRandom(ctx, req.(*MsgRequestRandom))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.random.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestRandom",
			Handler:    _Msg_RequestRandom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "random/tx.proto",
}

func (m *MsgRequestRandom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRandom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRandom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServiceFeeCap) > 0 {
		for iNdEx := len(m.ServiceFeeCap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceFeeCap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Oracle {
		i--
		if m.Oracle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Consumer) > 0 {
		i -= len(m.Consumer)
		copy(dAtA[i:], m.Consumer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Consumer)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockInterval != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.BlockInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestRandomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRandomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRandomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRequestRandom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockInterval != 0 {
		n += 1 + sovTx(uint64(m.BlockInterval))
	}
	l = len(m.Consumer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Oracle {
		n += 2
	}
	if len(m.ServiceFeeCap) > 0 {
		for _, e := range m.ServiceFeeCap {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgRequestRandomResponse) Size() (n int) {
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
func (m *MsgRequestRandom) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestRandom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRandom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInterval", wireType)
			}
			m.BlockInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consumer", wireType)
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
			m.Consumer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Oracle = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceFeeCap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceFeeCap = append(m.ServiceFeeCap, types.Coin{})
			if err := m.ServiceFeeCap[len(m.ServiceFeeCap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgRequestRandomResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestRandomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRandomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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