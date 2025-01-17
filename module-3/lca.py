import math

class LongestCommonAncester:
    def __init__(self, tree):
        numNodes = len(tree)
        self.maxJumpSize = int(math.ceil(math.log2(numNodes))) + 1
        self.tree = tree
        self.parent = [[-1]*self.maxJumpSize for _ in range(numNodes)]
        self.depth = [-1]*numNodes

        def preprocessImmediateAncestor(node, parent=-1, depth=0):
            self.parent[node][0] = parent
            self.depth[node] = depth
            for neighbor in self.tree[node]:
                if neighbor != parent:
                    preprocessImmediateAncestor(neighbor, node, depth + 1)
                
        preprocessImmediateAncestor(0)

        for k in range(self.maxJumpSize, -1, -1):
            for node in range(numNodes):
                if self.parent[node][k-1] != -1:
                     self.parent[node][k] = self.parent[self.parent[node][k-1]][k-1]
        

    def get_lca(self, u, v):
        if self.depth[u] < self.depth[v]:
            u, v = v, u
        
        for k in range(self.maxJumpSize-1, -1, -1):
            if self.depth[u] - (1 << k) >= self.depth[v]:
                u = self.parent[u][k]

        if u == v:
            return u
        
        for k in range(self.maxJumpSize-1, -1, -1):
            if self.parent[u][k] != self.parent[v][k]:
                u = self.parent[u][k]
                v = self.parent[v][k]

        return self.parent[u][0]


if __name__ == "__main__":

    tree = [
        [1, 2],
        [3, 4],
        [],
        [],
        [5],
        []
    ]

    solver = LongestCommonAncester(tree)

    queries = [(1, 2, 0), (3, 2, 0), (3, 5, 1), (2, 5, 0), (0, 3, 0)]
    for [a, b, expected] in queries:
        print(f"LCA({a}, {b}): {solver.get_lca(a, b)} expected {expected}")