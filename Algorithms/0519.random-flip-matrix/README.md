# [519. Random Flip Matrix](https://leetcode.com/problems/random-flip-matrix/)

## 题目

You are given the number of rows n_rowsand number of columns n_colsof a2Dbinary matrixwhere all values are initially 0.Write a function flipwhich choosesa 0 valueuniformly at random,changes it to 1,and then returns the position [row.id, col.id] of that value. Also, write a function reset which sets all values back to 0.Try to minimize the number of calls to system's Math.random() and optimize the time andspace complexity.

Note:

1. 1 <= n_rows, n_cols<= 10000
1. 0 <= row.id < n_rows and 0 <= col.id < n_cols
1. flipwill not be called when the matrix has no0 values left.
1. the total number of calls toflipand resetwill not exceed1000.

Example 1:

```text
Input:
["Solution","flip","flip","flip","flip"]
[[2,3],[],[],[],[]]
Output: [null,[0,1],[1,2],[1,0],[1,1]]
```

Example 2:

```text
Input:
["Solution","flip","flip","reset","flip"]
[[1,2],[],[],[],[]]
Output: [null,[0,0],[0,1],null,[0,0]]
```

Explanation of Input Syntax:

The input is two lists:the subroutines calledand theirarguments. Solution's constructorhas two arguments, n_rows and n_cols.flipand reset havenoarguments.Argumentsarealways wrapped with a list, even if there aren't any.

## 解题思路

见程序注释
