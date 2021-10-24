package router

import (
	"github.com/XGHXT/SYOJ-Backend/api/problem"
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

	// ---基础
	v1 := r.Group("/api/v1")
	// 用户
	v1.POST("/user/register", user.RegisterHandler)
	v1.POST("/user/login", user.LoginHandler)
	v1.GET ("/user/info", middleware.JWTAuthorMiddleware(), user.InfoHandler)
	v1.GET ("/user/detail/:id", user.GetUserDetailHandler)
	v1.POST("/user/update",middleware.JWTAuthorMiddleware(), user.UpdateUserHandler)
	// 题目
	//v1.POST("/problems", api.CreateProblemHandler)


	

	// ---管理端
	admin := r.Group("/api/admin")
	// 题目
	admin.POST("/problem/create", problem.CreateProblemHandler)
	admin.POST("/problem/testdata/update", problem.UpdateTestDataHandler)


	return r
}
