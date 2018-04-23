# [814. Binary Tree Pruning](https://leetcode.com/problems/binary-tree-pruning/)

## 题目

We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1.

Return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

(Recall that the subtree of a node X is X, plus every node that is a descendant of X.)

```text
Example 1:
Input: [1,null,0,0,1]
Output: [1,null,0,null,1]
Explanation:
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.
```

![1](1.png)

```text
Example 2:
Input: [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]
```

![2](2.png)

```text
Example 3:
Input: [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]
```

![3](3.png)

Note:

1. The binary tree will have at most 100 nodes.
1. The value of each node will only be 0 or 1.

## 解题思路

见程序注释
