class Solution(object):
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        
        hamming_distance = 0
        
        # fetching the binary list for both variables
        x_binary = list(bin(x)[2:])
        y_binary = list(bin(y)[2:])
        
        
        # if there is a variable that is smaller lenghtwise, appending 0s.
        if len(x_binary) > len(y_binary):
            for i in range(len(x_binary)):
                y_binary.insert(0,'0')
                
                if len(x_binary) == len(y_binary):
                    break                    
        elif len(y_binary) > len(x_binary):
            for i in range(len(y_binary)):
                x_binary.insert(0,'0')
                
                if len(y_binary) == len(x_binary):
                    break                        
        
        # fetching the different bits
        for i in range(len(x_binary)):
            if x_binary[i] != y_binary[i]:
                hamming_distance += 1
                
        
        return hamming_distance
        
