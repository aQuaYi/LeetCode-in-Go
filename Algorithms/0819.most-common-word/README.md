# [819. Most Common Word](https://leetcode.com/problems/most-common-word/)

## 题目

Given a paragraphand a list of banned words, return the most frequent word that is not in the list of banned words. It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation. Words in the paragraph are not case sensitive. The answer is in lowercase.

```text
Example:
Input:
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.
```

Note:

1. 1 <= paragraph.length <= 1000.
1. 1 <= banned.length <= 100.
1. 1 <= banned[i].length <= 10.
1. The answer is unique, and written in lowercase (even if its occurrences in paragraphmay haveuppercase symbols, and even if it is a proper noun.)
1. paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
1. Different words inparagraphare always separated by a space.
1. There are no hyphens or hyphenated words.
1. Words only consist of letters, never apostrophes or other punctuation symbols.

## 解题思路

见程序注释
