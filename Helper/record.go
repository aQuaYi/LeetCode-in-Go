package main

import "fmt"

type record struct {
	Easy, Medium, Hard, Total count
}

type count struct {
	Solved, Total int
}

func (r *record) progressTable() string {
	res := fmt.Sprintln("|     |Easy|Medium|Hard|Total|")
	res += fmt.Sprintln("|:---:|:---:|:---:|:---:|:---:|")

	res += fmt.Sprintf("|**Accepted**|%d|", r.Easy.Solved)
	res += fmt.Sprintf("%d|", r.Medium.Solved)
	res += fmt.Sprintf("%d|", r.Hard.Solved)
	res += fmt.Sprintf("%d|\n", r.Total.Solved)

	res += fmt.Sprintf("|**Total**|%d|", r.Easy.Total)
	res += fmt.Sprintf("%d|", r.Medium.Total)
	res += fmt.Sprintf("%d|", r.Hard.Total)
	res += fmt.Sprintf("%d|\n", r.Total.Total)

	return res
}

func (r *record) update(p problem) {
	if !p.isAvailble() {
		return
	}
	switch p.Difficulty {
	case 1:
		r.Easy.Total++
		if p.IsAccepted {
			r.Easy.Solved++
		}
	case 2:
		r.Medium.Total++
		if p.IsAccepted {
			r.Medium.Solved++
		}
	case 3:
		r.Hard.Total++
		if p.IsAccepted {
			r.Hard.Solved++
		}
	}
	r.Total.Total++
	if p.IsAccepted {
		r.Total.Solved++
	}
}
