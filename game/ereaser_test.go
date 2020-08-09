package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErease(t *testing.T) {
	tests := []struct {
		title    string
		puyos    []*Puyo
		cNum     int
		expected []*Puyo
	}{
		{
			"くっつき数 0 定義なので全部消える",
			[]*Puyo{
				{1, 10, 1}, {2, 10, 1}, {3, 10, 1},
			},
			0,
			[]*Puyo{},
		},
		{
			"くっつき数 1 定義なので全部消える",
			[]*Puyo{
				{1, 10, 1}, {2, 10, 1}, {3, 10, 1},
			},
			1,
			[]*Puyo{},
		},
		{
			"くっつき数 2 定義 / 前後にくっつきがなく消えない",
			[]*Puyo{
				{1, 9, 2}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 1}, {2, 10, 2}, {3, 10, 1},
			},
			2,
			[]*Puyo{
				{1, 9, 2}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 1}, {2, 10, 2}, {3, 10, 1},
			},
		},
		{
			"くっつき数 2 定義 / 縦にくっつきあり消える",
			[]*Puyo{
				{1, 9, 3}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 3}, {2, 10, 2}, {3, 10, 1},
			},
			2,
			[]*Puyo{
				{2, 9, 1}, {3, 9, 2},
				{2, 10, 2}, {3, 10, 1},
			},
		},
		{
			"くっつき数 2 定義 / 縦にくっつき（離れた位置で）あり消える",
			[]*Puyo{
				{1, 9, 3}, {2, 9, 1}, {3, 9, 2}, {4, 9, 3},
				{1, 10, 3}, {2, 10, 2}, {3, 10, 1}, {4, 10, 3},
			},
			2,
			[]*Puyo{
				{2, 9, 1}, {3, 9, 2},
				{2, 10, 2}, {3, 10, 1},
			},
		},
		{
			"くっつき数 2 定義 / 横にくっつきあり消える",
			[]*Puyo{
				{1, 9, 2}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 3}, {2, 10, 3}, {3, 10, 1},
			},
			2,
			[]*Puyo{
				{1, 9, 2}, {2, 9, 1}, {3, 9, 2},
				{3, 10, 1},
			},
		},
		{
			"くっつき数 2 定義 / 縦. 横にくっつきあり消える",
			[]*Puyo{
				{1, 9, 4}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 4}, {2, 10, 3}, {3, 10, 3},
			},
			2,
			[]*Puyo{
				{2, 9, 1}, {3, 9, 2},
			},
		},
		{
			"くっつき数 2 定義 / 全くっつき",
			[]*Puyo{
				{1, 9, 4}, {2, 9, 4}, {3, 9, 4},
				{1, 10, 4}, {2, 10, 4}, {3, 10, 4},
			},
			2,
			[]*Puyo{},
		},
		{
			"くっつき数 3 定義 / 縦. 横にくっつき 2 なので消えない",
			[]*Puyo{
				{1, 9, 4}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 4}, {2, 10, 3}, {3, 10, 3},
			},
			3,
			[]*Puyo{
				{1, 9, 4}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 4}, {2, 10, 3}, {3, 10, 3},
			},
		},
		{
			"くっつき数 3 定義 / 縦にくっつき消える",
			[]*Puyo{
				{1, 9, 4}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 4}, {2, 10, 3}, {3, 10, 3},
				{1, 11, 4}, {2, 11, 1}, {3, 11, 2},
			},
			3,
			[]*Puyo{
				{2, 9, 1}, {3, 9, 2},
				{2, 10, 3}, {3, 10, 3},
				{2, 11, 1}, {3, 11, 2},
			},
		},
		{
			"くっつき数 4 定義 / 縦にくっつき消える",
			[]*Puyo{
				{1, 9, 4}, {2, 9, 1}, {3, 9, 2},
				{1, 10, 4}, {2, 10, 3}, {3, 10, 3},
				{1, 11, 4}, {2, 11, 1}, {3, 11, 2},
				{1, 12, 4}, {2, 12, 1}, {3, 12, 2},
			},
			4,
			[]*Puyo{
				{2, 9, 1}, {3, 9, 2},
				{2, 10, 3}, {3, 10, 3},
				{2, 11, 1}, {3, 11, 2},
				{2, 12, 1}, {3, 12, 2},
			},
		},
		{
			"くっつき数 4 定義 / 縦横にくっつき消える（くっつきが定義に満たない箇所は消えない）",
			[]*Puyo{
				{1, 9, 4}, {2, 9, 1}, {3, 9, 2}, {4, 9, 4},
				{1, 10, 4}, {2, 10, 3}, {3, 10, 3}, {4, 10, 3},
				{1, 11, 4}, {2, 11, 4}, {3, 11, 2}, {4, 11, 4},
				{1, 12, 4}, {2, 12, 1}, {3, 12, 2}, {4, 12, 4},
			},
			4,
			[]*Puyo{
				{2, 9, 1}, {3, 9, 2}, {4, 9, 4},
				{2, 10, 3}, {3, 10, 3}, {4, 10, 3},
				{3, 11, 2}, {4, 11, 4},
				{2, 12, 1}, {3, 12, 2}, {4, 12, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			actual := Ereaser(tt.puyos, tt.cNum)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
