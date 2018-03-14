package problem0385

import (
	"strconv"
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger = kit.NestedInteger

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}

	if s[0] != '[' {
		return getValue(s)
	}

	stack := new(stackChars)

	var cur *NestedInteger

	si, ci := 0, 0
	for ; ci < len(s); ci++ {
		switch s[ci] {
		case '[':
			if cur != nil {
				stack.Push(cur)
			}

			cur = new(NestedInteger)
			si = ci + 1
		case ']':
			if ci > si {
				cur.Add(*getValue(s[si:ci]))
			}

			if !stack.Empty() {
				tmp := stack.Pop()
				tmp.Add(*cur)
				cur = tmp
			}

			si = ci + 1
		case ',':
			if s[ci-1] != ']' {
				cur.Add(*getValue(s[si:ci]))
			}
			si = ci + 1
		}
	}

	return cur
}

func getValue(s string) *NestedInteger {
	val, _ := strconv.Atoi(s)
	item := new(NestedInteger)
	item.SetInteger(val)
	return item
}

type stackChars struct {
	chars []*NestedInteger
}

func (s *stackChars) Push(nb *NestedInteger) {
	s.chars = append(s.chars, nb)
}

func (s *stackChars) Pop() *NestedInteger {
	slen := len(s.chars)
	rb := s.chars[slen-1]
	s.chars = s.chars[:slen-1]
	return rb
}

func (s *stackChars) Empty() bool {
	return len(s.chars) == 0
}

// func (s *stackChars) Peek() *NestedInteger {
// 	return s.chars[len(s.chars)-1]
// }
