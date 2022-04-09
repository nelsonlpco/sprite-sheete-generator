package imagemanager

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/nfnt/resize"
)

type ImageManager struct {
	Value *image.RGBA
}

func (i *ImageManager) Set(x, y int, c color.Color) {
	i.Value.Set(x, y, c)
}

func (i *ImageManager) ColorModel() color.Model {
	return i.Value.ColorModel()
}

func (i *ImageManager) At(x, y int) color.Color {
	return i.Value.At(x, y)
}

func (i *ImageManager) Bounds() image.Rectangle {
	return i.Value.Bounds()
}

func Width(img image.Image) int {
	return img.Bounds().Max.X - img.Bounds().Min.X
}

func Height(img image.Image) int {
	return img.Bounds().Max.Y - img.Bounds().Min.Y
}

func (i *ImageManager) DrawRaw(innerImg image.Image, sp image.Point, width uint, height uint) {
	resizedImg := resize.Resize(width, height, innerImg, resize.Bilinear)
	w := Width(resizedImg)
	h := Height(resizedImg)
	draw.Draw(i, image.Rectangle{sp, image.Point{sp.X + w, sp.Y + h}}, resizedImg, image.ZP, draw.Src)

}
