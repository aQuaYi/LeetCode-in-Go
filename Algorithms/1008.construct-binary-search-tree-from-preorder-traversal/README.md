# [1008. Construct Binary Search Tree from Preorder Traversal](https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/)

Return the root node of a binary search tree that matches the given preorder traversal.

(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)

Example 1:

![1266](1266.png)

```text
Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
```

Note:

- 1 <= preorder.length <= 100
- The values of preorder are distinct.
