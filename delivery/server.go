package delivery

import (

	// "final-project/delivery/middleware"
	"final-project-booking-room/config"
	"final-project-booking-room/delivery/controller"
	"final-project-booking-room/manager"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	auth       usecase.AuthUseCase
	engine     *gin.Engine
	host       string
	jwtService common.JwtToken
}

func (s *Server) setupControllers() {
	// authMiddlerware := middleware.NewAuthMiddleware(s.jwtService)
	rg := s.engine.Group("/final/v1")
	controller.NewRoomController(s.uc.RoomUsecase(), rg).Route()
	controller.NewUserController(s.uc.UserUseCase(), rg).Route()
	controller.NewRoomController(s.uc.RoomUsecase(), rg).Route()
	controller.NewBookingController(s.uc.BookingUsecase(), rg).Route()
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
