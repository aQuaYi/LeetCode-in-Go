package main

import "sort"

type problems []problem

func (ps *problems) add(p problem) {
	if len(*ps) <= p.ID {
		*ps = append(*ps, make([]problem, p.ID-len(*ps)+1)...)
	}
	(*ps)[p.ID] = p
}

func (ps problems) Len() int {
	return len(ps)
}

func (ps problems) Less(i, j int) bool {
	return ps[i].ID < ps[j].ID
}

func (ps problems) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps problems) accepted() problems {
	res := make([]problem, 0, len(ps))
	for _, p := range ps {
		if p.IsAccepted {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) available() problems {
	res := make([]problem, 0, len(ps))
	for _, p := range ps {
		if p.hasGoOption {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) unavailable() problems {
	res := make([]problem, 0, len(ps))
	for _, p := range ps {
		if !p.hasGoOption {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) table() string {
	sort.Sort(ps)
	res := "|题号|题目|通过率|难度|收藏|\n"
	res += "|:-:|:-|:-: | :-: | :-: |\n"
	for _, p := range ps {
		res += p.tableLine()
	}
	return res
}

func (ps problems) list() string {
	sort.Sort(ps)
	res := ""
	for _, p := range ps {
		res += p.listLine()
	}
	return res
}
