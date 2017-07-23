# [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

## 题目
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
```go
func convert(text string, nRows int) string
```
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

## 解题思路
输入"ABCDEFGHIJKLMNOPQRSTUVWXYZ"和参数4后，得到答案"AGMSYBFHLNRTXZCEIKOQUWDJPV"，
按照题目的摆放方法，可得：
```
A  G  M  S  Y
B FH LN RT XZ
CE IK OQ UW
D  J  P  V
```
这样更容易得到convert的处理过程：
1. 按照参数4，先准备4个存放字符串的容器，编号分别为0,1,2,3
1. 按照0,1,2,3,2,1顺序，把text中的字符，存入以上容器
1. 上一步没有存完的话，重复一遍，直到存完为止。
1. 把容器0,1,2,3中的内容，依次合并起来，输出。

## 总结


## Fastest Solution

```go
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	} 

	k := numRows*2 - 2
	n := len(s)
	buff := bytes.Buffer{}
	for i := 0; i < n; i += k {
		buff.WriteByte(s[i])
	}
	for i := 1; i < numRows-1; i++ {
		if i >= n {
			break
		}
		t := k - 2*i
		for j := i; j < n; j += k {
			buff.WriteByte(s[j])
			if j+t >= n {
				break
			}
			buff.WriteByte(s[j+t])
		}
	}
	for i := numRows - 1; i < n; i += k {
		buff.WriteByte(s[i])
	}
	return buff.String()
}
```