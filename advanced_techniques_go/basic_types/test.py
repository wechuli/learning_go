def doNothing(name: str,age: int)->str:
    return "nothing"


class Person:
    def __init__(self,name:str,age:int)->None:
        self.name = name
        self.age = age
    
    def sayHello(self)->str:
        return "Hello"
    
peter = Person("Peter", 20)

print(peter.sayHello())

city = {
    "name": "New York",
    "population": 10000000
}



mytuple = (1,2,3,4,5)
mylist = list(mytuple)



[i for i in range(10) if i%2==0]

a = 100

print(str(a))

print(type(str(a)))