# [830. Positions of Large Groups](https://leetcode.com/problems/positions-of-large-groups/)

## 题目

In a stringSof lowercase letters, these letters form consecutive groups of the same character.

For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and"yy".

Call a group large if it has 3 or more characters. We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.

Example 1:

```text
Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
```

Example 2:

```text
Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.
```

Example 3:

```text
Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
```

Note:

- 1 <= S.length <= 1000

## 解题思路

见程序注释
