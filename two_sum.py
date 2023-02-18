#https://leetcode.com/problems/two-sum/
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            newnum = nums[:i] + [target*20+20]+ nums[i+1:]
            try: return [i,newnum.index(target - nums[i])]
            except: continue

