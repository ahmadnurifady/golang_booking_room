package delivery

import (
	"final-project-booking-room/config"
	"final-project-booking-room/delivery/controller"
	"final-project-booking-room/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc     manager.UseCaseManager
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/final/v1")
	controller.NewRoomController(s.uc.RoomUsecase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := manager.NewRepoManager(infra)
	uc := manager.NewUseCaseManager(repo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		uc:     uc,
		engine: engine,
		host:   host,
	}
}
