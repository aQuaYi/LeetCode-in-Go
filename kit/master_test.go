package kit

import (
	"testing"
)

func Test_matches(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			"没有一个字母一样",
			args{
				"aaaaaaa",
				"bbbbbbb",
			},
			0,
		},

		{
			"a 和 b 只有 1 个字母一样",
			args{
				"aaaaaaa",
				"bbbbbba",
			},
			1,
		},

		{
			"a 和 b 只有 2 个字母一样",
			args{
				"aaaaaaa",
				"bbbbbaa",
			},
			2,
		},

		{
			"a 和 b 只有 3 个字母一样",
			args{
				"aaaaaaa",
				"bbbbaaa",
			},
			3,
		},

		{
			"a 和 b 只有 4 个字母一样",
			args{
				"aaaaaaa",
				"bbbaaaa",
			},
			4,
		},

		{
			"a 和 b 只有 5 个字母一样",
			args{
				"aaaaaaa",
				"bbaaaaa",
			},
			5,
		},

		{
			"a 和 b 只有 6 个字母一样",
			args{
				"aaaaaaa",
				"baaaaaa",
			},
			6,
		},

		{
			"一样的 a 和 b",
			args{
				"aaaaaaa",
				"aaaaaaa",
			},
			7,
		},

		// 添加参数
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matches(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaster_Guess(t *testing.T) {
	m := &Master{
		Secret:   "acckzz",
		WordList: []string{"acckzz", "ccbazz", "eiowzz", "abcczz"},
		Count:    5,
	}

	m.Update()

	tests := []struct {
		name  string
		m     *Master
		word  string
		want  int
		count int
	}{

		{
			"猜中了",
			m,
			"acckzz",
			6,
			4,
		},

		{
			"猜 word list 中的单词",
			m,
			"aaaaaa",
			-1,
			3,
		},

		{
			"猜中了 3 个字母",
			m,
			"ccbazz",
			3,
			2,
		},

		{
			"猜中了 2 个字母",
			m,
			"eiowzz",
			2,
			1,
		},

		{
			"猜中了 4 个字母",
			m,
			"abcczz",
			4,
			0,
		},

		// 添加参数
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Guess(tt.word); got != tt.want {
				t.Errorf("Master.Guess() = %v, want %v", got, tt.want)
			}
			if tt.m.Count != tt.count {
				t.Errorf("tt.m.Count = %v, but tt.count = %v", tt.m.Count, tt.count)
			}
		})
	}
}
