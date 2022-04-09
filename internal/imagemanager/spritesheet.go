package imagemanager

import (
	"image"
	"log"
	"math"

	"github.com/nelsonlpco/spritesheetgen/internal/file"
)

type SpriteSheet struct {
	Width             uint
	Height            uint
	SpriteByLine      uint
	TotalSprites      uint
	SpriteSize        uint
	fm                *file.Manager
	OutputSpriteSheet *ImageManager
}

func NewSpriteSheet(fm *file.Manager, width uint) *SpriteSheet {
	spriteSheet := &SpriteSheet{
		Width:        width,
		Height:       width / 2,
		TotalSprites: fm.Size(),
		fm:           fm,
	}
	spriteSheet.OutputSpriteSheet = &ImageManager{
		image.NewRGBA(
			image.Rectangle{
				Min: image.ZP,
				Max: image.Point{
					X: int(spriteSheet.Width),
					Y: int(spriteSheet.Height),
				},
			},
		),
	}
	spriteSheet.calcSpriteSize()
	spriteSheet.calMaxSpriteByLine()

	return spriteSheet
}

func (s *SpriteSheet) calcSpriteSize() {
	overflow := uint(4)
	maxWidth := uint(math.Round(float64(s.Width / s.TotalSprites)))
	maxHeight := uint(math.Round(float64(s.Height/2))) - overflow

	dif := maxHeight - maxWidth

	if dif > 0 {
		maxWidth += dif / 2
	}

	s.SpriteSize = maxWidth
}

func (s *SpriteSheet) calMaxSpriteByLine() {
	s.SpriteByLine = uint(math.Floor(float64(s.Width / s.SpriteSize)))
}

func (s *SpriteSheet) PlotSprites() {
	log.Printf("Ploting %v sprites of size: %v in sprite sheet of size %v\n", s.TotalSprites, s.SpriteSize, s.Width)
	var x, y uint
	for i := uint(0); i < s.TotalSprites; i++ {
		fileImg := s.fm.OpenImageByIndex(i)

		if i > 0 && i%s.SpriteByLine == 0 {
			y += s.SpriteSize
			x = 0
		}

		img, _, err := image.Decode(fileImg)
		if err != nil {
			log.Fatalf("error on plot image: %v error %v", i, err)
		}
		s.OutputSpriteSheet.DrawRaw(img, image.Point{X: int(x), Y: int(y)}, uint(s.SpriteSize), uint(s.SpriteSize))
		x += s.SpriteSize

		fileImg.Close()
	}
}

func (s *SpriteSheet) Save(outputPath string) {
	log.Printf("Saving %v file...\n\n", outputPath)
	s.fm.Save(outputPath, s.OutputSpriteSheet.Value)
	log.Println("Complete.")
}
