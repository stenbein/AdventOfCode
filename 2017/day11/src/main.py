#!/usr/bin/python3
'''Day 11 of the 2017 advent of code'''


class HexCounter():
    '''A hex maze walker object for
    keeping track of our position'''

    def __init__(self):

        self.x = 0
        self.y = 0
        self.z = 0

        self.furthest = 0

    def move(self, direction):
        '''map the direction to a state change'''

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
        if temp > self.furthest:
            self.furthest = temp

    def max(self):
        '''accounting for negative distance along the grid'''

        total = 0
        maxx = abs(self.x)
        maxy = abs(self.y)
        maxz = abs(self.z)

        total = abs(max(maxx, maxy, maxz))

        return total


def part_one(data):
    """Return the answer to part one of this day"""

    hexer = HexCounter()
    for coord in data:

        hexer.move(coord)

    return hexer.max()


def part_two(data):
    """Return the answer to part two of this day"""

    hexer = HexCounter()
    for coord in data:

        hexer.move(coord)

    return hexer.furthest


if __name__ == "__main__":

    DATA = ""
    with open("input", "r") as f:
        for line in f:
            DATA += line.rstrip() #hidden newline in file input

    COORDS = DATA.split(",")

    print("Part 1: {}".format(part_one(COORDS)))
    print("Part 2: {}".format(part_two(COORDS)))
