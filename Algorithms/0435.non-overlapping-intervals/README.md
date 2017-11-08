# [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)

## 题目

Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note:

1. You may assume the interval's end point is always bigger than its start point.
1. Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

Example 1:

```text
Input: [ [1,2], [2,3], [3,4], [1,3] ]

Output: 1

Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
```

Example 2:

```text
Input: [ [1,2], [1,2], [1,2] ]

Output: 2

Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
```

Example 3:

```text
Input: [ [1,2], [2,3] ]

Output: 0

Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
```

## 解题思路

见程序注释
