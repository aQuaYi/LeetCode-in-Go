package Problem0056

import "math/rand"

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	quickSort(intervals)

	res := make([]Interval, 0, len(intervals))
	temp := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if isOverlap(temp, intervals[i]) {
			temp = mergeInterval(temp, intervals[i])
			continue
		}
		res = append(res, temp)
		temp = intervals[i]
	}

	res = append(res, temp)

	return res
}

func isOverlap(a, b Interval) bool {
	return (a.Start <= b.Start && b.Start <= a.End) || (a.Start <= b.End && b.End <= a.End)
}

func mergeInterval(a, b Interval) Interval {
	return Interval{
		Start: min(a.Start, b.Start),
		End:   max(a.End, b.End),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func quickSort(is []Interval) {
	if len(is) <= 1 {
		return
	}

	j := rand.Intn(len(is))
	is[0], is[j] = is[j], is[0]
	j = partition(is)
	quickSort(is[0:j])
	quickSort(is[j+1:])
}

//partition 要考虑好，当a.Len()==2时，如何排好序
func partition(is []Interval) int {
	i, j := 1, len(is)-1
	for {
		for is[0].Start >= is[i].Start && i < len(is)-1 {
			i++
		}
		for is[j].Start >= is[0].Start && j > 0 {

			j--
		}
		if i >= j {
			break
		}
		is[i], is[j] = is[j], is[i]
	}
	is[0], is[j] = is[j], is[0]
	return j
}
