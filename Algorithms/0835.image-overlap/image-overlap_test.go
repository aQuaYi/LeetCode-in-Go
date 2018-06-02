package problem0835

import "testing"

type args struct {
	A [][]int
	B [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		"3X3 image",
		args{
			[][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}},
			[][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}},
		},
		3,
	},
}

func Test_largestOverlap(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestOverlap(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("largestOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Benchmark_largestOverlap(b *testing.B) {
	for i := 1; i < b.N; i++ {
		for _, tt := range tests {
			a, b := tt.args.A, tt.args.B
			largestOverlap(a, b)
		}
	}
}
