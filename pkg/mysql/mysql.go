package mysql

import (
	"fmt"
	"github.com/XGHXT/SYOJ-Backend/config"
	"github.com/XGHXT/SYOJ-Backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

var DB *gorm.DB

func Init(cfg *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		url.QueryEscape(cfg.Location),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&model.User{},&model.Submission{},&model.Category{},&model.Problem{},&model.Context{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	return err
}
