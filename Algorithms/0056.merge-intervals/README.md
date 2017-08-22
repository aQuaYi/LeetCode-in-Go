# [56. Merge Intervals](https://leetcode.com/problems/merge-intervals/)

## 题目
Given a collection of intervals, merge all overlapping intervals.

```
For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
```
## 解题思路
1. 先对 intervals 进行排序，按照 Start 递增
1. 依次处理重叠的情况。

## 附注
最快的答案和我的思路一样，只是使用的是 merge sort ，自我感觉我的答案更清晰，就没有改成他的。