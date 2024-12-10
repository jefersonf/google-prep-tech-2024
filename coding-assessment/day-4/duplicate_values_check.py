
class Solution:
    def solution(array) -> bool:
        max_value = 50000

        for i in range(len(array)):
            array[i] += max_value

        seen = [False]*(2*max_value + 1)

        for v in array:
            if seen[v]:
                return False
            seen[v] = True

        return True