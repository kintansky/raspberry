package utils

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Args struct {
	Pin      string
	WorkTemp int64
	LogPath  string
	LogLevel string
}

var logLevel = [...]string{
	"error",
	"warn",
	"info",
}

func (a *Args) cleanPin() (err error) {
	if a.Pin == "" {
		err = fmt.Errorf("未指定参数pin")
		return
	}
	return
}

func (a *Args) cleanLogLevel() (err error) {
	a.LogLevel = strings.ToLower(a.LogLevel)
	for _, v := range logLevel {
		if v == a.LogLevel {
			return
		}
	}
	return fmt.Errorf("loglevel参数错误，可接受" + strings.Join(logLevel[:], ","))
}

func ParseArgs() (*Args, error) {
	args := Args{}
	var err error
	dir, err := os.Executable()
	if err != nil {
		return &args, fmt.Errorf("os.Executable err:%s", err.Error())
	}
	baseDir := filepath.Dir(dir)
	flag.StringVar(&(args.Pin), "pin", "", "风扇使用的GPIO引脚")
	flag.Int64Var(&(args.WorkTemp), "worktemp", 70, "风扇开始工作的温度，默认70C")
	flag.StringVar(&(args.LogPath), "log", filepath.Join(baseDir, "/log/log.log"), "日志位置，默认项目log/log.log")
	flag.StringVar(&(args.LogLevel), "loglevel", "error", "log等级："+strings.Join(logLevel[:], ","))
	flag.Parse()
	err = args.cleanPin()
	if err != nil {
		return &args, err
	}
	err = args.cleanLogLevel()
	if err != nil {
		return &args, err
	}
	return &args, err
}
