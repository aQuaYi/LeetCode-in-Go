# [843. Guess the Word](https://leetcode.com/problems/guess-the-word/)

## 题目

This problem is aninteractive problemnew to the LeetCode platform.

We are given a word list of unique words, each word is 6 letters long, and one word in this list is chosen as secret.

You may call master.guess(word)to guess a word. The guessed word should havetype stringand must be from the original listwith 6 lowercase letters.

This function returns anintegertype, representingthe number of exact matches (value and position) of your guess to the secret word. Also, if your guess is not in the given wordlist, it will return -1 instead.

For each test case, you have 10 guesses to guess the word. At the end of any number of calls, if you have made 10 or less calls to master.guessand at least one of these guesses was the secret, you pass the testcase.

Besides the example test case below, there will be 5additional test cases, each with 100 words in the word list. The letters of each word in those testcases were chosenindependently at random from 'a' to 'z', such that every word in the given word lists is unique.

```text
Example 1:
Input:secret = "acckzz", wordlist = ["acckzz","ccbazz","eiowzz","abcczz"]

Explanation:

master.guess("aaaaaa") returns -1, because"aaaaaa"is not in wordlist.
master.guess("acckzz") returns 6, because"acckzz" is secret and has all 6matches.
master.guess("ccbazz") returns 3, because"ccbazz"has 3 matches.
master.guess("eiowzz") returns 2, because"eiowzz"has 2matches.
master.guess("abcczz") returns 4, because"abcczz" has 4 matches.

We made 5 calls tomaster.guess and one of them was the secret, so we pass the test case.
```

Note: Any solutions that attempt to circumvent the judgewill result in disqualification.

## 解题思路

见程序注释
