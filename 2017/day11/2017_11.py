#from pdb import set_trace as bp #use pb() #as breakpoint

class HexCounter(object):

    

    def __init__(self):

        self.x = 0
        self.y = 0
        self.z = 0

        self.final_max = 0

    def move(self, direction):

        if direction == "n":
            self.y += 1
            self.x -= 1
            
        elif direction == "s":
            self.y -= 1
            self.x += 1
            
        elif direction == "ne":
            self.z += 1
            self.x -= 1
            
        elif direction == "nw":
            self.z -= 1
            self.y += 1
            
        elif direction == "se":
            self.z += 1
            self.y -= 1
            
        elif direction == "sw":
            self.z -= 1
            self.x += 1
            
        else:
            raise ValueError("Undefined direction: ", direction)

        temp = self.max()
        if temp > self.final_max:
            self.final_max = temp

    def max(self):

        total = 0
        maxx = abs(self.x)
        maxy = abs(self.y)
        maxz = abs(self.z)

        #print("x: ", maxx, " y: ", maxy, " z: ", maxz)

        total = abs(max(maxx, maxy, maxz))

        return int(total)
        






lines = ""
#number 1
from collections import Counter

def test(coords, expected):

    coords = coords.split(",")
    hexer = HexCounter()
    for coord in coords:
        hexer.move(coord)
    if hexer.max() != expected:
        print("Result is: ", hexer.max(), " not: ", expected)
    return hexer.max() == expected

def tests():

    prams = ["ne,ne,ne", "ne,ne,sw,sw", "ne,ne,s,s", "se,sw,se,sw,sw"]
    #prams = ["ne,ne,s,s", "se,sw,se,sw,sw"]
    expected = [3, 0, 2, 3]
    #expected = [2, 3]

    results = [test(prams[i], expected[i]) for i in range(len(prams))]    

    if not all(results):
        raise ValueError("Wrong", results)

    return True

if tests():

    #f = open("Inputs/2017_11_test.txt", "r")
    f = open("Inputs/2017_11.txt", "r")
    for line in f:
        lines += line
    f.close()

    coords = line[:-1].split(",")
    #coords = Counter(coords)
    #print(coords)
    #print(adjust(coords))

    hexer = HexCounter()
    for coord in coords:

        hexer.move(coord)

    print(hexer.max())

#number 2

    print(hexer.final_max)








