package main

import (
	"fmt"

	"raspberry/fanAccordCPU/process"
	"raspberry/fanAccordCPU/utils"
)

func main() {
	args, err := utils.ParseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	settingProcessor := process.InitSetting(args.LogPath, args.LogLevel)

	platformProcessor := process.NewPlatform(settingProcessor.LOGGER)
	platformProcessor.InitPlatform()
	go platformProcessor.MonitorCPUTemp(args.WorkTemp)
	platformProcessor.ActiveFan(args.Pin)
}
