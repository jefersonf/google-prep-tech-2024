class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        memo = [0 for _ in range(amount+1)]
        memo[0] = 1
        for i in range(len(coins)):
            for x in range(coins[i], amount+1):
                memo[x] += memo[x - coins[i]]
        return memo[amount]