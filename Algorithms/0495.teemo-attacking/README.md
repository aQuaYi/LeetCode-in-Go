# [495. Teemo Attacking](https://leetcode.com/problems/teemo-attacking/)

## 题目
In LOL world, there is a hero called Teemo and his attacking can make his enemy Ashe be in poisoned condition. Now, given the Teemo's attacking ascending time series towards Ashe and the poisoning time duration per Teemo's attacking, you need to output the total time that Ashe is in poisoned condition.


You may assume that Teemo attacks at the very beginning of a specific time point, and makes Ashe be in poisoned condition immediately.

Example 1:
```
Input: [1,4], 2
Output: 4
Explanation: At time point 1, Teemo starts attacking Ashe and makes Ashe be poisoned immediately. This poisoned status will last 2 seconds until the end of time point 2. And at time point 4, Teemo attacks Ashe again, and causes Ashe to be in poisoned status for another 2 seconds. So you finally need to output 4.
```

Example 2:
```
Input: [1,2], 2
Output: 3
Explanation: At time point 1, Teemo starts attacking Ashe and makes Ashe be poisoned. This poisoned status will last 2 seconds until the end of time point 2. However, at the beginning of time point 2, Teemo attacks Ashe again who is already in poisoned status. Since the poisoned status won't add up together, though the second poisoning attack will still work at time point 2, it will stop at the end of time point 3. So you finally need to output 3.
```



Note:
1. You may assume the length of given time series array won't exceed 10000.
1. You may assume the numbers in the Teemo's attacking time series and his poisoning time duration per attacking are non-negative integers, which won't exceed 10,000,000.

## 解题思路
题目大意：给一个数组和一个数字，数组中元素是升序的，每一个元素代表释放毒气的时间点，数字duration表示释放毒气能够中毒的时间，比如时间点1开始攻击，持续2秒，则在第2秒结束后才会恢复不中毒，如果前一次中毒ing而后一次提早又被攻击，那就是后一次攻击后+持续时间才会恢复不中毒，现在要求这个人一共中毒了多少秒？
