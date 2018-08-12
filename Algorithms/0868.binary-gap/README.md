# [868. Binary Gap](https://leetcode.com/problems/binary-gap/)

## 题目

Given a positiveinteger N, find and return the longest distance between two consecutive 1's in the binary representation of N.

If there aren't two consecutive 1's, return 0.

Example 1:

```text
Input: 22
Output: 2
Explanation:
22 in binary is 0b10110.
In the binary representation of 22, there are three ones, and two consecutive pairs of 1's.
The first consecutive pair of 1's have distance 2.
The second consecutive pair of 1's have distance 1.
The answer is the largest of these two distances, which is 2.
```

Example 2:

```text
Input: 5
Output: 2
Explanation:
5 in binary is 0b101.
```

Example 3:

```text
Input: 6
Output: 1
Explanation:
6 in binary is 0b110.
```

Example 4:

```text
Input: 8
Output: 0
Explanation:
8 in binary is 0b1000.
There aren't any consecutive pairs of 1's in the binary representation of 8, so we return 0.
```

Note:

- 1 <= N <= 10^9

## 解题思路

见程序注释
