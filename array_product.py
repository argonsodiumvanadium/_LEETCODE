"""
This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
"""

def compute(arr):
    prod=1
    ret=[]
    for elem in arr:
        prod*=elem
    for i,elem in enumerate(arr[::-1]):
        ret.append(prod/elem)
    return ret[::-1]

print(compute([1,2,3,4,5]))
        
