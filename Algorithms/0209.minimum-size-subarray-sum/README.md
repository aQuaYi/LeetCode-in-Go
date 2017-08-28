# [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)

## 题目
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

For example, given the array `[2,3,1,2,4,3]` and s = `7`,
the subarray `[4,3]` has the minimal length under the problem constraint.

**More practice:**

If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).

## 解题思路
充分利用题目信息，切片中的数，全部是正数。

见程序注释
