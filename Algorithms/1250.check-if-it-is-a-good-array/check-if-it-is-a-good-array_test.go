package problem1250

import "testing"

func Test_isGoodArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{12, 5, 7, 23},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{29, 6, 10},
			},
			want: true,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{3, 6},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGoodArray(tt.args.nums); got != tt.want {
				t.Errorf("isGoodArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
