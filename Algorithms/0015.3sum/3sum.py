from typing import List

class Solution:
        def threeSum(self, nums: List[int]) -> List[List[int]]:
            """
            :type: nums: Listint]
            :rtype: List[List[int]]
            """
            # initialization
            nums, result, i = sorted(nums), [], 0
            while i < len(nums) -2:
                if i == 0 or nums[i-1] != nums[i]:
                    j, k = i+1, len(nums) - 1
                    while j < k: # keep i,j,k in order
                        if nums[i] + nums[j] + nums[k] < 0:
                            j += 1
                        elif nums[i] + nums[j] + nums[k] > 0:
                            k -= 1
                        else:
                            result.append([nums[i], nums[j], nums[k]])
                            j, k = j+1, k-1
                            while j < k and nums[j] == nums[j-1]:
                                j += 1
                            while j < k and nums[k] == nums[k+1]:
                                k -= 1
                i += 1
            return result


if __name__ == '__main__':
    sol = Solution()
    print(sol.threeSum([-1,0,1,2,-1,-4]))
