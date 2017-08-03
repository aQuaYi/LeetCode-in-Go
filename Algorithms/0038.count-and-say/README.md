# [38. Count and Say](https://leetcode.com/problems/count-and-say/)

## 题目
The count-and-say sequence is the sequence of integers with the first five terms as following:
```
1.     1
2.     11
3.     21
4.     1211
5.     111221
```
1 is read off as "one 1" or 11.

11 is read off as "two 1s" or 21.

21 is read off as "one 2, then one 1" or 1211.

Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:
```
Input: 1
Output: "1"
```
Example 2:
```
Input: 4
Output: "1211"
```

## 解题思路
利用[]byte记录上一次出现的次数和元素。

详见注释

## 总结
使用辅助数据，比如切片或者映射，记录信息，可以极大地加快程序。