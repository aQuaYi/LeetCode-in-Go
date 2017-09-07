# [87. Scramble String](https://leetcode.com/problems/scramble-string/)

## 题目
Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":
```
    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
```
To scramble the string, we may choose any non-leaf node and swap its two children.

For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".
```
    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
```
We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".
```
    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
```

We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.

## 解题思路

[来源](http://blog.sina.com.cn/s/blog_b9285de20101gy6n.html)

dp[i][j][k] 代表了s1从i开始，s2从j开始，长度为k的两个substring是否为scramble
string。

有三种情况需要考虑：
1. 如果两个substring相等的话，则为true
2. 如果两个substring中间某一个点，左边的substrings为scramble string，同时右边的substrings也为scramble string，则为true
3. 如果两个substring中间某一个点，s1左边的substring和s2右边的substring为scramble string, 同时s1右边substring和s2左边的substring也为scramble string，则为true

```go
func isScramble(a, b string) bool {
	if a == b {
		return true
	}
	// 题目规定了， s1 与 s2 等长
	n := len(a)

	dp := make([][][]bool, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = make([][]bool, n)
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = make([]bool, n+1)

			maxK := n - max(i, j)
			for k := 1; k <= maxK; k++ {
				if a[i:i+k] == b[j:j+k] {
					dp[i][j][k] = true
				} else {
					for d := 1; d < k; d++ {
						if (dp[i][j][d] && dp[i+d][j+d][k-d]) ||
							(dp[i][j+k-d][d] && dp[i+d][j][k-d]) {
							dp[i][j][k] = true
							break
						}
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```