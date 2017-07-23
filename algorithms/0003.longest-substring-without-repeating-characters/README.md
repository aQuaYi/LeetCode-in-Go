# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

## 解题思路
利用s[left:i+1]来表示s[:i+1]中的包含s[i]的最长子字符串。
location[s[i]]是字符s[i]在s[:i+1]中倒数第二次出现的序列号。
当left < location[s[i]]的时候，说明字符s[i]出现了两次。需要设置 left = location[s[i]] + 1,
保证字符s[i]只出现一次。

## 总结
利用Location保存字符上次出现的序列号，避免了查询工作。location和[Two Sum](./algorithms/0001.two-sum)中的m是一样的作用。
```go
// m 负责保存map[整数]整数的序列号
	m := make(map[int]int, len(nums))
```
