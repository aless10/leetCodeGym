class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = {}
        for i, num in enumerate(nums):
            value = d.get(num)
            if value is not None:
                if value[0] * 2 == target:
                    return [value[1], i]

            d[num] = (target - num, i)

        for k, v in d.items():
            complement, index = v
            if complement in d:
                complement_index = d[complement][1]
                if complement_index != index: 
                    return [index, complement_index]
