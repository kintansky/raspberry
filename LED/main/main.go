package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
)

// LedGroup LED灯组
type LedGroup struct {
	Red    gpio.PinIO
	Yellow gpio.PinIO
	Green  gpio.PinIO
}

func NewLedGroup(redPinName, yellowPinName, greenPinName string) *LedGroup {
	pinRed := gpioreg.ByName(redPinName)
	if pinRed == nil {
		log.Fatal("RedPin failed to find " + redPinName)
		return nil
	}
	pinRed.Out(gpio.Low)
	pinYellow := gpioreg.ByName(yellowPinName)
	if pinRed == nil {
		log.Fatal("YellowPin failed to find " + yellowPinName)
		return nil
	}
	pinYellow.Out(gpio.Low)
	pinGreen := gpioreg.ByName(greenPinName)
	if pinGreen == nil {
		log.Fatal("GreenPin failed to find " + greenPinName)
		return nil
	}
	pinGreen.Out(gpio.Low)
	return &LedGroup{
		Red:    pinRed,
		Yellow: pinYellow,
		Green:  pinGreen,
	}
}

func (l *LedGroup) ToArray() (ledArray []gpio.PinIO) {
	return []gpio.PinIO{l.Red, l.Yellow, l.Green}
}

func (l *LedGroup) LightenFlow(sep time.Duration) {
	for _, p := range l.ToArray() {
		p.Out(gpio.High)
		time.Sleep(sep)
		p.Out(gpio.Low)
	}
}

func (l *LedGroup) Reset() {
	for _, p := range l.ToArray() {
		p.Out(gpio.Low)
	}
}

func main() {
	// GPIO4、5、6分别对应红黄绿LED灯
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
		return
	}
	if rpi.Present() {
		fmt.Println("On Pi platform.")
	}

	ledGroup := NewLedGroup("GPIO4", "GPIO5", "GPIO6")
	func(sep time.Duration) {
		fmt.Println("lighten led group flow")
		for {
			ledGroup.LightenFlow(sep)
		}
	}(1 * time.Second)

}
