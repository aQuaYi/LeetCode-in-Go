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
		"zdz",
		"DDDDD!UUUUURRR!DDDDLLLD!",
	},

	{
		"zb",
		"DDDDD!UUUUUR!",
	},

	{
		"lazy",
		"RDD!LUU!DDDDD!RRRRU!",
	},

	{
		"leet",
		"RDD!RRRUU!!DDD!",
	},

	{
		"code",
		"RR!RRDD!LUU!R!",
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

func Test_move(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "原地移动",
			args: args{
				x1: 1,
				y1: 1,
				x2: 1,
				y2: 1,
			},
			want: "",
		},
		{
			name: "a to l",
			args: args{
				x1: 0,
				y1: 0,
				x2: 2,
				y2: 1,
			},
			want: "RDD",
		},
		{
			name: "f to j",
			args: args{
				x1: 1,
				y1: 0,
				x2: 1,
				y2: 4,
			},
			want: "RRRR",
		},
		{
			name: "j to u",
			args: args{
				x1: 1,
				y1: 4,
				x2: 4,
				y2: 0,
			},
			want: "LLLLDDD",
		},
		{
			name: "y to a",
			args: args{
				x1: 4,
				y1: 4,
				x2: 0,
				y2: 0,
			},
			want: "LLLLUUUU",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}
