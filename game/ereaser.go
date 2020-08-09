package game

type ePuyoGroup struct {
	color int

	pMap map[string]*Puyo
}

func (e *ePuyoGroup) IsMatchColor(color int) bool {
	return e.color == color
}

func (e *ePuyoGroup) Has(p *Puyo) bool {
	_, ok := e.pMap[p.ID()]
	return ok
}

func (e *ePuyoGroup) Add(p *Puyo) {
	e.pMap[p.ID()] = p
}

func (e *ePuyoGroup) CanErase(cNum int) bool {
	return len(e.pMap) >= cNum
}

func Ereaser(puyos []*Puyo, cNum int) []*Puyo {
	results := []*Puyo{}

	if cNum <= 1 || len(puyos) == 0 {
		return results
	}

	eGroupList := []*ePuyoGroup{}
	for _, p := range puyos {
		eGroup := &ePuyoGroup{color: p.Color, pMap: map[string]*Puyo{
			p.ID(): p,
		}}
		eGroupList = append(eGroupList, eGroup)
		for np := range nextPuyo(puyos, p) {
			for _, eg := range eGroupList {
				if eg.IsMatchColor(np.Color) && eg.Has(p) {
					eg.Add(np)
				}
			}
		}
	}

	deleteIDs := map[string]bool{}
	for _, eg := range eGroupList {
		if !eg.CanErase(cNum) {
			continue
		}
		for _, p := range eg.pMap {
			deleteIDs[p.ID()] = true
		}
	}

	for _, p := range puyos {
		_, ok := deleteIDs[p.ID()]
		if !ok {
			results = append(results, p)
		}
	}

	return results
}

func nextPuyo(puyos []*Puyo, puyo *Puyo) chan *Puyo {
	ch := make(chan *Puyo)

	go func() {
		defer close(ch)

		for _, p := range puyos {
			if p.X == puyo.X {
				if p.Y == puyo.Y+1 {
					ch <- p
				} else if p.Y == puyo.Y-1 {
					ch <- p
				}
			} else if p.Y == puyo.Y {
				if p.X == puyo.X+1 {
					ch <- p
				} else if p.X == puyo.X-1 {
					ch <- p
				}
			}
		}
	}()

	return ch
}
