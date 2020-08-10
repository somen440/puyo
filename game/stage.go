package game

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"strconv"
	"strings"

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

type PID string

type Color int

type Board [stageY][stageX]Color

type ColorFunc func(b *Board, x, y int, c Color)

func ToID(x, y int, c Color) PID {
	return PID(fmt.Sprintf("%d_%d_%d", x, y, c))
}

func FromID(id PID) (x, y int, c Color) {
	tmp := strings.Split(string(id), "_")
	x, _ = strconv.Atoi(tmp[0])
	y, _ = strconv.Atoi(tmp[1])
	tmpC, _ := strconv.Atoi(tmp[2])
	c = Color(tmpC)
	return
}

func newBoard() Board {
	b := [stageY][stageX]Color{}
	for i := 0; i < stageY; i++ {
		for j := 0; j < stageX; j++ {
			b[i][j] = 0
		}
	}
	return b
}
