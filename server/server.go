package server

import (
	"fmt"
	"net"

	"github.com/atmom/repo/config"
	repo "github.com/atmom/repo/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// start server.
func Start(config *config.RepoConfig) {
	if config.Port < 1024 || config.Port > 65536 {
		logrus.Fatalf("Repository server start error, invalid server port %d.", config.Port)
	}

	// listen tcp.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.Fatalf("Repository server start error %v.", err)
	}

	// create server.
	s := grpc.NewServer()
	repo.RegisterRepoServer(s, &service{Config: config})

	logrus.Infof("Repository server start success, listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.Fatalf("Repository server start error %v.", err)
	}
}
