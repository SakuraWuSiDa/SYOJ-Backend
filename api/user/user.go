package user

import (
	"errors"
	"github.com/XGHXT/SYOJ-Backend/dao"
	"github.com/XGHXT/SYOJ-Backend/handler"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/pkg"
	"github.com/XGHXT/SYOJ-Backend/pkg/errno"
	"github.com/XGHXT/SYOJ-Backend/router/middleware"
	"github.com/XGHXT/SYOJ-Backend/service"
	"github.com/XGHXT/SYOJ-Backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

// RegisterHandler 用户注册
func RegisterHandler(ctx *gin.Context) {
	req := new(model.CreateUserRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("Sign up Handler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler.SendBadRequest(ctx, errno.ErrValidation, err.Error())
		} else {
			handler.SendBadRequest(ctx, errno.ErrValidation, pkg.RemoveTopStruct(errs.Translate(pkg.Trans)))
		}
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}

	arg := &model.CreateUserParams{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := service.CreateUser(arg)
	if err != nil {
		zap.L().Error("server CreateUser function failed, ", zap.Error(err))
		if errors.Is(err, service.ErrorUserHasExists) {
			handler.SendForbidden(ctx, errno.ErrValidation, err.Error())
		} else {
			handler.SendError(ctx, errno.InternalServerError, err.Error())
		}
		return
	}

	token, err := middleware.ReleaseToken(user)
	if err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}
	res := model.CreateUserResponse{
		Username:  user.Username,
		UserID:    user.ID,
	}
	handler.SendResponse(ctx, nil, gin.H{
		"user":         res,
		"access_token": token,
	})
}

func LoginHandler(ctx *gin.Context) {
	req := new(model.LoginRequest)
	// 校验数据
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("Login Handler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler.SendBadRequest(ctx, errno.ErrValidation, err.Error())
		} else {
			handler.SendBadRequest(ctx, errno.ErrValidation, pkg.RemoveTopStruct(errs.Translate(pkg.Trans)))
		}
		return
	}

	arg := &model.LoginParams{
		Username:  req.Username,
		Password:  req.Password,
	}

	token, user, err := service.Login(arg)
	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler.SendBadRequest(ctx, errno.ErrValidation, err.Error())
		} else {
			handler.SendBadRequest(ctx, errno.ErrValidation, errs.Translate(pkg.Trans))
		}
		return
	}

	res := model.LoginResponse{
		User: model.CreateUserResponse{
			Username:  user.Username,
			UserID:    user.ID,
		},
		AccessToekn: token,
	}
	ctx.SetCookie("token", "true", 3600, "/", "localhost", false, true)
	handler.SendResponse(ctx, nil, gin.H{
		"user":  res,
	})
}

func InfoHandler(ctx *gin.Context) {
	username, _ := ctx.Get(middleware.ContextUserNameKey)
	user, err := dao.GetUser(username.(string))
	if err != nil {
		zap.L().Error("mysql.GetUser failed", zap.Error(err))
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}
	handler.SendResponse(ctx, nil, gin.H{
		"user":  user,
	})
}

func GetUserDetailHandler(ctx *gin.Context)  {
	uidStr := ctx.Param("id")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil || uid == 0 {
		handler.SendError(ctx, errno.ErrUserIDInvalid, err.Error())
		return
	}
	userDetail, err := service.GetUserDetails(uid)
	if err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}
	handler.SendResponse(ctx, nil, gin.H{
		"userDetail":  userDetail,
	})
}

func UpdateUserHandler(ctx *gin.Context) {
	req := new(model.UpdateUserParams)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("api.UpdateUser failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler.SendBadRequest(ctx, errno.ErrValidation, err.Error())
		} else {
			handler.SendBadRequest(ctx, errno.ErrValidation, pkg.RemoveTopStruct(errs.Translate(pkg.Trans)))
		}
		return
	}
	userID, _ := ctx.Get(middleware.ContextUserIDKey)
	if userID != req.UserID {
		handler.SendBadRequest(ctx, errno.ErrValidation, nil)
		return
	}
	user, err := service.UpdateUser(req)
	if err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}
	handler.SendResponse(ctx, nil, gin.H{
		"user": user,
	})
}