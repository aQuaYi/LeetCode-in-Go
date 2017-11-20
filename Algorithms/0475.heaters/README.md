# [475. Heaters](https://leetcode.com/problems/heaters/)

## 题目

Winter is coming! Your first job during the contest is to design a standard heater with fixed warm radius to warm all the houses.

Now, you are given positions of houses and heaters on a horizontal line, find out minimum radius of heaters so that all houses could be covered by those heaters.

So, your input will be the positions of houses and heaters seperately, and your expected output will be the minimum radius standard of heaters.

Note:

1. Numbers of houses and heaters you are given are non-negative and will not exceed 25000.
1. Positions of houses and heaters you are given are non-negative and will not exceed 10^9.
1. As long as a house is in the heaters' warm radius range, it can be warmed.
1. All the heaters follow your radius standard and the warm radius will the same.

Example 1:

```text
Input: [1,2,3],[2]
Output: 1
Explanation: The only heater was placed in the position 2, and if we use the radius 1 standard, then all the houses can be warmed.
```

Example 2:

```text
Input: [1,2,3,4],[1,4]
Output: 1
Explanation: The two heater was placed in the position 1 and 4. We need to use radius 1 standard, then all the houses can be warmed.
```

## 解题思路

见程序注释
