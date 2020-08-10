package game

type State int

const (
	Ready State = iota
	Playing
	GameOver
)

type Game struct {
	f int

	current  State
	board    Board
	playerPP *PuyoPair
	nextPP   *PuyoPair

	isFellOut bool
}

var G *Game

func (g *Game) Initialize() {
	g.f = 0
	g.current = Ready
	g.board = newBoard()
	g.playerPP = NewPuyoPair()
	g.nextPP = NewPuyoPair()
}

func init() {
	G = &Game{}
	G.Initialize()
}
