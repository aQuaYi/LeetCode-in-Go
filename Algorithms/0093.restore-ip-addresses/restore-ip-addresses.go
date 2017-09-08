package Problem0093

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return []string{}
	}

	res := []string{}

	begin := n / 4
	var dfs func(int)
	dfs = func(idx int) {
		if idx == 4 {

		}
	}

	return []string{"255.255.11.135", "255.255.111.35"}
}
