from typing import List


class Solution:
    def solution(mat) -> List[List[str]]:
        rows, cols = len(mat), len(mat[0])

        # transpose matrix
        for i in range(rows):
            for j in range(i + 1, cols):
                mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
        
        # reverse rows
        answer = [mat[rows-i-1] for i in range(rows)]
        
        return answer