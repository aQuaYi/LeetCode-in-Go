# [233. Number of Digit One](https://leetcode.com/problems/number-of-digit-one/)

## 题目
Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.

For example: 
```
Given n = 13,
Return 6, because digit 1 occurred in the following numbers: 1, 10, 11, 12, 13.
```

## 解题思路

以算百位上1为例子:   假设百位上是0, 1, 和 >=2 三种情况: 
1. n=41092, m=100, a=n/m=410, b=n%m=92. 计算百位上1的个数应该为 41 *100 次.
1. n=41192, m=100, a=n/m=411, b=n%m=92. 计算百位上1的个数应该为 41 *100 + (92+1) 次. 
1. n=41592, m=100, a=n/m=415, b=n%m=92. 计算百位上1的个数应该为 (41+1) *100 次. 
以上三种情况可以用 一个公式概括:
```
(a + 8) / 10 * m + (a % 10 == 1) * (b + 1);
```
见程序注释
