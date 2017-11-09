# [453. Minimum Moves to Equal Array Elements](https://leetcode.com/problems/minimum-moves-to-equal-array-elements/)

## 题目

Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

Example:

```text
Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
```

## 解题思路

n = len(nums)

经过 m 次移动后，每个数字都是 x，可得

```text
sum + m * (n - 1) = x * n
```

实际上，每次移动都会往最小的数字 min 上加一，所以

```text
x = min + m
```

通过以上两个等式，可得

```text
m = sum - min * n
```