# [1042. Flower Planting With No Adjacent](https://leetcode.com/problems/flower-planting-with-no-adjacent/)

You have N gardens, labelled 1 to N.  In each garden, you want to plant one of 4 types of flowers.

paths[i] = [x, y] describes the existence of a bidirectional path from garden x to garden y.

Also, there is no garden that has more than 3 paths coming into or leaving it.

Your task is to choose a flower type for each garden such that, for any two gardens connected by a path, they have different types of flowers.

Return any such a choice as an array answer, where answer[i] is the type of flower planted in the (i+1)-th garden.  The flower types are denoted 1, 2, 3, or 4.  It is guaranteed an answer exists.

Example 1:

```text
Input: N = 3, paths = [[1,2],[2,3],[3,1]]
Output: [1,2,3]
```

Example 2:

```text
Input: N = 4, paths = [[1,2],[3,4]]
Output: [1,2,1,2]
```

Example 3:

```text
Input: N = 4, paths = [[1,2],[2,3],[3,4],[4,1],[1,3],[2,4]]
Output: [1,2,3,4]
```

Note:

1. `1 <= N <= 10000`
1. `0 <= paths.size <= 20000`
1. `No garden has 4 or more paths coming into or leaving it.`
1. `It is guaranteed an answer exists.`
