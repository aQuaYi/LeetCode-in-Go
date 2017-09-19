 以下题目，暂时不能使用 Golang 解答
- [116. Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)
- [117. Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/)


## helper
[helper](./helper.v4)会帮助处理大部分琐碎的工作。

## notes
[notes](./notes)记录了我答题过程中，对知识点的总结。

## kit
针对 LeetCode 中经常出现的以下数据结构，在[kit](./kit)中进行了定义，并添加了与 []int 相互转换的函数。利用 Golang 1.9 新添加的 [type alias](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md)功能，易于添加单元测试。 
1. [*ListNode](./kit/ListNode.go)
1. [*TreeNode](./kit/TreeNode.go)
