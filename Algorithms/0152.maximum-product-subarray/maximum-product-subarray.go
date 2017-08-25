package Problem0152

func maxProduct(nums []int) int {
	max := nums[0]
	temps := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if max < 0 {
				max = 0
			}
			temps = make([]int, 0, len(nums))
			continue
		}

		mult(temps, nums[i])
		temps = append(temps, nums[i])
		tempMax := getMax(temps)
		if max < tempMax {
			max = tempMax
		}
	}

	return max
}

func mult(nums []int, n int) {
	for i := 0; i < len(nums); i++ {
		nums[i] *= n
	}
}

func getMax(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}
