package problem0736

import "strconv"

func evaluate(expression string) int {
	return helper(expression, nil)
}

func helper(exp string, s *scope) int {
	if exp[0] != '(' {
		num, err := strconv.Atoi(exp)
		if err != nil {
			return s.get(exp)
		}
		return num
	}

	// 删除外层的 "()"
	exp = exp[1 : len(exp)-1]
	var keyWord string
	keyWord, exp = split(exp)
	switch keyWord {
	case "add":
		a, b := split(exp)
		return helper(a, s) + helper(b, s)
	case "mult":
		a, b := split(exp)
		return helper(a, s) * helper(b, s)
	default:
		// 遇到 let 就意味着要生成新的 scope 了
		s = newScope(s)
		var key, val string
		for {
			key, exp = split(exp)
			if exp == "" {
				break
			}
			val, exp = split(exp)
			s.add(key, helper(val, s))
		}
		return helper(key, s)
	}
}

// split 把 exp 分割成 pre 和 post
// 其中 pre 是 exp 中的第一个独立的块，post 则是剩下的部分
// 比如
//     exp = "add 1 2"
//     则
//     pre  = "add"
//     post = "1 2"
// 又比如
//     exp = "(add x1 (add x2 x3))"
//     则
//     pre  = "(add x1 (add x2 x3))"
//     post = ""
func split(exp string) (pre, post string) {
	n := len(exp)
	i := 0
	if exp[0] == '(' {
		countLeft := 0
		for i < n {
			if exp[i] == '(' {
				countLeft++
			} else if exp[i] == ')' {
				countLeft--
			}
			if countLeft == 0 {
				break
			}
			i++
		}
	} else {
		for i+1 < n {
			if exp[i+1] == ' ' {
				break
			}
			i++
		}
	}

	pre = exp[:i+1]
	if i+1 == n {
		post = ""
	} else {
		post = exp[i+2:]
	}

	return
}

type scope struct {
	parent *scope
	m      map[string]int
}

func newScope(parent *scope) *scope {
	return &scope{parent: parent, m: make(map[string]int, 8)}
}

func (s *scope) get(key string) int {
	if val, ok := s.m[key]; ok {
		return val
	}
	return s.parent.get(key)
}

func (s *scope) add(key string, val int) {
	s.m[key] = val
}
