# [44. Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)

## 题目

Given an input string (`s`) and a pattern (`p`), implement wildcard pattern matching with support for `'?'` and `'*'`.

```text
'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
```

The matching should cover the entire input string (not partial).

Note:

- `s` could be empty and contains only lowercase letters `a-z`.
- `p` could be empty and contains only lowercase letters `a-z`, and characters like `?` or `*`.

Example 1:

```text
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

Example 2:

```text
Input:
s = "aa"
p = "*"
Output: true
Explanation:`'*' matches any sequence.
```

Example 3:

```text
Input:
s = "cb"
p = "?a"
Output: false
Explanation:`'?' matches 'c', but the second letter is 'a', which does not match 'b'.
```

Example 4:

```text
Input:
s = "adceb"
p = "*a*b"
Output: true
Explanation:`The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
```

Example 5:

```text
Input:
s = "acdcb"
p = "a*c?b"
Output: false
```

## 解题思路

注意审题：

1. '?' 可以匹配任意一个字符，但是不能匹配空字符""
1. '*' 可以任意多个字符，包括""

见程序注释
