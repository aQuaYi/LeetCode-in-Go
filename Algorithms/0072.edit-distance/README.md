# [72. Edit Distance](https://leetcode.com/problems/edit-distance/)

## 题目
Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)

You have the following 3 operations permitted on a word:
```
a) Insert a character
b) Delete a character
c) Replace a character
```

## 动态规划思想：

假设dp[i][j]表示以S[i]结尾的字符串和以T[j]结尾的字符串转换所需的最小操作数，考虑三种操作，然后取三者最小值：

1、替换：

假设S[i-1],T[j-1]已对齐，即dp[i-1][j-1]已知，则当S[i]==T[j]时，dp[i][j]=dp[i-1][j-1]，否则，dp[i][j]=dp[i-1][j-1]+1.

2、删除

假设S[i-1]，T[j]已对齐，即dp[i-1][j]已知，多出来的S[i]需删除，操作数+1，则dp[i][j]=dp[i-1][j]+1.

3、插入

假设S[i],T[j-1]已对齐，即dp[i][j-1]已知，需在S中插入S[i+1]=T[j]来匹配，操作数+1，则dp[i][j]=dp[i][j-1]+1.

状态转移方程：

dp[i][j]=min(dp[i-1][j-1]+(S[i]==T[j]?0,1),dp[i-1][j]+1,dp[i][j-1]+1)

初始值：

dp[i][0]=i

dp[0][j]=j

复杂度：

时间复杂度：O(m*n)

空间复杂度：O(m*n)

空间优化：

由状态转移方程可知，dp[i][j]与dp[i-1][j-1],dp[i-1][j],dp[i][j-1]有关，可以去掉一维，只留下dp[j]。

等式右边的dp[i-1][j]和dp[i][j-1]都可以直接改成dp[j]（旧的值）和dp[j-1]（已更新），只有dp[i-1][j-1]没有记录下来，通过某个变量保存起来之后就可以。

因此空间复杂度：O(n)

见程序注释
