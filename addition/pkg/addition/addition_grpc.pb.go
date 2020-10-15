// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package addition

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderClient interface {
	// Sends a greeting
	CalcAdd(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

var adderCalcAddStreamDesc = &grpc.StreamDesc{
	StreamName: "CalcAdd",
}

func (c *adderClient) CalcAdd(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/addition.Adder/CalcAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderService is the service API for Adder service.
// Fields should be assigned to their respective handler implementations only before
// RegisterAdderService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type AdderService struct {
	// Sends a greeting
	CalcAdd func(context.Context, *AddRequest) (*AddResponse, error)
}

func (s *AdderService) calcAdd(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CalcAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/addition.Adder/CalcAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CalcAdd(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterAdderService registers a service implementation with a gRPC server.
func RegisterAdderService(s grpc.ServiceRegistrar, srv *AdderService) {
	srvCopy := *srv
	if srvCopy.CalcAdd == nil {
		srvCopy.CalcAdd = func(context.Context, *AddRequest) (*AddResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CalcAdd not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "addition.Adder",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CalcAdd",
				Handler:    srvCopy.calcAdd,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "pkg/addition/addition.proto",
	}

	s.RegisterService(&sd, nil)
}