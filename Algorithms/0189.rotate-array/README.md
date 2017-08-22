# [189. Rotate Array](https://leetcode.com/problems/rotate-array/)

## 题目

Rotate an array of n elements to the right by k steps.
For example, with n = 7 and k = 3, the array `[1,2,3,4,5,6,7]` is rotated to `[5,6,7,1,2,3,4]`. 

Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.


[show hint]
Hint:
Could you do it in-place with O(1) extra space?


Related problem: [Reverse Words in a String II](https://leetcode.com/problems/reverse-words-in-a-string-ii/)

## 解题思路
###  三次翻转
1. nums 整体翻转
1. nums[:k] 翻转
1. nums[k:] 翻转
