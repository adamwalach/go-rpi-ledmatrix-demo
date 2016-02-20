package main

import (
	"log"
	"time"

	max7219 "github.com/adamwalach/go-max7219"
)

func main() {
	font := max7219.FontCP437
	mtx := max7219.NewMatrix(2)
	// Open SPI device with spibus and spidev parameters equal to 0 and 0.
	// Set LED matrix brightness is equal to 4
	err := mtx.Open(0, 0, 4)

	if err != nil {
		log.Fatal(err)
	}
	defer mtx.Close()
	mtx.OutputAsciiCode(0, font, 1, true)

	for {
		mtx.SlideMessage("Hello, IoT world!  ", max7219.FontCP437, true, 50*time.Millisecond)
		time.Sleep(1 * time.Second)
		mtx.SlideMessage("Hello, IoT world!  ", max7219.FontCP437, true, 50*time.Millisecond)
		for i := 0; i <= 64; i++ {
			mtx.OutputAsciiCode(0, font, i, true)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
