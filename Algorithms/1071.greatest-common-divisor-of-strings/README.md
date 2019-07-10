# [1071. Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/)

For strings S and T, we say "T divides S" if and only if S = T + ... + T  (T concatenated with itself 1 or more times)

Return the largest string X such that X divides str1 and X divides str2.

Example 1:

```text
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
```

Example 2:

```text
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
```

Example 3:

```text
Input: str1 = "LEET", str2 = "CODE"
Output: ""
```

Note:

1. `1 <= str1.length <= 1000`
1. `1 <= str2.length <= 1000`
1. `str1[i] and str2[i] are English uppercase letters.`
