# [823. Binary Trees With Factors](https://leetcode.com/problems/binary-trees-with-factors/)

## 题目

Given an array of unique integers, each integer is strictly greater than 1.

We make a binary tree using these integersand each number may be used for any number of times.

Each non-leaf node'svalue should be equal to the product of the values of it's children.

How many binary trees can we make? Return the answer modulo 10 ** 9 + 7.

Example 1:

```text
Input: A = [2, 4]
Output: 3
Explanation: We can make these trees: [2], [4], [4, 2, 2]
```

Example 2:

```text
Input: A = [2, 4, 5, 10]
Output: 7
Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
```

Note:

1. 1 <= A.length <=1000.
1. 2 <=A[i]<=10 ^ 9.

## 解题思路

见程序注释
