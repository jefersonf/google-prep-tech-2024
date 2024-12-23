class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        grid = [[1]*n, [1]*n]
        for i in range(1, m):
            for j in range(1, n):
                grid[1][j] = grid[1][j-1] + grid[0][j]
            for j in range(n):
                grid[0][j], grid[1][j] = grid[1][j], grid[0][j] 
        return grid[0][n-1]