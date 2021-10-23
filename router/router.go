package router

import (
	"github.com/XGHXT/SYOJ-Backend/api/user"
	"github.com/XGHXT/SYOJ-Backend/logger"
	"github.com/XGHXT/SYOJ-Backend/router/middleware"
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
	r.Use(middleware.CorsMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!")
	})
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 设置路由组
	v1 := r.Group("/api/v1")
	// 用户
	v1.POST("/user/register", user.RegisterHandler)
	v1.POST("/user/login", user.LoginHandler)
	v1.GET("/user/info", middleware.JWTAuthorMiddleware(), user.InfoHandler)
	v1.GET("/user/detail/:id", user.GetUserDetailHandler)
	v1.POST("/user/update",middleware.JWTAuthorMiddleware(), user.UpdateUserHandler)

	return r
}
