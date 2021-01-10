package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

type Button struct {
	pinName string
	pin     gpio.PinIO
}

func NewButton(pinName string) *Button {
	buttonPin := gpioreg.ByName(pinName)
	if buttonPin == nil {
		log.Fatal("ButtonPin failed to find " + pinName)
		return nil
	}
	// PIN常驻高电平，检测下降沿
	if err := buttonPin.In(gpio.PullUp, gpio.FallingEdge); err != nil {
		log.Fatal(err)
	}
	return &Button{
		pinName: pinName,
		pin:     buttonPin,
	}
}

func (b *Button) WaitForEdge() {
	for {
		b.pin.WaitForEdge(-1)
		fmt.Printf("pressed! %s went %s\n", b.pin, b.pin.Read())
	}
}

func (b *Button) State() gpio.Level {
	return b.pin.Read()
}

func main() {
	// Button 对应GPIO12、13
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
		return
	}
	button1 := NewButton("GPIO12")
	fmt.Println("Init state:", button1.State())
	// 展示状态变更
	go func() {
		lastState := button1.State()
		var nowState gpio.Level
		for {
			nowState = button1.State()
			if lastState != nowState {
				fmt.Println("State changed! Now:", nowState)
			}
			lastState = nowState
		}
	}()
	button1.WaitForEdge()
}
