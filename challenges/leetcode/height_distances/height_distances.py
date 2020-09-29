class Solution(object):
    def heightChecker(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        
        # sorting the list
        sorted_heights = sorted(heights)
        
        wrong_positions = 0
        
        # fetch the worng positions
        for i in range(len(heights)):
            if heights[i] != sorted_heights[i]:
                wrong_positions += 1
                
        return wrong_positions
