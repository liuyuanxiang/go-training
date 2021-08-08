package repo

import (
	"section03-project/internal/domain/address"
	"git.mysre.cn/yunlian-golang/go-hulk"
	"github.com/mitchellh/mapstructure"

	"context"
	"git.mysre.cn/yunlian-golang/go-hulk/dao"
	"github.com/google/uuid"
)

type AddressRepo struct {
	ctx context.Context
	db  *dao.MysqlDB
}

func NewAddressRepo(ctx context.Context) address.RepoInterface {
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
	address := new(address.Address)
	if !migrate.HasTable(address) {
		err := migrate.CreateTable(address)
		if err != nil {
			app.Log.Error(err)
		}
		err = migrate.RenameTable(address, address.TableName())
		if err != nil {
			app.Log.Error(err)
		}
	}
	return &AddressRepo{db: db, ctx: ctx}
}

func (a *AddressRepo) CreateAddr(address *address.Address) (id string, err error) {
	address.ID = uuid.Must(uuid.NewUUID()).String()
	err = a.db.UseGORM().Create(address).Error
	return address.ID, err
}

func (a *AddressRepo) UpdateAddr(id string, address *address.Address) error {
	panic("implement me")
}

func (a *AddressRepo) DeleteAddr(id string) error {
	panic("implement me")
}

func (a *AddressRepo) GetAddr(id string) (*address.Address, error) {
	panic("implement me")
}

func (a *AddressRepo) ListAddr(pageNum, pageSize int64) ([]*address.Address, error) {
	panic("implement me")
}
