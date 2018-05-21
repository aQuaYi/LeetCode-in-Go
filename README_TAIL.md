# 使用方法

```bash
$ go version
// 请确保版本号 >= 1.9
$ go env
...
GOPATH="第一个GOPATH目录:..."
...
$ go get github.com/aQuaYi/LeetCode-in-Go
// 下载位置： 第一个GOPATH目录/src/github.com/aQuaYi/LeetCode-in-Go
```

## helper

[helper](./helper) 会处理大部分琐碎的工作。

## notes

[notes](./notes) 记录了我答题过程中，对知识点的总结。

## kit

针对 LeetCode 中经常出现的以下数据结构，在 [kit](./kit) 中进行了定义，并添加了与 []int 相互转换的函数。利用 Go 1.9 新添加的 [type alias](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md) 功能，易于添加单元测试。

- [Heap](./kit/Heap.go)
- [Interval](./kit/Interval.go)
- [ListNode](./kit/ListNode.go)
- [NestedInteger](./kit/NestedInteger.go)
- [PriorityQueue](./kit/PriorityQueue.go)
- [Queue](./kit/Queue.go)
- [Stack](./kit/Stack.go)
- [TreeNode](./kit/TreeNode.go)