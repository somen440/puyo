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
}

type Puyo struct {
	X int
	Y int

	Color int
}

var (
	gamePuyoNum = 5
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

func (pp *PuyoPair) FirstColor() int {
	return pp.First.Color
}

func (pp *PuyoPair) FirstX() int {
	return pp.First.X
}

func (pp *PuyoPair) FirstY() int {
	return pp.First.Y
}

func (pp *PuyoPair) SecondColor() int {
	return pp.Second.Color
}

func (pp *PuyoPair) SecondX() int {
	return pp.Second.X
}

func (pp *PuyoPair) SecondY() int {
	return pp.Second.Y
}

func (pp *PuyoPair) Falling() {
	pp.First.Y++
	pp.Second.Y++
}

func (pp *PuyoPair) MovingLeft() {
	pp.First.X--
	pp.Second.X--
}

func (pp *PuyoPair) MovingRight() {
	pp.First.X++
	pp.Second.X++
}

func (pp *PuyoPair) Rolling() {
	isFirstMove := false
	tmpSec := *pp.Second
	if pp.First.X == pp.Second.X {
		if pp.Second.X == 0 {
			pp.Second.X = pp.First.X + 1
			pp.First.Y = pp.Second.Y
		} else if pp.Second.X == stageX-1 {
			pp.Second.X = pp.First.X - 1
			pp.First.Y = pp.Second.Y
		} else {
			isFirstMove = true
			if pp.First.Y < pp.Second.Y {
				pp.Second.X--
			} else {
				pp.Second.X++
			}
		}
	} else {
		if pp.Second.X == 0 {
			pp.Second.X = pp.First.X + 1
			pp.First.Y = pp.Second.Y
		} else if pp.Second.X == stageX-1 {
			pp.Second.X = pp.First.X - 1
			pp.First.Y = pp.Second.Y
		} else {
			isFirstMove = true
			if pp.First.X > pp.Second.X {
				pp.Second.Y--
			} else {
				pp.Second.Y++
			}
		}
	}
	if isFirstMove {
		pp.First.X = tmpSec.X
		pp.First.Y = tmpSec.Y
	}
}

func (pp *PuyoPair) BottomPuyo() *Puyo {
	if pp.First.Y > pp.Second.Y {
		return pp.First
	} else if pp.First.Y < pp.Second.Y {
		return pp.Second
	}
	return pp.First
}

func (pp *PuyoPair) NotBottomPuyo() *Puyo {
	if pp.First.Y < pp.Second.Y {
		return pp.First
	} else if pp.First.Y > pp.Second.Y {
		return pp.Second
	}
	return pp.Second
}

func (pp *PuyoPair) LeftPuyo() *Puyo {
	if pp.First.X < pp.Second.X {
		return pp.First
	} else if pp.First.X > pp.Second.X {
		return pp.Second
	}
	return pp.First
}

func (pp *PuyoPair) RightPuyo() *Puyo {
	if pp.First.X > pp.Second.X {
		return pp.First
	} else if pp.First.X < pp.Second.X {
		return pp.Second
	}
	return pp.First
}

type Move int

const (
	Left Move = iota
	Right
	Down
	Top
)

func hit(p *Puyo, move Move) bool {
	nextY := -1
	nextX := -1

	switch move {
	case Left:
		nextX = p.X - 1
		if nextX < 0 {
			return true
		}
	case Right:
		nextX = p.X + 1
		if nextX == stageX {
			return true
		}
		break
	case Down:
		nextY = p.Y + 1
		if nextY == stageY {
			return true
		}
		break
	}

	for _, fp := range fieldPuyos {
		if fp.X == p.X && fp.Y == nextY {
			return true
		} else if fp.Y == p.Y && fp.X == nextX {
			return true
		}
	}

	return false
}

var (
	currentPuyoPair *PuyoPair
	nextPuyoPair    *PuyoPair
	fieldPuyos      []*Puyo
)

func init() {
	currentPuyoPair = NewPuyoPair()
	nextPuyoPair = NewPuyoPair()
	fieldPuyos = []*Puyo{}
}
