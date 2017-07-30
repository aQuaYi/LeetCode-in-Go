# [21. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

## 题目
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

## 解题思路
把两个排序好的链合并，要求合并后依然是排序好的。结题步骤如下：
1. 先处理其中一条链为nil的情况，直接返回另一条链，这样可以简化后面的判断条件。
1. 设置好链接头head和用于移动节点指针node
1. 利用for循环反复比较，每次选取较小的节点，放在node.Next
1. 处理l1或l2中剩余的节点
## 总结
合理地安排步骤，可以有效地减轻后面的判断条件和处理步骤，让整个函数更清晰易懂。