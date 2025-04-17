package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"strconv"
)

type KriteriaController struct {
	service *services.KriteriaService
}

func NewKriteriaController() *KriteriaController {
	return &KriteriaController{
		service: services.NewKriteriaService(),
	}
}

func (c *KriteriaController) Create(ctx *gin.Context) {
	var req requests.CreateKriteriaRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}

	data := req.ToModelKritera()
	if err := c.service.Create(data); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success", req)
}

func (c *KriteriaController) Show(ctx *gin.Context) {
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
	response := resources.GetKriteriaResource(data)
	resources.Success(ctx, "success", response)
}

func (c *KriteriaController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req requests.UpdateKriteriaRequest
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
func (c *KriteriaController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := c.service.Delete(uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success")
}
