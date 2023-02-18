class Solution:
    def backspaceCompare(self, s: str, t: str) -> bool:
        return self.simplify(s) == self.simplify(t)
    def simplify (self,s):
        c1 = ""
        for letter in s:
            if letter == "#": c1 = c1[:-1]
            else: c1 += letter
        return c1
