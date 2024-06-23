package main

import (
	"context"
	"gitlab.wizmacau.com/jack/proxypool/internal/configs"
	"gitlab.wizmacau.com/jack/proxypool/internal/models"
	"gitlab.wizmacau.com/jack/proxypool/internal/server"
	zapLogger "gitlab.wizmacau.com/jack/proxypool/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.NewConfigs()
	if err != nil {
		panic("failed to load configs: " + err.Error())
	}

	logger := zapLogger.New(cfg.Logger)

	db, err := gorm.Open(sqlite.Open(cfg.Db.Dsn), &gorm.Config{})
	if err != nil {
		logger.Panic("failed to connect to database")
	}

	err = db.AutoMigrate(&models.IP{})
	if err != nil {
		logger.Panic("failed to migrate")
	}

	ctx, cancel := context.WithCancel(context.Background())

	httpServer, err := server.NewHttpServer(logger)
	if err != nil {
		logger.Panic("failed to create http server: " + err.Error())
	}

	HandleSignals(
		func() {
			httpServer.Stop()
			cancel()
		},
	)

	_ = httpServer.Run(ctx, 8888)
}
