from typing import List

class Solution:
    def maxChunksToSorted(self, arr: List[int]) -> int:
        """
        :type arr: List[int]
        :rtype: int
        """
        lastIdx, res = 0, 0
        for i in range(len(arr)):
            if lastIdx < arr[i]:
                lastIdx = arr[i]
                continue
            if i == lastIdx:
                res+=1
                lastIdx+=1
        return res



if __name__ == '__main__':
    sol = Solution()
    res = sol.maxChunksToSorted([4,3,2,1,0])
    print(res)
