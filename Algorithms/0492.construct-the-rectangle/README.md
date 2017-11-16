# [492. Construct the Rectangle](https://leetcode.com/problems/construct-the-rectangle/)

## 题目

For a web developer, it is very important to know how to design a web page's size. So, given a specific rectangular web page’s area, your job by now is to design a rectangular web page, whose length L and width W satisfy the following requirements:

1. The area of the rectangular web page you designed must equal to the given target area.
1. The width W should not be larger than the length L, which means L >= W.
1. The difference between length L and width W should be as small as possible.

You need to output the length L and the width W of the web page you designed in sequence.

Example:

```text
Input: 4
Output: [2, 2]
Explanation: The target area is 4, and all the possible ways to construct it are [1,4], [2,2], [4,1].
But according to requirement 2, [1,4] is illegal; according to requirement 3,  [4,1] is not optimal compared to [2,2]. So the length L is 2, and the width W is 2.
```

Note:

1. The given area won't exceed 10,000,000 and is a positive integer
1. The web page's width and length you designed must be positive integers.

## 解题思路

见程序注释
