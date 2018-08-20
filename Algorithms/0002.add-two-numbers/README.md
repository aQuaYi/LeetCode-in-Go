# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

## 题目

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

```text
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```

## 解题思路

```text
(2 -> 4 -> 3)是 342

(5 -> 6 -> 4)是 465

(7 -> 0 -> 8)是 807

342 + 465 = 807
```

所以，题目的本意是，把整数换了一种表达方式后，实现其加法。

设计程序时候，需要处理的点有

1. 位上的加法，需要处理进位问题
1. 如何进入下一位运算
1. 按位相加结束后，也还需要处理进位问题。

## 总结

读懂题意后，按步骤实现题目要求。