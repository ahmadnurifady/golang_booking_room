package controller

import (
	"final-project-booking-room/config"
	"final-project-booking-room/model"
	"final-project-booking-room/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	uc usecase.RoomUseCase
	rg *gin.RouterGroup
}

func (r *RoomController) createHandler(ctx *gin.Context) {
	var payload model.Room
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		panic(err)
	}

	createRoom, err := r.uc.RegisterNewRoom(payload)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusCreated, createRoom)
}

func (r *RoomController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id tidak ditemukan"})
	}

	rspPayload, err := r.uc.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "room yang dimaksud tidak ditemukan"})
	}

	ctx.JSON(http.StatusOK, rspPayload)
}

func (r *RoomController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id tidak ditemukan"})
	}

	err := r.uc.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "room yang dimaksud tidak ditemukan"})
	}

	ctx.JSON(http.StatusOK, "room telah dihapus")
}

func (r *RoomController) updateHandler(ctx *gin.Context) {
	var roomUpdate model.Room
	err := ctx.ShouldBind(&roomUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id tidak ditemukan"})
		return
	}

	err = r.uc.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product telah diupdate", "data": roomUpdate})
}

func (r *RoomController) Route() {
	br := r.rg.Group(config.RoomGroup)
	br.POST(config.RoomPost, r.createHandler)
	br.GET(config.RoomGet, r.getHandler)
	// br.DELETE("/:id", r.deleteHandler)
	// br.PUT(":id", r.updateHandler)
}

func NewRoomController(uc usecase.RoomUseCase, rg *gin.RouterGroup) *RoomController {
	return &RoomController{uc: uc, rg: rg}
}
