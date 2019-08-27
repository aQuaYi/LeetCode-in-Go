package problem1169

import (
	"sort"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	n := len(transactions)

	isInvalid := make([]bool, n)
	tps := make([]*trans, 0, n)
	for i, t := range transactions {
		tp := newTrans(i, t)
		if tp.amount > 1000 {
			isInvalid[i] = true
		}
		tps = append(tps, tp)
	}

	sort.Slice(tps, func(i int, j int) bool {
		return tps[i].time < tps[j].time
	})

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if tps[i].time+60 < tps[j].time {
				break
			}
			if tps[i].name == tps[j].name &&
				tps[i].city != tps[j].city {
				isInvalid[tps[i].index] = true
				isInvalid[tps[j].index] = true
			}
		}
	}

	res := make([]string, 0, n)
	for i := range isInvalid {
		if isInvalid[i] {
			res = append(res, transactions[i])
		}
	}

	return res
}

type trans struct {
	index, time, amount int
	name, city          string
}

func newTrans(i int, t string) *trans {
	name, time, amount, city := parse(t)
	return &trans{
		index:  i,
		time:   time,
		amount: amount,
		name:   name,
		city:   city,
	}
}

func parse(t string) (name string, time, amount int, city string) {
	items := strings.Split(t, ",")
	name = items[0]
	time, _ = strconv.Atoi(items[1])
	amount, _ = strconv.Atoi(items[2])
	city = items[3]
	return
}
