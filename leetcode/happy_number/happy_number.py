def happy_number(n):

    
    happy_number = 0
    counter = 0

    while n != 1:
        

        happy_number = 0

        for char in str(n):

            happy_number += int(char) ** 2

        n = happy_number
        
        counter += 1
        
        if counter > 100:
            n = 0
            break

    print(happy_number)
    return n

happy_number(19)