package service

import (
	"errors"
	"github.com/XGHXT/SYOJ-Backend/dao"
	"github.com/XGHXT/SYOJ-Backend/helper"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/router/middleware"
	"github.com/XGHXT/SYOJ-Backend/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	ErrorUserHasExists   = errors.New("该用户已经存在")
	ErrorUserNotExists   = errors.New("该用户不存在")
	ErrorInValidPassword = errors.New("密码错误")
)

func CreateUser(arg *model.CreateUserParams) (user *model.User, err error) {
	if ok := dao.CheckUserExists(arg.Username, arg.Email); ok {
		return nil, ErrorUserHasExists
	}
	u := &model.User{
		Gender:         2,
		StudentID:      "",
		Class:          "",
		CreatedAt:      time.Now(),
		Username:       arg.Username,
		Email:          arg.Email,
		Avatar:         "https://cn.gravatar.com/avatar/"+ helper.GetMD5OfStr(arg.Email) + "?s=200&d=mp&r=g",
		HashedPassword: arg.HashedPassword,
		IsAdmin:        0,
		IsSuperAdmin:   0,
	}
	user, err = dao.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Login(arg *model.LoginParams) (token string, user *model.User, err error) {
	user, err = dao.GetUser(arg.Username)

	// 1. 能否找到该用户
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, ErrorUserNotExists
		}
		return "", nil, err
	}

	// 2. 用户的密码是否正确
	err = util.CheckPassword(arg.Password, user.HashedPassword)
	if err != nil {
		return "", nil, ErrorInValidPassword
	}

	// 3. 发放 token
	token, err = middleware.ReleaseToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func UpdateUser(req *model.UpdateUserParams) (*model.User, error) {
	user, err := dao.GetUserByID(req.UserID)
	if err != nil {
		zap.L().Error("dao.GetUserByID failed", zap.Error(err))
		return nil, err
	}
	user.StudentID = req.StudentID
	user.Class = req.Class
	user.Gender = req.Gender
	user.Email = req.Email
	return dao.UpdateUser(user)
}

func GetUserDetails(userID int64) (*model.UserDetailResponse, error) {
	user, err := dao.GetUserByID(userID)
	if err != nil {
		zap.L().Error("dao.GetUserByID failed", zap.Error(err))
		return nil, err
	}

	accept, err := dao.GetPersonSolved(userID)
	if err != nil {
		zap.L().Error("dao.GetPersonSolved failed", zap.Error(err))
		return nil, err
	}

	submissionCount, err := dao.GetPersonSubmission(userID)
	if err != nil {
		zap.L().Error("dao.GetPersonSubmission failed", zap.Error(err))
		return nil, err
	}

	return &model.UserDetailResponse{
		UserID:          userID,
		Username:        user.Username,
		Email:           user.Email,
		StudentID:       user.StudentID,
		Class:           user.Class,
		Gender:          user.Gender,
		AcceptCount:     accept,
		SubmissionCount: submissionCount,
	}, nil
}
