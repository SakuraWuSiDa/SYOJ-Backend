package problem

import (
	"github.com/XGHXT/SYOJ-Backend/config"
	"github.com/XGHXT/SYOJ-Backend/handler"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/pkg"
	"github.com/XGHXT/SYOJ-Backend/pkg/errno"
	"github.com/XGHXT/SYOJ-Backend/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mime/multipart"
	"os"
	"path"
)

const (
	createProblemSuccessful = "创建题目成功"
	getProblemIdFailed      = "获取题目ID失败"
	getProblemSuccess       = "获取题目成功"
	invalidParams           = "请求参数错误"
	getProblemListSuccess   = "获取问题列表成功"
	deleteProblemSuccess    = "删除题目成功"
)

func CreateProblemHandler(ctx *gin.Context) {
	req := new(model.CreateProblemRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Error("api.CreateProblem failed", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			handler.SendBadRequest(ctx, errno.ErrValidation, err.Error())
		} else {
			handler.SendBadRequest(ctx, errno.ErrValidation, pkg.RemoveTopStruct(errs.Translate(pkg.Trans)))
		}
		return
	}
	problem, err := service.CreateProblem(req)
	if err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}
	handler.SendResponse(ctx, nil, gin.H{
		"problem": problem,
	})
}

func UpdateTestDataHandler(ctx *gin.Context) {
	index := ctx.DefaultPostForm("index", "")
	if index == "" {
		handler.SendBadRequest(ctx, errno.ErrValidation, nil)
		return
	}
	pd, err := service.GetProblemByIndex(index)
	if err != nil {
		handler.SendBadRequest(ctx, errno.ErrValidation, err.Error())
		return
	}
	var zipFile *multipart.FileHeader
	var errFile error
	if zipFile, errFile = ctx.FormFile("zip"); errFile != nil {
		handler.SendBadRequest(ctx, errno.ErrValidation, errFile.Error())
		return
	}
	dir := path.Join(config.Config.TestPath, index)
	os.MkdirAll(dir, os.ModePerm)
	// 保存压缩包
	zipPath := path.Join(dir, zipFile.Filename)
	err = ctx.SaveUploadedFile(zipFile, zipPath)
	if err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}

	if err := service.HandleTestDateZip(pd, dir, zipPath); err != nil {
		handler.SendError(ctx, errno.InternalServerError, err.Error())
		return
	}
	os.Remove(zipPath)
	handler.SendResponse(ctx, nil, gin.H{
		"problem": pd,
	})
}
