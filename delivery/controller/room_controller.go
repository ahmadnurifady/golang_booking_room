package controller

import (
	"final-project/model"
	"final-project/usecase"
	"fmt"
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

func (r *RoomController) getByRoomtypeHandler(ctx *gin.Context) {
	roomType := ctx.Query("roomtype")
	var rspPayload model.Room
	var err error

	if roomType != "" {
		rspPayload, err = r.uc.FindByRoomType(roomType)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "room yang dimaksud tidak ditemukan"})
		}
	}

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "roomType yang dimaksud tidak ditemukan"})
	}

	fmt.Println(rspPayload.RoomType)
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
	br := r.rg.Group("/rooms")
	br.POST("/create", r.createHandler)
	br.GET("/:id", r.getHandler)
	br.GET("/", r.getByRoomtypeHandler)
	br.DELETE("/:id", r.deleteHandler)
	br.PUT(":id", r.updateHandler)
}

func NewRoomController(uc usecase.RoomUseCase, rg *gin.RouterGroup) *RoomController {
	return &RoomController{uc: uc, rg: rg}
}
