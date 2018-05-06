# [816. Ambiguous Coordinates](https://leetcode.com/problems/ambiguous-coordinates/)

## 题目

We had some 2-dimensional coordinates, like "(1, 3)" or "(2, 0.5)". Then, we removedall commas, decimal points, and spaces, and ended up with the stringS. Return a list of strings representingall possibilities for what our original coordinates could have been.

Our original representation never had extraneous zeroes, so we never started with numbers like "00", "0.0", "0.00", "1.0", "001", "00.01", or any other number that can be represented withless digits. Also, a decimal point within a number never occurs without at least one digit occuring before it, so we never started with numbers like ".1".

The final answer list can be returned in any order. Also note that all coordinates in the final answerhave exactly one space between them (occurring after the comma.)

```text
Example 1:
Input: "(123)"
Output: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
```

```text
Example 2:
Input: "(00011)"
Output: ["(0.001, 1)", "(0, 0.011)"]
Explanation:
0.0, 00, 0001 or 00.01 are not allowed.
```

```text
Example 3:
Input: "(0123)"
Output: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]
```

```text
Example 4:
Input: "(100)"
Output: [(10, 0)]
Explanation:
1.0 is not allowed.
```

Note:

1. 4 <= S.length <= 12.
1. S[0] = "(", S[S.length - 1] = ")", and the other elements in S are digits.

## 解题思路

见程序注释
