# [935. Knight Dialer](https://leetcode.com/problems/knight-dialer/)

A chess knight can move as indicated in the chess diagram below:

![knight](knight.png)

![keypad](keypad.png)

This time, we place our chess knight on any numbered key of a phone pad (indicated above), and the knight makes `N-1` hops. Each hop must be from one key to another numbered key.

Each time it lands on a key (including the initial placement of the knight), it presses the number of that key, pressing N digits total.

How many distinct numbers can you dial in this manner?

Since the answer may be large, **output the answer modulo 10^9 + 7**.

Example 1:

```text
Input: 1
Output: 10
```

Example 2:

```text
Input: 2
Output: 20
```

Example 3:

```text
Input: 3
Output: 46
```

Note:

- `1 <= N <= 5000`

## 提醒

```text
1   2   3
4   5   6
7   8   9
*   0   #
```

4 可以到达 0， 4 -> 7 -> * -> 0。同理，6 也可以达到 0。
