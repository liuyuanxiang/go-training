package domain

import (
	"section03-project/internal/domain/address"
	"section03-project/internal/domain/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(user.NewUserUsecase, address.NewAddressUsecase, NewUserAddrService)
