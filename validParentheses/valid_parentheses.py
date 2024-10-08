class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        m = {
            "(": ")",
            "[": "]",
            "{": "}",
        }
        for ss in s:
            if ss in {"(", "{", "["}:
                stack.append(ss)
            else:
                try:
                    last = stack.pop()
                    if m[last] != ss:
                        return False
                except IndexError:
                    return False
        return True if not stack else False
