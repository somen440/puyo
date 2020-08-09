package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropper(t *testing.T) {
	tests := []struct {
		title    string
		puyos    []*Puyo
		expected []*Puyo
		bottom   int
	}{
		{
			"ないので何も起きない",
			[]*Puyo{},
			[]*Puyo{},
			4,
		},
		{
			"全部底にある",
			[]*Puyo{
				{0, 3, 1}, {1, 3, 1}, {2, 3, 1}, {2, 3, 1},
			},
			[]*Puyo{},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			actual := Dropper(tt.puyos, tt.bottom)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
