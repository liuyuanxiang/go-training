package repo

import (
	"section03-project/internal/domain/user"
	"git.mysre.cn/yunlian-golang/go-hulk"
	"github.com/mitchellh/mapstructure"
	"math"

	"context"
	"git.mysre.cn/yunlian-golang/go-hulk/dao"
	"github.com/google/uuid"
)

type UserRepo struct {
	ctx context.Context
	db  *dao.MysqlDB
}

func NewUserRepo(ctx context.Context) user.RepoInterface {
	app, _ := hulk.Facades(ctx).Context().GetGRPCApplication()
	mysqlConfig := app.Config().GetStringMap("mysql")
	app.Log.Debug("准备连接 Mysql：", mysqlConfig)
	opts := &dao.MysqlConfig{}
	mapstructure.Decode(mysqlConfig, opts)
	db, _, err := dao.NewGORM(opts, app.Log)
	if err != nil {
		return nil
	}
	migrate := db.UseGORM().Migrator()
	user := new(user.User)
	if !migrate.HasTable(user) {
		err := migrate.CreateTable(user)
		if err != nil {
			app.Log.Error(err)
		}
		err = migrate.RenameTable(user, user.TableName())
		if err != nil {
			app.Log.Error(err)
		}
	}
	return &UserRepo{db: db, ctx: ctx}
}

func (u *UserRepo) CreateUser(user *user.User) (string, error) {
	user.ID = uuid.Must(uuid.NewUUID()).String()
	err := u.db.UseGORM().Create(user).Error
	return user.ID, err
}

func (u *UserRepo) GetUser(id string) (user *user.User, err error) {
	u.db.UseGORM().Where("id = ?", id).First(&user)
	return user, nil
}

func (u *UserRepo) ListUser(pageNum, pageSize int) (users []*user.User, count int64, err error) {
	var user user.User
	offset := int(math.Ceil(float64((pageNum - 1) * pageSize)))
	u.db.UseGORM().Table(user.TableName()).Offset(offset).Limit(pageSize).Find(&users)
	u.db.UseGORM().Table(user.TableName()).Count(&count)
	return users, count, nil
}

func (u *UserRepo) UpdateUser(id string, user *user.User) error {
	panic("implement me")
}

func (u *UserRepo) DeleteUser(id string) error {
	panic("implement me")
}
