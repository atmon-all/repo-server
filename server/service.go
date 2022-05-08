package server

import (
	"context"

	"github.com/atmom/repo/config"
	repo "github.com/atmom/repo/grpc"
	"github.com/sirupsen/logrus"
)

type service struct {
	repo.UnimplementedRepoServer
	Config *config.RepoConfig
}

// Get components from repository server.
func (service *service) GetCompt(ctx context.Context, request *repo.GetComptRequest) (*repo.GetComptResponse, error) {
	logrus.Infof("Received: %v", request)
	return &repo.GetComptResponse{Code: 0, Message: "success"}, nil
}
