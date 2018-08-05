# [528. Random Pick with Weight](https://leetcode.com/problems/random-pick-with-weight/)

## 题目

Given an array w of positive integers, where w[i] describes the weight of index i,write a function pickIndexwhich randomlypicks an indexin proportionto its weight.

Note:

1. 1 <= w.length <= 10000
1. 1 <= w[i] <= 10^5
1. pickIndex will be called at most 10000 times.

Example 1:

```text
Input:
["Solution","pickIndex"]
[[[1]],[]]
Output: [null,0]
```

Example 2:

```text
Input:
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
Output: [null,0,1,1,1,0]
```

Explanation of Input Syntax:

The input is two lists:the subroutines calledand theirarguments.Solution'sconstructor has one argument, thearray w. pickIndex has no arguments.Argumentsarealways wrapped with a list, even if there aren't any.

## 解题思路

见程序注释
