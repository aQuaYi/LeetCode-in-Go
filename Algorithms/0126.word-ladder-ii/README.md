# [126. Word Ladder II](https://leetcode.com/problems/word-ladder-ii/)

## 题目
Given two words (beginWord and endWord), and a dictionary's word list, find **all shortest** transformation sequence(s) from beginWord to endWord, such that:
1. Only one letter can be changed at a time
1. Each transformed word must exist in the word list. Note that beginWord is not a transformed word.

For example,
```
Given:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log","cog"]

Return
  [
    ["hit","hot","dot","dog","cog"],
    ["hit","hot","lot","log","cog"]
  ]
```

Note:
1. Return an empty list if there is no such transformation sequence.
1. All words have the same length.
1. All words contain only lowercase alphabetic characters.
1. You may assume no duplicates in the word list.
1. You may assume beginWord and endWord are non-empty and are not the same.

UPDATE (2017/1/20):
The wordList parameter had been changed to a list of strings (instead of a set of strings). Please reload the code definition to get the latest changes.

## 解题思路

见程序注释

这一次，我不想感谢服务器。
![100%](126.png)
