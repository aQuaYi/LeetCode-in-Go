# [478. Generate Random Point in a Circle](https://leetcode.com/problems/generate-random-point-in-a-circle/)

## 题目

Given the radius and x-y positions of the center of a circle, write a function `randPoint` which generates a uniform randompoint in the circle.

Note:

1. input and output values are in [floating-point](https://www.webopedia.com/TERM/F/floating_point_number.html).
1. radius and x-y position of the center of the circle is passed into the class constructor.
1. a point on the circumference of the circle is considered to bein the circle.
1. randPointreturnsa size 2 array containing x-position and y-position of the random point, in that order.

Example 1:

```text
Input:
["Solution","randPoint","randPoint","randPoint"]
[[1,0,0],[],[],[]]
Output: [null,[-0.72939,-0.65505],[-0.78502,-0.28626],[-0.83119,-0.19803]]
```

Example 2:

```text
Input:
["Solution","randPoint","randPoint","randPoint"]
[[10,5,-7.5],[],[],[]]
Output: [null,[11.52438,-8.33273],[2.46992,-16.21705],[11.13430,-12.42337]]
```

Explanation of Input Syntax:

The input is two lists:the subroutines calledand theirarguments.Solution'sconstructor has three arguments, the radius, x-position of the center, and y-position of the center of the circle. randPoint has no arguments.Argumentsarealways wrapped with a list, even if there aren't any.

## 解题思路

见程序注释
