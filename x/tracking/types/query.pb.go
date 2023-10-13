// DONTCOVER
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/tracking/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryBlockGasTrackingRequest is the request for Query.BlockGasTracking.
type QueryBlockGasTrackingRequest struct {
}

func (m *QueryBlockGasTrackingRequest) Reset()         { *m = QueryBlockGasTrackingRequest{} }
func (m *QueryBlockGasTrackingRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasTrackingRequest) ProtoMessage()    {}
func (*QueryBlockGasTrackingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65deeabf437120fa, []int{0}
}
func (m *QueryBlockGasTrackingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasTrackingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasTrackingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasTrackingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasTrackingRequest.Merge(m, src)
}
func (m *QueryBlockGasTrackingRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasTrackingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasTrackingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasTrackingRequest proto.InternalMessageInfo

// QueryBlockGasTrackingResponse is the response for Query.BlockGasTracking.
type QueryBlockGasTrackingResponse struct {
	Block BlockTracking `protobuf:"bytes,1,opt,name=block,proto3" json:"block"`
}

func (m *QueryBlockGasTrackingResponse) Reset()         { *m = QueryBlockGasTrackingResponse{} }
func (m *QueryBlockGasTrackingResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasTrackingResponse) ProtoMessage()    {}
func (*QueryBlockGasTrackingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65deeabf437120fa, []int{1}
}
func (m *QueryBlockGasTrackingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasTrackingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasTrackingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasTrackingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasTrackingResponse.Merge(m, src)
}
func (m *QueryBlockGasTrackingResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasTrackingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasTrackingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasTrackingResponse proto.InternalMessageInfo

func (m *QueryBlockGasTrackingResponse) GetBlock() BlockTracking {
	if m != nil {
		return m.Block
	}
	return BlockTracking{}
}

func init() {
	proto.RegisterType((*QueryBlockGasTrackingRequest)(nil), "archway.tracking.v1.QueryBlockGasTrackingRequest")
	proto.RegisterType((*QueryBlockGasTrackingResponse)(nil), "archway.tracking.v1.QueryBlockGasTrackingResponse")
}

func init() { proto.RegisterFile("archway/tracking/v1/query.proto", fileDescriptor_65deeabf437120fa) }

var fileDescriptor_65deeabf437120fa = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0x4a, 0xce,
	0x28, 0x4f, 0xac, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4,
	0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0x2a, 0xd0,
	0x83, 0x29, 0xd0, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xeb, 0x83, 0x58,
	0x10, 0xa5, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89,
	0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x50, 0x59, 0x25, 0x6c, 0x36, 0xc1,
	0x0d, 0x05, 0xab, 0x51, 0x92, 0xe3, 0x92, 0x09, 0x04, 0xd9, 0xed, 0x94, 0x93, 0x9f, 0x9c, 0xed,
	0x9e, 0x58, 0x1c, 0x02, 0x95, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x8a, 0xe7, 0x92,
	0xc5, 0x21, 0x5f, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xc7, 0xc5, 0x9a, 0x04, 0x92, 0x93,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd2, 0xc3, 0xe2, 0x7a, 0x3d, 0xb0, 0x6e, 0x98, 0x56,
	0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x20, 0xda, 0x8c, 0xb6, 0x30, 0x72, 0xb1, 0x82, 0x6d,
	0x10, 0x5a, 0xc5, 0xc8, 0x25, 0x80, 0x6e, 0x8d, 0x90, 0x21, 0x56, 0xf3, 0xf0, 0x39, 0x59, 0xca,
	0x88, 0x14, 0x2d, 0x10, 0x5f, 0x28, 0xe9, 0x37, 0x5d, 0x7e, 0x32, 0x99, 0x49, 0x53, 0x48, 0x5d,
	0x1f, 0x5b, 0x98, 0x81, 0x5d, 0x1a, 0x9f, 0x9e, 0x58, 0x1c, 0x0f, 0x13, 0x75, 0xf2, 0x3d, 0xf1,
	0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8,
	0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xe3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24,
	0xbd, 0xe4, 0xfc, 0x5c, 0x98, 0x61, 0xba, 0x79, 0xa9, 0x25, 0xe5, 0xf9, 0x45, 0xd9, 0x70, 0xc3,
	0x2b, 0x10, 0xc6, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x63, 0xc3, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x68, 0xeb, 0x07, 0x96, 0x1d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// BlockGasTracking returns block gas tracking for the current block
	BlockGasTracking(ctx context.Context, in *QueryBlockGasTrackingRequest, opts ...grpc.CallOption) (*QueryBlockGasTrackingResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) BlockGasTracking(ctx context.Context, in *QueryBlockGasTrackingRequest, opts ...grpc.CallOption) (*QueryBlockGasTrackingResponse, error) {
	out := new(QueryBlockGasTrackingResponse)
	err := c.cc.Invoke(ctx, "/archway.tracking.v1.Query/BlockGasTracking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// BlockGasTracking returns block gas tracking for the current block
	BlockGasTracking(context.Context, *QueryBlockGasTrackingRequest) (*QueryBlockGasTrackingResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) BlockGasTracking(ctx context.Context, req *QueryBlockGasTrackingRequest) (*QueryBlockGasTrackingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGasTracking not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_BlockGasTracking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockGasTrackingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockGasTracking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archway.tracking.v1.Query/BlockGasTracking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockGasTracking(ctx, req.(*QueryBlockGasTrackingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "archway.tracking.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockGasTracking",
			Handler:    _Query_BlockGasTracking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archway/tracking/v1/query.proto",
}

func (m *QueryBlockGasTrackingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasTrackingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasTrackingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBlockGasTrackingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasTrackingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasTrackingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryBlockGasTrackingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBlockGasTrackingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Block.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBlockGasTrackingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBlockGasTrackingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasTrackingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBlockGasTrackingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBlockGasTrackingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasTrackingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
