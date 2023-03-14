from collections import defaultdict

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        s_counter = defaultdict(int)
        for letter in s:
            s_counter[letter] += 1

        t_counter = defaultdict(int)
        for letter in t:
            t_counter[letter] += 1

        
        for letter, count in t_counter.items():
            if letter not in s_counter or s_counter[letter] != count:
                return False
        return True