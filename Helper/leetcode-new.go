package main

import (
	"time"
)

func getLeetCode() *leetcode {
	probs, record := parseAlgs(getAlgorithms())
	lc := &leetcode{
		Username: getConfig().Username,

		Record:   record,
		Problems: *probs,

		Ranking: getRanking(),
		updated: time.Now(),
	}

	// 每更新一次，就保存一次
	lc.save()

	return lc
}

func parseAlgs(algs *algorithms) (*problems, record) {
	probs := &problems{}
	r := record{}

	for _, ps := range algs.Problems {
		p := newProblem(ps)
		probs.add(p)
		r.update(p)
	}

	return probs, r
}
