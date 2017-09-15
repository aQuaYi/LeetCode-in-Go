### Leetcode 暂时不提供以下题目的 Golang 解答方式
1. [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)
1. [117. Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/)


## helper
[helper](./helper.v4)会帮助处理大部分琐碎的工作。

## notes
[notes](./notes)记录了我答题过程中，对知识点的总结。

## kit
在[kit](./kit)中添加了LeetCode常用的数据结构和处理函数：
1. 为[*ListNode](./kit/ListNode.go)添加了与[]int相互转换的函数，方便添加单元测试。使用方式可以参考[21. Merge Two Sorted Lists](./Algorithms/0021.merge-two-sorted-lists/merge-two-sorted-lists_test.go)
1. 为[*TreeNode](./kit/TreeNode.go)添加了与[]int相互转换的函数，方便添加单元测试。