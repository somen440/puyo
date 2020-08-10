package game

func moveLeft(board Board, playerPP *PuyoPair) (Board, *PuyoPair) {
	fx, fy, fc := playerPP.First.X, playerPP.First.Y, playerPP.First.Color
	sx, sy, sc := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	board[fy][fx] = 0
	board[sy][sx] = 0

	playerPP.First.X--
	playerPP.Second.X--

	board[fy][fx-1] = Color(fc)
	board[sy][sx-1] = Color(sc)

	return board, playerPP
}

func moveRight(board Board, playerPP *PuyoPair) (Board, *PuyoPair) {
	fx, fy, fc := playerPP.First.X, playerPP.First.Y, playerPP.First.Color
	sx, sy, sc := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	board[fy][fx] = 0
	board[sy][sx] = 0

	playerPP.First.X++
	playerPP.Second.X++

	board[fy][fx+1] = Color(fc)
	board[sy][sx+1] = Color(sc)

	return board, playerPP
}

func moveDown(board Board, playerPP *PuyoPair) (Board, *PuyoPair) {
	fx, fy, fc := playerPP.First.X, playerPP.First.Y, playerPP.First.Color
	sx, sy, sc := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	board[fy][fx] = 0
	board[sy][sx] = 0

	playerPP.First.Y++
	playerPP.Second.Y++

	board[fy+1][fx] = Color(fc)
	board[sy+1][sx] = Color(sc)

	return board, playerPP
}
