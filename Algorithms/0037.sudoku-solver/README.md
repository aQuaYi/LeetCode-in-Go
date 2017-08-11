# [37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)

## 题目
Write a program to solve a Sudoku puzzle by filling the empty cells.

Empty cells are indicated by the character `'.'`.

You may assume that there will be only one unique solution.

![problem](1.png)

A sudoku puzzle...

![solution](2.png)

...and its solution numbers marked in red.

## 解题思路
我把数独框中，需要保证数字不重复的3×3小块，称为一个block。
由于数独需要保持每行，每列，每个block中的数字不重复。解题思路如下：
1. 依次往9个block中，分别填写1~9。
1. 如果block中已经存在n了，去填写下一个数。
1. 在可行的空位填好 n 后
    1. 如果后面的填写没有问题，返回true
    1. 如果后面的填写有问题，把n移入下一个可行的位置。
        1. n在这个block中，没有位置放了，返回false
1. 1～9都填写完了，就去填写下一个block分别填写1～9
1. 所有的block都填写完了，结束。

具体过程和细节，见程序及注释。
## 总结
虽然，我一直主张编写不超过20行的函数，但是匿名函数确实好用。