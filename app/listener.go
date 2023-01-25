package app

import (
	"fmt"
	"net"

	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pingme/rpc"
)

type svc struct {
	rpc.UnimplementedPingMeServer
}

func (svc) PingMe(c context.Context, p *rpc.PingArg) (*rpc.PingArg, error) {
	return p, nil
}

// edit
func StartGrpc(addr string, port int) error {
	addr = fmt.Sprintf(":%d", port)
	listen, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	var s = grpc.NewServer()
	rpcServer := &svc{}
	rpc.RegisterPingMeServer(s, rpcServer)
	reflection.Register(s)
	return s.Serve(listen)
}
