from typing import List


class Solution:
    def countPairs(self, nums: List[int], target: int) -> int:
        nums = sorted(nums)
        count = 0
        for i in range(len(nums)):
            left = i + 1
            right = len(nums) - 1
            while left <= right:
                mid = (left + right) // 2
                if nums[i] + nums[mid] >= target:
                    right = mid - 1
                else:
                    left = mid + 1
            count += (right - i)
        return count