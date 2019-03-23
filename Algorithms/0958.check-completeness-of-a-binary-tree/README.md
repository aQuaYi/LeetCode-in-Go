# [958. Check Completeness of a Binary Tree](https://leetcode.com/problems/check-completeness-of-a-binary-tree/)

Given a binary tree, determine if it is a complete binary tree.

Definition of a complete binary tree from [Wikipedia](http://en.wikipedia.org/wiki/Binary_tree#Types_of_binary_trees):

> In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Example 1:

![ex1](ex1.png)

```text
Input: [1,2,3,4,5,6]
Output: true
Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
```

Example 2:

![ex2](ex2.png)

```text
Input: [1,2,3,4,5,null,7]
Output: false
Explanation: The node with value 7 isn't as far left as possible.
```

Note:

1. The tree will have between 1 and 100 nodes.