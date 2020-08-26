package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/olablt/tiny32/font"
	"tinygo.org/x/drivers/ssd1306"
)

var display ssd1306.Device

func main() {

	println("[STM32] starting oled script")

	// init i2c
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})
	// machine.I2C0.Configure(machine.I2CConfig{SDA: machine.PB7, SCL: machine.PB6, Frequency: 400000})

	handleDisplay()
	// handleFont()
}

func handleDisplay() {
	const oled_address = 60
	// const oled_address = 0x3C
	// const oled_address = 0x78
	display = ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: oled_address,
		// Address: ssd1306.Address_128_32,
		Width:  128,
		Height: 64,
	})

	// font := NewFontDraw(&Regular9pt7b, &display)
	TestFont := font.NewDraw(&font.Picopixel, &display)
	GoFont := font.NewDraw(&font.Regular58pt, &display)

	display.ClearBuffer()
	display.ClearDisplay()

	// black := color.RGBA{1, 1, 1, 255}
	// white := color.RGBA{255, 255, 255, 255}

	println("[STM32] starting OLED loop")
	for {
		display.ClearBuffer()

		// TestFont.Info()
		TestFont.WriteTL([]byte("TL"))
		TestFont.WriteTR([]byte("TR"))
		TestFont.WriteBL([]byte("BL"))
		TestFont.WriteBR([]byte("BR"))
		// TestFont.WriteC([]byte("13:34:15"))
		GoFont.WriteC([]byte("AC"))

		display.Display()

		time.Sleep(1000 * time.Millisecond)
	}
}

func Rect(x int16, y int16, w int16, h int16, c color.RGBA) {
	for i := x; i < x+w; i++ {
		display.SetPixel(i, y, c)
		display.SetPixel(i, y+h, c)
	}
	for j := y; j < y+h; j++ {
		display.SetPixel(x, j, c)
		display.SetPixel(x+w, j, c)
	}

	// // filled rect
	// for i := x; i < x+w; i++ {
	// 	for j := y; j < y+h; j++ {
	// 		display.SetPixel(i, j, c)
	// 	}
	// }
}
