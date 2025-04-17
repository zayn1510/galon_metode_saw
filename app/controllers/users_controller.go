package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"strconv"
)

type UsersController struct {
	service *services.UsersService
}

func NewUsersController() *UsersController {
	return &UsersController{
		service: services.NewUsersService(),
	}
}

func (c *UsersController) Create(ctx *gin.Context) {
	var req requests.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}

	data, err := req.ToUser()
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	if err := c.service.Create(data); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success", req)
}

func (c *UsersController) Show(ctx *gin.Context) {
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
	response := resources.GetUsersResource(data)
	resources.Success(ctx, "success", response)
}

func (c *UsersController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req requests.UpdateUserRequest
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
func (c *UsersController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	if err := c.service.Delete(uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success")
}

func (c *UsersController) DataAlternatif(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	alternatif, err := services.NewAlternatifService().ShowAlternatif(uint(id))
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success", alternatif)
}
