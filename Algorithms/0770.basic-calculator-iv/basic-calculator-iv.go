package Problem0770

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {

	return nil
}

type nums []num

type num struct {
	vars []string // 单个变量
	c    int      // 系数
}

func (n *num) update(m map[string]int) {
	for i, s := range n.vars {
		if v, ok := m[s]; ok {
			n.c *= v
			n.vars[i] = ""
		}
	}

}
