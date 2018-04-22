# [809. Expressive Words](https://leetcode.com/problems/expressive-words/)

## 题目

Sometimes people repeat letters to represent extra feeling, such as "hello" -> "heeellooo", "hi" -> "hiiii". Here, we havegroups, of adjacent letters that are all the same character, and adjacent characters tothe group are different. A groupis extended if that group is length 3 or more, so "e" and "o" would be extended in the first example, and "i" would be extended in the second example. As another example, the groups of "abbcccaaaa" would be "a", "bb", "ccc", and "aaaa"; and "ccc" and "aaaa" are the extended groups of that string.

For some given string S, a query word is stretchy if it can be made to be equal to S by extending some groups. Formally, we are allowed to repeatedly choose a group(as defined above) of characters c, and add some number of thesame character c to it so that the length of the group is 3 or more. Note that we cannot extend a group of size one like "h" to a group of size two like "hh" - all extensions must leave the group extended - ie., at least 3 characters long.

Given a list of query words, return the number of words that are stretchy.

```text
Example:
Input:
S = "heeellooo"
words = ["hello", "hi", "helo"]
Output: 1
Explanation:
We can extend "e" and "o" in the word "hello" to get "heeellooo".
We can't extend "helo" to get "heeellooo" because the group "ll" is not extended.
```

Notes:

1. 0 <= len(S) <= 100.
1. 0 <= len(words) <= 100.
1. 0 <= len(words[i]) <= 100.
1. S and all words in wordsconsist only oflowercase letters

## 解题思路

见程序注释
