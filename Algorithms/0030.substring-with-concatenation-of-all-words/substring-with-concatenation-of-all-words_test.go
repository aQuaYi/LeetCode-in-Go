package Problem0030

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one string
	two []string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0030(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{"barfoothefoobarman", []string{"foo", "bar"}},
			ans{[]int{0, 9}},
		},

		question{
			para{"barbarfoothefoobarman", []string{"foo", "bar"}},
			ans{[]int{3, 12}},
		},

		question{
			para{"attoinattoin", []string{"at", "to", "in"}},
			ans{[]int{0, 2, 4, 6}},
		},

		question{
			para{"abc", []string{""}},
			ans{[]int{0, 1, 2, 3}},
		},

		question{
			para{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}},
			ans{[]int{8}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		res := findSubstring(p.one, p.two)
		sort.Ints(res)
		ast.Equal(a.one, res, "输入:%v", p)
	}
}

func Test_getRecord(t *testing.T) {
	ast := assert.New(t)
	record := getRecord("aAABBAAaaAAaaBBaaBBAABB", []string{"AA", "BB"})
	ans := map[int][]int{
		1:  []int{0, 1, 0},
		17: []int{1, 0, 1},
	}
	for k, r := range record {
		v, ok := ans[k]
		ast.True(ok, "ans中没有键为%d", k)
		ast.Equal(v, r, "收集的数据不对")
	}
}

func Test_addRecord(t *testing.T) {
	ast := assert.New(t)
	record := map[int][]int{}

	addRecord("AABBAAaaAAaaBBaaBBAABB", []string{"AA", "BB"}, 0, &record)
	ans := map[int][]int{
		0:  []int{0, 1, 0},
		16: []int{1, 0, 1},
	}
	for k, r := range record {
		v, ok := ans[k]
		ast.True(ok, "ans中没有键为%d", k)
		ast.Equal(v, r, "收集的数据不对")
	}
}

func Test_isFound(t *testing.T) {
	ast := assert.New(t)

	ok := isFound([]int{0, 2, 1, 3})
	ast.True(ok, "的确发现了")

	ok = isFound([]int{0, 3, 0, 3})
	ast.False(ok, "确实没有发现")
}
func Test_getIndex(t *testing.T) {
	ast := assert.New(t)

	id := getIndex([]string{"ab", "cd"}, "cd")
	ast.Equal(1, id, "id不正确")

	id = getIndex([]string{"ab", "cd"}, "ef")
	ast.Equal(-1, id, "id不正确")
}
