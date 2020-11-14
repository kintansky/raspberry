package common

import (
	"fmt"
	"testing"
)

// 数据库配置结构体

// DBConfigBase 生产配置基类
type DBConfigBase struct {
	IP      string `ini:"ip"`
	Port    int    `ini:"port"`
	User    string `ini:"user"`
	Pwd     string `ini:"pwd"`
	CnxType string `ini:"cnxType"`
}

func newDBConfigBase() *DBConfigBase {
	return &DBConfigBase{}
}

// 其他配置结构体

// IPMANConfig 结构体
type IPMANConfig struct {
	DBName   string `ini:"db"`
	CmdTable string `ini:"cmdTable"`
}

func newIPMANConfig() *IPMANConfig {
	return &IPMANConfig{}
}

// Config 配置处理器
type Config struct {
	IPMAN   IPMANConfig  `ini:"ipman"`
	DB      DBConfigBase `ini:"db"`
	DBLocal DBConfigBase `ini:"dblocal"`
}

func NewConfig() *Config {
	return &Config{
		IPMAN:   *newIPMANConfig(),
		DB:      *newDBConfigBase(),
		DBLocal: *newDBConfigBase(),
	}
}

func TestLoadINI(t *testing.T) {
	cfg := NewConfig()
	err := LoadINI("./conf.ini", cfg)
	if err != nil {
		fmt.Printf("load ini failed, %v\n", err)
	}
	fmt.Printf("%#v", cfg)
}
