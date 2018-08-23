# [888. Fair Candy Swap](https://leetcode.com/problems/fair-candy-swap/)

## 题目

Alice and Bob have candy bars of different sizes: A[i] is the size of the i-th bar of candy that Alice has, and B[j] is the size of the j-th bar of candy that Bob has.

Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same totalamount of candy. (The total amount of candya person has is the sum of the sizes of candybars they have.)

Return an integer array answhere ans[0] is the size of the candy bar that Alice must exchange, and ans[1] is the size of the candy bar that Bob must exchange.

If there are multiple answers, you may return any one of them. It is guaranteed an answer exists.

Example 1:

```text
Input: A = [1,1], B = [2,2]
Output: [1,2]
```

Example 2:

```text
Input: A = [1,2], B = [2,3]
Output: [1,2]
```

Example 3:

```text
Input: A = [2], B = [1,3]
Output: [2,3]
```

Example 4:

```text
Input: A = [1,2,5], B = [2,4]
Output: [5,4]
```

Note:

1. 1 <= A.length <= 10000
1. 1 <= B.length <= 10000
1. 1 <= A[i] <= 100000
1. 1 <= B[i] <= 100000
1. It is guaranteed that Alice and Bob have different total amounts ofcandy.
1. It is guaranteed there exists ananswer.

## 解题思路

见程序注释
