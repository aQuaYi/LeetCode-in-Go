# [230. Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)

## 题目
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note: 
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Follow up:
What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

## 解题思路
以下是[BST](https://zh.wikipedia.org/zh-cn/%E4%BA%8C%E5%85%83%E6%90%9C%E5%B0%8B%E6%A8%B9)的性质：
1. 若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
1. 若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
1. 任意节点的左、右子树也分别为二叉查找树；
1. 没有键值相等的节点。

见程序注释
