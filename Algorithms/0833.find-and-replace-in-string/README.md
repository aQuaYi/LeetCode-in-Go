# [833. Find And Replace in String](https://leetcode.com/problems/find-and-replace-in-string/)

## 题目

To some string S, we will perform somereplacementoperations that replace groups of letters with new ones (not necessarily the same size).

Each replacement operation has 3 parameters: a starting index i, a source wordxand a target wordy. The rule is that if xstarts at position iin the original string S, then we will replace that occurrence ofxwithy. If not, we do nothing.

For example, if we haveS = "abcd"and we have some replacement operationi = 2, x = "cd", y = "ffff", then because"cd"starts at position 2in the original string S, we will replace it with "ffff".

Using another example on S = "abcd", if we have both the replacement operation i = 0, x = "ab", y = "eee", as well as another replacement operationi = 2, x = "ec", y = "ffff", this second operation does nothing because in the original stringS[2] = 'c', which doesn't matchx[0] = 'e'.

All these operations occur simultaneously. It's guaranteed that there won't be any overlap in replacement: for example,S = "abc", indexes = [0, 1],sources = ["ab","bc"] is not a valid test case.

Example 1:

```text
Input: S = "abcd", indexes = [0,2], sources = ["a","cd"], targets = ["eee","ffff"]
Output: "eeebffff"
Explanation: "a" starts at index 0 in S, so it's replaced by "eee".
"cd" starts at index 2 in S, so it's replaced by "ffff".
```

Example 2:

```text
Input: S = "abcd", indexes = [0,2], sources = ["ab","ec"], targets = ["eee","ffff"]
Output: "eeecd"
Explanation: "ab" starts at index 0 in S, so it's replaced by "eee".
"ec" doesn't starts at index 2 in the original S, so we do nothing.
```

Notes:

1. 0 <=indexes.length =sources.length =targets.length <= 100
1. 0<indexes[i]< S.length <= 1000
1. All characters in given inputs are lowercase letters.

## 解题思路

见程序注释
