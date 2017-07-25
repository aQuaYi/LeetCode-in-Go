# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## 题目
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

就是说，x轴上在1,2,...,n点上有许多垂直的线段，长度依次是a1, a2, ..., an。找出两条线段，使他们和x抽围成的面积最大。面积公式是 Min(ai, aj) X |j - i|

## 解题思路
穷举法是O(n^2)的复杂度，会触发leetcode的时间限制。

 O(n)的复杂度的解法是，保持两个指针i,j；分别指向长度数组的首尾。如果ai 小于aj，则移动i向后（i++）。反之，移动j向前（j--）。如果当前的area大于了所记录的area，替换之。这个想法的基础是，如果i的长度小于j，无论如何移动j，短板在i，不可能找到比当前记录的area更大的值了，只能通过移动i来找到新的可能的更大面积。
