class Solution:
    def firstUniqChar(self, s: str) -> int:
        co = Counter(s)
        for i,j in co.items():
            if j == 1:
                return s.index(i)
        return -1
        