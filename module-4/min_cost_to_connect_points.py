class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:

        V = len(points)
        parent = list(range(V))
        edges = []
        
        def findSet(u):
            if parent[u] == u:
                return u
            parent[u] = findSet(parent[u])
            return parent[u]

        def unionSet(u, v):
            parent[findSet(u)] = findSet(v)
        
        for i in range(V):
            for j in range(i + 1, V):
                x0, y0 = points[i]
                x1, y1 = points[j]
                dist = abs(x0 - x1) + abs(y0 - y1)
                edges.append((dist, i, j))
        
        edges.sort()
        minCost = 0
        minEdges = V - 1

        for [w, s, t] in edges:
            if findSet(s) == findSet(t):
                continue
            unionSet(s, t)
            minCost += w
            minEdges -= 1
            if minEdges == 0:
                break

        return minCost
