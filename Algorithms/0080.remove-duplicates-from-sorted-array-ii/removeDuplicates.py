from typing import List

class Solution:
        def removeDuplicates(self, nums: List[int]) -> int:
            """
            @params: nums a int list
            @output: int -> length of the new list
            """
            size = len(nums)

            i = 2
            for j in range(i, size):
                if nums[i-2] != nums[j]:
                    nums[i] = nums[j]
                    i += 1
            return i
