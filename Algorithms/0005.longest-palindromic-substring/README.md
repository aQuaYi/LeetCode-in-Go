# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## 题目
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example:
```
Input: "cbbd"
Output: "bb"
```
## 解题思路
题目要求寻找字符串中的最长回文。
当然，我们可以使用下面的程序判断字符串s[i:j+1]是否是回文。
```go
func isPalindromic(s *string, i, j int ) bool {
    for  i< j {
        if (*s)[i] != (*s)[j] {
            return false
        } 
        i++
        j--
    }
    return true
}
```
但是，这样就没有利用回文的一下特点，假定回文的长度为l，x为任意字符
1. 当l为奇数时，回文的`正中间段`只会是，“x”，或“xxx”，或“xxxxx”，或...
1. 当l为偶数时，回文的`正中间段`只会是，“xx”，或“xxxx”，或“xxxxxx”，或...
1. 同一个字符串的任意两个回文substring的`正中间段`，不会重叠。

## 总结
充分利用查找对象的特点，可以加快查找速度。