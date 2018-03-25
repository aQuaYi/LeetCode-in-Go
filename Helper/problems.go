package main

type problems []problem

func (ps *problems) add(p problem) {
	if len(*ps) <= p.ID {
		*ps = append(*ps, make([]problem, p.ID-len(*ps)+1)...)
	}
	(*ps)[p.ID] = p
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
		if !p.HasNoGoOption {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) unavailable() problems {
	res := make([]problem, 0, len(ps))
	for _, p := range ps {
		if p.HasNoGoOption {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) table() string {
	res := "|题号|题目|通过率|难度|收藏|\n"
	res += "|:-:|:-|:-: | :-: | :-: |\n"
	for _, p := range ps {
		res += p.tableLine()
	}
	return res
}

func (ps problems) list() string {
	res := ""
	for _, p := range ps {
		res += p.listLine()
	}
	return res
}
