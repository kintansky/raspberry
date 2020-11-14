package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	// "periph.io/x/periph/host/rpi"
)

func main() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	// 确认平台是否是PI
	// pi := rpi.Present()
	// fmt.Println(pi)
	// 注册GPIO引脚，可以用rpi或者gpioreg
	// p := rpi.P1_32
	// 对PWM引脚GPIO12进行操作，高电平3V3，低电平0V
	p := gpioreg.ByName("GPIO12")
	var sw bool
	for {
		if sw {
			err := p.Out(gpio.High)
			if err != nil {
				log.Fatal(err)
				continue
			}
			fmt.Printf("%s up\n", p.Name())
			sw = false
		} else {
			err := p.Out(gpio.Low)
			if err != nil {
				log.Fatal(err)
				continue
			}
			fmt.Printf("%s down\n", p.Name())
			sw = true
		}
		time.Sleep(5 * time.Second)
	}
}
