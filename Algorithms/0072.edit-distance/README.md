# [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

## 题目
Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)

You have the following 3 operations permitted on a word:
```
a) Insert a character
b) Delete a character
c) Replace a character
```
## 解题思路
变换单词有三种操作
1. 插入
1. 删除
1. 替换
可以看到，`替换` = `删除` + `插入`，所以，`替换`的效率更高，应该优先使用。但是，有些时候`替换`的效率会更低，例如
1. 从“gd”变换到“bad”，在“gd”的g后插入a，可以减少一次`替换`
1. 从“gaod”变换到“bad”，`删除`o，可以减少一次`替换`

所以，程序的执行步骤应该是
1. 通过`插入`或者`删除`，让两个word一样长。
1. 统计需要替换的次数。

见程序注释
