# [984. String Without AAA or BBB](https://leetcode.com/problems/string-without-aaa-or-bbb/)

Given two integers A and B, return any string S such that:

S has length A + B and contains exactly A 'a' letters, and exactly B 'b' letters;
The substring 'aaa' does not occur in S;
The substring 'bbb' does not occur in S.

Example 1:

```text
Input: A = 1, B = 2
Output: "abb"
Explanation: "abb", "bab" and "bba" are all correct answers.
```

Example 2:

```text
Input: A = 4, B = 1
Output: "aabaa"

Note:

- 0 <= A <= 100
- 0 <= B <= 100
- It is guaranteed such an S exists for the given A and B.
