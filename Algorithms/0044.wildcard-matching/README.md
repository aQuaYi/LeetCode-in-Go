# [44. Wildcard Matching](https://leetcode.com/problems/wildcard-matching/)

## 题目
Implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).

The function prototype should be:
bool isMatch(const char *s, const char *p)

Some examples:
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "c*a*b") → false

## 解题思路
注意审题：
1. '?' 可以匹配任意一个字符，但是不能匹配空字符""
1. '*' 可以任意多个字符，包括""

见程序注释
