# [710. Random Pick with Blacklist](https://leetcode.com/problems/random-pick-with-blacklist/)

## 题目

Given a blacklistB containing unique integersfrom [0, N), write a function to return a uniform random integer from [0, N) which is NOTin B.

Optimize it such that it minimizes the call to system&rsquo;s Math.random().

Note:

- 1 <= N <= 1000000000
- 0 <= B.length < min(100000, N)
- [0, N)does NOT include N. See interval notation.

Example 1:

```text
Input:
["Solution","pick","pick","pick"]
[[1,[]],[],[],[]]
Output: [null,0,0,0]
```

Example 2:

```text
Input:
["Solution","pick","pick","pick"]
[[2,[]],[],[],[]]
Output: [null,1,1,1]
```

Example 3:

```text
Input:
["Solution","pick","pick","pick"]
[[3,[1]],[],[],[]]
Output: [null,0,0,2]
```

Example 4:

```text
Input:
["Solution","pick","pick","pick"]
[[4,[2]],[],[],[]]
Output: [null,1,3,1]
```

Explanation of Input Syntax:

The input is two lists:the subroutines calledand theirarguments.Solution'sconstructor has two arguments,N and the blacklist B. pick has no arguments.Argumentsarealways wrapped with a list, even if there aren't any.

## 解题思路

见程序注释
