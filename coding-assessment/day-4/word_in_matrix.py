class Solution:

    FOUND = False
    dx = [1, 0, -1,  0, 1, -1,  1, -1]
    dy = [0, 1,  0, -1, 1, -1, -1,  1]

    def dfs(mat, vis, i, j, word, idx):
        rows, cols = len(mat), len(mat[0])
        if i < 0 or i >= rows or j < 0 or j >= cols:
            return
        
        if vis[i][j]:
            return
        
        vis[i][j] = True
        
        if idx >= len(word):
            global FOUND
            FOUND = True
            return
        
        if mat[i][j] == word[idx]:
            idx += 1
            for k in range(len(dx)):
                dfs(mat, vis, i + dx[k], j + dy[k], word, idx)

        vis[i][j] = True
        
    def solution(mat, word):
        rows, cols = len(mat), len(mat[0])
        for i in range(rows):
            for j in range(cols):
                visited = [[False for _ in range(cols)] for _ in range(rows)]
                dfs(mat, visited, i, j, word, 0)
        return FOUND
    
