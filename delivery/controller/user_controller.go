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

func (u *UserController) getByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	rspPayload, err := u.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) Route() {
	ur := u.rg.Group(config.UserGroup)
	ur.POST(config.UserPost, u.createHandler)
	ur.GET(config.UserGet, u.getByIdHandler)
}
func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{uc: uc, rg: rg}
}
