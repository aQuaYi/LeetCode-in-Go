package Problem0282

func addOperators(num string, target int) []string {
	res := []string{}
	if len(num) == 0 {
		return res
	}

	if len(num) == 1 && i(num) == target {
		res = append(res, num)
		return res
	}

	if len(num) == 2 {
		a, b := i(num[0:1]), i(num[1:])
		if a+b == target {
			res = append(res, num[0:1]+"+"+num[1:])
		}

		if a-b == target {
			res = append(res, num[0:1]+"-"+num[1:])
		}

		if a*b == target {
			res = append(res, num[0:1]+"*"+num[1:])
		}
	}

	if num[0] == '0' {

	}
	
	return
}

// 把 string 转换成 int
func i(s string) int {
	bs := []byte(s)
	res := int(bs[0] - '0')
	for i := 1; i < len(bs); i++ {
		res = res*10 + int(bs[i]-'0')
	}
	return res
}

// 给 ss 中的每个字符串加上 prefix 作为前缀
func add(prefix string, ss []string) []string {
	res := make([]string, len(ss))
	for i, s := range ss {
		res[i] = prefix + s
	}
	return res
}
