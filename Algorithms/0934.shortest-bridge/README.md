# [934. Shortest Bridge](https://leetcode.com/problems/shortest-bridge/)

In a given 2D binary array `A`, there are two islands.  (An island is a 4-directionally connected group of `1s` not connected to any other 1s.)

Now, we may change `0s` to `1s` so as to connect the two islands together to form 1 island.

Return the smallest number of `0s` that must be flipped.  (It is guaranteed that the answer is at least 1.)

Example 1:

```text
Input: [[0,1],[1,0]]
Output: 1
```

Example 2:

```text
Input: [[0,1,0],[0,0,0],[0,0,1]]
Output: 2
```

Example 3:

```text
Input: [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
Output: 1
```

Note:

1. `1 <= A.length = A[0].length <= 100`
1. `A[i][j] == 0` or `A[i][j] == 1`
