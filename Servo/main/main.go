package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/gpio"
	"periph.io/x/conn/gpio/gpioreg"
	"periph.io/x/conn/physic"
	"periph.io/x/host"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	p := gpioreg.ByName("GPIO20")
	if p == nil {
		log.Fatal("Failed to find GPIO20")
	}
	step := 0.5
	clockWidth := 5 // 20ms
	for i := 0; i < 10; i++ {
		// 范围0.5~3.5ms=-90~+90度
		d, err := gpio.ParseDuty(fmt.Sprintf("%.f%%", float64(100*(0.5+step*float64(i%4))/float64(clockWidth))))
		// d, err := gpio.ParseDuty(fmt.Sprintf("%.f%%", 100*(0.5)/20))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("duty:", d.String())
		if err := p.PWM(d, physic.Frequency(1000/clockWidth)*physic.Hertz); err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
	err := p.Halt()
	if err != nil {
		fmt.Println("p.Halt err:", err)
		return
	}
}
