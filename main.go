package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var gray = color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}
var red = color.RGBA{0xFF, 0x00, 0x00, 0xAA}
var blue = color.RGBA{0x00, 0x00, 0xFF, 0xAA}

var xpos, ypos int
var hero image.Rectangle

func main() {
	fmt.Println("Hello World!")
	renderTest()
}

func renderTest() {
	hero.Min.X = 100
	hero.Min.Y = 100

	hero.Max.X = 150
	hero.Max.Y = 150

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer w.Release()

		var b screen.Buffer
		defer func() {
			if b != nil {
				b.Release()
			}
		}()

		for {

			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
				switch e.Code {
				case key.CodeEscape:
					return
				case key.CodeA:
					xpos = xpos - 5
				case key.CodeD:
					xpos = xpos + 5
				case key.CodeW:
					ypos = ypos - 5
				case key.CodeS:
					ypos = ypos + 5
				}
				updateHero()

			case paint.Event:
				// Gray backgroundut
				w.Fill(b.Bounds(), red, screen.Src)
				w.Fill(hero, blue, screen.Src)

				// dst := buffer.RGBA()
				// draw.Draw(dst, dst.Bounds(), etc)
				w.Upload(image.Point{0, 0}, b, b.Bounds())
				w.Publish()
			case size.Event:
				if b != nil {
					b.Release()
				}
				b, err = s.NewBuffer(e.Size())
				fmt.Println(b.Size().X)
				fmt.Println(b.Size().Y)
				if err != nil {
					log.Fatal(err)
				}
			case error:
				log.Print(e)
			}

			// case mouse.Event:
			// 	draw(b.RGBA(), 0, 0, 10, 10)
			// 	w.Send(paint.Event{})
		}
	})
}

// The drawLight method draws a light as position i, j to m, with the given width and height.
func updateHero() {
	hero.Min.X = xpos
	hero.Min.Y = ypos
	hero.Max.X = xpos + 10
	hero.Max.Y = ypos + 10

	fmt.Println("X Pos: ", hero.Min.X)
	fmt.Println("Y Pos: ", hero.Min.Y)

	//	w.Fill(hero, blue, screen.Src)
}
