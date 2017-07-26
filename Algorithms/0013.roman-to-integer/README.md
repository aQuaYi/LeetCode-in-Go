# [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

## 题目
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.

## 解题思路
这一题是[12. Integer to Roman](./Algorithms/0012.integer-to-roman)的一个逆转换。~~一样的解题思路~~

此题，最关键的信息是
> 右加左减，左减数字必须为一位，比如8写成VIII，而非IIX。

解题思路
1. 从右往左处理字符串。
1. 当前字符代表的数字，小于右边字符的时候，总体减去当前字符代表的数字。
1. 否则，总体加上当前字符代表的数字。
## 总结
抓住关键信息，避免思维定式。