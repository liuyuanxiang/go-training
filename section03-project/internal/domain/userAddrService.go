package domain

import (
	"context"
	"section03-project/internal/domain/address"
	"section03-project/internal/domain/user"
)

// 涉及到多个实体的行为（无法放在某一实体）时放在领域服务内
type UserAddrService struct {
	ctx         context.Context
	userUsecase *user.Usecase
	addrUsecase *address.Usecase
}

func NewUserAddrService(ctx context.Context, userUsecase *user.Usecase, addrUsecase *address.Usecase) *UserAddrService {
	uaService := &UserAddrService{ctx: ctx}
	uaService.userUsecase = userUsecase
	uaService.addrUsecase = addrUsecase
	return uaService
}

// 检查用户及地址信息合法性
func (ua *UserAddrService) CheckUserAddress(userId, AddressId string) error {
	_, err := ua.userUsecase.Get(ua.ctx, userId)
	if err != nil {
		return err
	}

	_, err = ua.addrUsecase.Get(ua.ctx, AddressId)
	if err != nil {
		return err
	}
	return nil
}
