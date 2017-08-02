# [29. Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)

## 题目
Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.

## 解题思路
不能使用乘法，除法和取余运算，编写一个除法函数。大的步骤是
1. 取出两个数的符号和绝对值
1. 使用绝对值做除法
1. 还原结果的符号
1. 检查是否溢出

其中，使用绝对值做除法是重点，思路参考程序注释。

## 总结
递归果然简洁有力