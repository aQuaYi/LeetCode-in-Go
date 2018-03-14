package problem0352

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type Interval = kit.Interval

type SummaryRanges struct {
	is []Interval
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (sr *SummaryRanges) Addnum(val int) {
	if sr.is == nil {
		sr.is = []Interval{
			Interval{
				Start: val,
				End:   val,
			},
		}
		return
	}

	lo, hi := 0, len(sr.is)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if sr.is[mid].Start <= val && val <= sr.is[mid].End {
			return
		} else if val < sr.is[mid].Start {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if lo == 0 {
		if sr.is[0].Start-1 == val {
			sr.is[0].Start--
			return
		}
		ni := Interval{Start: val, End: val}
		sr.is = append(sr.is, ni)
		copy(sr.is[1:], sr.is)
		sr.is[0] = ni
		return
	}

	if lo == len(sr.is) {
		if sr.is[lo-1].End+1 == val {
			sr.is[lo-1].End++
			return
		}
		sr.is = append(sr.is, Interval{Start: val, End: val})
		return
	}

	if sr.is[lo-1].End+1 < val && val < sr.is[lo].Start-1 {
		sr.is = append(sr.is, Interval{})
		copy(sr.is[lo+1:], sr.is[lo:])
		sr.is[lo] = Interval{Start: val, End: val}
		return
	}

	if sr.is[lo-1].End == val-1 && val+1 == sr.is[lo].Start {
		sr.is[lo-1].End = sr.is[lo].End
		n := len(sr.is)
		copy(sr.is[lo:], sr.is[lo+1:])
		sr.is = sr.is[:n-1]
		return
	}

	if sr.is[lo-1].End == val-1 {
		sr.is[lo-1].End++
	} else {
		sr.is[lo].Start--
	}
}

func (sr *SummaryRanges) Getintervals() []Interval {
	return sr.is
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */
