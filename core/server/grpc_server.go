package server

import (
	"errors"
	"sync"
	"net"
	"strconv"
	"fmt"

	context "golang.org/x/net/context"
	"github.com/lucasmbaia/tdbas/core/proto"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
)

type ServerGRPC struct {
	sync.RWMutex

	server	    *grpc.Server
	port	    int
	certificate string
	key	    string
}

type ServerGRPCConfig struct {
	Port	    int
	Certificate string
	Key	    string
}

func NewServerGRPC(cfg ServerGRPCConfig) (s ServerGRPC, err error) {
	if cfg.Port == 0 {
		err = errors.New("Port must be specified")
		return
	}

	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key
	return
}

func (s *ServerGRPC) Listen() (err error) {
	var (
		l	  net.Listener
		grpcOpts  = []grpc.ServerOption{}
		grpcCreds credentials.TransportCredentials
	)

	if l, err = net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(s.port))); err != nil {
		return
	}

	if s.certificate != "" && s.key != "" {
		if grpcCreds, err = credentials.NewServerTLSFromFile(s.certificate, s.key); err != nil {
			return err
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	s.server = grpc.NewServer(grpcOpts...)
	tdbas.RegisterTdbasServiceServer(s.server, s)

	if err = s.server.Serve(l); err != nil {
		return
	}

	return
}

func (s *ServerGRPC) CallTask(ctx context.Context, task *tdbas.Task) (status *tdbas.TaskStatus, err error) {
	return
}

func (s *ServerGRPC) Close() {
	if s.server != nil {
		s.server.Stop()
	}

	return
}

