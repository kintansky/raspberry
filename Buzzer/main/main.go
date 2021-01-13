package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	// 有源蜂鸣器使用的是GPIO16
	p := gpioreg.ByName("GPIO16")
	if p == nil {
		log.Fatal("Failed to find GPIO16")
	}
	// 有源蜂鸣器之间使用了一个PNP三极管，基极为GPIO16，高电平截止，低电平导通
	p.Out(gpio.High)
	p.Out(gpio.Low)
	time.Sleep(5)
	err := p.Halt()
	if err != nil {
		fmt.Println("p.Halt err:", err)
		return
	}
}
