package image

import (
	"fmt"
	"image"
	"math"
	"os"
)

func OpenImage(key string) (image.Image, error) {
	path := fmt.Sprintf("./static/file/%s", key)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return src, err
}

func GetSize(img image.Image) (int, int) {
	size := img.Bounds().Size()
	return size.X, size.Y
}

func RGBA2RGB(r, g, b, a uint32) (r2, g2, b2 uint32) {
	r2 = (1-a)*255 + a*r
	g2 = (1-a)*255 + a*g
	b2 = (1-a)*255 + a*b
	return
}

func Color2RGBA(r, g, b, a uint32) (r2, g2, b2, a2 uint32) {
	if 0 >= a {
		return 0, 0, 0, 0
	}
	r2 = uint32(float64(r) / float64(a) * 255)
	g2 = uint32(float64(g) / float64(a) * 255)
	b2 = uint32(float64(b) / float64(a) * 255)
	a2 = uint32(float64(a) / float64(a))
	return
}

func GetImageColorSummary(img image.Image, startX int, endX int, startY int, endY int, details int) (r, g, b float64) {
	r, g, b = 0, 0, 0
	num := 0.0
	for y := startY; y <= endY; y += details {
		for x := startX; x <= endX; x += details {
			r2, g2, b2, a2 := img.At(x, y).RGBA()
			r3, g3, b3, a3 := Color2RGBA(r2, g2, b2, a2)
			r4, g4, b4 := RGBA2RGB(r3, g3, b3, a3)
			r += float64(r4) * float64(r4)
			g += float64(g4) * float64(g4)
			b += float64(b4) * float64(b4)
			num++
		}
	}
	return math.Sqrt(r / num), math.Sqrt(g / num), math.Sqrt(b / num)
}

func GetImageGridColorSummary(img image.Image, x int, y int, sizeX int, sizeY int) (r, g, b float64) {
	width, height := GetSize(img)
	gridSizeX, gridSizeY := width/sizeX-1, height/sizeY-1
	r, g, b = GetImageColorSummary(img, x*gridSizeX, (x+1)*gridSizeX, y*gridSizeY, (y+1)*gridSizeY, 2)
	return r, g, b
}

type Color struct {
	R int
	G int
	B int
}

func RGB2Color(r, g, b float64) Color {
	return Color{
		R: int(r),
		G: int(g),
		B: int(b),
	}
}

func GetImageSummary(key string, x int, y int) ([]Color, error) {
	img, err := OpenImage(key)
	if err != nil {
		return nil, err
	}
	var colors []Color
	for i := 0; i <= (x*y)-1; i++ {
		colors = append(colors, RGB2Color(GetImageGridColorSummary(img, i%x, i/x, x, y)))
	}
	return colors, nil
}
