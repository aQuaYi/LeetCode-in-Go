# [882. Reachable Nodes In Subdivided Graph](https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/)

## 题目

Starting with anundirected graph (the "original graph") with nodes from 0 to N-1, subdivisions are made to some of the edges.

The graph is given as follows: edges[k] is a list of integer pairs (i, j, n) such that (i, j) is an edge of the original graph,

and n is the total number of new nodes on that edge.

Then, the edge (i, j) is deleted from the original graph,nnew nodes (x_1, x_2, ..., x_n) are added to the original graph,

and n+1 newedges (i, x_1), (x_1, x_2), (x_2, x_3), ..., (x_{n-1}, x_n), (x_n, j)are added to the originalgraph.

Now, you start at node 0from the original graph, and in each move, you travel along oneedge.

Return how many nodes you can reach in at most M moves.

Example 1:

```text
Input: edges = [[0,1,10],[0,2,1],[1,2,2]], M = 6, N = 3
Output: 13
Explanation:
The nodes that are reachable in the final graph after M = 6 moves are indicated below.
```

![original-to-final](origfinal.png)

Example 2:

```text
Input: edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], M = 10, N = 4
Output: 23
```

Note:

1. 0 <= edges.length <= 10000
1. 0 <= edges[i][0] <edges[i][1] < N
1. There does not exist anyi != j for which edges[i][0] == edges[j][0] and edges[i][1] == edges[j][1].
1. The original graphhas no parallel edges.
1. 0 <= edges[i][2] <= 10000
1. 0 <= M <= 10^9
1. 1 <= N <= 3000
1. A reachable node is a node that can be travelled tousing at mostM moves starting fromnode 0.

## 解题思路

见程序注释
