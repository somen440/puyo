package game

func Dropper(board Board, playerPP *PuyoPair) (Board, *PuyoPair, bool) {
	isFellOut := false

	xLen := len(board[0])
	yLen := len(board)

	for x := xLen - 1; x >= 0; x-- {
		for y := yLen - 2; y >= 0; y-- {
			ny := y + 1

			mc := board[y][x]
			sc := board[ny][x]
			if sc != 0 {
				continue
			}
			board[ny][x] = mc
			board[y][x] = 0

			pp := playerPP.GetMatchPuyo(x, y)
			if pp != nil {
				pp.Y = ny

				if pp.Y == yLen-1 {
					isFellOut = true
					continue
				}
				nny := ny + 1
				if playerPP.First.X == pp.X && playerPP.First.Y == nny ||
					playerPP.Second.X == pp.X && playerPP.Second.Y == nny {
					continue
				}
				tc := board[nny][x]
				if tc != 0 {
					isFellOut = true
				}
			}
		}
	}

	return board, playerPP, isFellOut
}
