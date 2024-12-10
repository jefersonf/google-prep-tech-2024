
class Solution:
    def solution(array) -> int:
        array = sorted(array)
        n = len(array)
        max_prod = float('-inf')
        for i in range(n):
            for j in range(n):
                for k in range(n):
                    if i != j and i != k and j != k:
                        prod = array[i] * array[j] * array[k]
                        max_prod = max(max_prod, prod)
        return max_prod