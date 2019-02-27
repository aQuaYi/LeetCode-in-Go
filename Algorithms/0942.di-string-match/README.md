# [942. DI String Match](https://leetcode.com/problems/di-string-match/)

Given a string `S` that **only** contains "I" (increase) or "D" (decrease), let N = S.length.

Return **any** permutation `A` of `[0, 1, ..., N]` such that for all `i = 0, ..., N-1`:

- If `S[i] == "I"`, then `A[i] < A[i+1]`
- If `S[i] == "D"`, then `A[i] > A[i+1]`

Example 1:

```text
Input: "IDID"
Output: [0,4,1,3,2]
```

Example 2:

```text
Input: "III"
Output: [0,1,2,3]
```

Example 3:

```text
Input: "DDI"
Output: [3,2,0,1]
```

Note:

1. `1 <= S.length <= 10000`
1. `S` only contains characters "I" or "D".
