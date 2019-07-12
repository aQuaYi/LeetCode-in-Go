# [1078. Occurrences After Bigram](https://leetcode.com/problems/occurrences-after-bigram/)

Given words first and second, consider occurrences in some text of the form "first second third", where second comes immediately after first, and third comes immediately after second.

For each such occurrence, add "third" to the answer, and return the answer.

Example 1:

```text
Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
Output: ["girl","student"]
```

Example 2:

```text
Input: text = "we will we will rock you", first = "we", second = "will"
Output: ["we","rock"]
```

Note:

1. `1 <= text.length <= 1000`
1. `text consists of space separated words, where each word consists of lowercase English letters.`
1. `1 <= first.length, second.length <= 10`
1. `first and second consist of lowercase English letters.`
