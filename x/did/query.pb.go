// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/v1/query.proto

package did

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.
type QueryDidDocumentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryDidDocumentRequest) Reset()         { *m = QueryDidDocumentRequest{} }
func (m *QueryDidDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentRequest) ProtoMessage()    {}
func (*QueryDidDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1fa9bb626e2869, []int{0}
}
func (m *QueryDidDocumentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentRequest.Merge(m, src)
}
func (m *QueryDidDocumentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentRequest proto.InternalMessageInfo

func (m *QueryDidDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method
type QueryDidDocumentResponse struct {
	// Returns a did document
	DidDocument DidDocument `protobuf:"bytes,1,opt,name=didDocument,proto3" json:"didDocument"`
}

func (m *QueryDidDocumentResponse) Reset()         { *m = QueryDidDocumentResponse{} }
func (m *QueryDidDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentResponse) ProtoMessage()    {}
func (*QueryDidDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1fa9bb626e2869, []int{1}
}
func (m *QueryDidDocumentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentResponse.Merge(m, src)
}
func (m *QueryDidDocumentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentResponse proto.InternalMessageInfo

func (m *QueryDidDocumentResponse) GetDidDocument() DidDocument {
	if m != nil {
		return m.DidDocument
	}
	return DidDocument{}
}

func init() {
	proto.RegisterType((*QueryDidDocumentRequest)(nil), "elestodao.elesto.did.v1.QueryDidDocumentRequest")
	proto.RegisterType((*QueryDidDocumentResponse)(nil), "elestodao.elesto.did.v1.QueryDidDocumentResponse")
}

func init() { proto.RegisterFile("did/v1/query.proto", fileDescriptor_ae1fa9bb626e2869) }

var fileDescriptor_ae1fa9bb626e2869 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x2f, 0x45, 0x05, 0x53, 0x10, 0x09, 0x4a, 0x4b, 0xd1, 0x53, 0x8a, 0x82, 0x0a, 0x26,
	0x5e, 0x75, 0x17, 0x4a, 0x47, 0x17, 0x3b, 0xba, 0xe5, 0x9a, 0x70, 0x0d, 0xb4, 0x79, 0xd7, 0x26,
	0x77, 0x28, 0xe2, 0xe2, 0x27, 0x50, 0x1c, 0xfd, 0x42, 0x1d, 0x0b, 0x2e, 0x4e, 0x22, 0xad, 0x1f,
	0x44, 0x72, 0xa9, 0x58, 0x90, 0x1b, 0x1c, 0x02, 0x8f, 0xc7, 0xff, 0xf7, 0x0b, 0xff, 0x87, 0x89,
	0x50, 0x82, 0xe5, 0x11, 0x1b, 0x65, 0x72, 0x7c, 0x47, 0xd3, 0x31, 0x58, 0x20, 0x35, 0x39, 0x90,
	0xc6, 0x82, 0xe0, 0x40, 0xfd, 0x44, 0x85, 0x12, 0x34, 0x8f, 0x1a, 0x3b, 0x09, 0x40, 0x32, 0x90,
	0x8c, 0xa7, 0x8a, 0x71, 0xad, 0xc1, 0x72, 0xab, 0x40, 0x1b, 0x8f, 0x35, 0x4e, 0x7a, 0x60, 0x86,
	0x60, 0x58, 0xcc, 0x8d, 0xf4, 0x3e, 0x96, 0x47, 0xb1, 0xb4, 0x3c, 0x62, 0x29, 0x4f, 0x94, 0x2e,
	0xc2, 0x8b, 0xec, 0xe6, 0xe2, 0x5b, 0x27, 0xf6, 0x9b, 0xad, 0x04, 0x12, 0x28, 0x46, 0xe6, 0x26,
	0xbf, 0x6d, 0x1e, 0xe3, 0xda, 0xb5, 0x33, 0x75, 0x94, 0xe8, 0x40, 0x2f, 0x1b, 0x4a, 0x6d, 0xbb,
	0x72, 0x94, 0x49, 0x63, 0xc9, 0x06, 0xae, 0x28, 0x51, 0x47, 0xfb, 0xe8, 0x68, 0xbd, 0x5b, 0x51,
	0xa2, 0xd9, 0xc7, 0xf5, 0xbf, 0x51, 0x93, 0x82, 0x36, 0x92, 0x5c, 0xe1, 0xaa, 0xf8, 0x5d, 0x17,
	0x50, 0xb5, 0x75, 0x40, 0x4b, 0x7a, 0xd2, 0x25, 0x45, 0x7b, 0x65, 0xf2, 0xb1, 0x17, 0x74, 0x97,
	0xf1, 0xd6, 0x2b, 0xc2, 0xab, 0xc5, 0x57, 0xe4, 0x19, 0xe1, 0xea, 0x52, 0x98, 0x9c, 0x95, 0x2a,
	0x4b, 0x5a, 0x34, 0xa2, 0x7f, 0x10, 0xbe, 0x4c, 0x73, 0xf7, 0xf1, 0xed, 0xeb, 0xa5, 0x52, 0x23,
	0xdb, 0xcc, 0x03, 0xee, 0x88, 0xee, 0x19, 0x76, 0xaf, 0xc4, 0x43, 0xfb, 0x72, 0x32, 0x0b, 0xd1,
	0x74, 0x16, 0xa2, 0xcf, 0x59, 0x88, 0x9e, 0xe6, 0x61, 0x30, 0x9d, 0x87, 0xc1, 0xfb, 0x3c, 0x0c,
	0x6e, 0x0e, 0x13, 0x65, 0xfb, 0x59, 0x4c, 0x7b, 0x30, 0x5c, 0xa0, 0xa7, 0x82, 0xc3, 0x8f, 0x25,
	0xbf, 0x60, 0xb7, 0x4e, 0x13, 0xaf, 0x15, 0xa7, 0x3f, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0x0a, 0x5e, 0x95, 0x1b, 0x02, 0x00, 0x00,
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
	// DidDocument queries a did documents with an id.
	DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error) {
	out := new(QueryDidDocumentResponse)
	err := c.cc.Invoke(ctx, "/elestodao.elesto.did.v1.Query/DidDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DidDocument queries a did documents with an id.
	DidDocument(context.Context, *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DidDocument(ctx context.Context, req *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocument not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DidDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elestodao.elesto.did.v1.Query/DidDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocument(ctx, req.(*QueryDidDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elestodao.elesto.did.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DidDocument",
			Handler:    _Query_DidDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/v1/query.proto",
}

func (m *QueryDidDocumentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDidDocumentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DidDocument.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryDidDocumentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocumentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DidDocument.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDidDocumentRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDidDocumentResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocument", wireType)
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
			if err := m.DidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
