# problem url : https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

class Solution:
    def __init__ (self):
        self._min=10001
        self.diff=0
    def maxProfit(self, prices) -> int:
        for i in prices:
            if i - self._min > self.diff:
                self.diff=i-self._min
            if self._min > i:
                self._min = i
        return self.diff

print(Solution().maxProfit([7,1,5,6,4]))
