package main



type leetcode struct {
	Username string
	Ranking  int

	Categories map[string]category
	Problems   map[int]problem
}

type category struct {
	Name  string
	Count map[string]count
}

type count struct {
	Solved, Total int
}

type problem struct {
	ID                                 int
	Title, TitleSlug                   string
	PassRate                           string
	Difficulty                         int
	IsAccepted, IsFavor, IsPaid, IsNew bool
}
