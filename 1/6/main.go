package main

/*
go build 1/6/main.go
mv main 1/6/lissa
./1/6/lissa > 1/6/out.gif

// or

make -C 1/6
*/

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 0, 255}}
var palette2 = []color.Color{color.RGBA{255, 255, 255, 255}, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}}

const (
	whitelndex = 0 // Первый цвет палитры
	blacklndex = 1 // Следующий цвет палитры
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles = 5
		// Количество полных колебаний x
		res  = 0.001 // Угловое разрешение
		size = 500
		// Канва изображения охватывает [size..+size]
		nframes = 256
		// Количество кадров анимации
		delay = 64
	// Задержка между кадрами (единица - 10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	// Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, []color.Color{color.RGBA{0, 0, 0, 255}, color.RGBA{uint8(i * 213 % 255), uint8(i * 88 % 255), uint8(i * 12 % 255), 255}})
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blacklndex)
		}
		phase += 0.25
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	// Примечание: игнорируем ошибки
}
