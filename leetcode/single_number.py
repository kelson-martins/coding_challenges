
def singleNumber(nums):
    """
    :type nums: List[int]
    :rtype: int
    """
    
    n_map = {}
    n = 0
    
    for number in nums:            
        
        if number in n_map:
            del n_map[number]
        else:
            n_map[number] = number
            n = number
        
    return n_map.values()[0]

print(singleNumber([1,2,3,2,7,3,1]))
