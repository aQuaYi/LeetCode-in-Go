# [817. Linked List Components](https://leetcode.com/problems/linked-list-components/)

## 题目

We are given head, the head node of a linked list containing unique integer values.

We are also given the list G, a subset of the values in the linked list.

Return the number of connected components in G, where two values are connected if they appear consecutively in the linked list.

Example 1:

```text
Input:
head: 0->1->2->3
G = [0, 1, 3]
Output: 2
Explanation:
0 and 1 are connected, so [0, 1] and [3] are the two connected components.
```

Example 2:

```text
Input:
head: 0->1->2->3->4
G = [0, 3, 1, 4]
Output: 2
Explanation:
0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
```

Note:

1. If N is the length of the linked list given by head, 1 <= N <= 10000.
1. The value of each node in the linked list will be in the range [0, N - 1].
1. 1 <= G.length <= 10000.
1. G is a subset of all values in the linked list.

## 解题思路

见程序注释
