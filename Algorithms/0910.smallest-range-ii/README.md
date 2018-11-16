# [910. Smallest Range II](https://leetcode.com/problems/smallest-range-ii/)

Given an array `A` of integers, for each integer `A[i]` we need to choose **either** `x = -K` **or** `x = K`, and add `x` to `A[i] (only once)`.

After this process, we have some array `B`.

Return the smallest possible difference between the maximum value of `B` and the minimum value of `B`.

Example 1:

```text
Input: A = [1], K = 0
Output: 0
Explanation: `B` = [1]
```

Example 2:

```text
Input: A = [0,10], K = 2
Output: 6
Explanation: `B` = [2,8]
```

Example 3:

```text
Input: A = [1,3,6], K = 3
Output: 3
Explanation: `B` = [4,6,3]
```

Note:

- `1 <= A.length <= 10000`
- `0 <= A[i] <= 10000`
- `0 <= K <= 10000`