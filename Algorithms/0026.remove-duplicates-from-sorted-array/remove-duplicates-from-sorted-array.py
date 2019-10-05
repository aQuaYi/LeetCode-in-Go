from typing import List

class Solution:
        def removeDuplicates(self, nums: List[int]) -> int:
            if not nums:
                return 0
            index_i = 0
            for index_j in range(1, len(nums)):
                if nums[index_i] != nums[index_j]:
                    index_i += 1
                    nums[index_i] = nums[index_j]
            print(nums[:index_i+1])
            return index_i + 1

if __name__ == '__main__':
    sol = Solution()
    print(sol.removeDuplicates([0,0,1,1,1,2,2,3,3,4]))
