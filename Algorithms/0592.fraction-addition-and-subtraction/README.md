# [592. Fraction Addition and Subtraction](https://leetcode.com/problems/fraction-addition-and-subtraction/)

## 题目

Given a string representing an expression of fraction addition and subtraction, you need to return the calculation result in string format. The final result should be irreducible fraction. If your final result is an integer, say 2, you need to change it to the format of fraction that has denominator 1. So in this case, 2 should be converted to 2/1.

Example 1:

```text
Input:"-1/2+1/2"
Output: "0/1"
```

Example 2:

```text
Input:"-1/2+1/2+1/3"
Output: "1/3"
```

Example 3:

```text
Input:"1/3-1/2"
Output: "-1/6"
```

Example 4:

```text
Input:"5/3+1/3"
Output: "2/1"
```

Note:

1. The input string only contains '0' to '9', '/', '+' and '-'. So does the output.
1. Each fraction (input and output) has format ±numerator/denominator. If the first input fraction or the output is positive, then '+' will be omitted.
1. The input only contains valid irreducible fractions, where the numerator and denominator of each fraction will always be in the range [1,10]. If the denominator is 1, it means this fraction is actually an integer in a fraction format defined above.
1. The number of given fractions will be in the range [1,10].
1. The numerator and denominator of the final result are guaranteed to be valid and in the range of 32-bit int.

## 解题思路

见程序注释
