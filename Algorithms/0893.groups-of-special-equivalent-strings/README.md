# [893. Groups of Special-Equivalent Strings](https://leetcode.com/problems/groups-of-special-equivalent-strings/)

## 题目

You are given an array A of strings.

Two strings S and T arespecial-equivalentif after any number of moves, S == T.

A move consists of choosing two indices i and j with i % 2 == j % 2, and swapping S[i] with S[j].

Now, a group of special-equivalent strings from Ais anon-empty subset S of Asuch that any string not in Sis not special-equivalent with any string in S.

Return the number of groups of special-equivalent strings from A.

Example 1:

```text
Input: ["a","b","c","a","c","c"]
Output: 3
Explanation: 3 groups ["a","a"], ["b"], ["c","c","c"]
```

Example 2:

```text
Input: ["aa","bb","ab","ba"]
Output: 4
Explanation: 4 groups ["aa"], ["bb"], ["ab"], ["ba"]
```

Example 3:

```text
Input: ["abc","acb","bac","bca","cab","cba"]
Output: 3
Explanation: 3 groups ["abc","cba"], ["acb","bca"], ["bac","cab"]
```

Example 4:

```text
Input: ["abcd","cdab","adcb","cbad"]
Output: 1
Explanation: 1 group ["abcd","cdab","adcb","cbad"]
```

Note:

1. 1 <= A.length <= 1000
1. 1 <= A[i].length <= 20
1. All A[i] have the same length.
1. All A[i] consist of only lowercase letters.

## 解题思路

见程序注释
