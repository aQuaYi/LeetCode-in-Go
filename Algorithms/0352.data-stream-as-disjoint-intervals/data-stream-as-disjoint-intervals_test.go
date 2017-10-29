package Problem0352

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	sr := Constructor()

	sr.Addnum(1)
	ast.Equal([]Interval{Interval{Start: 1, End: 1}}, sr.Getintervals())

	sr.Addnum(3)
	ast.Equal([]Interval{
		Interval{Start: 1, End: 1},
		Interval{Start: 3, End: 3},
	},
		sr.Getintervals(),
	)

	sr.Addnum(7)
	ast.Equal([]Interval{
		Interval{Start: 1, End: 1},
		Interval{Start: 3, End: 3},
		Interval{Start: 7, End: 7},
	},
		sr.Getintervals(),
	)

	sr.Addnum(2)
	ast.Equal([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 7, End: 7},
	},
		sr.Getintervals(),
	)

	sr.Addnum(6)
	ast.Equal([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 6, End: 7},
	},
		sr.Getintervals(),
	)
}
