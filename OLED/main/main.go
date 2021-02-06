package main

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
	"time"

	"periph.io/x/conn/i2c/i2creg"
	"periph.io/x/devices/ssd1306"
	"periph.io/x/host"
)

func main() {
	// Load all the drivers:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
		return
	}

	// Open a handle to the first available I²C bus:
	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer bus.Close()

	// Open a handle to a ssd1306 connected on the I²C bus:
	dev, err := ssd1306.NewI2C(bus, &ssd1306.DefaultOpts)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dev.Halt()
	fmt.Println("device bounds:", dev.Bounds())
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	g, err := gif.DecodeAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("image bounds:", g.Image[0].Bounds())

	topLeftPoint := dev.Bounds().Min
	var targetRect image.Rectangle
	for i := 0; ; i++ {
		index := i % len(g.Image)
		c := time.After(time.Duration(10*g.Delay[index]) * time.Millisecond)
		img := g.Image[index]
		targetRect = image.Rectangle{
			Min: image.Point{X: (topLeftPoint.X + i) % dev.Bounds().Max.X, Y: topLeftPoint.Y},
			Max: image.Point{X: (topLeftPoint.X+i)%dev.Bounds().Max.X + g.Image[index].Bounds().Max.X, Y: topLeftPoint.Y + g.Image[index].Bounds().Max.Y},
		}
		fmt.Println("targetRect:", targetRect)
		dev.Draw(targetRect, img, image.Point{})
		<-c
	}
}
