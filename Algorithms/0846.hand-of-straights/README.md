# [846. Hand of Straights](https://leetcode.com/problems/hand-of-straights/)

## 题目

Alice has a hand of cards, given as an array of integers.

Now she wants to rearrange the cards into groups so that each group is size W, and consists of W consecutive cards.

Return true if and only if she can.

Example 1:

```text
Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
```

Example 2:

```text
Input: hand = [1,2,3,4,5], W = 4
Output: false
Explanation: Alice's hand can't be rearranged into groups of 4.
```

Note:

1. 1 <= hand.length <= 10000
1. 0 <= hand[i]<= 10^9
1. 1 <= W <= hand.length

## 解题思路

见程序注释
