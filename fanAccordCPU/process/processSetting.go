package process

import (
	log "raspberry/common/log"
)

// SettingProcessor 配置处理器
type SettingProcessor struct {
	LOGGER *log.Logger
}

// InitSetting 初始化配置
func InitSetting(logLevel string, logPath string) *SettingProcessor {
	s := &SettingProcessor{
		LOGGER: log.NewLogger(logPath, logLevel),
	}
	return s
}
