# [220. Contains Duplicate III](https://leetcode.com/problems/contains-duplicate-iii/)

## 题目
Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

## 解题思路
```
存在 | i - j | <= k 使得 | nums[i] - nums[j] | <= t
```
见程序注释
