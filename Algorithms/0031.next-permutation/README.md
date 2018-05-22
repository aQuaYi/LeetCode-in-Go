# [31. Next Permutation](https://leetcode.com/problems/next-permutation/)

## 题目

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

```text
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

## 解题思路

1　　2　　7　　4　　3　　1

下一个排列为：

1　　3　　1　　2　　4　　7

那么是如何得到的呢，步骤如下：

从后往前，找到最长的降序排列

1　　2　　`7　　4　　3　　1`

把这个降序排列，转换成升序排列

1　　2　　`1　　3　　4　　7`

把序列前的元素，与序列中，第一个大于他的元素互换。

1　　`3`　　1　　`2`　　4　　7
