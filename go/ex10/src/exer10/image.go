package exer10

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func Circle(rgba *image.RGBA, out, in int) {
	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			r := math.Sqrt(math.Pow(float64(i)-100, 2) + math.Pow(float64(j)-100, 2))
			if r < float64(out) && r >= float64(in) {
				rgba.Set(i, j, color.RGBA{0, 0, 0, 255})
			} else {
				rgba.Set(i, j, color.RGBA{255, 255, 255, 255})
			}
		}
	}
}

func DrawCircle(outerRadius, innerRadius int, outputFile string) {
	rect := image.Rect(0, 0, 200, 200)
	rgba := image.NewRGBA(rect)
	Circle(rgba, outerRadius, innerRadius)

	file, err := os.Create(outputFile)
	if err != nil {
		panic("Creating a file failed!")
	}
	defer file.Close()

	pngerr := png.Encode(file, rgba)
	if pngerr != nil {
		panic("Encoding image to a file failed!")
	}
	cerr := file.Close()
	if cerr != nil {
		panic("Closing file failed!")
	}

}
