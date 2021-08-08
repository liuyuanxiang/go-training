package application

import (
	"context"
	"section03-project/internal/domain"
	"section03-project/internal/domain/address"
	"section03-project/internal/domain/user"
	"section03-project/internal/interfaces/assembler"
	"section03-project/internal/interfaces/dto"
	"github.com/mitchellh/mapstructure"
	"math"
)

// 应用服务自身不包含业务逻辑，通过调度用例及领域服务间接实现业务
type ApplicationService struct {
	ctx context.Context
	uu  *user.Usecase
	au  *address.Usecase
	uas *domain.UserAddrService
}

func NewApplicationService(ctx context.Context, userUsecase *user.Usecase, addressUsercase *address.Usecase, userAddrService *domain.UserAddrService) *ApplicationService {
	return &ApplicationService{
		ctx: ctx,
		uu:  userUsecase,
		au:  addressUsercase,
		uas: userAddrService,
	}
}

func (a *ApplicationService) CreateUser(userInfo *dto.UserDTO) (string, error) {
	userDO := assembler.ToUserDO(userInfo)
	return a.uu.Create(a.ctx, userDO)
}

func (a *ApplicationService) GetUser(id string) (dto.UserDTO, error) {
	user, err := a.uu.Get(a.ctx, id)
	if err != nil {
		return dto.UserDTO{}, nil
	}
	return dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (a *ApplicationService) ListUser(pageNum, pageSize int) (dto.UserListDTO, error) {
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 15
	}

	userList, count, error := a.uu.List(a.ctx, pageNum, pageSize)
	if len(userList) == 0 {
		return dto.UserListDTO{}, nil
	}
	var list []dto.UserDTO
	mapstructure.Decode(userList, &list)
	return dto.UserListDTO{
		List:        list,
		TotalCount:  count,
		CurrentPage: int64(pageNum),
		TotalPage:   int64(math.Ceil(float64(count) / float64(pageSize))),
		PageSize:    int64(pageSize),
	}, error
}

func (a *ApplicationService) CheckUserAddress(userId, AddressId string) error {
	return a.uas.CheckUserAddress(userId, AddressId)
}
