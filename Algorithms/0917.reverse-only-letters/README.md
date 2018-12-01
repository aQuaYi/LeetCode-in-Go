# [917. Reverse Only Letters](https://leetcode.com/problems/reverse-only-letters/)

Given a string `S`, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

Example 1:

```text
Input: "ab-cd"
Output: "dc-ba"
```

Example 2:

```text
Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"
```

Example 3:

```text
Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"
```

Note:

1. `S.length <= 100`
1. `33 <= S[i].ASCIIcode <= 122`
1. `S` doesn't contain `\` or `"`
