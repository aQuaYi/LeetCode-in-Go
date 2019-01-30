# [932. Beautiful Array](https://leetcode.com/problems/beautiful-array/)

For some fixed `N`, an array `A` is *beautiful* if it is a permutation of the integers `1, 2, ..., N`, such that:

For every `i < j`, there is **no** `k` with `i < k < j` such that `A[k] * 2 = A[i] + A[j]`.

Given `N`, return **any** beautiful array `A`.  (It is guaranteed that one exists.)

Example 1:

```text
Input: 4
Output: [2,1,4,3]
```

Example 2:

```text
Input: 5
Output: [3,1,2,5,4]
```

Note:

- `1 <= N <= 1000`
