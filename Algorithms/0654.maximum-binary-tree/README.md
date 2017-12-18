# [654. Maximum Binary Tree](https://leetcode.com/problems/maximum-binary-tree/)

## 题目

Given an integer array with no duplicates. A maximum tree building on this array is defined as follow:

1. The root is the maximum number in the array.
1. The left subtree is the maximum tree constructed from left part subarray divided by the maximum number.
1. The right subtree is the maximum tree constructed from right part subarray divided by the maximum number.

Construct the maximum tree by the given array and output the root node of this tree.

Example 1:

```text
Input: [3,2,1,6,0,5]
Output: return the tree root node representing the following tree:

      6
    /   \
   3     5
    \    /
     2  0
       \
        1
```

Note:

1. The size of the given array will be in the range [1,1000].

## 解题思路

见程序注释
