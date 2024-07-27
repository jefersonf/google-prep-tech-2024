from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:

        currentIndex = 0
        left = 0
        while left < len(nums):
            if nums[left] != 0:
                nums[currentIndex] = nums[left]
                currentIndex += 1
            left += 1
            
        while currentIndex < len(nums):
            nums[currentIndex] = 0
            currentIndex += 1