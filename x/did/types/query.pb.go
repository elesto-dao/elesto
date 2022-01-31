// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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
type QueryDidDocumentsRequest struct {
	// status enables to query for validators matching a given status.
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDidDocumentsRequest) Reset()         { *m = QueryDidDocumentsRequest{} }
func (m *QueryDidDocumentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentsRequest) ProtoMessage()    {}
func (*QueryDidDocumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{0}
}
func (m *QueryDidDocumentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentsRequest.Merge(m, src)
}
func (m *QueryDidDocumentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentsRequest proto.InternalMessageInfo

func (m *QueryDidDocumentsRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *QueryDidDocumentsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method
type QueryDidDocumentsResponse struct {
	// validators contains all the queried validators.
	DidDocuments []DidDocument `protobuf:"bytes,1,rep,name=didDocuments,proto3" json:"didDocuments"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDidDocumentsResponse) Reset()         { *m = QueryDidDocumentsResponse{} }
func (m *QueryDidDocumentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentsResponse) ProtoMessage()    {}
func (*QueryDidDocumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{1}
}
func (m *QueryDidDocumentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentsResponse.Merge(m, src)
}
func (m *QueryDidDocumentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentsResponse proto.InternalMessageInfo

func (m *QueryDidDocumentsResponse) GetDidDocuments() []DidDocument {
	if m != nil {
		return m.DidDocuments
	}
	return nil
}

func (m *QueryDidDocumentsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.
type QueryDidDocumentRequest struct {
	// status enables to query for validators matching a given status.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryDidDocumentRequest) Reset()         { *m = QueryDidDocumentRequest{} }
func (m *QueryDidDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentRequest) ProtoMessage()    {}
func (*QueryDidDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{2}
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
	// validators contains all the queried validators.
	DidDocument DidDocument `protobuf:"bytes,1,opt,name=didDocument,proto3" json:"didDocument"`
	DidMetadata DidMetadata `protobuf:"bytes,2,opt,name=didMetadata,proto3" json:"didMetadata"`
}

func (m *QueryDidDocumentResponse) Reset()         { *m = QueryDidDocumentResponse{} }
func (m *QueryDidDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentResponse) ProtoMessage()    {}
func (*QueryDidDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31228b4ee4821623, []int{3}
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

func (m *QueryDidDocumentResponse) GetDidMetadata() DidMetadata {
	if m != nil {
		return m.DidMetadata
	}
	return DidMetadata{}
}

func init() {
	proto.RegisterType((*QueryDidDocumentsRequest)(nil), "allinbits.cosmoscash.did.QueryDidDocumentsRequest")
	proto.RegisterType((*QueryDidDocumentsResponse)(nil), "allinbits.cosmoscash.did.QueryDidDocumentsResponse")
	proto.RegisterType((*QueryDidDocumentRequest)(nil), "allinbits.cosmoscash.did.QueryDidDocumentRequest")
	proto.RegisterType((*QueryDidDocumentResponse)(nil), "allinbits.cosmoscash.did.QueryDidDocumentResponse")
}

func init() { proto.RegisterFile("did/query.proto", fileDescriptor_31228b4ee4821623) }

var fileDescriptor_31228b4ee4821623 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x8a, 0xd4, 0x30,
	0x14, 0xc6, 0x9b, 0xaa, 0x0b, 0xa6, 0xab, 0x42, 0xfc, 0x57, 0xab, 0xd4, 0x52, 0x50, 0x47, 0xc1,
	0x84, 0xe9, 0xbe, 0xc1, 0xb2, 0x28, 0x5e, 0x2c, 0x6a, 0x2f, 0xbd, 0x4b, 0x9b, 0xd0, 0x0d, 0xcc,
	0x34, 0xdd, 0x4d, 0xba, 0xb8, 0x8a, 0x37, 0x3e, 0x81, 0xa0, 0xf8, 0x26, 0xa2, 0x8f, 0xb0, 0x97,
	0x0b, 0xde, 0x78, 0x25, 0x32, 0xe3, 0x83, 0x48, 0x9b, 0x74, 0xcd, 0xea, 0x8c, 0xee, 0xdc, 0x25,
	0x87, 0xf3, 0x7d, 0xdf, 0xef, 0x9c, 0xa6, 0xf0, 0x12, 0x13, 0x8c, 0xec, 0xb6, 0x7c, 0xef, 0x00,
	0x37, 0x7b, 0x52, 0x4b, 0x14, 0xd2, 0xc9, 0x44, 0xd4, 0x85, 0xd0, 0x0a, 0x97, 0x52, 0x4d, 0xa5,
	0x2a, 0xa9, 0xda, 0xc1, 0x4c, 0xb0, 0xe8, 0x56, 0x25, 0x65, 0x35, 0xe1, 0x84, 0x36, 0x82, 0xd0,
	0xba, 0x96, 0x9a, 0x6a, 0x21, 0x6b, 0x65, 0x74, 0xd1, 0x03, 0xd3, 0x4d, 0x0a, 0xaa, 0xb8, 0x31,
	0x24, 0xfb, 0xe3, 0x82, 0x6b, 0x3a, 0x26, 0x0d, 0xad, 0x44, 0xdd, 0x37, 0xdb, 0xde, 0x0b, 0x5d,
	0x28, 0x13, 0xcc, 0x5e, 0xaf, 0x54, 0xb2, 0x92, 0xfd, 0x91, 0x74, 0x27, 0x53, 0x4d, 0x5f, 0xc1,
	0xf0, 0x79, 0x67, 0xb3, 0x25, 0xd8, 0x96, 0x2c, 0xdb, 0x29, 0xaf, 0xb5, 0xca, 0xf9, 0x6e, 0xcb,
	0x95, 0x46, 0xd7, 0xe0, 0x9a, 0xd2, 0x54, 0xb7, 0x2a, 0x04, 0x09, 0x18, 0x9d, 0xcf, 0xed, 0x0d,
	0x3d, 0x82, 0xf0, 0x77, 0x58, 0xe8, 0x27, 0x60, 0x14, 0x64, 0x77, 0xed, 0x1c, 0xb8, 0x23, 0xc3,
	0x66, 0x54, 0x4b, 0x86, 0x9f, 0xd1, 0x8a, 0x5b, 0xcf, 0xdc, 0x51, 0xa6, 0x9f, 0x00, 0xbc, 0xb1,
	0x20, 0x5c, 0x35, 0xb2, 0x56, 0x1c, 0x3d, 0x85, 0xeb, 0xcc, 0xa9, 0x87, 0x20, 0x39, 0x33, 0x0a,
	0xb2, 0x3b, 0x78, 0xd9, 0xe6, 0xb0, 0xe3, 0xb2, 0x79, 0xf6, 0xf0, 0xfb, 0x6d, 0x2f, 0x3f, 0x61,
	0x80, 0x1e, 0x2f, 0xc0, 0xbe, 0xf7, 0x5f, 0x6c, 0x43, 0x73, 0x82, 0xfb, 0x3e, 0xbc, 0xfe, 0x27,
	0xf6, 0xb0, 0xb2, 0x8b, 0xd0, 0x17, 0xcc, 0xae, 0xcb, 0x17, 0x2c, 0xfd, 0x02, 0xfe, 0xde, 0xef,
	0xf1, 0x84, 0xdb, 0x30, 0x70, 0x00, 0x7b, 0xd5, 0x8a, 0x03, 0xba, 0x7a, 0x6b, 0xb7, 0xcd, 0x35,
	0x65, 0x54, 0x53, 0x3b, 0xe0, 0xbf, 0xed, 0x86, 0x66, 0xc7, 0x6e, 0x28, 0x65, 0x9f, 0x7d, 0x78,
	0xae, 0x47, 0x47, 0x1f, 0x00, 0x5c, 0x77, 0x3f, 0x11, 0xca, 0x96, 0x9b, 0x2e, 0x7b, 0x4c, 0xd1,
	0xc6, 0x4a, 0x1a, 0xb3, 0xa1, 0xf4, 0xe6, 0xdb, 0xaf, 0x3f, 0xdf, 0xfb, 0x57, 0xd1, 0x65, 0x72,
	0x2c, 0x26, 0xf6, 0x55, 0x2b, 0xf4, 0x11, 0xc0, 0xc0, 0x51, 0xa1, 0xf1, 0xe9, 0x13, 0x06, 0xa8,
	0x6c, 0x15, 0x89, 0x65, 0x4a, 0x7a, 0xa6, 0x08, 0x85, 0x0b, 0x98, 0xc8, 0x6b, 0xc1, 0xde, 0x6c,
	0x3e, 0x39, 0x9c, 0xc5, 0xe0, 0x68, 0x16, 0x83, 0x1f, 0xb3, 0x18, 0xbc, 0x9b, 0xc7, 0xde, 0xd1,
	0x3c, 0xf6, 0xbe, 0xcd, 0x63, 0xef, 0x05, 0xa9, 0x84, 0xde, 0x69, 0x0b, 0x5c, 0xca, 0xa9, 0xa3,
	0x36, 0xc9, 0x0f, 0xbb, 0x68, 0xb2, 0x9f, 0x91, 0x97, 0xbd, 0x9d, 0x3e, 0x68, 0xb8, 0x2a, 0xd6,
	0xfa, 0xbf, 0x74, 0xe3, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x2c, 0x3c, 0x3d, 0x41, 0x04,
	0x00, 0x00,
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
	// DidDocuments queries all did documents that match the given status.
	DidDocuments(ctx context.Context, in *QueryDidDocumentsRequest, opts ...grpc.CallOption) (*QueryDidDocumentsResponse, error)
	// DidDocument queries a did documents with an id.
	DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDocuments(ctx context.Context, in *QueryDidDocumentsRequest, opts ...grpc.CallOption) (*QueryDidDocumentsResponse, error) {
	out := new(QueryDidDocumentsResponse)
	err := c.cc.Invoke(ctx, "/allinbits.cosmoscash.did.Query/DidDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error) {
	out := new(QueryDidDocumentResponse)
	err := c.cc.Invoke(ctx, "/allinbits.cosmoscash.did.Query/DidDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DidDocuments queries all did documents that match the given status.
	DidDocuments(context.Context, *QueryDidDocumentsRequest) (*QueryDidDocumentsResponse, error)
	// DidDocument queries a did documents with an id.
	DidDocument(context.Context, *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DidDocuments(ctx context.Context, req *QueryDidDocumentsRequest) (*QueryDidDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocuments not implemented")
}
func (*UnimplementedQueryServer) DidDocument(ctx context.Context, req *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocument not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DidDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allinbits.cosmoscash.did.Query/DidDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocuments(ctx, req.(*QueryDidDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/allinbits.cosmoscash.did.Query/DidDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocument(ctx, req.(*QueryDidDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allinbits.cosmoscash.did.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DidDocuments",
			Handler:    _Query_DidDocuments_Handler,
		},
		{
			MethodName: "DidDocument",
			Handler:    _Query_DidDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/query.proto",
}

func (m *QueryDidDocumentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDidDocumentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DidDocuments) > 0 {
		for iNdEx := len(m.DidDocuments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidDocuments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
		size, err := m.DidMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
func (m *QueryDidDocumentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocumentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DidDocuments) > 0 {
		for _, e := range m.DidDocuments {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
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
	l = m.DidMetadata.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDidDocumentsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDidDocumentsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocuments", wireType)
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
			m.DidDocuments = append(m.DidDocuments, DidDocument{})
			if err := m.DidDocuments[len(m.DidDocuments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidMetadata", wireType)
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
			if err := m.DidMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
