# [201. Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/)

## 题目
Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

For example, given the range [5, 7], you should return 4.

Credits:Special thanks to @amrsaqr for adding this problem and creating all test cases.

## 解题思路

在草稿纸上，演算以下两组参数，就可以明白程序的用意了。
1.
```
m = 5, n = 7
5 : 0101
6 : 0110
7 : 0111
结果
4 : 0100
```
2.
```
m = 6, n = 8
6 : 0110
7 : 0111
8 : 1000
结果
0 : 0000
```