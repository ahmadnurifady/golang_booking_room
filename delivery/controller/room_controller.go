package controller

import (
	"final-project-booking-room/model"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"

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
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createRoom, err := r.uc.RegisterNewRoom(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "ok", createRoom)
}

func (r *RoomController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id cant be empty")
		return
	}

	rspPayload, err := r.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "ok", rspPayload)
}

func (r *RoomController) getByRoomtypeHandler(ctx *gin.Context) {
	roomType := ctx.Query("roomtype")
	// var rspPayload model.Room
	// var err error

	if roomType == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "roomtype cant be empty")
		return
	}
	rspPayload, err := r.uc.FindByRoomType(roomType)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	// fmt.Println(rspPayload.RoomType)
	common.SendSingleResponse(ctx, "Ok", rspPayload)

}

func (r *RoomController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	err := r.uc.DeleteById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", err)
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
