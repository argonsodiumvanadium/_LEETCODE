"""
Given a string and a substring and a positive integer n.

Remove the n occurences of a substring from the end of the originally given string. If the substring is not present in the string or the number of times the substring appears is less than n, then the output is the original string. No space should be present in place of the removed substring.
"""
lambda s,t,n:for i in range(n):s=s[:s.rfind(t)]+string[s.rfind(t)+len(t):]

rem_last_n('Cats do meow-meow while lions do not meow.','meow',2)
