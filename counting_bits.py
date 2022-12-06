# https://leetcode.com/problems/counting-bits/submissions/
class Solution:
    def __init__ (self):
        self.mem = []
    def countBits(self, n: int):
        for number in range(0,n+1):
            self.mem = self.mem + [self.get_bin(number)]
        return self.mem
    def get_bin (self, n:int):
        print(f'finding binary of {n}')
        _sum = 0
        l=len(self.mem)
        while n:
            if l > n:
                print(f'replacing {n} with {self.mem[n]}')
                return _sum + self.mem[n]
            _sum = _sum + (n&1)
            n=n>>1
        return _sum

print(Solution().countBits(32))
