package common

import (
	"section03-project/internal/common/repo"
	"github.com/google/wire"
)

var RepoProviderSet = wire.NewSet(repo.NewUserRepo, repo.NewAddressRepo)
