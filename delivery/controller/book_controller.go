package controller

import (
	// "final-project/model/dto"
	"final-project/usecase"
	"final-project/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	uc usecase.BookingUseCase
	rg *gin.RouterGroup
}

// func (b *BookingController) createHandler(ctx *gin.Context) {
// 	var payload dto.BookingRequestDto
// 	if err := ctx.ShouldBindJSON(&payload); err != nil {
// 		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	rspPayload, err := b.uc.RegisterNewBooking(payload)
// 	if err != nil {
// 		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	common.SendCreateResponse(ctx, "Ok", rspPayload)
// }

func (b *BookingController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Booking ID can't be empty")
		return
	}
	userId := ctx.Param("userId")

	rspPayload, err := b.uc.FindById(id, userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (b *BookingController) Route() {
	br := b.rg.Group("/booking")
	// br.POST("/", b.createHandler)
	br.GET("/:id", b.getHandler)
}

func NewBookingController(
	uc usecase.BookingUseCase,
	rg *gin.RouterGroup,
) *BookingController {
	return &BookingController{
		uc: uc,
		rg: rg}
}
