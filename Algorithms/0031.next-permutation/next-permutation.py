from typing import List

class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        k, l = -1, 0
        # look for the last index of a decreasing sequence starting from the beginning
        for i in range(len(nums)-1):
            if nums[i] < nums[i+1]:
                k = i
        if k == -1:
            nums.reverse()

        for i in range(k+1, len(nums)):
            if nums[i] > nums[k]:
                l = i
        # switch the last element of the decreasing sequence with the candidate, i.e. nums[l]
        nums[k], nums[l] = nums[l], nums[k]
        nums[k+1:] = nums[:k:-1]
