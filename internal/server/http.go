package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.wizmacau.com/jack/proxypool/internal/server/handlers"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HttpServer struct {
	engine *gin.Engine
	server *http.Server
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

func (s *HttpServer) Run(ctx context.Context, port int) error {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", strconv.Itoa(port)),
		Handler: s.engine,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	return nil
}

func (s *HttpServer) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}
	log.Println("server stopped")
}
