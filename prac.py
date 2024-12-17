# # #artificial intelligence and machine learning

# # #Tensorflow: 
# # #made by google for numerical computation and ml. 
# # #Can be used to deploy ml models

# # #Keras
# # #api for building and training deep learning models

# # #Pytorch


# # #Web application
# # #Django
# # #Flask

# # #oop based language ( classes and objects )
# # #Polymorphism , encapsulation , abstraction , inheritence


# # class Animal:
# #     def __init__(self,age):
# #         self.age = age

# #     def getAge(self):
# #         return self.age * 10

# # class Dog(Animal):
# #     pass
# #     def getAge(self):
# #         return self.age * 20
    
    
# #     ##Protected -> Can be used by class and subclass
# #     ##Private -> can be used only by class
# #     ##Public -> Can be used by class, subclass, object


# # d = Dog(15)
# # print(d.getAge(4))

# # #Abstraction :- 

# # #list , tuple

# # # List can be changed []
# # # tuple cannot be changed ()
# # # {}
# # # Dictionary { key : value }

# # # integer , character , string 

# # # by using with , file need not be closed

# lst = [1,2,3]

# lst.extend([4,5])
# print(lst)

# pyhon is a high-level interpreted language known for its reliability and ease of use
# It is mostly used in web application and ai and ml

MOD = 10**9 + 7

def GetAnswer(N, A):
    total_subsets = 1 << N  # This is 2^N (all subsets)
    good_colorings = 0
    
    # Iterate over all possible 3-colorings of the array using bitmasking.
    # For each mask, assign some elements to white, some to black, some to uncolored.
    for mask in range(1, total_subsets):
        x, y = 0, 0
        
        for i in range(N):
            if mask & (1 << i):
                # Assign to white (x)
                x ^= A[i]
            else:
                # Assign to black (y)
                y ^= A[i]
        
        # Check the conditions for good coloring
        if x > 0 and y > 0 and (x % y == 0 or y % x == 0):
            good_colorings += 1
    
    return good_colorings % MOD

# Driver code to read input and run the function
if __name__ == "__main__":
    import sys
    input = sys.stdin.read
    data = input().splitlines()
    
    N = int(data[0])
    A = list(map(int, data[1:N+1]))
    
    result = GetAnswer(N, A)
    print(result)

