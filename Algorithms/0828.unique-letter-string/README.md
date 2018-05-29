# [828. Unique Letter String](https://leetcode.com/problems/unique-letter-string/)

## 题目

A character is unique in string S if it occurs exactly once in it.

For example, in string S = "LETTER", the only unique characters are "L" and "R".

Let's define UNIQ(S) as the number of unique characters in string S.

For example, UNIQ("LETTER") = 2.

Given a string S with only uppercases, calculate the sum of UNIQ(substring) over all non-empty substrings of S.

If there are two or more equal substrings at different positions in S, we consider them different.

Since the answer can be very large, retrun the answermodulo `10^9+7`.

Example 1:

```text
Input: "ABC"
Output: 10
Explanation: All possible substrings are: "A","B","C","AB","BC" and "ABC".
Evey substring is composed with only unique letters.
Sum of lengths of all substring is 1 + 1 + 1 + 2 + 2 + 3 = 10
```

Example 2:

```text
Input: "ABA"
Output: 8
Explanation: The same as example 1, except uni("ABA") = 1.
```

Note: 0 <= S.length <= 10000.

## 解题思路

见程序注释
