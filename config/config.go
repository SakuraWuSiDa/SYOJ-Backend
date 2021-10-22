package config

import (
	"fmt"
	"github.com/XGHXT/SYOJ-Backend/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量，用来保存所有配置信息
var Config = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	JwtSecret    string `mapstructure:"jwt_secret"`
	Port         string    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
	*JudgeServer `mapstructure:"judger"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	Location     string `mapstructure:"location"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"pool_size"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}

type JudgeServer struct {
	TestPath string `mapstructure:"path"`
}

type ConfigYaml struct {
	Name string
}

func Init() error {
	c := ConfigYaml{
		Name: "config_test",
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	return nil
}

func (c *ConfigYaml) initConfig() error {
	viper.SetConfigName(c.Name)
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AddConfigPath(util.GetProjectAbsPath() + "/conf")
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return err
	}
	// 把读取到的配置信息反序列化到 Config 变量中
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper.Runmarshal() failed, err: %v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config.yaml has been changed...")
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		}
	})

	return nil
}