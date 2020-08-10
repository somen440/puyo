package game

type connMap struct {
	color  Color
	mainID PID
	idMap  map[PID]bool
}

func (cm *connMap) IsMatchColor(color Color) bool {
	return cm.color == color
}

func (cm *connMap) Has(id PID) bool {
	_, ok := cm.idMap[id]
	return ok
}

func (cm *connMap) Add(id PID) {
	cm.idMap[id] = true
}

func (cm *connMap) CanErase(cNum int) bool {
	return len(cm.idMap) >= cNum
}

func Ereaser(board Board, playerPP *PuyoPair, cNum int) (Board, bool) {
	if cNum <= 1 {
		return newBoard(), true
	}
	xLen := len(board[0])
	yLen := len(board)

	cMS := []*connMap{}

	isMoving := func(x, y int) bool {
		if y == yLen-1 {
			return false
		}
		return board[y+1][x] == 0
	}

	for y, rows := range board {
		for x, c := range rows {
			if c == 0 || playerPP.Has(x, y) || isMoving(x, y) {
				continue
			}

			mPID := ToID(x, y, c)
			cm := &connMap{color: c, idMap: map[PID]bool{
				mPID: true,
			}, mainID: mPID}
			cMS = append(cMS, cm)

			addCM := func(nx, ny int) {
				if !((0 <= nx && nx < xLen) && (0 <= ny && ny < yLen)) {
					return
				}
				if playerPP.Has(nx, ny) || isMoving(nx, ny) {
					return
				}
				sc := board[ny][nx]
				for _, cm := range cMS {
					if cm.Has(mPID) && cm.IsMatchColor(sc) {
						cm.Add(ToID(nx, ny, sc))
					}
				}
			}
			addCM(x, y-1)
			addCM(x, y+1)
			addCM(x-1, y)
			addCM(x+1, y)
		}
	}

	deletableIDM := map[PID]bool{}
	for _, cm := range cMS {
		if !cm.CanErase(cNum) {
			continue
		}
		for id := range cm.idMap {
			deletableIDM[id] = true
		}
	}

	isErased := false
	for y, rows := range board {
		for x, c := range rows {
			pid := ToID(x, y, c)
			_, ok := deletableIDM[pid]
			if !ok {
				continue
			}
			board[y][x] = 0
			isErased = true
		}
	}

	return board, isErased
}
