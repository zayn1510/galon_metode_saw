package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/controllers"
	"github.com/zayn1510/goarchi/app/middleware"
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
	router.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/pung", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pung",
		})
	})

}

func SetUpRouterKriteria(router *gin.RouterGroup) {

	kriteria := router.Group("kriteria")
	kriteria.Use(middleware.JWTMiddleware())
	kriteria.GET("", controllers.NewKriteriaController().Show)
	kriteria.POST("", controllers.NewKriteriaController().Create)
	kriteria.PUT("/:id", controllers.NewKriteriaController().Update)
	kriteria.DELETE("/:id", controllers.NewKriteriaController().Delete)
}

func SetUpRouterDepot(router *gin.RouterGroup) {
	depot := router.Group("depot")
	depot.Use(middleware.JWTMiddleware())
	depot.GET("", controllers.NewDepotController().Show)
	depot.GET("/preview/:filename", controllers.NewDepotController().PreviewFile)
	depot.GET("/:id", controllers.NewDepotController().DetailDepotById)
	depot.POST("", controllers.NewDepotController().Create)
	depot.POST("/update/:id", controllers.NewDepotController().Update)
	depot.DELETE("/:id", controllers.NewDepotController().Delete)

}

func SetUpRouterUsers(router *gin.RouterGroup) {
	users := router.Group("users")
	users.Use(middleware.JWTMiddleware())
	users.GET("", controllers.NewUsersController().Show)
	users.GET("/:id", controllers.NewUsersController().UserDetail)
	users.GET("/by/:username", controllers.NewUsersController().UserDetailByUsername)
	users.POST("", controllers.NewUsersController().Create)
	users.PUT("/:id", controllers.NewUsersController().Update)
	users.DELETE("/:id", controllers.NewUsersController().Delete)
	users.GET("alternatif/:id", controllers.NewUsersController().DataAlternatif)
}
func SetUpRouterUsersLocation(router *gin.RouterGroup) {
	users := router.Group("user-locations")
	users.Use(middleware.JWTMiddleware())
	users.GET("", controllers.NewUserlocationController().Show)
	users.GET("/:id", controllers.NewUserlocationController().CheckUserLocation)
	users.POST("", controllers.NewUserlocationController().Create)
	users.PUT("/:id", controllers.NewUserlocationController().Update)
	users.DELETE("/:id", controllers.NewUserlocationController().Delete)
}
func SetUpRouterKecamatan(router *gin.RouterGroup) {
	kecamatan := router.Group("kecamatan")
	kecamatan.Use(middleware.JWTMiddleware())
	kecamatan.GET("", controllers.NewKecamatanController().Show)
	kecamatan.POST("", controllers.NewKecamatanController().Create)
	kecamatan.PUT("/:id", controllers.NewKecamatanController().Update)
	kecamatan.DELETE("/:id", controllers.NewKecamatanController().Delete)
	kecamatan.GET("/jumlah-depot", controllers.NewKecamatanController().JumlahDepot)
}

func SetUpRouterRating(router *gin.RouterGroup) {
	rating := router.Group("rating")
	rating.Use(middleware.JWTMiddleware())
	rating.GET("", controllers.NewRatingController().Show)
	rating.POST("", controllers.NewRatingController().Create)
	rating.PUT("/:id", controllers.NewRatingController().Update)
	rating.DELETE("/:id", controllers.NewRatingController().Delete)
}

func SetUpRouterStat(router *gin.RouterGroup) {
	stat := router.Group("statistik")
	stat.Use(middleware.JWTMiddleware())
	stat.GET("", controllers.NewStatsController().GetStats)
}
func SetUpRouterAuth(router *gin.RouterGroup) {
	auth := router.Group("auth")
	auth.POST("/login", controllers.NewAuthController().Login)
	auth.Use(middleware.JWTMiddleware())
	{
		auth.PUT("/update-password", controllers.NewAuthController().UpdateNewPassword)
	}
}
func setUpRouterLoginLogs(router *gin.RouterGroup) {
	loginLogs := router.Group("login-logs")
	loginLogs.Use(middleware.JWTMiddleware())
	loginLogs.GET("", controllers.NewAuthController().GetAllLoginLogs)
}
func SetUpRouterUserRating(router *gin.RouterGroup) {
	rating := router.Group("user")
	rating.GET("/rating", controllers.NewRatingController().Show)
	rating.GET("/depot", controllers.NewDepotController().Show)
	rating.POST("/signup", controllers.NewUsersController().Create)
	rating.GET("/alternatif/:id", controllers.NewUsersController().DataAlternatif)
	depot := router.Group("depot-user")
	depot.GET("/preview/:filename", controllers.NewDepotController().PreviewFile)
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/")
	setUpRouterPing(api)
	SetUpRouterKriteria(api)
	SetUpRouterDepot(api)
	SetUpRouterUsers(api)
	SetUpRouterUsersLocation(api)
	SetUpRouterKecamatan(api)
	SetUpRouterRating(api)
	SetUpRouterStat(api)
	SetUpRouterAuth(api)
	setUpRouterLoginLogs(api)
	SetUpRouterUserRating(api)
}
