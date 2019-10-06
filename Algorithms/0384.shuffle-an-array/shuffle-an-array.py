from typing import List
import random
import copy

class Solution:

    def __init__(self, nums: List[int]):
        self.origNums = copy.deepcopy(nums)
        self.nums = copy.deepcopy(nums)


    def reset(self) -> List[int]:
        """
        Resets the array to its original configuration and return it.
        """
        return self.origNums
                                                    

    def shuffle(self) -> List[int]:
        """
        Returns a random shuffling of the array.
        """
        i, j = len(self.nums), 0
        for k in range(i-1, 1, -1):
            j = random.randint(0,k)
            self.nums[k], self.nums[j] = self.nums[j], self.nums[k]
        return self.nums

# Your Solution object will be instantiated and called as such:
obj = Solution([1,2,3,4,5,6,7,8,9,10])
param_1 = obj.reset()
param_2 = obj.shuffle()
print("param_1: %s" % param_1)
print("param_2: %s" % param_2)
