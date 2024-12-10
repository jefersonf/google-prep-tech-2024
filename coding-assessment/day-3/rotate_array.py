from typing import List


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        
        left, right = 0, len(nums) - (k % len(nums)) - 1
        self.reverse(nums, left, right)
        
        left, right = len(nums) - (k % len(nums)), len(nums) - 1
        self.reverse(nums, left, right)

        left, right = 0, len(nums) - 1
        self.reverse(nums, left, right)

    def reverse(self, nums: List[int], left: int, right: int) -> None:
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1