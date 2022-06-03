// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: github.com/aperturerobotics/protobuf-project/example/example.proto

package example

import (
	context "context"
	errors "errors"

	drpc1 "github.com/planetscale/vtprotobuf/codec/drpc"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto struct{}

func (drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return drpc1.Marshal(msg)
}

func (drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return drpc1.Unmarshal(buf, msg)
}

type DRPCEchoerClient interface {
	DRPCConn() drpc.Conn

	Echo(ctx context.Context, in *EchoMsg) (*EchoMsg, error)
	EchoServerStream(ctx context.Context, in *EchoMsg) (DRPCEchoer_EchoServerStreamClient, error)
	EchoClientStream(ctx context.Context) (DRPCEchoer_EchoClientStreamClient, error)
	EchoBidiStream(ctx context.Context) (DRPCEchoer_EchoBidiStreamClient, error)
}

type drpcEchoerClient struct {
	cc drpc.Conn
}

func NewDRPCEchoerClient(cc drpc.Conn) DRPCEchoerClient {
	return &drpcEchoerClient{cc}
}

func (c *drpcEchoerClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcEchoerClient) Echo(ctx context.Context, in *EchoMsg) (*EchoMsg, error) {
	out := new(EchoMsg)
	err := c.cc.Invoke(ctx, "/example.Echoer/Echo", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcEchoerClient) EchoServerStream(ctx context.Context, in *EchoMsg) (DRPCEchoer_EchoServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, "/example.Echoer/EchoServerStream", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcEchoer_EchoServerStreamClient{stream}
	if err := x.MsgSend(in, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return nil, err
	}
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DRPCEchoer_EchoServerStreamClient interface {
	drpc.Stream
	Recv() (*EchoMsg, error)
}

type drpcEchoer_EchoServerStreamClient struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoServerStreamClient) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcEchoer_EchoServerStreamClient) RecvMsg(m *EchoMsg) error {
	return x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

func (c *drpcEchoerClient) EchoClientStream(ctx context.Context) (DRPCEchoer_EchoClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, "/example.Echoer/EchoClientStream", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcEchoer_EchoClientStreamClient{stream}
	return x, nil
}

type DRPCEchoer_EchoClientStreamClient interface {
	drpc.Stream
	Send(*EchoMsg) error
	CloseAndRecv() (*EchoMsg, error)
}

type drpcEchoer_EchoClientStreamClient struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoClientStreamClient) Send(m *EchoMsg) error {
	return x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

func (x *drpcEchoer_EchoClientStreamClient) CloseAndRecv() (*EchoMsg, error) {
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EchoMsg)
	if err := x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcEchoer_EchoClientStreamClient) CloseAndRecvMsg(m *EchoMsg) error {
	if err := x.CloseSend(); err != nil {
		return err
	}
	return x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

func (c *drpcEchoerClient) EchoBidiStream(ctx context.Context) (DRPCEchoer_EchoBidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, "/example.Echoer/EchoBidiStream", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcEchoer_EchoBidiStreamClient{stream}
	return x, nil
}

type DRPCEchoer_EchoBidiStreamClient interface {
	drpc.Stream
	Send(*EchoMsg) error
	Recv() (*EchoMsg, error)
}

type drpcEchoer_EchoBidiStreamClient struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoBidiStreamClient) Send(m *EchoMsg) error {
	return x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

func (x *drpcEchoer_EchoBidiStreamClient) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcEchoer_EchoBidiStreamClient) RecvMsg(m *EchoMsg) error {
	return x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

type DRPCEchoerServer interface {
	Echo(context.Context, *EchoMsg) (*EchoMsg, error)
	EchoServerStream(*EchoMsg, DRPCEchoer_EchoServerStreamStream) error
	EchoClientStream(DRPCEchoer_EchoClientStreamStream) error
	EchoBidiStream(DRPCEchoer_EchoBidiStreamStream) error
}

type DRPCEchoerUnimplementedServer struct{}

func (s *DRPCEchoerUnimplementedServer) Echo(context.Context, *EchoMsg) (*EchoMsg, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCEchoerUnimplementedServer) EchoServerStream(*EchoMsg, DRPCEchoer_EchoServerStreamStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCEchoerUnimplementedServer) EchoClientStream(DRPCEchoer_EchoClientStreamStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCEchoerUnimplementedServer) EchoBidiStream(DRPCEchoer_EchoBidiStreamStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCEchoerDescription struct{}

func (DRPCEchoerDescription) NumMethods() int { return 4 }

func (DRPCEchoerDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/example.Echoer/Echo", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCEchoerServer).
					Echo(
						ctx,
						in1.(*EchoMsg),
					)
			}, DRPCEchoerServer.Echo, true
	case 1:
		return "/example.Echoer/EchoServerStream", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCEchoerServer).
					EchoServerStream(
						in1.(*EchoMsg),
						&drpcEchoer_EchoServerStreamStream{in2.(drpc.Stream)},
					)
			}, DRPCEchoerServer.EchoServerStream, true
	case 2:
		return "/example.Echoer/EchoClientStream", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCEchoerServer).
					EchoClientStream(
						&drpcEchoer_EchoClientStreamStream{in1.(drpc.Stream)},
					)
			}, DRPCEchoerServer.EchoClientStream, true
	case 3:
		return "/example.Echoer/EchoBidiStream", drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCEchoerServer).
					EchoBidiStream(
						&drpcEchoer_EchoBidiStreamStream{in1.(drpc.Stream)},
					)
			}, DRPCEchoerServer.EchoBidiStream, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterEchoer(mux drpc.Mux, impl DRPCEchoerServer) error {
	return mux.Register(impl, DRPCEchoerDescription{})
}

type DRPCEchoer_EchoStream interface {
	drpc.Stream
	SendAndClose(*EchoMsg) error
}

type drpcEchoer_EchoStream struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoStream) SendAndClose(m *EchoMsg) error {
	if err := x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCEchoer_EchoServerStreamStream interface {
	drpc.Stream
	Send(*EchoMsg) error
}

type drpcEchoer_EchoServerStreamStream struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoServerStreamStream) Send(m *EchoMsg) error {
	return x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

type DRPCEchoer_EchoClientStreamStream interface {
	drpc.Stream
	SendAndClose(*EchoMsg) error
	Recv() (*EchoMsg, error)
}

type drpcEchoer_EchoClientStreamStream struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoClientStreamStream) SendAndClose(m *EchoMsg) error {
	if err := x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

func (x *drpcEchoer_EchoClientStreamStream) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcEchoer_EchoClientStreamStream) RecvMsg(m *EchoMsg) error {
	return x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

type DRPCEchoer_EchoBidiStreamStream interface {
	drpc.Stream
	Send(*EchoMsg) error
	Recv() (*EchoMsg, error)
}

type drpcEchoer_EchoBidiStreamStream struct {
	drpc.Stream
}

func (x *drpcEchoer_EchoBidiStreamStream) Send(m *EchoMsg) error {
	return x.MsgSend(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}

func (x *drpcEchoer_EchoBidiStreamStream) Recv() (*EchoMsg, error) {
	m := new(EchoMsg)
	if err := x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcEchoer_EchoBidiStreamStream) RecvMsg(m *EchoMsg) error {
	return x.MsgRecv(m, drpcEncoding_File_github_com_aperturerobotics_protobuf_project_example_example_proto{})
}
