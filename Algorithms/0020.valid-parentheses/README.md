# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## 题目
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

## 解题思路
[栈](https://zh.wikipedia.org/wiki/%E5%A0%86%E6%A0%88)是一个后进先出的队列，用在这里可以避免复杂的判断结构。但是，Go语言的标准库没有栈这种结构，我就手动实现了一个。

## 总结
选用合适的数据结构，可以让程序清晰。