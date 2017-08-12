# [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/)

## 题目
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

For example, 
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
![rainwatertrap](rainwatertrap.png)

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

## 解题思路
解题的关键有两点
1. 理解 i 点的存水量为 min(max(height[:i+1]...), max(height[i:]...)) - height[i]
1. 高效地找到 max(height[:i+1]...) 和 max(height[i:]...)

## 总结
利用记录，比每次现找要省时的多。