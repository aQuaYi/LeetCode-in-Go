# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```text
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 解题思路

```go
a + b = target
```

也可以看成是

```go
a = target - b
```

在`map[整数]整数的序号`中，可以查询到a的序号。这样就不用嵌套两个for循环了。