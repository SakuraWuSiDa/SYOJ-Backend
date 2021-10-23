package middleware

import (
	"errors"
	"github.com/XGHXT/SYOJ-Backend/handler"
	"github.com/XGHXT/SYOJ-Backend/pkg/errno"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	ContextUserNameKey = "username"
	ContextUserIDKey = "user_id"
)

var (
	ErrorNeedLogin    = errors.New("需要进行登录")
	ErrorInvalidToken = errors.New("不正确的 Token")
)

// JWTAuthorMiddleware 基于 JWT 认证的中间件
func JWTAuthorMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带 Token 的三种方式：1. 放在请求头里面，2. 放在请求体里面，3. 放在 URI
		// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			handler.SendError(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}

		// 按空格进行分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			handler.SendError(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString，我们使用之前定义好的解析 JWT 的函数来解析
		claims, err := ParseToekn(parts[1])
		if err != nil {
			handler.SendError(ctx, errno.ErrTokenInvalid, err.Error())
			ctx.Abort()
			return
		}

		// 将当前请求的 username 信息保存到请求的上下文中
		ctx.Set(ContextUserNameKey, claims.UserName)
		ctx.Set(ContextUserIDKey, claims.UserID)
		ctx.Next() // 后续的处理函数可以通过 ctx.Get("username") 来获取当前请求的用户信息
	}
}
