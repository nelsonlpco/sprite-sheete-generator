package cli

import (
	"flag"

	"github.com/nfnt/resize"
)

type Cli struct {
	filePath        string
	outputPath      string
	resizeAlgorithm int
	size            uint
	spriteSheetSize uint
}

func New() *Cli {
	filePath := flag.String("p", "", "images root path (require)")
	resizeAlgorithm := flag.Int("a", 2, `
Interpolation algorithm
		1 - Nearest-neighbor
		2 - Bilinear
		3 - Bicubic
		4 - Mitchell-Netravali
		5 - Lanczos(2)
		6 - Lanczos(3)
`)
	spriteSheetSize := flag.Int("s", 1024, "Sprite Sheet size when receive 0 calculates to base 2")
	output := flag.String("o", "./spriteSheet.png", "output sprite sheet png path")

	flag.Parse()

	return &Cli{
		filePath:        *filePath,
		outputPath:      *output,
		resizeAlgorithm: *resizeAlgorithm,
		spriteSheetSize: uint(*spriteSheetSize),
	}
}

func (c Cli) FilePath() string {
	if c.filePath == "" {
		flag.PrintDefaults()
	}
	return c.filePath
}

func (c Cli) Output() string {
	return c.outputPath
}

func (c Cli) InterpolationFunction() resize.InterpolationFunction {
	return resize.InterpolationFunction(c.resizeAlgorithm)
}

func (c Cli) SpriteSheetWidth() uint {
	return c.spriteSheetSize
}
