# -*- coding: utf-8 -*-
"""
--- Day 13: Packet Scanners ---



"""

'''
class FireWall(list):

    current_step = 0
    count = 0

    def add_layer(self, index, size):

        if len(self) > index:
            self[index] = size
        else:
            while len(self) < index:
                self.append(0)
            self.append(size)

    def increment(self):

        
'''
    
def counts(walls, offset):
    
    counts = 0
    step = 0
    for wall in walls:

        '''print("wall: ", wall)
        print("step: ", step)
        print("offset: ", offset)
        print((step + offset) % ((2*wall)-2) == 0)
        input("----------------")'''

        if wall != 0:
            if (step + offset) % ((2*wall)-2) == 0:        
                counts += 1


        step += 1
    
    '''for step in range(len(walls)-1):

        if walls[step] != 0:
            if (step + offset) % ((2*walls[step] - 2)) == 0:

                counts += 1'''

    return counts
    
def process_line(line):

    return int(line.split(": ")[0]), int(line.split(": ")[1][:-1])


walls = []

f = open("Inputs/2017_13.txt", "r")
for line in f:
    wall = process_line(line)
    while len(walls) < wall[0]:
        walls.append(0)
    walls.append(wall[1])

f.close()

#walls = [2,0,2]

offset = -1
results = -1
while results != 0:

    offset += 1
    results = counts(walls, offset)
    if results == 0:
        print(results, offset)

    if offset > 200000001: break



