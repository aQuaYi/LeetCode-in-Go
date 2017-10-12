# helper

[helper](./helper.v5) 会帮助处理大部分琐碎的工作。

## notes

[notes](./notes) 记录了我答题过程中，对知识点的总结。

## kit

针对 LeetCode 中经常出现的以下数据结构，在 [kit](./kit) 中进行了定义，并添加了与 []int 相互转换的函数。利用 Go 1.9 新添加的 [type alias](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md) 功能，易于添加单元测试。

1. [ListNode](./kit/ListNode.go)
1. [TreeNode](./kit/TreeNode.go)
1. [Queue](./kit/Queue.go)
1. [Stack](./kit/Stack.go)
1. [Heap](./kit/Heap.go)