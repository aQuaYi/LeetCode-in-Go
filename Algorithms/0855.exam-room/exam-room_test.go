package problem0855

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExamRoom(t *testing.T) {
	ast := assert.New(t)

	r := Constructor(10)

	seats := []int{0, 9, 4, 2}

	for i := 0; i < len(seats); i++ {
		ast.Equal(seats[i], r.Seat(), "第 %d 次，入座了 %d", i, seats[i])
	}

	r.Leave(4)

	ast.Equal(5, r.Seat(), "4 号位离座后，应该是 5 号位入座")

	r.Leave(0)

	ast.Equal(0, r.Seat(), "0 号位离座后，应该是 0 号位入座")
}

func Test_ExamRoom_2(t *testing.T) {
	ast := assert.New(t)
	r := Constructor(1000000000)

	for i := 0; i < 10; i++ {
		ast.Equal(0, r.Seat(), "第 %d 次，入座了 %d", i, 0)
		r.Leave(0)
	}
}
