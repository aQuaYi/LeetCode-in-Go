# [162. Find Peak Element](https://leetcode.com/problems/find-peak-element/)

## 题目
A peak element is an element that is greater than its neighbors.

Given an input array where num[i] ≠ num[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that num[-1] = num[n] = -∞.

For example, in array [1, 2, 3, 1], 3 is a peak element and your function should return the index number 2.

click to show spoilers.

Note:
Your solution should be in logarithmic complexity.


## 解题思路
注意理解题意，题目要求，
1. 返回任意一个局部高点即可
1. 要求程序的复杂度为O(logN)
见程序注释
