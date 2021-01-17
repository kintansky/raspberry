package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"periph.io/x/conn/gpio"
	"periph.io/x/conn/gpio/gpioreg"
	"periph.io/x/conn/physic"
	"periph.io/x/host"
)

// var clockWidth float64 = 5 * 1000 // us,PWM基波间隔，越小分辨率越高,测试最小约10ms

func testServo(p gpio.PinIO, pulse int, clockWidth int) {
	d, err := gpio.ParseDuty(fmt.Sprintf("%.f%%", float64(100*pulse)/float64(clockWidth)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s duty:%s\n", p.Name(), d.String())
	if err := p.PWM(d, physic.Frequency(1000000/clockWidth)*physic.Hertz); err != nil {
		log.Fatal(err)
	}
}

func releasePin(p gpio.PinIO) {
	err := p.Halt()
	if err != nil {
		fmt.Printf("%s Halt err:%s\n", p.Name(), err.Error())
		return
	}
}

type args struct {
	pin        string
	pos1       int
	pos2       int
	clockWidth int
}

func parseArg() (ags *args, err error) {
	ags = &args{}
	flag.StringVar(&ags.pin, "pin", "", "测试的GPIO引脚")
	flag.IntVar(&ags.pos1, "pos1", 500, "起始位置单位us")
	flag.IntVar(&ags.pos2, "pos2", 4000, "结束位置单位us")
	flag.IntVar(&ags.clockWidth, "clockWidth", 20*1000, "PWM基波间隔")
	flag.Parse()
	if ags.pin == "" {
		err = fmt.Errorf("pin不能为空")
		return
	}
	if float64(ags.clockWidth) <= math.Max(float64(ags.pos1), float64(ags.pos2)) {
		err = fmt.Errorf("PWM基波间隔%dus过小,无法覆盖测试范围[%d~%d]us", ags.clockWidth, ags.pos1, ags.pos2)
		return
	}
	return
}

func main() {
	ags, err := parseArg()
	if err != nil {
		fmt.Println("parseArg err:", err.Error())
		return
	}
	fmt.Printf("%#v\n", ags)
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	// Use gpioreg GPIO pin registry to find a GPIO pin by name.
	testPin := gpioreg.ByName(ags.pin)
	if testPin == nil {
		log.Fatal("Failed to find ", ags.pin)
	}
	defer releasePin(testPin)
	// 最小分辨率大概为50us，国华Servo测试结果：3700us=+90度 750us=-90度
	for i := 0; i < 5; i++ {
		testServo(testPin, ags.pos1, ags.clockWidth)
		time.Sleep(2 * time.Second)
		testServo(testPin, ags.pos2, ags.clockWidth)
		time.Sleep(5 * time.Second)
	}
}
