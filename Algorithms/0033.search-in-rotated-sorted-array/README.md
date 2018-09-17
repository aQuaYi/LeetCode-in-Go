# [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## 题目

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

## 解题思路

先假设old = [0, 1, 2, 4, 5, 6, 7]，利用二分查找法，很容易可以`5`的索引号，当 old 变换成了 new = [4, 5, 6, 7, 0, 1, 2]以后，同样可以使用二分查找法，因为 old 和 new 中的元素有明确的对应关系

old[i] == new[j]，只要i和j满足关系式

```go
j=i+4
if j > len(old) {
    j -= len(old)
}
```

其中，4 = old 中的最大值在new中的索引号 + 1

所以，如果我们手中只有new，我们可以假装自己还是在对old使用二分查找法，当需要获取old[i]的值进行比较判断的时候，使用new[j]的值替代即可。

## 总结

本题是二分查找法的升级版