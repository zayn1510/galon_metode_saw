package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
)

type StatsController struct {
	service *services.StatsService
}

func NewStatsController() *StatsController {
	return &StatsController{
		service: services.NewStatsService(),
	}
}

func (c *StatsController) GetStats(ctx *gin.Context) {
	stats := c.service.GetStatistik()
	resources.Success(ctx, "success", stats)
}
