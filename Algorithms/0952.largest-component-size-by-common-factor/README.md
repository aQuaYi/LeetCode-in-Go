# [952. Largest Component Size by Common Factor](https://leetcode.com/problems/largest-component-size-by-common-factor/)

Given a non-empty array of unique positive integers `A`, consider the following graph:

- There are `A.length` nodes, labelled `A[0]` to `A[A.length - 1]`;
- There is an edge between `A[i]` and `A[j]` if and only if `A[i]` and `A[j]` share a common factor greater than 1.

Return the size of the largest connected component in the graph.

Example 1:

```text
Input: [4,6,15,35]
Output: 4
```

![example1](1.png)

Example 2:

```text
Input: [20,50,9,63]
Output: 2
```

![example2](2.png)

Example 3:

```text
Input: [2,3,6,7,4,12,21,39]
Output: 8
```

![example3](3.png)

Note:

1. `1 <= A.length <= 20000`
1. `1 <= A[i] <= 100000`
