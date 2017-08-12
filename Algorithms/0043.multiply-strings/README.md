# [43. Multiply Strings](https://leetcode.com/problems/multiply-strings/)

## 题目
Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`.

Note:
1. The length of both `num1` and `num2` is < 110.
1. Both `num1` and `num2` contains only digits `0-9`.
1. Both `num1` and `num2` does not contain any leading zero.
1. You must not use any `built-in BigInteger library` or `convert the inputs to integer directly`.

## 解题思路
1. 利用[]int{}记录中间计算过程，此时不考虑进位，例如，"123" 和 "65" 的记录为 []int{6,17,28,15}
1. 统一处理进位，记录变为[]int{7,9,9,5}
1. 转换为string，返回

细节见程序注释。
## 总结


