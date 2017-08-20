# [448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)

## 题目
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:
```
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
```

## 解题思路
利用切片的元素的特性，不断地把 a[i] 直接放入 a[i]-1 位，这样就可以很快地排序好切片。
然后，检查切片各个元素的值是否符合属性，不符合的，就放入答案中。 

## 总结


