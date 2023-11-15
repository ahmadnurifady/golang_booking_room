package controller

import (
	"final-project-booking-room/config"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	uc usecase.BookingUseCase
	rg *gin.RouterGroup
}

func (b *BookingController) createHandler(ctx *gin.Context) {
	var payload dto.BookingRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("ini payload booking details :", payload.BoookingDetails)

	rspPayload, err := b.uc.RegisterNewBooking(payload)
	fmt.Println(rspPayload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Booking ID can't be empty")
		return
	}

	userId := ctx.MustGet(config.UserSesion).(string)
	rspPayload, err := b.uc.FindById(id, userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) Route() {
	bc := b.rg.Group(config.BookingGroup)
	bc.POST(config.BookingPost, b.createHandler)
	bc.GET(config.BookingGet, b.getHandler)
}

func NewBookingController(
	uc usecase.BookingUseCase,
	rg *gin.RouterGroup,
) *BookingController {
	return &BookingController{
		uc: uc,
		rg: rg}
}
