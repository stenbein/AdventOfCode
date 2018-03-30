#from pdb import set_trace as bp #use pb() #as breakpoint

input_string = '''..##.##.######...#.######
##...#...###....##.#.#.##
###.#.#.#..#.##.####.#.#.
..##.##...#..#.##.....##.
##.##...#.....#.#..#.####
.###...#.........###.####
#..##....###...#######..#
###..#.####.###.#.#......
.#....##..##...###..###.#
###.#..#.##.###.#..###...
####.#..##.#.#.#.#.#...##
##.#####.#......#.#.#.#.#
..##..####...#..#.#.####.
.####.####.####...##.#.##
#####....#...#.####.#..#.
.#..###..........#..#.#..
.#.##.#.#.##.##.#..#.#...
..##...#..#.....##.####..
..#.#...######..##..##.#.
.####.###....##...####.#.
.#####..#####....####.#..
###..#..##.#......##.###.
.########...#.#...###....
...##.#.##.#####.###.####
.....##.#.#....#..#....#.'''

input_string_t = '''..#
#..
...'''

class Cleaner(object):

    def __init__(self):
        self.position = (0,0)
        self.heading = (0,1) #north via x,y coords
        self.count = 0

    def step(self):

        self.position = (self.position[0] + self.heading[0], self.position[1] + self.heading[1])
        
    def turn_clockwise(self):

        '''self.heading = (self.heading[0] * 0 + self.heading[1] * 1,
             self.heading[0] * -1 + self.heading[1] * 0)'''
        self.heading = (self.heading[1] * 1, self.heading[0] * -1)


    def turn_c_clockwise(self):

        self.heading = (self.heading[1] * -1, self.heading[0] * 1)

    def cycle_1(self, steps_to_take):
		
        for step in range(steps_to_take):
	    #turn
            if grid.get(self.position, ".") == ".":
                self.turn_c_clockwise()
                grid[self.position] = "#"
                self.step()
                self.count += 1
			
            else:
                self.turn_clockwise()
                grid[self.position] = "."
                self.step()

    def cycle_2(self, steps_to_take):
		
        for step in range(steps_to_take):

            current = grid.get(self.position, ".")
            if current == ".":
                self.turn_c_clockwise()
                grid[self.position] = "W"
                self.step()

            elif current == "W":
                grid[self.position] = "#"
                self.step()
                self.count += 1

            elif current == "#":
                self.turn_clockwise()
                grid[self.position] = "F"
                self.step()
                
            else:
                self.turn_clockwise()
                self.turn_clockwise()
                grid[self.position] = "."
                self.step()
	
            #take action
            #move

#CLOCKWISE_TURN = [[0, -1],[1,0]]
#COUNTER_CLOCKWISE_TURN = [[0, 1],[-1,0]]

grid = {}

inputs = input_string.split("\n")
size = len(inputs) #input expected to be square
if size %2 != 1: raise ValueError("Expect input to be uneven grid.")

top = int(size / 2)

for line in inputs:

    left = -int(size / 2)
    for char in line:
		
        grid[(left, top)] = char
        left += 1

    top -= 1

    

    
#print(grid)
#print(len(grid))


cleaner = Cleaner()
cleaner.cycle_2(10000000)
print(cleaner.count)

