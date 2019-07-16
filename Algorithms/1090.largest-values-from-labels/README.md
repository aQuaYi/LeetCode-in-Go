# [1090. Largest Values From Labels](https://leetcode.com/problems/largest-values-from-labels/)

We have a set of items: the i-th item has value values[i] and label labels[i].

Then, we choose a subset S of these items, such that:

- |S| <= num_wanted
- For every label L, the number of items in S with label L is <= use_limit.

Return the largest possible sum of the subset S.

Example 1:

```text
Input: values = [5,4,3,2,1], labels = [1,1,2,2,3], num_wanted = 3, use_limit = 1
Output: 9
Explanation: The subset chosen is the first, third, and fifth item.
```

Example 2:

```text
Input: values = [5,4,3,2,1], labels = [1,3,3,3,2], num_wanted = 3, use_limit = 2
Output: 12
Explanation: The subset chosen is the first, second, and third item.
```

Example 3:

```text
Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 1
Output: 16
Explanation: The subset chosen is the first and fourth item.
```

Example 4:

```text
Input: values = [9,8,8,7,6], labels = [0,0,0,1,1], num_wanted = 3, use_limit = 2
Output: 24
Explanation: The subset chosen is the first, second, and fourth item.
```

Note:

1. `1 <= values.length == labels.length <= 20000`
1. `0 <= values[i], labels[i] <= 20000`
1. `1 <= num_wanted, use_limit <= values.length`
