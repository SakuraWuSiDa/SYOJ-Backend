package router

import (
	"github.com/XGHXT/SYOJ-Backend/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// r.Use(middleware.CorsMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})
	r.Use(logger.GinLogger(), logger.GinRecovery(true))



	return r
}
