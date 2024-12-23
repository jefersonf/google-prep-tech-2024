class Solution:
    def minimumSteps(self, s: str) -> int:
        leftIndex = 0
        steps = 0
        for i in range(len(s)):
            if s[i] == '0':
                steps += i - leftIndex
                leftIndex += 1
        return steps