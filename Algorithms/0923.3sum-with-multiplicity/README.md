# [923. 3Sum With Multiplicity](https://leetcode.com/problems/3sum-with-multiplicity/)

Given an integer array `A`, and an integer `target`, return the number of tuples `i, j, k`  such that `i < j < k` and `A[i] + A[j] + A[k] == target`.

**As the answer can be very large, return it modulo** `10^9 + 7`.

Example 1:

```text
Input: A = [1,1,2,2,3,3,4,4,5,5], target = 8
Output: 20
Explanation:
Enumerating by the values (A[i], A[j], A[k]):
(1, 2, 5) occurs 8 times;
(1, 3, 4) occurs 8 times;
(2, 2, 4) occurs 2 times;
(2, 3, 3) occurs 2 times.
```

Example 2:

```text
Input: A = [1,1,2,2,2,2], target = 5
Output: 12
Explanation:
A[i] = 1, A[j] = A[k] = 2 occurs 12 times:
We choose one 1 from [1,1] in 2 ways,
and two 2s from [2,2,2,2] in 6 ways.
```

Note:

1. `3 <= A.length <= 3000`
1. `0 <= A[i] <= 100`
1. `0 <= target <= 300`
