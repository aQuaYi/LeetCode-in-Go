# [941. Valid Mountain Array](https://leetcode.com/problems/valid-mountain-array/)

Given an array `A` of integers, return `true` if and only if it is a valid mountain array.

Recall that A is a mountain array if and only if:

- `A.length >= 3`
- There exists some i with `0 < i < A.length - 1` such that:
  - `A[0] < A[1] < ... A[i-1] < A[i]`
  - `A[i] > A[i+1] > ... > A[B.length - 1]`

Example 1:

```text
Input: [2,1]
Output: false
```

Example 2:

```text
Input: [3,5,5]
Output: false
```

Example 3:

```text
Input: [0,3,2,1]
Output: true
```

Note:

1. `0 <= A.length <= 10000`
1. `0 <= A[i] <= 10000`