package controller

import (
	"final-project/config"
	"final-project/model"
	"final-project/usecase"
	"final-project/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (u *UserController) getAllHandler(ctx *gin.Context) {

	rspPayload, err := u.uc.ViewAllUser()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) createHandler(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := u.uc.RegisterNewUser(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) UpdateUserHandler(ctx *gin.Context) {

	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := u.uc.UpdateUserById(payload.Id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	rspPayload, err := u.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) DeleteByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	_, err := u.uc.DeleteUser(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", nil)
}

func (u *UserController) Route() {
	ur := u.rg.Group(config.UserGroup)
	ur.POST(config.UserPost, u.createHandler)
	ur.PUT(config.UserUpdate, u.UpdateUserHandler)
	ur.GET(config.UserGet, u.getByIdHandler)
	ur.DELETE(config.UserDelete, u.DeleteByIdHandler)
	ur.GET(config.UserGetAll, u.getAllHandler)

}
func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{uc: uc, rg: rg}
}
