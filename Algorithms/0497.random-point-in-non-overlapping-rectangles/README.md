# [497. Random Point in Non-overlapping Rectangles](https://leetcode.com/problems/random-point-in-non-overlapping-rectangles/)

## 题目

Given a list of non-overlapping axis-aligned rectangles rects, write a function pick which randomly and uniformily picks an integer point in the spacecovered by the rectangles.

Note:

1. An integer pointis a point that has integer coordinates.
1. A pointon the perimeterof a rectangle isincluded in the space covered by the rectangles.
1. ith rectangle = rects[i] =[x1,y1,x2,y2], where [x1, y1]are the integer coordinates of the bottom-left corner, and [x2, y2]are the integer coordinates of the top-right corner.
1. length and width of each rectangle does not exceed 2000.
1. 1 <= rects.length<= 100
1. pick return a point as an array of integer coordinates[p_x, p_y]
1. pick is called at most 10000times.

Example 1:

```text
Input:
["Solution","pick","pick","pick"]
[[[[1,1,5,5]]],[],[],[]]
Output:
[null,[4,1],[4,1],[3,3]]
```

Example 2:

```text
Input:
["Solution","pick","pick","pick","pick","pick"]
[[[[-2,-2,-1,-1],[1,0,3,0]]],[],[],[],[],[]]
Output:
[null,[-1,-2],[2,0],[-2,-1],[3,0],[-2,-2]]
```

Explanation of Input Syntax:

The input is two lists:the subroutines calledand theirarguments.Solution'sconstructor has one argument, the array of rectangles rects. pickhas no arguments.Argumentsarealways wrapped with a list, even if there aren't any.

## 解题思路

见程序注释
