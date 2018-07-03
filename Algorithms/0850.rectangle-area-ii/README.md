# [850. Rectangle Area II](https://leetcode.com/problems/rectangle-area-ii/)

## 题目

We are given a list of (axis-aligned)rectangles. Eachrectangle[i] = [x1, y1, x2, y2], where (x1, y1) are the coordinates of the bottom-left corner, and (x2, y2) are the coordinates of the top-right corner of the ith rectangle.

Find the total area covered by all rectangles in the plane. Since the answermay be too large, return it modulo 10^9 + 7.

![pic](rectangle_area_ii_pic.png)

Example 1:

```text
Input: [[0,0,2,2],[1,0,2,3],[1,0,3,1]]
Output: 6
Explanation: As illustrated in the picture.
```

Example 2:

```text
Input: [[0,0,1000000000,1000000000]]
Output: 49
Explanation: The answer is 10^18 modulo (10^9 + 7), which is (10^9)^2 = (-7)^2 = 49.
```

Note:

1. 1 <= rectangles.length <= 200
1. rectanges[i].length = 4
1. 0 <= rectangles[i][j] <= 10^9
1. The total area covered by all rectangles will never exceed2^63 - 1and thus will fit in a 64-bit signed integer.

## 解题思路

见程序注释
