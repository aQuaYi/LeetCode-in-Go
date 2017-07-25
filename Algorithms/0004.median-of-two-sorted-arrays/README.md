# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## 题目
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
```
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
```
Example 2:
```
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
```
## 解题思路
输入两个已经排序好的整数切片，合并两个切片，返回合并后切片的中位数。

所以，题目的关键是，如何高效地合并切片。

## 总结
分步骤完成程序

