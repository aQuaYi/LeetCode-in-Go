package Problem0753

func openLock(deadends []string, target string) int {
	dp := [10][10][10][10]int{}

for _,d:=range deadends {
dp[d[0]-'0'][d[1]-'0'][d[2]-'0'][d[3]-'0']	= -1
}

	res := 0

	return res
}


func getIndex(s string) []int {
res := make([]int, 4)	
for i:=0; i<4; i++{
res[i] = s[i]-'0')	
}
}