package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.wizmacau.com/jack/proxypool/internal/server/handlers"
	"gitlab.wizmacau.com/jack/proxypool/pkg/logger"
	"net/http"
	"strconv"
	"time"
)

type HttpServer struct {
	engine *gin.Engine
	server *http.Server
	logger logger.Logger
}

func NewHttpServer(logger logger.Logger) (*HttpServer, error) {
	s := &HttpServer{
		logger: logger,
	}
	err := s.init()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *HttpServer) init() error {
	var (
		router = gin.New()
	)
	router.GET("/health", handlers.Health(time.Now().UTC()))
	s.engine = router
	return nil
}

func (s *HttpServer) Run(ctx context.Context, port int) error {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", strconv.Itoa(port)),
		Handler: s.engine,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			s.logger.Error("server stopped", logger.String("error", err.Error()))
		}
	}()

	<-ctx.Done()
	return nil
}

func (s *HttpServer) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		s.logger.Error("server forced to shutdown", logger.String("error", err.Error()))
	}
	s.logger.Info("server stopped")
}
