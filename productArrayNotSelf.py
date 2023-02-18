class Solution:
    def productExceptSelf (self, nums) :
        answer = [1]*len(nums)
        print(answer)
        for i, val in enumerate(nums):
            print(answer)
            for j, _ in enumerate(answer):
                answer[j] *= val*(i!=j) + (i==j) 
        return answer

print(Solution().productExceptSelf([1,2,3,4]))
