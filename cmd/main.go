package main

import (
	"gitlab.wizmacau.com/jack/proxypool/internal/configs"
	"gitlab.wizmacau.com/jack/proxypool/internal/models"
	"gitlab.wizmacau.com/jack/proxypool/internal/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.NewConfigs()
	if err != nil {
		panic("failed to load configs: " + err.Error())
	}

	db, err := gorm.Open(sqlite.Open(cfg.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.IP{})
	if err != nil {
		panic("failed to migrate")
	}

	httpServer, err := server.NewHttpServer()
	if err != nil {
		panic("failed to start http server: " + err.Error())
	}

	_ = httpServer.Run(8888)
}
