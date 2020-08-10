package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropper(t *testing.T) {
	tests := []struct {
		title             string
		board             Board
		playerPP          *PuyoPair
		expectedBoard     Board
		expectedPlayerPP  *PuyoPair
		expectedIsFellOut bool
	}{
		{
			"全部底にある",
			Board{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{2, 1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1, 2},
			},
			&PuyoPair{
				First: &Puyo{
					X:     2,
					Y:     11,
					Color: 1,
				},
				Second: &Puyo{
					X:     2,
					Y:     10,
					Color: 2,
				},
			},
			Board{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{2, 1, 2, 1, 2, 1},
				{1, 2, 1, 2, 1, 2},
			},
			&PuyoPair{
				First: &Puyo{
					X:     2,
					Y:     11,
					Color: 1,
				},
				Second: &Puyo{
					X:     2,
					Y:     10,
					Color: 2,
				},
			},
			false,
		},
		{
			"落ちる",
			Board{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0},
				{2, 0, 2, 0, 2, 1},
				{1, 2, 1, 0, 1, 2},
			},
			&PuyoPair{
				First: &Puyo{
					X:     1,
					Y:     6,
					Color: 2,
				},
				Second: &Puyo{
					X:     1,
					Y:     7,
					Color: 1,
				},
			},
			Board{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{2, 0, 2, 1, 2, 1},
				{1, 2, 1, 0, 1, 2},
			},
			&PuyoPair{
				First: &Puyo{
					X:     1,
					Y:     7,
					Color: 2,
				},
				Second: &Puyo{
					X:     1,
					Y:     8,
					Color: 1,
				},
			},
			false,
		},
		{
			"落ちた（PlayerPP も落ち切った）",
			Board{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0},
				{0, 1, 0, 1, 0, 0},
				{2, 0, 2, 0, 2, 1},
				{1, 2, 1, 0, 1, 2},
			},
			&PuyoPair{
				First: &Puyo{
					X:     1,
					Y:     8,
					Color: 2,
				},
				Second: &Puyo{
					X:     1,
					Y:     9,
					Color: 1,
				},
			},
			Board{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0},
				{2, 1, 2, 1, 2, 1},
				{1, 2, 1, 0, 1, 2},
			},
			&PuyoPair{
				First: &Puyo{
					X:     1,
					Y:     9,
					Color: 2,
				},
				Second: &Puyo{
					X:     1,
					Y:     10,
					Color: 1,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			actualBoard, actualPlayerPP, actualIsFellOut := Dropper(tt.board, tt.playerPP)
			assert.Equal(t, tt.expectedBoard, actualBoard)
			assert.Equal(t, tt.expectedPlayerPP, actualPlayerPP)
			assert.Equal(t, tt.expectedIsFellOut, actualIsFellOut)
		})
	}
}
