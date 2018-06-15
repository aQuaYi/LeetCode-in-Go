# [844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)

## 题目

Given twostringsSand T,return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

```text
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
```

Example 2:

```text
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
```

Example 3:

```text
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
```

Example 4:

```text
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
```

Note:

1. 1 <= S.length <= 200
1. 1 <= T.length <= 200
1. Sand T only containlowercase letters and '#' characters.

Follow up:

1. Can you solve it in O(N) time and O(1) space?

## 解题思路

见程序注释
