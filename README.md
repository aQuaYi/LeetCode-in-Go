# [LeetCode](https://leetcode.com)习题的Golang解答。
[![Build Status](https://www.travis-ci.org/aQuaYi/LeetCode-in-Golang.svg?branch=master)](https://www.travis-ci.org/aQuaYi/LeetCode-in-Golang)
[![codecov](https://codecov.io/gh/aQuaYi/LeetCode-in-Golang/branch/master/graph/badge.svg)](https://codecov.io/gh/aQuaYi/LeetCode-in-Golang)

[本人的LeetCode主页](https://leetcode.com/aQuaYi/)

## 统计
||Easy|Medium|Hard|Total|
|:--|:--:|:--:|:--:|:--:|
|**Algorithms**|7/155|13/307|2/115|22/577|
|**Draft**|0/2|0/5|0/1|0/8|
|**Total**|7/157|13/312|2/116|22/585|


## 已添加的题目
|题号|题目|难度|总体通过率|收藏|
|:-:|:-|:-: | :-: | :-: |
|1|[Two Sum](./Algorithms/0001.two-sum)|☆|34%||
|2|[Add Two Numbers](./Algorithms/0002.add-two-numbers)|☆ ☆|27%||
|3|[Longest Substring Without Repeating Characters](./Algorithms/0003.longest-substring-without-repeating-characters)|☆ ☆|24%||
|4|[Median of Two Sorted Arrays](./Algorithms/0004.median-of-two-sorted-arrays)|☆ ☆ ☆|21%||
|5|[Longest Palindromic Substring](./Algorithms/0005.longest-palindromic-substring)|☆ ☆|25%||
|6|[ZigZag Conversion](./Algorithms/0006.zigzag-conversion)|☆ ☆|26%||
|7|[Reverse Integer](./Algorithms/0007.reverse-integer)|☆|24%||
|8|[String to Integer (atoi)](./Algorithms/0008.string-to-integer-atoi)|☆ ☆|13%||
|9|[Palindrome Number](./Algorithms/0009.palindrome-number)|☆|35%||
|10|[Regular Expression Matching](./Algorithms/0010.regular-expression-matching)|☆ ☆ ☆|24%|❤|
|11|[Container With Most Water](./Algorithms/0011.container-with-most-water)|☆ ☆|36%||
|12|[Integer to Roman](./Algorithms/0012.integer-to-roman)|☆ ☆|44%||
|13|[Roman to Integer](./Algorithms/0013.roman-to-integer)|☆|45%||
|14|[Longest Common Prefix](./Algorithms/0014.longest-common-prefix)|☆|31%||
|15|[3Sum](./Algorithms/0015.3sum)|☆ ☆|21%||
|16|[3Sum Closest](./Algorithms/0016.3sum-closest)|☆ ☆|31%||
|17|[Letter Combinations of a Phone Number](./Algorithms/0017.letter-combinations-of-a-phone-number)|☆ ☆|34%||
|18|[4Sum](./Algorithms/0018.4sum)|☆ ☆|26%||
|19|[Remove Nth Node From End of List](./Algorithms/0019.remove-nth-node-from-end-of-list)|☆ ☆|33%||
|20|[Valid Parentheses](./Algorithms/0020.valid-parentheses)|☆|33%||
|21|[Merge Two Sorted Lists](./Algorithms/0021.merge-two-sorted-lists)|☆|38%||
|22|[Generate Parentheses](./Algorithms/0022.generate-parentheses)|☆ ☆|44%|❤|

## helper
[helper](./helper)会帮助处理大部分琐碎的工作。

## kit
在[kit](./kit)中添加了LeetCode常用的数据结构和处理函数：
1. 为[*ListNode](./kit/ListNode.go)添加了与[]int相互转换的函数，方便添加单元测试。使用方式可以参考[21. Merge Two Sorted Lists](./Algorithms/0021.merge-two-sorted-lists/merge-two-sorted-lists_test.go)
