package dao

import (
	"fmt"
	"github.com/XGHXT/SYOJ-Backend/model"
	"github.com/XGHXT/SYOJ-Backend/pkg/mysql"
)

func CheckUserExists(username string, email string) bool {
	var user model.User
	mysql.DB.Where("username = ? ", username).First(&user)
	fmt.Println(user.ID)
	if user.ID != 0 {
		return true
	}
	mysql.DB.Where("email = ? ", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func CreateUser(user *model.User) (*model.User, error) {
	err := mysql.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(Username string) (*model.User, error) {
	user := new(model.User)
	err := mysql.DB.Where("username = ?", Username).First(&user).Error
	return user, err
}

func GetUserByID(id int64) (*model.User, error) {
	user := new(model.User)
	err := mysql.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUserList(offset, limit int) ([]*model.User, error) {
	users := []*model.User{}
	err := mysql.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func GetUserSize() int64 {
	var total int64
	mysql.DB.Model(&model.User{}).Count(&total)
	return total
}

func UpdateUser(u *model.User) (*model.User, error) {
	err := mysql.DB.Save(&u).Error
	return u, err
}

func DeleteUser(userID int64) error {
	err := mysql.DB.Delete(&model.User{}, userID).Error
	return err
}

