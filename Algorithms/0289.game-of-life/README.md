# [289. Game of Life](https://leetcode.com/problems/game-of-life/)

## 题目
According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."



Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

1. Any live cell with fewer than two live neighbors dies, as if caused by under-population.
1. Any live cell with two or three live neighbors lives on to the next generation.
1. Any live cell with more than three live neighbors dies, as if by over-population..
1. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Write a function to compute the next state (after one update) of the board given its current state.

Follow up: 
1. Could you solve it in-place? Remember that the board needs to be updated at the same time: You cannot update some cells first and then use their updated values to update other cells.
1. In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches the border of the array. How would you address these problems?


## 解题思路
这道题是有名的康威生命游戏, 这是一种细胞自动机，每一个位置有两种状态，1为活细胞，0为死细胞，对于每个位置都满足如下的条件：

1. 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡
2. 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活
3. 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡
4. 如果死细胞周围正好有三个活细胞，则该位置死细胞复活

由于题目中要求我们用置换方法in-place来解题，所以我们就不能新建一个相同大小的数组，那么我们只能更新原有数组，但是题目中要求所有的位置必须被同时更新，但是在循环程序中我们还是一个位置一个位置更新的，那么当一个位置更新了，这个位置成为其他位置的neighbor时，我们怎么知道其未更新的状态呢，我们可以使用状态机转换：

- 状态0： 死细胞转为死细胞
- 状态1： 活细胞转为活细胞
- 状态2： 活细胞转为死细胞
- 状态3： 死细胞转为活细胞

最后我们对所有状态对2取余，那么状态0和2就变成死细胞，状态1和3就是活细胞，达成目的。我们先对原数组进行逐个扫描，对于每一个位置，扫描其周围八个位置，如果遇到状态1或2，就计数器累加1，扫完8个邻居，如果少于两个活细胞或者大于三个活细胞，而且当前位置是活细胞的话，标记状态2，如果正好有三个活细胞且当前是死细胞的话，标记状态3。完成一遍扫描后再对数据扫描一遍，对2取余变成我们想要的结果。
