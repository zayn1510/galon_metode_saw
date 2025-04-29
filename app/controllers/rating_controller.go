package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"strconv"
)

type RatingController struct {
	service *services.RatingService
}

func NewRatingController() *RatingController {
	return &RatingController{
		service: services.NewRatingService(),
	}
}

func (c *RatingController) Create(ctx *gin.Context) {
	var req requests.CreateRatingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}

	data := req.ToModelRating()
	if err := c.service.Create(data); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success", req)
}

func (c *RatingController) Show(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	depotStr := ctx.DefaultQuery("depot", "0")
	depot, _ := strconv.Atoi(depotStr)
	data, err := c.service.FindAll(offset, limit, uint(depot))

	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	response := resources.GetRatingResource(data)
	resources.Success(ctx, "success", response)
}

func (c *RatingController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req requests.UpdateRatingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}
	if err := c.service.Update(&req, uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success", req)
}
func (c *RatingController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := c.service.Delete(uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success")
}
