# [770. Basic Calculator IV](https://leetcode.com/problems/basic-calculator-iv/)

## 题目

Given an `expression` such as `expression = "e + 8 - a + 5"` and an evaluation map such as {"e": 1} (given in terms of evalvars = ["e"] and evalints = [1]), return a list of tokens representing the simplified expression, such as ["-1*a","14"]

- An expression alternates chunks and symbols, with a space separating each chunk and symbol.
- A chunk is either an expression in parentheses, a variable, or a non-negative integer.
- A variable is a string of lowercase letters (not including digits.) Note that variables can be multiple letters, and note that variables never have a leading coefficient or unary operator like "2x" or "-x".

Expressions are evaluated in the usual order: brackets first, then multiplication, then addition and subtraction. For example, expression = "1 + 2 * 3" has an answer of ["7"].

The format of the output is as follows:

- For each term of free variables with non-zero coefficient, we write the free variables within a term in sorted order lexicographically. For example, we would never write a term like `"b*a*c"`, only `"a*b*c"`.
- Terms have degree equal to the number of free variables being multiplied, counting multiplicity. (For example, `"a*a*b*c"` has degree 4.) We write the largest degree terms of our answer first, breaking ties by lexicographic order ignoring the leading coefficient of the term.
- The leading coefficient of the term is placed directly to the left with an asterisk separating it from the variables (if they exist.) A leading coefficient of 1 is still printed.
- An example of a well formatted answer is [`"-2*a*a*a"`, `"3*a*a*b"`, `"3*b*b"`, `"4*a"`, `"5*c"`, `"-6"`]
- Terms (including constant terms) with coefficient 0 are not included. For example, an expression of "0" has an output of [].

Examples:

```text
Input: expression = "e + 8 - a + 5", evalvars = ["e"], evalints = [1]
Output: ["-1*a","14"]

Input: expression = "e - 8 + temperature - pressure",
evalvars = ["e", "temperature"], evalints = [1, 12]
Output: ["-1*pressure","5"]

Input: expression = "(e + 8) * (e - 8)", evalvars = [], evalints = []
Output: ["1*e*e","-64"]

Input: expression = "7 - 7", evalvars = [], evalints = []
Output: []

Input: expression = "a * b * c + b * a * c * 4", evalvars = [], evalints = []
Output: ["5*a*b*c"]

Input: expression = "((a - b) * (b - c) + (c - a)) * ((a - b) + (b - c) * (c - a))",
evalvars = [], evalints = []
Output: ["-1*a*a*b*b","2*a*a*b*c","-1*a*a*c*c","1*a*b*b*b","-1*a*b*b*c","-1*a*b*c*c","1*a*c*c*c","-1*b*b*b*c","2*b*b*c*c","-1*b*c*c*c","2*a*a*b","-2*a*a*c","-2*a*b*b","2*a*c*c","1*b*b*b","-1*b*b*c","1*b*c*c","-1*c*c*c","-1*a*a","1*a*b","1*a*c","-1*b*c"]
```

Note:

1. expression will have length in range [1, 250].
1. evalvars, evalints will have equal lengths in range [0, 100].

## 解题思路

见程序注释
