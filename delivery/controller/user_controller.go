package controller

import (
	"net/http"
	"project-final/config"
	"project-final/delivery/middleware"
	"project-final/model"
	"project-final/usecase"
	"project-final/utils/common"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc             usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
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
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (u *UserController) UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payload.Id = id

	rspPayload, err := u.uc.UpdateUserById(payload.Id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
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
	ur.POST(config.UserPost, u.authMiddleware.RequireToken("admin"), u.createHandler)
	ur.PUT(config.UserUpdate, u.authMiddleware.RequireToken("admin", "employee"), u.UpdateUserHandler)
	ur.GET(config.UserGet, u.authMiddleware.RequireToken("admin"), u.getByIdHandler)
	ur.DELETE(config.UserDelete, u.authMiddleware.RequireToken("admin"), u.DeleteByIdHandler)
	ur.GET(config.UserGetAll, u.authMiddleware.RequireToken("admin"), u.getAllHandler)

}
func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup, authmiddleware middleware.AuthMiddleware) *UserController {
	return &UserController{uc: uc, rg: rg, authMiddleware: authmiddleware}
}
