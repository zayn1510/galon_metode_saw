package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/middleware"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"github.com/zayn1510/goarchi/core/tools"
	"log"
	"math"
	"strconv"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		service: services.NewAuthService(),
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req requests.AuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	err, validateErr := requests.Validate(req)
	if err != nil {
		resources.BadRequest(ctx, validateErr)
		return
	}

	user, err := c.service.Login(&req)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}

	token, errToken := middleware.GenerateJWT(user.Username, int64(user.ID))
	if errToken != nil {
		resources.InternalError(ctx, errToken)
		return
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Panic di goroutine log:", r)
			}
		}()
		ip := ctx.ClientIP()

		userAgent := ctx.GetHeader("User-Agent")
		geo, err := tools.GetIPDetails(userAgent)
		if err != nil {
			log.Println("Error getting IP details:", err)
			return
		}
		result := &requests.LoginlogsRequest{
			UserId:    user.ID,
			Nama:      user.Nama,
			Role:      user.Role,
			Username:  user.Username,
			IPAddress: ip,
			ISP:       geo.ISP,
			Device:    geo.Device,
			Browser:   geo.Browser,
			Country:   geo.Country,
			Platform:  geo.OS,
			City:      geo.City,
			Token:     token,
		}
		if err := c.service.SaveLoginLogs(result.ToModel()); err != nil {
			log.Println("Gagal simpan login log:", err)
		}
	}()
	response := &resources.AuthResource{
		Userid:   user.ID,
		Username: user.Username,
		Token:    token,
		Role:     user.Role,
	}
	resources.Success(ctx, "success", response)
}

func (c *AuthController) GetAllLoginLogs(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	filterStr := ctx.DefaultQuery("filter", "")
	data, err := c.service.GetLoginLogs(offset, limit, filterStr)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	response := resources.GetLoginLogsResource(data)
	totaldata, err := c.service.CountLoginLogs()
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	total := int(math.Ceil(float64(totaldata) / float64(limit)))
	resources.SuccessWithPaginaition(ctx, "success", response, &total, &page, &limit)
}
