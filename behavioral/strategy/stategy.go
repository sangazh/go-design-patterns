package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

var output = flag.String("output", "console", "'console' or 'image'")
var activeStrategy PrintStrategy

func main() {
	flag.Parse()

	switch *output {
	case "console":
		activeStrategy = new(ConsoleSquare)
	case "image":
		activeStrategy = &ImageSquare{"image.jpg"}
	default:
		activeStrategy = new(ConsoleSquare)
	}

	err := activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}

}

type PrintStrategy interface {
	Print() error
}
type ConsoleSquare struct{}

func (*ConsoleSquare) Print() error {
	fmt.Println("Square")
	return nil
}

type ImageSquare struct {
	DestinationFilePath string
}

func (s *ImageSquare) Print() error {
	width := 800
	height := 600
	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	bgColor := image.Uniform{C: color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{C: color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(s.DestinationFilePath)
	if err != nil {
		return errors.New("error opening image")
	}
	defer w.Close()

	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return errors.New("error writing image to disk")
	}
	return nil
}
