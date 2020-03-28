package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)





var palette = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},  
	color.RGBA{0x00, 0x00, 0x00, 0xFF}, /
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // red
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // blue
	color.RGBA{0xFF, 0xCC, 0x00, 0xFF}, // yellow
	color.RGBA{0x00, 0x80, 0xFF, 0x80}, // cyan
	color.RGBA{0x80, 0x00, 0xFF, 0xFF}, // magenta
}

func main() {
	lissajous(os.Stdout)
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
