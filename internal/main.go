package main

import (
	"github.com/nelsonlpco/spritesheetgen/internal/cli"
	"github.com/nelsonlpco/spritesheetgen/internal/file"
	"github.com/nelsonlpco/spritesheetgen/internal/imagemanager"
)

func main() {
	args := cli.New()
	if args.FilePath() == "" {
		return
	}

	fm := file.New(args.FilePath())

	spriteSheet := imagemanager.NewSpriteSheet(fm, args.SpriteSheetWidth())
	spriteSheet.PlotSprites()
	spriteSheet.Save(args.Output())
}
