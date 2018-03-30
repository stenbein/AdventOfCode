
from time import sleep

def cls(): print("\n" * 100)

'''     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 

Starting at the only line touching the top of the diagram, it must go down,
pass through A, and continue onward to the first +.
Travel right, up, and right, passing through B in the process.
Continue down (collecting C), right, and up (collecting D).
Finally, go all the way left through E and stopping at F.

'''


map_str ='''     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+'''


SOUTH = (1, 0)
NORTH = (-1, 0)
EAST = (0, 1)
WEST = (0, -1)


class TextMap(object):

    def __init__(self, map_str):

        self.map_list = []
        
        for row in map_str.split("\n"):
            self.map_list.append([char for char in row])

        self.entry_point = (0, self.map_list[0].index("|"))
        self.location = self.entry_point #current location
        
        #row,col coord as direction traveled in to reach current square
        self.heading = SOUTH
        self.output = []
        self.steps = 0

    def get_current_tile(self):
        try:
            return self.map_list[self.location[0]][self.location[1]]
        except:
            raise ValueError("Args: ", self.location[0],self.location[1])

    def parse_location(self):

        current_tile = self.get_current_tile()
        
        if current_tile == "+":
            #maybe change direction change
            if self.look_ahead(self.heading):
                self.step()
            else:
                clockwise_heading = self.get_new_heading(self.heading, 1)
                c_clockwise_heading = self.get_new_heading(self.heading, -1)
                if self.look_ahead(clockwise_heading):
                    self.heading = clockwise_heading
                    self.step()
                elif self.look_ahead(c_clockwise_heading):
                    self.heading = c_clockwise_heading
                    self.step()
                else:
                    raise ValueError("Current location has no exit.", self.location)

        elif current_tile in "|-":
            self.step()
        elif current_tile in [chr(letter) for letter in range(ord("A"), ord("Z")+1)]:
            self.record(current_tile)
            self.step()
        else:
            self.display()
            raise ValueError("Current tile is unknown character input: ", current_tile)

    def step(self):
        self.steps += 1
        self.location = (self.location[0] + self.heading[0],
                             self.location[1] + self.heading[1])

    def look_ahead(self, heading):
        
        try:
            next_tile = self.map_list[self.location[0] + heading[0]][self.location[1] + heading[1]]
            if next_tile != " ":
                return True
            else:
                return False
        except:
            return False

    def get_new_heading(self, heading, direction):
        
        return (heading[1]*-direction,
            heading[0]*direction)

    def record(self, value):
        self.output.append(value)

    def walk_map(self, display):

        #check for exit condition
        while self.get_current_tile() != "F":
            if display:
                cls()
                self.display()
                sleep(0.5)
            self.parse_location()
        
        return

        
    def display(self):

        print("-*" * int(len(self.map_list[0])/2), end='')
        row_index = 0
        for row in self.map_list:
            if self.location[0] == row_index:
                copy = ''.join(row)
                copy = copy[:self.location[1]] + "@" + copy[self.location[1]+1:]
                print(''.join(copy))
            else:
                print(''.join(row))
            row_index += 1
        print("Output: ", "".join(self.output))
        print("Steps: ", self.steps)
        


map_str = ""
f = open("Inputs/2017_19.txt", "r")
for line in f:
    map_str += line
print(map_str)


tm = TextMap(map_str)
tm.walk_map(False)


