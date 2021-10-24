package dao

import (
	"github.com/XGHXT/SYOJ-Backend/helper"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/pkg/mysql"
	"time"
)


func CreateProblem(req *model.CreateProblemRequest) (*model.Problem, error) {
	var creatIndex string
	var pd model.Problem
	mysql.DB.Where("is_open = ?", req.IsOpen).Order("id DESC").First(&pd)
	if pd.ID != 0 {
		creatIndex = helper.NextIndex(pd.Index)
	}else {
		if req.IsOpen {
			creatIndex = "P1000"
		}else {
			creatIndex = "C1000"
		}
	}

	problem := &model.Problem{
		Index:       creatIndex,
		Title:       req.Title,
		TimeLimit:   req.TimeLimit,
		MemoryLimit: req.MemoryLimit,
		Author:      req.Author,
		CreatedAt:   time.Now(),
		Source:      req.Source,
		Background:  req.Background,
		Statement:   req.Statement,
		Input:       req.Input,
		Output:      req.Output,
		ExamplesIn:  req.ExamplesIn,
		ExamplesOut: req.ExamplesOut,
		Hint:        req.Hint,
		IsOpen:      req.IsOpen,
		Tags:        req.Tags,
		Testdatas:   nil,
	}
	err := mysql.DB.Create(problem).Error
	if err != nil {
		return problem, err
	}
	return problem, err
}

func GetProblemByIndex(index string) (*model.Problem, error) {
	problem := new(model.Problem)
	err := mysql.DB.Where("index = ?", index).First(problem).Error
	return problem, err

}

func UpdateProblem(pd *model.Problem) (*model.Problem, error) {
	err := mysql.DB.Save(&pd).Error
	return pd, err
}

func GetProblemList(offset, limit int) ([]*model.Problem, error) {
	problems := []*model.Problem{}
	err := mysql.DB.Offset(offset).Limit(limit).Find(&problems).Error
	return problems, err
}

func GetProblemSize() int64 {
	var count int64
	mysql.DB.Model(&model.Problem{}).Count(&count)
	return count
}

func DeleteProblem(id int64) error {
	err := mysql.DB.Delete(&model.Problem{}, id).Error
	return err
}
