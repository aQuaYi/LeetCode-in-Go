from typping import List

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        """
        @param nums: a list of integer
        @output: an integer
        """
        if len(nums) <= 1:
            return len(nums)

        nums = sorted(nums)

        max_num, temp = 0, 1
        for i in range(len(nums)):
            if nums[i] - 1 == nums[i-1]:
                temp+=1
            elif nums[i] != nums[i-1]:
                temp = 1
            if max_num < temp:
                max_num = temp
        return max_num
