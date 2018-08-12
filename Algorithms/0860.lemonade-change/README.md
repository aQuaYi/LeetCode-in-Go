# [860. Lemonade Change](https://leetcode.com/problems/lemonade-change/)

## 题目

At a lemonade stand, each lemonade costs $5.

Customers are standing in a queue to buy from you, and order one at a time (in the order specified by bills).

Each customer will only buy one lemonade andpay with either a $5, $10, or $20 bill. You must provide the correct change to each customer, so that the net transaction is that the customer pays $5.

Note that you don't have any changein hand at first.

Return trueif and only if you can provide every customer with correct change.

Example 1:

```text
Input: [5,5,5,10,20]
Output: true
Explanation:
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.
```

Example 2:

```text
Input: [5,5,10]
Output: true
```

Example 3:

```text
Input: [10,10]
Output: false
```

Example 4:

```text
Input: [5,5,10,10,20]
Output: false
Explanation:
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can't give change of $15 back because we only have two $10 bills.
Since not every customer received correct change, the answer is false.
```

Note:

1. 0 <= bills.length <= 10000
1. bills[i]will be either5, 10, or 20.

## 解题思路

见程序注释
