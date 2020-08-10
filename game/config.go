package game

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 480

	puyoImgWidth  = 40
	puyoImgHeight = 40

	puyoColorMax = 5

	stageX = 6
	stageY = 12

	connectionNum = 4

	puyoInitX = 3
	puyoInitY = 0

	fontSize = 32
)

var (
	arcadeFont font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)
