package problem0649

func predictPartyVictory(senate string) string {
	n := len(senate)

	// qr 和 qd 按顺序保存了 r 和 d 的议员的索引号
	qr := make([]int, 0, n)
	qd := make([]int, 0, n)
	for i, b := range senate {
		if b == 'R' {
			qr = append(qr, i)
		} else {
			qd = append(qd, i)
		}
	}

	// 对议员来说，最优策略就是，把下一个对方党派的议员 ban 掉
	for len(qr) > 0 && len(qd) > 0 {
		ri := qr[0]
		qr = qr[1:]
		di := qd[0]
		qd = qd[1:]
		if ri < di {
			// ri ban 掉了 di
			// ri 的索引号 +n
			qr = append(qr, ri+n)
		} else {
			qd = append(qd, di+n)
		}
	}

	if len(qr) > 0 {
		return "Radiant"
	}
	return "Dire"
}
