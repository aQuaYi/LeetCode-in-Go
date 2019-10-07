from operator import itemgetter
from typing import List

class Solution:
    def maxChunksToSorted(self, arr: List[int]) -> int:

        def helper(arr: List[int]) -> int:
            lastIdx, res = 0, 0
            for i in range(len(arr)):
                if lastIdx < arr[i]:
                    lastIdx = arr[i]
                    continue

                if i == lastIdx:
                    res += 1
                    lastIdx += 1
            return res

        def convert(arr: List[int]) -> List[int]:
            a = list()
            for i in range(len(arr)):
                a.append([arr[i], i])
            print(a)
            b = sorted(a, key=itemgetter(0))
            print(b)
            res = [0] * len(a)
            for i in range(len(b)):
                res[b[i][1]] = i
            print(res)
            return res

        return helper(convert(arr))

if __name__ == '__main__':
    sol = Solution()
    print(sol.maxChunksToSorted([5,4,3,2,1]))
    print(sol.maxChunksToSorted([6,4,5]))
