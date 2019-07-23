# [1104. Path In Zigzag Labelled Binary Tree](https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/)

In an infinite binary tree where every node has two children, the nodes are labelled in row order.

In the odd numbered rows (ie., the first, third, fifth,...), the labelling is left to right, while in the even numbered rows (second, fourth, sixth,...), the labelling is right to left.

![tree](tree.png)

Given the label of a node in this tree, return the labels in the path from the root of the tree to the node with that label.

Example 1:

```text
Input: label = 14
Output: [1,3,4,14]
```

Example 2:

```text
Input: label = 26
Output: [1,2,6,10,26]
```

Constraints:

1. `1 <= label <= 10^6`
