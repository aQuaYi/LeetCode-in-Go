# [831. Masking Personal Information](https://leetcode.com/problems/masking-personal-information/)

## 题目

We are given apersonal information string S, which may representeither an email address or a phone number.

We would like to mask thispersonal information according to thefollowing rules:

1.Email address:

We define aname to be a string of length >= 2 consistingof only lowercase lettersa-z or uppercaselettersA-Z.

An email address starts with a name, followed by thesymbol '@', followed by a name, followed by thedot'.'andfollowed by a name.

All email addresses areguaranteed to be valid and in the format of"name1@name2.name3".

To mask an email, all names must be converted to lowercase and all letters between the first and last letter of the first name must be replaced by 5 asterisks '*'.

2.Phone number:

A phone number is a string consisting ofonly the digits 0-9 or the characters from the set {'+', '-', '(', ')', ''}.You may assume a phonenumber contains10 to 13 digits.

The last 10 digits make up the localnumber, while the digits before those make up the country code. Note thatthe country code is optional. We want to expose only the last 4 digitsand mask all otherdigits.

The localnumbershould be formatted and masked as "***-***-1111",where 1 represents the exposed digits.

To mask a phone number with country code like "+111 111 111 1111", we write it in the form "+***-***-***-1111". The '+'sign and the first '-'sign before the local number should only exist if there is a country code. For example, a 12 digit phone number maskshould startwith "+**-".

Note that extraneous characters like "(", ")", " ", as well asextra dashes or plus signs not part of the above formatting scheme should be removed.

Return the correct "mask" of the information provided.

Example 1:

```text
Input: "LeetCode@LeetCode.com"
Output: "l*****e@leetcode.com"
Explanation:All names are converted to lowercase, and the letters between the
            first and last letter of the first name is replaced by 5 asterisks.
            Therefore, "leetcode" -> "l*****e".
```

Example 2:

```text
Input: "AB@qq.com"
Output: "a*****b@qq.com"
Explanation:There must be 5 asterisks between the first and last letter
            of the first name "ab". Therefore, "ab" -> "a*****b".
```

Example 3:

```text
Input: "1(234)567-890"
Output: "***-***-7890"
Explanation:10 digits in the phone number, which means all digits make up the local number.
```

Example 4:

```text
Input: "86-(10)12345678"
Output: "+**-***-***-5678"
Explanation:12 digits, 2 digits for country code and 10 digits for local number.
```

Notes:

1. S.length<=40.
1. Emails have length at least 8.
1. Phone numbers have length at least 10.

## 解题思路

见程序注释
