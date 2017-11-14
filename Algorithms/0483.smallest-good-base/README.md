# [483. Smallest Good Base](https://leetcode.com/problems/smallest-good-base/)

## 题目

For an integer n, we call k>=2 a good base of n, if all digits of n base k are 1.
Now given a string representing n, you should return the smallest good base of n in string format.

把 n 转换成 k 进制后，新的数全是 1

Example 1:

```text
Input: "13"
Output: "3"
Explanation: 13 base 3 is 111.
```

Example 2:

```text
Input: "4681"
Output: "8"
Explanation: 4681 base 8 is 11111.
```

Example 3:

```text
Input: "1000000000000000000"
Output: "999999999999999999"
Explanation: 1000000000000000000 base 999999999999999999 is 11.
```

Note:

1. The range of n is [3, 10^18].
1. The string representing n is always valid and will not have leading zeros.

## 解题思路

见程序注释
