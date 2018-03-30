#from pdb import set_trace as bp #use pb() #as breakpoint

class HexGraph(object):

    def __init__(self):

        self.north_south = 0
        self.nwest_seast = 0
        self.neast_swest = 0

    def move(self, direction):

        if direction == "n":
            self.nwest_seast += 1
            self.neast_swest += 1
        elif direction == "s":
            self.nwest_seast -= 1
            self.neast_swest -= 1
        elif direction == "ne":
            self.north_south += 1
            self.nwest_seast -= 1
        elif direction == "nw":
            self.north_south += 1
            self.neast_swest -= 1
        elif direction == "se":
            self.north_south -= 1
            self.neast_swest += 1
        elif direction == "sw":
            self.north_south -= 1
            self.nwest_seast += 1
        else:
            raise ValueError("Undefined direction: ", direction)

    def max(self):

        total = 0
        maxx = self.north_south
        maxy = self.nwest_seast
        maxz = self.neast_swest

        print("x: ", maxx, " y: ", maxy, " z: ", maxz)

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









