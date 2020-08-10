package game

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/somen440/golang-puyo/resources/images"
)

var (
	puyoImgs [puyoColorMax]*ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Puyo_png))
	if err != nil {
		log.Fatal(err)
	}
	puyoImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	for i := 0; i < puyoColorMax; i++ {
		tx := i * puyoImgWidth
		puyoImgs[i] = puyoImage.SubImage(image.Rect(tx, 0, tx+puyoImgWidth, puyoImgHeight)).(*ebiten.Image)
	}
	puyoImgs[0] = puyoImgs[len(puyoImgs)-1]
}

type Puyo struct {
	X int
	Y int

	Color int
}

var (
	gamePuyoNum = 4
)

func NewPuyo(x, y int) *Puyo {
	return &Puyo{
		X:     x,
		Y:     y,
		Color: rand.Intn(gamePuyoNum) + 1,
	}
}

func (p *Puyo) ID() string {
	return fmt.Sprintf("%d%d%d", p.X, p.Y, p.Color)
}

type PuyoPair struct {
	First  *Puyo
	Second *Puyo
}

func NewPuyoPair() *PuyoPair {
	return &PuyoPair{
		First:  NewPuyo(puyoInitX, puyoInitY),
		Second: NewPuyo(puyoInitX, puyoInitY+1),
	}
}

func (pp *PuyoPair) NotBottomPuyo() *Puyo {
	if pp.First.Y < pp.Second.Y {
		return pp.First
	} else if pp.First.Y > pp.Second.Y {
		return pp.Second
	}
	return pp.Second
}

func (pp *PuyoPair) GetMatchPuyo(x, y int) *Puyo {
	if pp.First.X == x && pp.First.Y == y {
		return pp.First
	} else if pp.Second.X == x && pp.Second.Y == y {
		return pp.Second
	}
	return nil
}

func (pp *PuyoPair) Has(x, y int) bool {
	return (pp.First.X == x && pp.First.Y == y) || (pp.Second.X == x && pp.Second.Y == y)
}
