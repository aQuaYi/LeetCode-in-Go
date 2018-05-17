# [818. Race Car](https://leetcode.com/problems/race-car/)

## 题目

Your car starts at position 0 and speed +1 on an infinite number line. (Your car can go into negative positions.)

Your car drives automatically according to a sequence of instructions A (accelerate) and R (reverse).

When you get an instruction "A", your car does the following:position += speed, speed *= 2.

When you get an instruction "R", your car does the following: if your speed is positive thenspeed = -1, otherwisespeed = 1. (Your position stays the same.)

For example, after commands "AAR", your car goes to positions 0->1->3->3, and your speed goes to 1->2->4->-1.

Now for some target position, say the length of the shortest sequence of instructions to get there.

```text
Example 1:
Input:
target = 3
Output: 2
Explanation:
The shortest instruction sequence is "AA".
Your position goes from 0->1->3.
```

```text
Example 2:
Input:
target = 6
Output: 5
Explanation:
The shortest instruction sequence is "AAARA".
Your position goes from 0->1->3->7->7->6.
```

Note:

1. 1 <= target <= 10000.

## 解题思路

见程序注释
