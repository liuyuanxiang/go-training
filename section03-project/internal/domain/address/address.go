package address

import (
	"context"
)

type Address struct {
	ID       string `gorm:"id;primary_key" json:"id"`
	UserID   string `gorm:"id;primary_key" json:"user_id"`
	Name     string `json:"name"`
	Mobile   int64  `json:"mobile"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Street   string `json:"street"`
	PostCode string `json:"post_code"`
}

// 实体行为通过用例表示
type Usecase struct {
	repo RepoInterface
}

func NewAddressUsecase(repo RepoInterface) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Address) TableName() string {
	return "user_address"
}

func (uc *Usecase) Create(ctx context.Context, addr *Address) (string, error) {
	return uc.repo.CreateAddr(addr)
}

func (uc *Usecase) Update(ctx context.Context, id string, addr *Address) error {
	return uc.repo.UpdateAddr(id, addr)
}

func (uc *Usecase) Delete(ctx context.Context, id string) error {
	return uc.repo.DeleteAddr(id)
}

func (uc *Usecase) Get(ctx context.Context, id string) (*Address, error) {
	return uc.repo.GetAddr(id)
}

func (uc *Usecase) List(ctx context.Context, pageNum, pageSize int64) ([]*Address, error) {
	return uc.repo.ListAddr(pageNum, pageSize)
}
