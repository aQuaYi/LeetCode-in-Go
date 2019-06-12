# [1017. Convert to Base -2](https://leetcode.com/problems/convert-to-base-2/)

Given a number N, return a string consisting of "0"s and "1"s that represents its value in base -2 (negative two).

The returned string must have no leading zeroes, unless the string is "0".

Example 1:

```text
Input: 2
Output: "110"
Explanation: (-2) ^ 2 + (-2) ^ 1 = 2
```

Example 2:

```text
Input: 3
Output: "111"
Explanation: (-2) ^ 2 + (-2) ^ 1 + (-2) ^ 0 = 3
```

Example 3:

```text
Input: 4
Output: "100"
Explanation: (-2) ^ 2 = 4
```

Note:

- `0 <= N <= 10^9`
