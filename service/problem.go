package service

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"github.com/XGHXT/SYOJ-Backend/dao"
	"github.com/XGHXT/SYOJ-Backend/helper"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/pkg"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	ErrorProblemNotExist = errors.New("该题目不存在")
)

func CreateProblem(req *model.CreateProblemRequest) (*model.Problem, error) {
	rsp, err := dao.CreateProblem(req)
	if err != nil {
		return nil, err
	}

	return rsp, err
}

func GetProblemByIndex(index string) (*model.Problem, error) {
	problem, err := dao.GetProblemByIndex(index)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrorProblemNotExist
		} else {
			return nil, err
		}
	}
	return problem, nil
}

func HandleTestDateZip(problem *model.Problem, dir string, zipPath string) error {
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	input := make(map[int]*zip.File)
	output := make(map[int]*zip.File)
	for _, f := range zipReader.File {
		var id, base = 0, 1
		for i := strings.LastIndex(f.Name, ".") - 1; i >= 0; i-- {
			ch := int(f.Name[i]) - int('0')
			if ch < 0 || ch > 9 {
				break
			}
			id += base * ch
			base *= 10
		}
		if path.Ext(f.Name) == ".in" {
			input[id] = f
		} else if path.Ext(f.Name) == ".out" || path.Ext(f.Name) == ".ans" {
			output[id] = f
		}
	}
	testdatas, tmp := make(pkg.H), make(pkg.H)
	for k, in := range input {
		if out, ok := output[k]; ok {
			inPath := path.Join(dir, in.Name)
			outPath := path.Join(dir, out.Name)
			if err := helper.StoreZipFile(in, inPath); err != nil {
				return err
			}
			if err := helper.StoreZipFile(out, outPath); err != nil {
				return err
			}
			md5, err := helper.MD5(outPath)
			if err != nil {
				return err
			}
			tmp[strconv.Itoa(k)] = pkg.H{
				"input_name":          in.Name,
				"output_name":         out.Name,
				"input_size":          helper.FileSize(inPath),
				"output_size":         helper.FileSize(outPath),
				"stripped_output_md5": md5,
			}
		}
	}
	testdatas["test_cases"] = tmp
	js, _ := json.Marshal(testdatas)
	problem.Testdatas = js
	if _, err := dao.UpdateProblem(problem); err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(dir, "info"), js, os.ModePerm); err != nil {
		return err
	}
	return nil
}


// GetProblemList 获取 problem 的列表
func GetProblemList(offset, limit int) ([]*model.Problem, int64, error) {
	problems, err := dao.GetProblemList(offset, limit)
	total := dao.GetProblemSize()

	return problems, total, err
}

// DeleteProblem 删除题目
func DeleteProblem(id int64) error {
	return dao.DeleteProblem(id)
}
