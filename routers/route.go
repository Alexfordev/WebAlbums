package routers

import (
	"net/http"
	"time"

	"github.com/alexfordev/WebAlbums/middlewares"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "API is working fine",
	})
}

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(middlewares.Logger(), gin.Recovery())

	config := cors.Config{
		AllowMethods:     []string{"GET", "OPTIONS", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Create a limiter with a rate limit of 5 requests per second
	rateLimit := tollbooth.NewLimiter(1000, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second})
	// Limit requests by IP address
	rateLimit.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})

	router.Use(middlewares.LimitHandler(rateLimit))

	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	// m.SetMetricPath("/metrics")
	m.SetMetricPath("/bgrgb/v1/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

	healthV2 := router.Group("/bgrgb/v1")
	{
		healthV2.GET("/", HealthCheck)
	}

	// configsV1 := router.Group("/bgrgb/v1/configs").Use(middlewares.CheckToken())
	// {
	// 	configsController := controllers.NewConfigsController()
	// 	configsV1.POST("", configsController.Create)
	// 	configsV1.GET("", configsController.GetAll)
	// }

	// usersV1 := router.Group("/bgrgb/v1/users")
	// {
	// 	usersController := controllers.NewUsersController()
	// 	usersV1.POST("/action/register", usersController.Create)
	// 	usersV1.GET("", usersController.GetAll)
	// }

	return
}
