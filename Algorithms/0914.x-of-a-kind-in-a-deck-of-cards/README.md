# [914. X of a Kind in a Deck of Cards](https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/)

In a deck of cards, each card has an integer written on it.

Return `true` if and only if you can choose `X >= 2` such that it is possible to split the entire deck into 1 or more groups of cards, where:

- Each group has exactly X cards.
- All the cards in each group have the same integer.

Example 1:

```text
Input: [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4]
```

Example 2:

```text
Input: [1,1,1,2,2,2,3,3]
Output: false
Explanation: No possible partition.
```

Example 3:

```text
Input: [1]
Output: false
Explanation: No possible partition.
```

Example 4:

```text
Input: [1,1]
Output: true
Explanation: Possible partition [1,1]
```

Example 5:

```text
Input: [1,1,2,2,2,2]
Output: true
Explanation: Possible partition [1,1],[2,2],[2,2]
```

Note:

- `1 <= deck.length <= 10000`
- `0 <= deck[i] < 10000`