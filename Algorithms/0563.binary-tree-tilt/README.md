# [563. Binary Tree Tilt](https://leetcode.com/problems/binary-tree-tilt/)

## 题目

Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:

```text
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1
```

Note:

1. The sum of node values in any subtree won't exceed the range of 32-bit integer.
1. All the tilt values won't exceed the range of 32-bit integer.

## 解题思路

见程序注释
