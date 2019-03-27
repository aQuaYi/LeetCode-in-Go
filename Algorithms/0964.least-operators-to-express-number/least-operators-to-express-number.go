package problem0964

// ref: https://leetcode.com/problems/least-operators-to-express-number/discuss/208376/python2-O(log-target)-chinese

// 由于不能有括号，所以每一位（x进制）必须是由自己组成或者由下一位减自己余数组成,所以首先短除法求出每一位的余数
// 在个位的时候，必须用x/x表示1，所以*2，但其它位不用，所以单独先提出
// 正数表示时，可用自己+剩下的正数表示或者多加一个本位然后减去上一位的余数表示
// 负数表示时，可用自己的余数加上剩下的正数表示或者用自己的余数+剩下的余数，但此时可以合并同级项减少运算符
// 如在10进制下，86可表示为
// 80 + 6 （6 为下一位正数表示
// 80 + 10 - 4 （4 为下一位负数表示）
// 100 - 20 + 6 （100-20为本位余数表示，6为下一位正数表示
// 100 - 20 + 10 - 4 （100-20为本位余数表示，10 -4 为下一位余数表示
// 在此时， 20 和 10注定为同一位且符号相反，可以省去两个符号（一个符号在该位用i 个符号表示（如100为第二位，所以表示为+10 * 10，用两个符号，在此时所有符号均带自己的正负号
// 因为在前面算法中所有位都带自己的正负号，所以第一位应该减去自己的符号，所以总量减1
// 或者在余数表示法中，加上一个更高位的减去自己表示本位余数
// 所以此题归根结底还是考察对进制的理解而不是简单的理解bfs, dfs，那样复杂度远远高于此，但是是对惯性思维者的一种挑战

func leastOpsExpressTarget(x int, target int) int {
	target, r := target/x, target%x
	pos, neg := r*2, (x-r)*2 // 处理个位上的数

	bit := 1
	for target > 0 {
		target, r = target/x, target%x
		pos, neg = min(r*bit+pos, (r+1)*bit+neg), min((x-r)*bit+pos, (x-r-1)*bit+neg)
		bit++
	}

	return min(pos, bit+neg) - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
