package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
)

type State int

const (
	Ready State = iota
	Playing
	GameOver
)

type Game struct {
	current State

	f int
}

var G *Game

func init() {
	G = &Game{
		current: Playing,
	}
}

func (g *Game) inputKey() {
	isLeftHit := hit(currentPuyoPair.LeftPuyo(), Left)
	isRightHit := hit(currentPuyoPair.RightPuyo(), Right)
	isDownHit := hit(currentPuyoPair.BottomPuyo(), Down)

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		if !isLeftHit {
			currentPuyoPair.MovingLeft()
		}
		break

	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		if !isRightHit {
			currentPuyoPair.MovingRight()
		}
		break

	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		if !isDownHit {
			currentPuyoPair.Falling()
		}
		break

	case inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeySpace):
		if !isDownHit {
			currentPuyoPair.Rolling()
		}
		break
	}
}

func (g *Game) falling() bool {
	if g.f%30 != 0 {
		return true
	}

	if !hit(currentPuyoPair.First, Down) && !hit(currentPuyoPair.Second, Down) {
		currentPuyoPair.Falling()
		return true
	}

	return false
}

func (g *Game) redraw(f int) {
	if g.f%f != 0 {
		return
	}

	InitializeBoard()

	board[currentPuyoPair.FirstY()][currentPuyoPair.FirstX()] = currentPuyoPair.FirstColor()
	board[currentPuyoPair.SecondY()][currentPuyoPair.SecondX()] = currentPuyoPair.SecondColor()

	for _, p := range fieldPuyos {
		board[p.Y][p.X] = p.Color
	}
}

func (g *Game) changePuyo() bool {
	bottom := currentPuyoPair.BottomPuyo()
	nBottom := currentPuyoPair.NotBottomPuyo()

	if nBottom.X == puyoInitX && nBottom.Y == puyoInitY {
		return false
	}

	currentPuyoPair = nextPuyoPair
	nextPuyoPair = NewPuyoPair()

	if hit(bottom, Down) {
		fieldPuyos = append(fieldPuyos, bottom)
		for {
			if hit(nBottom, Down) {
				fieldPuyos = append(fieldPuyos, nBottom)
				break
			}
			nBottom.Y++
		}
	} else {
		fieldPuyos = append(fieldPuyos, nBottom)
		for {
			if hit(bottom, Down) {
				fieldPuyos = append(fieldPuyos, bottom)
				break
			}
			bottom.Y++
		}
	}
	return true
}

func (g *Game) Update(screen *ebiten.Image) error {
	switch g.current {
	case Ready:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.current = Playing
		}
		break

	case Playing:
		g.f++

		g.inputKey()
		isFalled := g.falling()
		g.redraw(30)

		if !isFalled {
			if ok := g.changePuyo(); !ok {
				g.current = GameOver
			}
			fieldPuyos = Ereaser(fieldPuyos, connectionNum)
		}
		g.redraw(10)
		break

	case GameOver:
		break
	}

	return nil
}

func (g *Game) drawField(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(stageImage, op)
}

func (g *Game) drawPuyo(screen *ebiten.Image) {
	for i, rows := range board {
		for j, row := range rows {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(j*puyoImgWidth), float64(i*puyoImgHeight))
			if row != 0 {
				screen.DrawImage(puyoImgs[row-1], op)
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.current == Ready {
		s1 := "Press SPACE Key Start"
		x := (ScreenWidth - len(s1)*fontSize) / 2
		text.Draw(screen, s1, arcadeFont, x, ScreenHeight/2, color.White)
		return
	}

	g.drawField(screen)
	g.drawPuyo(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"first: x: %d\nfirst: y: %d\nsecond: x: %d\nsecond: y: %d\n\nbottom: x: %d\nbottom: y: %d\n",
		currentPuyoPair.FirstX(),
		currentPuyoPair.FirstY(),
		currentPuyoPair.SecondX(),
		currentPuyoPair.SecondY(),
		currentPuyoPair.BottomPuyo().X,
		currentPuyoPair.BottomPuyo().Y,
	))

	if g.current == GameOver {
		str := "Game Over"
		x := (ScreenWidth - len(str)*fontSize) / 2
		text.Draw(screen, str, arcadeFont, x, ScreenHeight/2, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
