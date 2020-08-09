package game

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/somen440/golang-puyo/resources/images"
)

var (
	stageImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Stage_png))
	if err != nil {
		log.Fatal(err)
	}
	stageImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

var (
	board [stageY][stageX]int
)

func InitializeBoard() {
	for i := 0; i < stageY; i++ {
		for j := 0; j < stageX; j++ {
			board[i][j] = 0
		}
	}
}

func init() {
	InitializeBoard()
}
