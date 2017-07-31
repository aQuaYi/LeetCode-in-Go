# [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

## 题目
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

## 解题思路
最先想到的方案是，依次合并res=merge(res, lists[i])，这个方案的坏处是需要合并两个长度相差很大的链。
LeetCode上最快的解决方法是，借鉴了[归并排序](https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F)的思想，让lists中临近的list先两两合并，再让新生成的lists中的临近list两两合并，直到合并完成。

## 总结
