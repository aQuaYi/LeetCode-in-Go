# [963. Minimum Area Rectangle II](https://leetcode.com/problems/minimum-area-rectangle-ii/)

Given a set of points in the xy-plane, determine the minimum area of any rectangle formed from these points, with sides not necessarily parallel to the x and y axes.

If there isn't any rectangle, return 0.

Example 1:

![1](1.png)

```text
Input: [[1,2],[2,1],[1,0],[0,1]]
Output: 2.00000
Explanation: The minimum area rectangle occurs at [1,2],[2,1],[1,0],[0,1], with an area of 2.
```

Example 2:

![2](2.png)

```text
Input: [[0,1],[2,1],[1,1],[1,0],[2,0]]
Output: 1.00000
Explanation: The minimum area rectangle occurs at [1,0],[1,1],[2,1],[2,0], with an area of 1.
```

Example 3:

![3](3.png)

```text
Input: [[0,3],[1,2],[3,1],[1,3],[2,1]]
Output: 0
Explanation: There is no possible rectangle to form from these points.
```

Example 4:

![4](4.png)

```text
Input: [[3,1],[1,1],[0,1],[2,1],[3,3],[3,2],[0,2],[2,3]]
Output: 2.00000
Explanation: The minimum area rectangle occurs at [2,1],[2,3],[3,3],[3,1], with an area of 2.
```

Note:

- 1 <= points.length <= 50
- 0 <= points[i][0] <= 40000
- 0 <= points[i][1] <= 40000
- All points are distinct.
- Answers within 10^-5 of the actual value will be accepted as correct.