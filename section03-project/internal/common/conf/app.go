package conf

import "git.mysre.cn/yunlian-golang/go-hulk/dao"

type AppConf struct {
	App struct {
		Name string `yaml:"name"`
		Env  string `yaml:"env"`
	} `yaml:"app"`
	Grpc struct {
		Port string `yaml:"port"`
	} `yaml:"grpc"`
	Http struct {
		Port string `yaml:"port"`
	} `yaml:"http"`
	Mysql dao.MysqlConfig `json:"mysql"`
	Log   struct {
		Path string `yaml:"path"`
	} `yaml:"log"`
}
