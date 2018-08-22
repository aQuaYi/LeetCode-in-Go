# [886. Possible Bipartition](https://leetcode.com/problems/possible-bipartition/)

## 题目

Given a set of Npeople (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.

Each person may dislike some other people, and they should not go into the same group.

Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.

Return trueif and only if it is possible to split everyone into two groups in this way.

Example 1:

```text
Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
```

Example 2:

```text
Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
```

Example 3:

```text
Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false
```

Note:

1. 1 <= N <= 2000
1. 0 <= dislikes.length <= 10000
1. 1 <= dislikes[i][j] <= N
1. dislikes[i][0] < dislikes[i][1]
1. There does not exist i != j for which dislikes[i] == dislikes[j].

## 解题思路

见程序注释
