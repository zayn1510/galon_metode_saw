package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/controllers"
)

func setUpRouterPing(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome",
		})
	})

}

func SetUpRouterKriteria(router *gin.RouterGroup) {
	kriteria := router.Group("kriteria")
	kriteria.GET("", controllers.NewKriteriaController().Show)
	kriteria.POST("", controllers.NewKriteriaController().Create)
	kriteria.PUT("/:id", controllers.NewKriteriaController().Update)
	kriteria.DELETE("/:id", controllers.NewKriteriaController().Delete)
}

func SetUpRouterDepot(router *gin.RouterGroup) {
	depot := router.Group("depot")
	depot.GET("", controllers.NewDepotController().Show)
	depot.GET("/preview/:filename", controllers.NewDepotController().PreviewFile)
	depot.GET("/:id", controllers.NewDepotController().DetailDepotById)
	depot.POST("", controllers.NewDepotController().Create)
	depot.POST("/update/:id", controllers.NewDepotController().Update)
	depot.DELETE("/:id", controllers.NewDepotController().Delete)

}

func SetUpRouterUsers(router *gin.RouterGroup) {
	users := router.Group("users")
	users.GET("", controllers.NewUsersController().Show)
	users.GET("/:id", controllers.NewUsersController().UserDetail)
	users.POST("", controllers.NewUsersController().Create)
	users.PUT("/:id", controllers.NewUsersController().Update)
	users.DELETE("/:id", controllers.NewUsersController().Delete)
	users.GET("alternatif/:id", controllers.NewUsersController().DataAlternatif)
}
func SetUpRouterUsersLocation(router *gin.RouterGroup) {
	users := router.Group("user-locations")
	users.GET("", controllers.NewUserlocationController().Show)
	users.POST("", controllers.NewUserlocationController().Create)
	users.PUT("/:id", controllers.NewUserlocationController().Update)
	users.DELETE("/:id", controllers.NewUserlocationController().Delete)
}
func SetUpRouterKecamatan(router *gin.RouterGroup) {
	users := router.Group("kecamatan")
	users.GET("", controllers.NewKecamatanController().Show)
	users.POST("", controllers.NewKecamatanController().Create)
	users.PUT("/:id", controllers.NewKecamatanController().Update)
	users.DELETE("/:id", controllers.NewKecamatanController().Delete)
	users.GET("/jumlah-depot", controllers.NewKecamatanController().JumlahDepot)
}

func SetUpRouterRating(router *gin.RouterGroup) {
	users := router.Group("rating")
	users.GET("", controllers.NewRatingController().Show)
	users.POST("", controllers.NewRatingController().Create)
	users.PUT("/:id", controllers.NewRatingController().Update)
	users.DELETE("/:id", controllers.NewRatingController().Delete)
}

func SetUpRouterStat(router *gin.RouterGroup) {
	stat := router.Group("statistik")
	stat.GET("", controllers.NewStatsController().GetStats)
}
func SetUpRouterAuth(router *gin.RouterGroup) {
	users := router.Group("auth")
	users.POST("/login", controllers.NewAuthController().Login)
}
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	setUpRouterPing(api)
	SetUpRouterKriteria(api)
	SetUpRouterDepot(api)
	SetUpRouterUsers(api)
	SetUpRouterUsersLocation(api)
	SetUpRouterKecamatan(api)
	SetUpRouterRating(api)
	SetUpRouterStat(api)
	SetUpRouterAuth(api)
}
