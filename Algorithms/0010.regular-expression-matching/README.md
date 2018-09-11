# [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)

## 题目

<<<<<<< HEAD
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.
=======
Implement regular expression matching with support for '.' and '*'.
>>>>>>> 0c654ed44b072a904e502ab100e6984a5b0e55e1

```text
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```

The matching should cover the entire input string (not partial).

Note:

- s could be empty and contains only lowercase letters a-z.
- p could be empty and contains only lowercase letters a-z, and characters like . or *.

Example 1:

```text
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

Example 2:

<<<<<<< HEAD
```text
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

Example 3:
=======
## 总结
>>>>>>> 0c654ed44b072a904e502ab100e6984a5b0e55e1

```text
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

Example 4:

```text
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
```

Example 5:

```text
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```
