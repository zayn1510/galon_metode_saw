package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"strconv"
)

type DepotController struct {
	service *services.DepotService
}

func NewDepotController() *DepotController {
	return &DepotController{
		service: services.NewDepotService(),
	}
}

func (c *DepotController) Create(ctx *gin.Context) {
	var req requests.CreateDepotRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}

	data := req.ToDepot()
	if err := c.service.Create(data); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success", req)
}

func (c *DepotController) Show(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	data, err := c.service.FindAll(offset, limit)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	response := resources.GetDepotResource(data)
	resources.Success(ctx, "success", response)
}

func (c *DepotController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req requests.UpdateDepotRequest
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
func (c *DepotController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := c.service.Delete(uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success")
}
