package main

import (
	"flag"
	"section03-project/internal/common/conf"
	"section03-project/internal/common/logger"
	"section03-project/internal/interfaces"

	"git.mysre.cn/yunlian-golang/go-hulk"
	"git.mysre.cn/yunlian-golang/go-hulk/boot"
	"git.mysre.cn/yunlian-golang/go-hulk/config"
	"git.mysre.cn/yunlian-golang/go-hulk/config/file"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "../config", "config path, eg: -conf config.yaml")
}

func initApp(app *boot.GRPCApplication) error {
	// 设置整个工程项目可直接使用的 Log 处理器
	logger.InitLogger(app.Logger())

	// 构建应用各层级服务的依赖关联，依赖构建完成返回顶层接口层的实现实例
	serve, err := MakeServerWithDI(app.Context())
	if err != nil {
		return err
	}

	// 注册服务端的业务实现
	return interfaces.RegisterServer(app, serve)
}

func main() {

	//flag.Parse()
	cfg := config.New(config.WithSource(file.NewSource("./config/app.yaml")))
	if err := cfg.Load(); err != nil {
		panic(err)
	}

	ac := new(conf.AppConf)
	if err := cfg.Unmarshal(ac); err != nil {
		panic(err)
	}

	app := hulk.NewGRPCApplication(cfg, boot.WithGRPCGateway(true))
	if err := initApp(app); err != nil {
		app.Logger().Error("应用", app.Name, "注册引导流程执行失败，请检查 err:", err)
	}

	if err := app.Run(); err != nil {
		app.Logger().Error("应用", app.Name, "启动失败 err:", err)
	}
}
