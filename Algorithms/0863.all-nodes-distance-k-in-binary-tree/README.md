# [863. All Nodes Distance K in Binary Tree](https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/)

## 题目

We are given a binary tree (with root noderoot), a target node, and an integer value K.

Return a list of the values of allnodes that have a distance K from the target node. The answer can be returned in any order.

Example 1:

```text
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

Output: [7,4,1]

Explanation:
The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.

Note that the inputs "root" and "target" are actually TreeNodes.
The descriptions of the inputs above are just serializations of these objects.
```

![pic](pic.png)

Note:

1. The given tree is non-empty.
1. Each node in the tree has unique values0 <= node.val <= 500.
1. The targetnode is a node in the tree.
1. 0 <= K <= 1000.

## 解题思路

见程序注释
