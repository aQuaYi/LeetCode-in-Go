# [483. Smallest Good Base](https://leetcode.com/problems/smallest-good-base/)

## 题目

For an integer n, we call k>=2 a good base of n, if all digits of n base k are 1.
Now given a string representing n, you should return the smallest good base of n in string format.

把 n 转换成 k 进制后，新的数全是 1

Example 1:

```text
Input: "13"
Output: "3"
Explanation: 13 base 3 is 111.
```

Example 2:

```text
Input: "4681"
Output: "8"
Explanation: 4681 base 8 is 11111.
```

Example 3:

```text
Input: "1000000000000000000"
Output: "999999999999999999"
Explanation: 1000000000000000000 base 999999999999999999 is 11.
```

Note:

1. The range of n is [3, 10^18].
1. The string representing n is always valid and will not have leading zeros.

## 解题思路

见程序注释

```golang
func smallestGoodBase(n string) string {
    num, _ := strconv.Atoi(n)

    mMax := int(math.Log2(float64(num)))

    for m := mMax; 1 < m; m-- {
        k := int(math.Pow(float64(num), 1.0/float64(m)))

        tmp := int(math.Pow(float64(k), float64(m)+1)-1) / (k - 1)

        if tmp == num {
            return strconv.Itoa(k)
        }
    }

    return strconv.Itoa(num - 1)
}
```

以上程序在逻辑和数学上都是完全正确的，但是无法得出正确的解。
例如，当输入"16035713712910627"，其解为"502"，但是当　for 循环中的 k == 502 时，temp == 16035713712910626 与 num 不相等。这是由于精度不够引起的错误。
