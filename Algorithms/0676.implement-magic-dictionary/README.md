# [676. Implement Magic Dictionary](https://leetcode.com/problems/implement-magic-dictionary/)

## 题目

Implement a magic directory with buildDict, and search methods.

For the method buildDict, you'll be given a list of non-repetitive words to build a dictionary.

For the method search, you'll be given a word, and judge whether if you modify exactly one character into another character in this word, the modified word is in the dictionary you just built.

Example 1:

```text
Input: buildDict(["hello", "leetcode"]), Output: Null
Input: search("hello"), Output: False
Input: search("hhllo"), Output: True
Input: search("hell"), Output: False
Input: search("leetcoded"), Output: False
```

Note:

1. You may assume that all the inputs are consist of lowercase letters a-z.
1. For contest purpose, the test data is rather small by now. You could think about highly efficient algorithm after the contest.
1. Please remember to RESET your class variables declared in class MagicDictionary, as static/class variables are persisted across multiple test cases. Please see here for more details.

## 解题思路

见程序注释
