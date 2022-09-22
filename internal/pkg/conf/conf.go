package conf

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	Base       Base       `yaml:"base"`
	JWT        JWT        `yaml:"jwt"`
	Http       Http       `yaml:"http"`
	GRPC       GRPC       `yaml:"grpc"`
	MySQL      MySQL      `yaml:"mysql"`
	MongoDB    MongoDB    `yaml:"mongodb"`
	Redis      Redis      `yaml:"redis"`
	Prometheus Prometheus `yaml:"prometheus"`
	Proof      Proof      `yaml:"proof"`
}

type Base struct {
	IsDebug bool `mapstructure:"is_debug"`
}

type Http struct {
	Port int `yaml:"port"`
}

type GRPC struct {
	Port int `yaml:"port"`
}

type MySQL struct {
	Username string `yaml:"username"`
	Passport string `yaml:"passport"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type JWT struct {
	Issuer string `yaml:"issuer"`
	Secret string `yaml:"secret"`
}

type MongoDB struct {
	DSN      string `yaml:"dsn"`
	Database string `yaml:"database"`
	PoolSize uint64 `mapstructure:"pool_size"`
}

type Prometheus struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Proof struct {
	InfoLog string `mapstructure:"info_log"`
	ErrLog  string `mapstructure:"err_log"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DataBase int    `yaml:"data_base"`
	UserName string `yaml:"user_name"`
	Password string `yaml:"password"`
}

// MustInitConf 初始化配置文件
func MustInitConf() {
	var configFile string
	flag.StringVar(&configFile, "c", "./etc/configs/local.yaml", "your app config file.")
	if !flag.Parsed() {
		flag.Parse()
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	//Unmarshal 将配置解组为 Struct。确保正确设置结构字段上的标签。
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal error config file: %w", err))
	}

	fmt.Println("config initialized")
}
