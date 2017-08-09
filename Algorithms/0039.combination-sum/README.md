# [39. Combination Sum](https://leetcode.com/problems/combination-sum/)

## 题目
Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7, 
A solution set is:
``` 
[
  [7],
  [2, 2, 3]
]
```
## 解题思路
本题就是SICP中1.2.2节换零钱的另一种说法。请参考我当时的笔记。
## 关于换硬币例子的思考
有一美元想要换成硬币，硬币有5种，分别是1、 5、 10、 25、 50美分。问一共有多少种不同的换法？
现在分析如下，总数应该等于
1. 100-10=90美分，用5种硬币的换法数，加上
1. 100美分，用除了10美分外的4种硬币的换法数

以上划分方法，是对换钱方式的一个**划分**，划分的标准是：换出的零钱中，是否包含10美分。划分后的子集
1. 减少了钱数
1. 减少了硬币的种数

要求就是说，子集简化了全集。同时，子集还可以按照全集的划分方法，继续划分下去
1. 100-10=90美分，用5种硬币的换法数，等于
    1. 90-10=80美分，用5种硬币的换法数，加上
    1. 90美分，用除了10美分外的4种硬币的换法数
1. 100美分，用除了10美分外的4种硬币的换法数，等于
    1. 100-5=95美分，用除了10美分外的4种硬币的换法数，加上
    1. 100美分，用除了10美分和5美分外的3种硬币的换法数

但是，子集的划分不可能无限继续下去。很快就会遇到以下终止条件。
1. 0美分需要换零钱，显然，这个划分路径已经找到了一种可行的换法，应该返回1
1. <0美分需要换零钱，显然，这个划分路径前面减多了，这不是一种可行的换法，应该返回0
1. 要剩下的0种硬币来换零钱，显然，已经没有可供选择的硬币了，这也不是一种可行的换法，应该返回0

**注意**，以上只是用10美分和5美分做例子说明，先选哪一种硬币，不影响结果。

### 关于递归的终结
1. 递归调用必须有终止条件。
1. 递归调用总是去尝试解决规模更小的子问题，这样递归才能收敛于终止条件。
1. 递归调用的子集必须是对父集的一个划分，不然，就会漏掉或者出现重复的结果。

## 总结
