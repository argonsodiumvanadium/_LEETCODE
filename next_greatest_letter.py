# https://leetcode.com/problems/find-smallest-letter-greater-than-target/

class Solution:
    def nextGreatestLetter (self, chars, target) -> str:
        _min=None
        for char in chars:
            if ord(char) > ord(target):
                if ord(_min if _min else 'z')  >= ord(char):
                    _min=char
        return _min if _min else chars[0]
print(Solution().nextGreatestLetter([char for char in ["a","b","c","d","e","f","g","h","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"]],'y'))
