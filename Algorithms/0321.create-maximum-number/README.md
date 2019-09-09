# [321. Create Maximum Number](https://leetcode.com/problems/create-maximum-number/)

Given two arrays of length m and n with digits 0-9 representing two numbers.
Create the maximum number of length k <= m + n from digits of the two. The relative order of the digits
from the same array must be preserved. Return an array of the k digits. You should try to optimize your time and space complexity.

Example 1:

```text
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
return [9, 8, 6, 5, 3]
```

Example 2:

```text
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
return [6, 7, 6, 0, 4]
```

Example 3:

```text
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
return [9, 8, 9]
```

dits:Special thanks to @dietpepsi for adding this problem and creating all test cases.
