package Problem0095

import (
	"fmt"
	"github.com/aQuaYi/LeetCode-in-Golang/kit"
)
type TreeNode = kit.TreeNode

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode

	if n <=0{
	return  res  
}

in := make([]int , n)
for i:=range in  {
	in[i] = i+1
}

pres := getPres(in )

for _, pre :=range pres  {
	temp := preIn2Tree(pre, in )
	res = append(res , temp )
}


	return res 
}

func getPres(in []int ) [][]int  {
	size := len(in )
	if size  <= 1 {
		return [][]int{in }
	}

	if size  == 2{
		return [][]int {
			[]int{in[1],in[0]},
			[]int{in[0],in[1]},
		}
	}

	res := [][]int{}
	for i:=0; i<size  ; i++{
		 ls := getPres(in[:i])
		 rs := getPres(in[i+1:])
		 for _,l :=range ls  {
			 for _,r :=range rs  {
				 temp := make([]int , 1, size )
				 temp[0] = in[i]
				 temp = append(temp , l...)
				 temp = append(temp, r...)
				 res = append(res , temp)
			 }
		 }
	}

	return res 
}


func preIn2Tree(pre, in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = preIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = preIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	msg := fmt.Sprintf("%d 不存在于 %v 中", val, nums)
	panic(msg)
}
