package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (g *Game) inputKey() {
	if g.f%15 == 0 || g.isFellOut {
		return
	}

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		if hitLeft(g.board, g.playerPP) {
			break
		}
		g.board, g.playerPP = moveLeft(g.board, g.playerPP)
		break

	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		if hitRight(g.board, g.playerPP) {
			break
		}
		g.board, g.playerPP = moveRight(g.board, g.playerPP)
		break

	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		if hitDown(g.board, g.playerPP) {
			break
		}
		g.board, g.playerPP = moveDown(g.board, g.playerPP)
		break

	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		break

	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		g.board, g.playerPP = Rolling(g.board, g.playerPP)
		break
	}

	g.isFellOut = hitDown(g.board, g.playerPP)
}

func (g *Game) deploy() {
	f, s := g.playerPP.First, g.playerPP.Second
	g.board[f.Y][f.X] = Color(f.Color)
	g.board[s.Y][s.X] = Color(s.Color)
}

func (g *Game) falling() {
	if g.f%15 != 0 || g.isFellOut {
		return
	}

	var isFellOut bool
	g.board, g.playerPP, isFellOut = Dropper(g.board, g.playerPP)
	g.isFellOut = isFellOut
}

func (g *Game) judgeGameOver() bool {
	if g.f%15 == 0 || !g.isFellOut {
		return false
	}
	nbp := g.playerPP.NotBottomPuyo()
	return nbp.X == puyoInitX && nbp.Y == puyoInitY
}

func (g *Game) fellOutProc() (isEreased bool) {
	if g.f%15 == 0 {
		return
	}

	g.board, isEreased = Ereaser(g.board, g.playerPP, connectionNum)
	if !g.isFellOut {
		return
	}

	g.playerPP = g.nextPP
	g.nextPP = NewPuyoPair()

	g.isFellOut = false
	return
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

		isEreased := g.fellOutProc()
		g.deploy()

		g.inputKey()
		g.deploy()

		g.falling()
		g.deploy()

		if g.judgeGameOver() && !isEreased {
			g.current = GameOver
		}
		break

	case GameOver:
		for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
			if ebiten.IsKeyPressed(k) {
				g.Initialize()
			}
		}
		break
	}

	return nil
}
