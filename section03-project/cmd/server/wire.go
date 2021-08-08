// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"context"
	"section03-project/internal/application"
	"section03-project/internal/common"
	"section03-project/internal/domain"
	"section03-project/internal/interfaces"
	"github.com/google/wire"
)

func MakeServerWithDI(ctx context.Context) (*interfaces.GRPCServer, error) {
	wire.Build(
		common.RepoProviderSet,
		domain.ProviderSet,
		application.NewApplicationService,
		interfaces.NewGRPCServer,
	)
	return &interfaces.GRPCServer{}, nil
}
