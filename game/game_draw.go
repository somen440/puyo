package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
)

func (g *Game) drawField(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(stageImage, op)
}

func (g *Game) drawPuyo(screen *ebiten.Image) {
	for y, rows := range g.board {
		for x, c := range rows {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*puyoImgWidth), float64(y*puyoImgHeight))
			if c != 0 {
				screen.DrawImage(puyoImgs[c-1], op)
			}
		}
	}
}

func (g *Game) drawNextPuyo(screen *ebiten.Image) {
	padding := 15

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(stageX*puyoImgWidth+padding), float64(padding))
	screen.DrawImage(puyoImgs[g.nextPP.First.Color-1], op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(stageX*puyoImgWidth+padding), float64(padding+puyoImgHeight))
	screen.DrawImage(puyoImgs[g.nextPP.Second.Color-1], op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	switch g.current {
	case Ready:
		g.drawField(screen)

		s1 := "Press SPACE Key Start"
		x := (ScreenWidth - len(s1)*fontSize) / 2
		text.Draw(screen, s1, arcadeFont, x, ScreenHeight/2, color.White)
		break

	case Playing:
		g.drawField(screen)
		g.drawNextPuyo(screen)
		g.drawPuyo(screen)
		break

	case GameOver:
		g.drawField(screen)
		g.drawNextPuyo(screen)
		g.drawPuyo(screen)

		str := "Game Over"
		x := (ScreenWidth - len(str)*fontSize) / 2
		text.Draw(screen, str, arcadeFont, x, ScreenHeight/2, color.White)
		break
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"first: x: %d\nfirst: y: %d\nsecond: x: %d\nsecond: y: %d\n\nstate: %v\n",
		g.playerPP.First.X,
		g.playerPP.First.Y,
		g.playerPP.Second.X,
		g.playerPP.Second.Y,
		g.current,
	))
}
