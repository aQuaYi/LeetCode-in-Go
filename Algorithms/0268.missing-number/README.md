# [268. Missing Number](https://leetcode.com/problems/missing-number/)

## 题目

        
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.


For example,
Given nums = [0, 1, 3] return 2.



Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?


## 解题思路
`^`：按位异或运算符，是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。

特别地，`^`支持交换律
```
1^1^2^2^3^3 == 1^2^2^3^3^1 == 1^3^2^3^1^2
```

