# [875. Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)

## 题目

Koko loves to eat bananas. There are Npiles of bananas, the i-thpile has piles[i] bananas. The guards have gone and will come back in H hours.

Koko can decide her bananas-per-hour eating speed of K. Each hour, she chooses some pile of bananas, and eats K bananas from that pile. If the pile has less than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.

Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back.

Return the minimum integer K such that she can eat all the bananas within H hours.

Example 1:

```text
Input: piles = [3,6,7,11], H = 8
Output: 4
```

Example 2:

```text
Input: piles = [30,11,23,4,20], H = 5
Output: 30
```

Example 3:

```text
Input: piles = [30,11,23,4,20], H = 6
Output: 23
```

Note:

1. 1 <= piles.length <= 10^4
1. piles.length <= H <= 10^9
1. 1 <= piles[i] <= 10^9

## 解题思路

见程序注释
