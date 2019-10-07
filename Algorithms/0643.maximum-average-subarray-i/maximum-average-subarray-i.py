class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        """
        @params: nums -> a list of integers; k -> number of integers
        @output: the maximum average of k integers
        """
        avgs = []
        for i in range(len(nums)-k):
            avgs.append(sum(nums[i:i+k]) / k)
        return max(avgs)

    def findMaxAverage1(self, nums: List[int], k: int) -> float:
        temp = 0
        for i in range(k):
            temp += nums[i]

        Max = temp

        for i in range(k,len(nums)):
            temp = temp - nums[i-k] + nums[i]
            if Max < temp:
                Max = temp
        return Max / k
