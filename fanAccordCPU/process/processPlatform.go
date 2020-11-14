package process

import (
	"fmt"
	"os/exec"
	log "raspberry/common/log"
	"strconv"
	"strings"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

type PlatformProcessor struct {
	logger      *log.Logger
	monitorChan chan bool
}

func NewPlatform(lg *log.Logger) *PlatformProcessor {
	return &PlatformProcessor{
		logger:      lg,
		monitorChan: make(chan bool, 0),
	}
}

func (p *PlatformProcessor) InitPlatform() (pi bool) {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		p.logger.LogError("host.Init", "error", "", "", err)
		return
	}
	// 确认平台是否是PI
	pi = rpi.Present()
	return
}

func (p *PlatformProcessor) MonitorCPUTemp(workTemp int64) {
	var nowTemp float64
	var err error
	for {
		nowTemp, err = p.GetCPUTemp()
		if err != nil {
			p.logger.LogError("p.GetCPUTemp", "error", "", "", err)
			close(p.monitorChan)
			return
		}
		if nowTemp > float64(workTemp) {
			p.logger.LogInfo("MonitorCPUTemp", "active", fmt.Sprintf("now temp: %.2f(>%d), active fan", nowTemp, workTemp))
			p.monitorChan <- true
		} else {
			p.logger.LogInfo("MonitorCPUTemp", "deactive", fmt.Sprintf("now temp: %.2f(<=%d), deactive fan", nowTemp, workTemp))
			p.monitorChan <- false
		}
		time.Sleep(5 * time.Second)
	}
}

func (p *PlatformProcessor) GetCPUTemp() (temp float64, err error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	ret, err := cmd.Output()
	if err != nil {
		p.logger.LogError("cmd.Output", "error", "cat", "/sys/class/thermal/thermal_zone0/temp", err)
		return
	}
	temp, _ = strconv.ParseFloat(strings.TrimSpace(string(ret)), 64)
	temp = temp / 1000
	return
}

func (p *PlatformProcessor) ActiveFan(pin string) {
	fanPin := gpioreg.ByName(pin)
	for active := range p.monitorChan {
		if active {
			if fanPin.Read() {
				p.logger.LogInfo("fanPin.Read", "success", "fan actived, no action")
				continue
			}
			err := fanPin.Out(gpio.High)
			if err != nil {
				p.logger.LogError("fanPin.Out", "error", fmt.Sprintf("set %s output", fanPin.Name()), gpio.High, err)
				continue
			}
			p.logger.LogInfo("fanPin.Out", "success", fmt.Sprintf("set %s output to high", fanPin.Name()))
		} else {
			if !fanPin.Read() {
				p.logger.LogInfo("fanPin.Read", "success", "fan deactived, no action")
				continue
			}
			err := fanPin.Out(gpio.Low)
			if err != nil {
				p.logger.LogError("fanPin.Out", "error", fmt.Sprintf("set %s output", fanPin.Name()), gpio.Low, err)
				continue
			}
			p.logger.LogInfo("fanPin.Out", "success", fmt.Sprintf("set %s output to low", fanPin.Name()))
		}
	}
}
