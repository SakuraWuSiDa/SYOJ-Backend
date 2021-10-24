package helper

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// NextIndex 获取下一个题目index
func NextIndex(index string) string {
	s := StrToInt64(index[1:]) + 1
	return index[0:1] + strconv.FormatInt(s, 10)
}

// GetMD5OfStr 获取一个字符串的MD5
func GetMD5OfStr(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// RemoveTestDatas 删除一个文件夹下的.in .out .ans
func RemoveTestDatas(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		ext := filepath.Ext(name)
		if ext == ".in" || ext == ".out" || ext == ".ans" {
			os.Remove(filepath.Join(dir, name))
		}
	}
	return nil
}

// StoreZipFile 保存zip文件到指定文件夹
func StoreZipFile(zipFIle *zip.File, dest string) error {
	inFile, err := zipFIle.Open()
	if err != nil {
		return err
	}
	defer inFile.Close()
	outFile, err2 := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFIle.Mode())
	if err2 != nil {
		return err2
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, inFile)
	if err != nil {
		return err
	}
	return nil
}

// MD5 求文件的md5(每行去掉首尾的空格)
func MD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	var content string
	for {
		line, err := rd.ReadString('\n')
		content += strings.TrimSpace(line)
		if err != nil || io.EOF == err {
			break
		}
	}
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil)), nil
}

// FileSize 求文件大小
func FileSize(path string) uint {
	fileInfo, _ := os.Stat(path)
	return uint(fileInfo.Size())
}