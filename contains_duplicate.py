class Solution :
	def containsDuplicate (self, nums):#:List[int]) -> bool:
		seen={}
		for i in nums:
			if i in seen:
				return False
			seen[i] = True
		return True

s = Solution()
print(s.containsDuplicate([1,2,3,1]))
print(s.containsDuplicate([1,2,3,4]))
print(s.containsDuplicate([1,1,1]))


