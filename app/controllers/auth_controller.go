package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"github.com/zayn1510/goarchi/core/tools"
	"log"
	"net/http"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController() AuthController {
	return AuthController{
		service: services.NewAuthService(),
	}
}

func (c AuthController) Login(ctx *gin.Context) {
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
	ip := ctx.ClientIP()

	userAgent := ctx.GetHeader("User-Agent")
	geo, err := tools.GetIPDetails(userAgent)
	if err != nil {
		log.Println("Error getting IP details:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch IP details"})
		return
	}

	result := &requests.LoginlogsRequest{
		UserId:    (user.ID),
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
	}
	toModelLoginLogs := result.ToModel()
	if err := c.service.SaveLoginLogs(toModelLoginLogs); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success", result)

}
