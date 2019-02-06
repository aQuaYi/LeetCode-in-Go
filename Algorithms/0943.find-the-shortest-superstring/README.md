# [943. Find the Shortest Superstring](https://leetcode.com/problems/find-the-shortest-superstring/)

Given an array A of strings, find any smallest string that contains each string in A as a substring.

We may assume that no string in A is substring of another string in A.

Example 1:

```text
Input: ["alex","loves","leetcode"]
Output: "alexlovesleetcode"
Explanation: All permutations of "alex","loves","leetcode" would also be accepted.
```

Example 2:

```text
Input: ["catg","ctaagt","gcta","ttca","atgcatc"]
Output: "gctaagttcatgcatc"
```

Note:

1. 1 <= A.length <= 12
1. 1 <= A[i].length <= 20
