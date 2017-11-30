package Problem0632

import (
	"sort"
)

func smallestRange(intss [][]int) []int {
	// 把 intss 中的数字连同其队伍信息，整合成 num 结构体放入 ns
	ns := makeNums(intss)
	// 让 ns 根据其数值降序排列
	sort.Sort(ns)

	s := newStatus(ns, len(intss))
	s.check()
	return s.res
}

type status struct {
	res       []int
	ns        nums
	i, j      int
	isGetAll  bool
	numCount  []int
	teamCount int
	min       int
}

func newStatus(ns nums, teamCount int) *status {
	return &status{
		res:       make([]int, 2),
		ns:        ns,
		i:         0,
		j:         -1,
		isGetAll:  false,
		numCount:  make([]int, teamCount),
		teamCount: teamCount,
		min:       1<<31 - 1,
	}
}

func (s *status) check() {
	for s.j < len(s.ns) {
		for s.j < len(s.ns) && !s.isGetAll {
			s.expend()
		}
		for s.isGetAll {
			s.updateRes()
			s.shrink()
		}
	}
}

// 让 s.j++ 并修改相关的状态值
func (s *status) expend() {
	s.j++
	if s.j >= len(s.ns) {
		return
	}
	if s.numCount[s.ns[s.j].team] == 0 {
		s.teamCount--
	}
	s.numCount[s.ns[s.j].team]++
	if s.teamCount == 0 {
		s.isGetAll = true
	}
}

// 让 s.i-- 并修改相关的状态值
func (s *status) shrink() {
	if s.numCount[s.ns[s.i].team] == 1 {
		s.teamCount++
	}
	s.numCount[s.ns[s.i].team]--

	if s.teamCount > 0 {
		s.isGetAll = false
	}
	s.i++
}

// 更新 s.res
func (s *status) updateRes() {
	beg, end := s.ns[s.j].n, s.ns[s.i].n
	if s.min >= end-beg {
		s.res[0] = beg
		s.res[1] = end
		s.min = end - beg
	}
}

func makeNums(intss [][]int) nums {
	ns := make(nums, 0, len(intss)*3)
	for i := range intss {
		temp := make(nums, len(intss[i]))
		for idx, n := range intss[i] {
			temp[idx] = num{
				n:    n,
				team: i,
			}
		}
		ns = append(ns, temp...)
	}
	return ns
}

// nums 实现了 sort.Interface 接口
type nums []num
type num struct {
	n, team int
}

func (n nums) Len() int { return len(n) }

func (n nums) Less(i, j int) bool { return n[i].n > n[j].n }

func (n nums) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
