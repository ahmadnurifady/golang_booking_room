package controller

import (
	"final-project-booking-room/config"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/usecase"
	"final-project-booking-room/utils/common"
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

	rspPayload, err := b.uc.RegisterNewBooking(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) UpdateStatusHandler(ctx *gin.Context) {
	var payload dto.Approval
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rspPayload, err := b.uc.UpdateStatusBookAndRoom(payload.BookingDetailId, payload.Approval)
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

	rspPayload, err := b.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getByStatusHandler(ctx *gin.Context) {
	status := ctx.Param("status")
	if status == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Status can't be empty")
		return
	}
	rspPayload, err := b.uc.ViewAllBookingByStatus(status)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) getAllHandler(ctx *gin.Context) {
	rspPayload, err := b.uc.ViewAllBooking()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) Route() {
	bc := b.rg.Group(config.BookingGroup)
	bc.POST(config.BookingPost, b.createHandler)
	bc.PUT(config.Approval, b.UpdateStatusHandler)
	bc.GET(config.BookingGetAll, b.getAllHandler)
	bc.GET(config.BookingGet, b.getHandler)
	bc.GET(config.BookingGetAllByStatus, b.getByStatusHandler)
}

func NewBookingController(
	uc usecase.BookingUseCase,
	rg *gin.RouterGroup,
) *BookingController {
	return &BookingController{
		uc: uc,
		rg: rg}
}
