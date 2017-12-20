# [672. Bulb Switcher II](https://leetcode.com/problems/bulb-switcher-ii/)

## 题目

There is a room with n lights which are turned on initially and 4 buttons on the wall. After performing exactly m unknown operations towards buttons, you need to return how many different kinds of status of the n lights could be.

Suppose n lights are labeled as number [1, 2, 3 ..., n], function of these 4 buttons are given below:

1. Flip all the lights.
1. Flip lights with even numbers.
1. Flip lights with odd numbers.
1. Flip lights with (3k + 1) numbers, k = 0, 1, 2, ...

Example 1:

```text
Input: n = 1, m = 1.
Output: 2
Explanation: Status can be: [on], [off]
```

Example 2:

```text
Input: n = 2, m = 1.
Output: 3
Explanation: Status can be: [on, off], [off, on], [off, off]
```

Example 3:

```text
Input: n = 3, m = 1.
Output: 4
Explanation: Status can be: [off, on, off], [on, off, on], [off, off, off], [off, on, on].
```

Note: n and m both fit in range [0, 1000].

## 解题思路

见程序注释
