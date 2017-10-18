# [319. Bulb Switcher](https://leetcode.com/problems/bulb-switcher/)

## 题目

There are n bulbs that are initially off. You first turn on all the bulbs. Then, you turn off every second bulb. On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on). For the ith round, you toggle every i bulb. For the nth round, you only toggle the last bulb.

Find how many bulbs are on after n rounds.

Example:

```text
Given n = 3.

At first, the three bulbs are [off, off, off].
After first round, the three bulbs are [on, on, on].
After second round, the three bulbs are [on, off, on].
After third round, the three bulbs are [on, off, off].
So you should return 1, because there is only one bulb is on.
```

## 解题思路

大概意思是，要在[1,n]里头找哪些灯泡被执行了奇数次操作

1. 对于一个非平方数来说（比如12），因为有成对的补数（1vs12; 2vs6），所以总是会按下2的倍数次
1. 但是对于一个平方数来说（比如36），因为其中有个数(6)它的补数是自己，所以会按被下奇数次

然后这道题就变成了找[1,n]中间有几个平方数了
