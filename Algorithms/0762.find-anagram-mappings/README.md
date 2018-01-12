# [762. Find Anagram Mappings](https://leetcode.com/problems/find-anagram-mappings/)

## 题目

Given two lists Aand B, and B is an anagram of A. B is an anagram of A means B is made by randomizing the order of the elements in A.

We want to find an index mapping P, from A to B. A mapping P[i] = j means the ith element in A appears in B at index j.

These lists A and B may contain duplicates. If there are multiple answers, output any of them.

For example, given

```text
A = [12, 28, 46, 32, 50]
B = [50, 12, 32, 46, 28]
```

We should return

```text
[1, 4, 3, 2, 0]
```

as P[0] = 1 because the 0th element of A appears at B[1], and P[1] = 4 because the 1st element of A appears at B[4], and so on.
Note:

1. A, B have equal lengths in range [1, 100].
1. A[i], B[i] are integers in range [0, 10^5].

## 解题思路

见程序注释

我把问题想复杂了，以下是想复杂以后的解法。

```text
输入
[47 34 51 47 47 34]
[47 51 34 34 47 47]
输出
[0, 2, 1, 5, 4, 3]
```

复杂的地方在于，输出包含了 [0,len(A)) 之间的所有整数

```golang
func anagramMappings(A, B []int) []int {
    n := len(A)
    indexs := make(map[int]int, n)
    values := make(map[int]int, n)
    res := make([]int, n)

    for i := range res {
        if A[i] == B[i] {
            res[i] = i
            continue
        }

        if val, ok := values[A[i]]; ok {
            res[i] = val
            delete(values, A[i])
        } else {
            indexs[A[i]] = i
        }

        if idx, ok := indexs[B[i]]; ok {
            res[idx] = i
            delete(indexs, B[i])
        } else {
            values[B[i]] = i
        }
    }

    return res
}
```
