# [LeetCode](https://leetcode.com) 的 Go 解答 {{- /* 本文件是用来生成 README.md 的模板 */}}

[![LeetCode 排名](https://img.shields.io/badge/{{.Username}}-{{.Ranking}}-blue.svg)](https://leetcode.com/{{.Username}}/)
[![codecov](https://codecov.io/gh/aQuaYi/LeetCode-in-Go/branch/master/graph/badge.svg)](https://codecov.io/gh/aQuaYi/LeetCode-in-Go)
[![Build Status](https://www.travis-ci.org/aQuaYi/LeetCode-in-Go.svg?branch=master)](https://www.travis-ci.org/aQuaYi/LeetCode-in-Go)
 [![Go](https://img.shields.io/badge/Go-1.13-blue.svg)](https://golang.google.cn)

## 进度

> 统计规则：1.免费题，2.算法题，3.能提交 Go 解答

{{.ProgressTable}}

## 题解

{{.AvailableTable}}
以下免费的算法题，暂时不能提交 Go 解答

{{.UnavailableList}}

## helper

[helper](./Helper) 会处理大部分琐碎的工作。

## notes

[notes](./notes) 记录了我答题过程中，对知识点的总结。

## kit

针对 LeetCode 中经常出现的以下数据结构，在 [kit](./kit) 中进行了定义，并添加了与 []int 相互转换的函数。利用 Go 1.9 添加的 [type alias](https://github.com/golang/proposal/blob/master/design/18130-type-alias.md) 功能，易于添加单元测试。

- [Heap](./kit/Heap.go)
- [Interval](./kit/Interval.go)
- [ListNode](./kit/ListNode.go)
- [NestedInteger](./kit/NestedInteger.go)
- [PriorityQueue](./kit/PriorityQueue.go)
- [Queue](./kit/Queue.go)
- [Stack](./kit/Stack.go)
- [TreeNode](./kit/TreeNode.go)
- [Master](./kit/master.go)

## 致谢

感谢所有贡献者的辛苦付出

<a href="https://github.com/aQuaYi/LeetCode-in-Go/graphs/contributors"><img src="https://opencollective.com/leetcode-in-go/contributors.svg?width=890&button=false"/></a>

感谢 JetBrains

[![GoLand](GoLand.png)]( https://www.jetbrains.com/?from=LeetCode-in-Go)
