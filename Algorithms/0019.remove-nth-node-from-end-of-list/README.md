# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 题目
Given a linked list, remove the nth node from the end of list and return its head.

For example,
```
Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end, the linked list becomes 1->2->3->5.
```
**Note**:
1. Given n will always be valid. 不存在链的长度<n的情况
1. Try to do this in one pass.

## 解题思路
1. 获取倒数第n个节点的父节点d
1. 再根据父节点d是否为head节点，选取不同的删除方式。
