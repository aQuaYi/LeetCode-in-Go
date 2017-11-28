package Problem0491

import (
	"fmt"
)

func findSubsequences(nums []int) [][]int {
	size := len(nums)
	temp := make([][]int, 1, (size-1)*(size-1))
	temp[0] = []int{}

	for _, n := range nums {
		sizeRes := len(temp)
		for i := 0; i < sizeRes; i++ {
			if len(temp[i]) == 0 ||
				temp[i][len(temp[i])-1] <= n {
				t := make([]int, len(temp[i])+1)
				copy(t, temp[i])
				t[len(temp[i])] = n
				temp = append(temp, t)
			}
		}
	}

	isAdded := make(map[string]bool)
	res := make([][]int, 0, len(temp))
	for _, t := range temp {
		key := fmt.Sprint(t)
		if len(t) >= 2 && !isAdded[key] {
			res = append(res, t)
			isAdded[key] = true
		}
	}

	return res
}

// class Solution {
//     public List<List<Integer>> findSubsequences(int[] nums) {
//         List<List<Integer>> result = new ArrayList<>();
//         List<Integer> list = new ArrayList<>();
//         findSubseq(nums, 0, result, list);
//         return result;
//     }
//     private void findSubseq(int[] nums, int index, List<List<Integer>> result, List<Integer> list) {
//         if (index == nums.length) {
//             if (list.size() > 1) {
//                 result.add(new ArrayList<>(list));
//             }
//             return ;
//         }
//         if (list.isEmpty() || list.get(list.size() - 1) <= nums[index]) {
//             list.add(nums[index]);
//             findSubseq(nums, index + 1, result, list);
//             list.remove(list.size() - 1);
//         }
//         if (list.isEmpty() || list.get(list.size() - 1) != nums[index]){
//             findSubseq(nums, index + 1, result, list);
//         }
//     }
// }
