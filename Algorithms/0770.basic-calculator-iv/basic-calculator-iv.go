package problem0770

import (
	"sort"
	"strconv"
	"strings"
)

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	m := make(map[string]int, len(evalvars))
	for i := range evalvars {
		m[evalvars[i]] = evalints[i]
	}

	expression = strings.Replace(expression, " ", "", -1)

	numbers := parse(expression, m)

	return getResult(numbers)
}

func parse(expression string, m map[string]int) nums {
	expression += "+" // 为了方便 for 循环
	// n1 opt1 n2 opt2 n3 opt3 ...
	// 根据 opt2 是否为 * 决定运算顺序
	n1, n2 := nums{num{c: 0}}, nums{num{c: 0}}
	opt1, opt2 := byte('+'), byte('+')
	for len(expression) > 0 {
		var n3 nums
		var i int
		// 获取 n3
		if expression[0] == '(' {
			// 遇见括号，就把括号中的内容取出
			// 进行递归
			i = indexOfCounterParentheses(expression)
			n3 = parse(expression[1:i], m)
			// 此时，i 是右括号的索引值
			i++
			// 此时，i 是右括号右边的运算符号的索引值
		} else {
			i = indexOfNextOpt(expression)
			n3raw := expression[:i]
			v, ok := m[n3raw]
			if ok {
				n3 = nums{num{c: v}}
			} else {
				i3, err := strconv.Atoi(n3raw)
				if err != nil {
					n3 = nums{num{vars: []string{n3raw}, c: 1}}
				} else {
					n3 = nums{num{c: i3}}
				}
			}
		}

		// 根据 opt2
		if opt2 == '*' {
			n2 = operate(opt2, n2, n3)
		} else {
			n1 = operate(opt1, n1, n2)
			n2 = n3
			opt1 = opt2
		}

		opt2 = expression[i]
		expression = expression[i+1:]
	}

	res := operate(opt1, n1, n2)

	return res
}

func operate(opt byte, a, b nums) nums {
	var res nums
	switch opt {
	case '+':
		res = add(a, b)
	case '-':
		res = minus(a, b)
	case '*':
		res = mult(a, b)
	}
	return res
}

func add(a, b nums) nums {
	return append(a, b...)
}

func minus(a, b nums) nums {
	for i := range b {
		b[i].c *= -1
	}
	return append(a, b...)
}

func mult(a, b nums) nums {
	res := make(nums, 0, len(a)*len(b))
	for i := range a {
		for j := range b {
			vars := make([]string, 0, len(a[i].vars)+len(b[j].vars))
			vars = append(vars, a[i].vars...)
			vars = append(vars, b[j].vars...)
			res = append(res, num{vars: vars, c: a[i].c * b[j].c})
		}
	}

	return res
}

func indexOfCounterParentheses(expression string) int {
	i := 1
	count := 1
	for ; i < len(expression); i++ {
		switch expression[i] {
		case '(':
			count++
		case ')':
			count--
		}
		if count == 0 {
			break
		}
	}
	return i
}

func indexOfNextOpt(expression string) int {
	var i int
	for i = 1; i < len(expression); i++ {
		if expression[i] == '+' ||
			expression[i] == '-' ||
			expression[i] == '*' {
			break
		}
	}
	return i
}

func getResult(numbers nums) []string {
	numbers = update(numbers)
	res := make([]string, 0, len(numbers))

	for _, n := range numbers {
		if n.c == 0 {
			continue
		}
		temp := strconv.Itoa(n.c)
		if n.key != "" {
			temp = temp + "*" + n.key
		}
		res = append(res, temp)
	}

	return res
}

type nums []num

type num struct {
	vars []string // 变量
	key  string   // 变量排序后，使用 "*" 连接，成为 key
	c    int      // 系数
}

func update(ns nums) nums {
	// 更新每个 num 的 key
	for i := range ns {
		sort.Strings(ns[i].vars)
		ns[i].key = strings.Join(ns[i].vars, "*")
	}

	sort.Slice(ns, func(i int, j int) bool {
		leni := len(ns[i].vars)
		lenj := len(ns[j].vars)
		// 首先，变量多的放在前面
		if leni != lenj {
			return leni > lenj
		}
		// 变量一样多的时候
		// 不同变量，小的放在前面
		for k := 0; k < leni; k++ {
			if ns[i].vars[k] == ns[j].vars[k] {
				continue
			}
			return ns[i].vars[k] < ns[j].vars[k]
		}
		// i，j 所有变量都一样的时候
		// 返回 false 表示不用交换
		return false
	})

	res := make(nums, 1, len(ns))
	res[0] = ns[0]
	i, j := 0, 1
	for j < len(ns) {
		if res[i].key != ns[j].key {
			res = append(res, ns[j])
			i++
		} else {
			res[i].c += ns[j].c
		}
		j++
	}

	return res
}
