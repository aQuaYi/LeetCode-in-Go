package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aQuaYi/GoKit"
)

const (
	unavailableFile = "unavailable.json"
)

func getLeetCode() *leetcode {
	probs, record := parseAlgs(getAlgorithms())
	lc := &leetcode{
		Username: getConfig().Username,

		Record:   record,
		Problems: *probs,

		Ranking: getRanking(),
		Updated: time.Now(),
	}

	return lc
}

func parseAlgs(alg *algorithms) (*problems, record) {
	hasNoGoOption := readUnavailable()
	probs := &problems{}
	r := record{}

	for _, ps := range alg.Problems {
		p := newProblem(ps)
		if hasNoGoOption[p.ID] {
			p.HasNoGoOption = true
		}
		probs.add(p)
		r.update(p)
	}

	return probs, r
}

func readUnavailable() map[int]bool {
	type unavailable struct {
		List []int
	}

	if !GoKit.Exist(unavailableFile) {
		log.Panicf("%s 不存在，没有不能解答的题目", unavailableFile)
	}

	raw := read(unavailableFile)
	u := unavailable{}
	if err := json.Unmarshal(raw, &u); err != nil {
		log.Panicf("获取 %s 失败：%s", unavailableFile, err)
	}

	res := make(map[int]bool, len(u.List))
	for i := range u.List {
		res[u.List[i]] = true
	}

	return res
}
