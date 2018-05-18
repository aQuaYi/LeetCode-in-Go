# [822. Card Flipping Game](https://leetcode.com/problems/card-flipping-game/)

## 题目

On a table are N cards, with a positive integer printed on the front and back of each card (possibly different).

We flip any number of cards, and after we choose onecard.

If the number X on the back of the chosencard is not on the front of any card, then this number X is good.

What is the smallest number that is good? If no number is good, output 0.

Here, fronts[i] and backs[i] represent the number on the front and back of card i.

Aflip swaps the front and back numbers, so the value on the front is now on the back and vice versa.

Example:

Input: fronts = [1,2,4,4,7], backs = [1,3,4,1,3]
Output: 2
Explanation: If we flip the second card, the fronts are [1,3,4,4,7] and the backs are [1,2,4,1,3].
We choose the second card, which has number 2 on the back, and it isn't on the front of any card, so 2 is good.

Note:

1. 1 <= fronts.length == backs.length<=1000.
1. 1 <=fronts[i]<= 2000.
1. 1 <= backs[i]<= 2000.

## 解题思路

见程序注释
