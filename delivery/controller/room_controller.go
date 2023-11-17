package controller

import (
	"final-project-booking-room/config"
	"final-project-booking-room/delivery/middleware"
	"final-project-booking-room/model"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"

	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	uc             usecase.RoomUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
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

func (r *RoomController) getAllRoom(ctx *gin.Context) {
	rspPayload, err := r.uc.ViewAllRooms()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
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
		getAll, err := r.uc.ViewAllRooms()
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		common.SendSingleResponse(ctx, "Ok", getAll)
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

	_, err := r.uc.DeleteById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", nil)
}

func (r *RoomController) changeStatusHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	err := r.uc.ChangeRoomStatus(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Room status has changed to available", err)
}

func (r *RoomController) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	var roomUpdate model.Room

	roomUpdate.Id = id

	err := ctx.ShouldBindJSON(&roomUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomUpdate, err = r.uc.UpdateById(roomUpdate.Id, roomUpdate)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", roomUpdate)
}

func (r *RoomController) Route() {
	br := r.rg.Group(config.RoomGroup)
	br.POST(config.RoomPost, r.createHandler)
	br.GET(config.RoomGetByroomType, r.getByRoomtypeHandler)
	br.GET(config.RoomGetAll, r.authMiddleware.RequireToken("admin"), r.getAllRoom)
	br.GET(config.RoomGetById, r.getHandler)
	br.DELETE(config.RoomDelete, r.authMiddleware.RequireToken("admin"), r.deleteHandler)
	br.PUT(config.RoomUpdate, r.authMiddleware.RequireToken("admin"), r.updateHandler)
	br.PUT(config.RoomUpdateStatus, r.changeStatusHandler)
}

func NewRoomController(uc usecase.RoomUseCase, rg *gin.RouterGroup, authmiddleware middleware.AuthMiddleware) *RoomController {
	return &RoomController{uc: uc, rg: rg, authMiddleware: authmiddleware}
}
