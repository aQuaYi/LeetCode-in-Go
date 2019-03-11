# [948. Bag of Tokens](https://leetcode.com/problems/bag-of-tokens/)

You have an initial power P, an initial score of 0 points, and a bag of tokens.

Each token can be used at most once, has a value token[i], and has potentially two ways to use it.

- If we have at least token[i] power, we may play the token face up, losing token[i] power, and gaining 1 point.
- If we have at least 1 point, we may play the token face down, gaining token[i] power, and losing 1 point.

Return the largest number of points we can have after playing any number of tokens.

Example 1:

```text
Input: tokens = [100], P = 50
Output: 0
```

Example 2:

```text
Input: tokens = [100,200], P = 150
Output: 1
```

Example 3:

```text
Input: tokens = [100,200,300,400], P = 200
Output: 2
```

Note:

1. tokens.length <= 1000
1. 0 <= tokens[i] < 10000
1. 0 <= P < 10000