package problem0636

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	size := len(logs)

	res := make([]int, n)
	// stack 用于存放　start 的 function ID
	stack := make([]int, 0, size)
	// 注意，题目中 start:2 和 end:2 相隔了一秒
	// endPoint = 3 相当于题目中的　end:2
	endPoint := 0
	lastEndPoint := 0

	for _, log := range logs {
		fid, soe, ts := analyze(log)

		if soe == "start" {
			// start:ts 相当于 endPoint = ts
			endPoint = ts
			if len(stack) > 0 {
				// 遇到了新的 start
				// 上一个 start 的 function 就需要即使统计运行时间
				res[stack[len(stack)-1]] += endPoint - lastEndPoint
			}
			stack = append(stack, fid)
			lastEndPoint = endPoint
		} else {
			// end:ts 相当于 endPoint = ts + 1
			endPoint = ts + 1
			res[fid] += endPoint - lastEndPoint
			stack = stack[:len(stack)-1]
			lastEndPoint = endPoint
		}

	}

	return res
}

// 把 "function_id:start_or_end:timestamp" 解析后
// 返回 function_id, start_or_end, timestamp
func analyze(log string) (int, string, int) {
	ss := strings.Split(log, ":")
	funcID, _ := strconv.Atoi(ss[0])
	timeStamp, _ := strconv.Atoi(ss[2])
	return funcID, ss[1], timeStamp
}
