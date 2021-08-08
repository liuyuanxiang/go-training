package user

import (
	"context"
)

type User struct {
	ID       string `gorm:"id;primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// 实体行为通过用例表示
type Usecase struct {
	repo RepoInterface
}

func NewUserUsecase(repo RepoInterface) *Usecase {
	return &Usecase{repo: repo}
}

func (u *User) TableName() string {
	return "user"
}

func (uc *Usecase) Create(ctx context.Context, user *User) (string, error) {
	return uc.repo.CreateUser(user)
}

func (uc *Usecase) Update(ctx context.Context, id string, user *User) error {
	return uc.repo.UpdateUser(id, user)
}

func (uc *Usecase) Delete(ctx context.Context, id string) error {
	return uc.repo.DeleteUser(id)
}

func (uc *Usecase) Get(ctx context.Context, id string) (*User, error) {
	return uc.repo.GetUser(id)
}

func (uc *Usecase) List(ctx context.Context, pageNum, pageSize int) ([]*User, int64, error) {
	return uc.repo.ListUser(pageNum, pageSize)
}
