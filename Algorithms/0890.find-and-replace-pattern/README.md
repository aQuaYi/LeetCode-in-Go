# [890. Find and Replace Pattern](https://leetcode.com/problems/find-and-replace-pattern/)

## 题目

You have a list ofwords and a pattern, and you want to know which words in words matches the pattern.

A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

(Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)

Return a list of the words in wordsthat match the given pattern.

You may return the answer in any order.

Example 1:

```text
Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
Output: ["mee","aqq"]
Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
since a and b map to the same letter.
```

Note:

1. 1 <= words.length <= 50
1. 1 <= pattern.length = words[i].length<= 20

## 解题思路

见程序注释
