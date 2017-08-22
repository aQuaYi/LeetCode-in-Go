package Problem0056

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

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
	return (a.Start <= b.Start && b.Start <= a.End) || (a.Start <= b.End && b.End <= a.End) || (b.Start <= a.Start && a.Start <= b.End)
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
