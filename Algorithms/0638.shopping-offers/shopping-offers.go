package problem0638

func shoppingOffers(price []int, specials [][]int, needs []int) int {
	// 题目规定了
	// 最多买 6 种产品，每种产品最多买 6 个
	var dp [7][7][7][7][7][7]int
	for len(needs) < 6 {
		needs = append(needs, 0)
	}
	for len(price) < 6 {
		price = append(price, 0)
	}

	for i := 0; i <= needs[0]; i++ {
		for j := 0; j <= needs[1]; j++ {
			for k := 0; k <= needs[2]; k++ {
				for l := 0; l <= needs[3]; l++ {
					for m := 0; m <= needs[4]; m++ {
						for n := 0; n <= needs[5]; n++ {
							dp[i][j][k][l][m][n] = price[0]*i + price[1]*j + price[2]*k + price[3]*l + price[4]*m + price[5]*n
						}
					}
				}
			}
		}
	}

	for t := 0; t < len(specials); t++ {
		p, s := extend(specials[t])

		for i := s[0]; i <= needs[0]; i++ {
			for j := s[1]; j <= needs[1]; j++ {
				for k := s[2]; k <= needs[2]; k++ {
					for l := s[3]; l <= needs[3]; l++ {
						for m := s[4]; m <= needs[4]; m++ {
							for n := s[5]; n <= needs[5]; n++ {
								dp[i][j][k][l][m][n] = min(dp[i][j][k][l][m][n], dp[i-s[0]][j-s[1]][k-s[2]][l-s[3]][m-s[4]][n-s[5]]+p)
							}
						}
					}
				}
			}
		}
	}

	return dp[needs[0]][needs[1]][needs[2]][needs[3]][needs[4]][needs[5]]
}

// 获取 special 中的价格，并把 special 中的长度扩充为 6
func extend(special []int) (int, []int) {
	n := len(special)
	prices := special[n-1]
	special = special[:n-1]
	for len(special) < 6 {
		special = append(special, 0)
	}
	return prices, special
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
