package game

func hitDown(board Board, playerPP *PuyoPair) bool {
	yLen := len(board)

	fx, fy, _ := playerPP.First.X, playerPP.First.Y, playerPP.First.Color
	sx, sy, _ := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	if fy == yLen-1 || sy == yLen-1 {
		return true
	}

	fny := fy + 1
	if fx == sx && fny == sy {
		return board[sy+1][sx] != 0
	}

	return board[fy+1][fx] != 0
}

func hitLeft(board Board, playerPP *PuyoPair) bool {
	fx, fy, _ := playerPP.First.X, playerPP.First.Y, playerPP.First.Color
	sx, sy, _ := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	if fx == 0 || sx == 0 {
		return true
	}

	fnx := fx - 1
	if fnx == sx && fy == sy {
		return board[sy][sx-1] != 0
	}

	return board[fy][fnx] != 0
}

func hitRight(board Board, playerPP *PuyoPair) bool {
	xLen := len(board[0])

	fx, fy, _ := playerPP.First.X, playerPP.First.Y, playerPP.First.Color
	sx, sy, _ := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	if fx == xLen-1 || sx == xLen-1 {
		return true
	}

	fnx := fx + 1
	if fnx == sx && fy == sy {
		return board[sy][sx+1] != 0
	}

	return board[fy][fnx] != 0
}
