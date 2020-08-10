package game

func Rolling(board Board, playerPP *PuyoPair) (Board, *PuyoPair) {
	fx, fy := playerPP.First.X, playerPP.First.Y
	sx, sy, sc := playerPP.Second.X, playerPP.Second.Y, playerPP.Second.Color

	board[sy][sx] = 0

	//     3, 0
	// 2, 1    4, 1
	//     3, 2
	switch {
	// 上
	case fx == sx && fy > sy:
		if !hitRight(board, playerPP) {
			sx++
			sy++
		}
		break

	// 下
	case fx == sx && fy < sy:
		if !hitLeft(board, playerPP) {
			sx--
			sy--
		}
		break

	// 左
	case fx > sx && fy == sy:
		sx++
		sy--
		break

	// 右
	case fx < sx && fy == sy:
		if !hitDown(board, playerPP) {
			sx--
			sy++
		}
		break
	}

	board[sy][sx] = Color(sc)
	playerPP.Second.X, playerPP.Second.Y = sx, sy

	return board, playerPP
}
