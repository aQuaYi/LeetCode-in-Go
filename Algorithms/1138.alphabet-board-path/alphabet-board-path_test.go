package problem1138

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target string
	ans    string
}{

	{
		"dfdsfkjdfleiwllsgoidkjflsdfjdfierwsd",
		"RRR!LLLD!URRR!DDD!LLLUU!D!URRRR!LU!LLLD!DR!UURRR!LD!LDDD!LUU!!DRR!LLUU!DRRR!LU!U!LLLDD!URRRR!LLLL!DR!DRR!UUU!LLLD!RRRR!LU!LLLD!RRR!UR!LLDDD!D!UR!UUU!",
	},

	{
		"zz",
		"DDDDD!!",
	},

	{
		"zdz",
		"DDDDD!UUUUURRR!LLLDDDDD!",
	},

	{
		"zb",
		"DDDDD!UUUUUR!",
	},

	{
		"lazy",
		"DDR!LUU!DDDDD!URRRR!",
	},

	{
		"leet",
		"DDR!UURRR!!DDD!",
	},

	{
		"code",
		"RR!DDRR!LUU!R!",
	},

	// 可以有多个 testcase
}

func Test_alphabetBoardPath(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, alphabetBoardPath(tc.target), "输入:%v", tc)
	}
}

func Benchmark_alphabetBoardPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			alphabetBoardPath(tc.target)
		}
	}
}

func Test_coordinate(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{
			"a 的坐标",
			args{
				r: 'a',
			},
			0, 0,
		},
		{
			"z 的坐标",
			args{
				r: 'z',
			},
			5, 0,
		},
		{
			"i 的坐标",
			args{
				r: 'i',
			},
			1, 3,
		},
		{
			"m 的坐标",
			args{
				r: 'm',
			},
			2, 2,
		},
		{
			"q 的坐标",
			args{
				r: 'q',
			},
			3, 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := coordinate(tt.args.r)
			if gotX != tt.wantX {
				t.Errorf("coordinate() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("coordinate() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
