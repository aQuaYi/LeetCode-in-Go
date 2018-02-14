# [745. Prefix and Suffix Search](https://leetcode.com/problems/prefix-and-suffix-search/)

## 题目

Given many words, words[i] has weight i.

Design a class WordFilter that supports one function, WordFilter.f(String prefix, String suffix). It will return the word with given prefix and suffix with maximum weight. If no word exists, return -1.

Examples:

```text
Input:
WordFilter(["apple"])
WordFilter.f("a", "e") // returns 0
WordFilter.f("b", "") // returns -1
```

Note:

1. words has length in range [1, 15000].
1. For each test case, up to words.length queries WordFilter.f may be made.
1. words[i] has length in range [1, 10].
1. prefix, suffix have lengths in range [0, 10].
1. words[i] and prefix, suffix queries consist of lowercase letters only.

## 解题思路

见程序注释
