package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"strconv"
	"strings"
)

type UserlocationController struct {
	service *services.UserlocationService
}

func NewUserlocationController() *UserlocationController {
	return &UserlocationController{
		service: services.NewUserLocationService(),
	}
}

func (c *UserlocationController) Create(ctx *gin.Context) {
	var req requests.CreateUserLocationRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}

	data := req.ToUserLocation()
	if err := c.service.Create(data); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success", req)
}

func (c *UserlocationController) Show(ctx *gin.Context) {
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
	response := resources.GetUserLocationResource(data)
	resources.Success(ctx, "success", response)
}
func (c *UserlocationController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req requests.UpdateUserLocationRequest
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
func (c *UserlocationController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := c.service.Delete(uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success")
}
func (c *UserlocationController) CheckUserLocation(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	result, err := c.service.IsExistIdUser(uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "ID tidak ditemukan") {
			resources.NotFound(ctx, err)
			return
		}
		resources.InternalError(ctx, err)
		return
	}
	response := resources.NewUserLocationResource(*result)
	resources.Success(ctx, "success", response)
}
