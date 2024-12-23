import heapq
from collections import defaultdict

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def treeQueries(self, root: Optional[TreeNode], queries: List[int]) -> List[int]:
        depth = {}
        height = {}
        layer = defaultdict(list)
        def solve(node: Optional[TreeNode], curDepth: int) -> int:
            if not node:
                return -1
            k = node.val
            layer[curDepth].append(k)
            depth[k] = curDepth
            height[k] = 1 + max(solve(node.left, curDepth + 1), solve(node.right, curDepth + 1))
            return height[k]

        solve(root, 0)
        treeHeight = defaultdict(int)
        for i, neighbors in layer.items():
            layer_pq = []
            for node in neighbors:
                heapq.heappush(layer_pq, (-(depth[node] + height[node]), node))

            top1 = heapq.heappop(layer_pq)
            top2 = (0, 0)
            if len(layer_pq) > 0:
                top2 = heapq.heappop(layer_pq)
            # print("1st: ", top1)
            # print("2nd: ", top2)
            for node in layer[depth[neighbors[0]]]:
                if node == top1[1]:
                    treeHeight[node] = max(depth[node]-1, -top2[0])
                elif node == top2[1]:
                    treeHeight[node] = max(depth[node]-1, -top1[0])
                else:
                    treeHeight[node] = max(depth[node]-1, -top1[0])
        answers = []
        for node in queries:
            answers.append(treeHeight[node])
        return answers