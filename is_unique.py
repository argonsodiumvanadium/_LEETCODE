#  check if a string s2 is a rotation of string s1

def is_rotation (s1, s2):
    active_itrs = dict({})
    for n in range(len(s2)*3):
        i = n%len(s2)
        for key, val in active_itrs.items():
            if s1[val%len(s1)] == s2[i]:
                active_itrs[key] += 1 
            else:
                del active_itrs[key]
            if active_itrs[key] == len(s1):
                return (key, val)
        if s2[i] == s1[0]:
            active_itrs[i] = 1

print(is_rotation("waterloo","loowate"))
