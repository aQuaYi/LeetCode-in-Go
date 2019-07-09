# [307. Range Sum Query - Mutable](https://leetcode.com/problems/range-sum-query-mutable/)

## 题目

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Example:

```text
Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
```

Note:

1. The array is only modifiable by the update function.
1. You may assume the number of calls to update and sumRange function is distributed evenly.
