package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.wizmacau.com/jack/proxypool/internal/server/handlers"
	"strconv"
	"time"
)

type HttpServer struct {
	engine *gin.Engine
}

func NewHttpServer() (*HttpServer, error) {
	s := &HttpServer{}
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

func (s *HttpServer) Run(port int) error {
	err := s.engine.Run(fmt.Sprintf(":%s", strconv.Itoa(port)))
	if err != nil {
		return err
	}
	return nil
}
