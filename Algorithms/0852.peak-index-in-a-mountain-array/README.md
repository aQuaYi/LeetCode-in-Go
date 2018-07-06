# [852. Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/)

## 题目

Let's call an array A a mountainif the following properties hold:

- A.length >= 3
- There exists some 0 < i< A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]

Given an array that is definitely a mountain, return anyisuch thatA[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

```text
Input: [0,1,0]
Output: 1
```

Example 2:

```text
Input: [0,2,1,0]
Output: 1
```

Note:

1. 3 <= A.length <= 10000
1. 0 <= A[i] <= 10^6
1. Ais a mountain, as defined above.

## 解题思路

见程序注释
