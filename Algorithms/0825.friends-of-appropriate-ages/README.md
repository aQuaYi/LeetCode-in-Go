# [825. Friends Of Appropriate Ages](https://leetcode.com/problems/friends-of-appropriate-ages/)

## 题目

Some people will make friend requests. Thelist of their ages is given andages[i]is the age of theith person.

Person A will NOT friend request person B (B != A) if any of the following conditions are true:

- age[B]<= 0.5 * age[A]+ 7
- age[B]> age[A]
- age[B]> 100 &&age[A]< 100

Otherwise, A will friend request B.

Note that ifA requests B, B does not necessarily request A. Also, people will not friend request themselves.

How many total friend requests are made?

Example 1:

```text
Input: [16,16]
Output: 2
Explanation: 2 people friend request each other.
```

Example 2:

```text
Input: [16,17,18]
Output: 2
Explanation: Friend requests are made 17 -> 16, 18 -> 17.
```

Example 3:

```text
Input: [20,30,100,110,120]
Output:
Explanation: Friend requests are made 110 -> 100, 120 -> 110, 120 -> 100.
```

Notes:

1. 1 <= ages.length<= 20000.
1. 1 <= ages[i] <= 120.

## 解题思路

见程序注释
