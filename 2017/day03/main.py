#!/usr/bin/python3
'''Day 3 of the 2017 advent of code

    This problem gave me considerable trouble
    my approach to solve part 2 was similar to
    how I solved part 1 and I got bogged down
    into the details. What a mess. I paused work
    on that one and came back later on, near
    the end of the whole advent.
'''

def get_neighbours(coords):
    '''return a tuple for each of the cardnial
    directions away from each coord'''

    return [
        (coords[0], coords[1] + 1)
        , (coords[0] + 1, coords[1] + 1)
        , (coords[0] + 1, coords[1])
        , (coords[0] + 1, coords[1] - 1)
        , (coords[0], coords[1] - 1)
        , (coords[0] - 1, coords[1] - 1)
        , (coords[0] - 1, coords[1])
        , (coords[0] - 1, coords[1] + 1)
        ]

def counterclockwise(heading):
    """transforms an (x,y) coord heading
    into the next clockwise itteration"""

    #heading up, move left
    if heading == (0, 1):
        return (-1, 0)
    #heading left, move down
    elif heading == (-1, 0):
        return (0, -1)
    #heading down, move right
    elif heading == (0, -1):
        return (1, 0)
    #heading right, move up
    elif heading == (1, 0):
        return (0, 1)
    else:
        raise ValueError("Corrds.")

class RingGrid():
    '''Helper class to store the grid lookup
    and keep track of values.'''

    def __init__(self):

        self.grid = {(0, 0):1, (1, 0):1}
        self.curr_index = (1, 0)
        self.heading = (1, 0)

    def get_value(self, coord):
        '''the value at the coord of the grid'''

        if coord in self.grid:
            return self.grid[coord]
        return 0

    def get_next_index(self):
        '''next index is imagined as the same value
        pattern from problem 1, incrementing 1 by 1
        counter clockwise outward'''

        '''try to turn counter clockwise and take
        a step, if that value is 0 we didn't visit
        that grid point before, else revert to the
        prior heading'''
        heading = counterclockwise(self.heading)
        test_index = (self.curr_index[0] + heading[0]
                      , self.curr_index[1]  + heading[1])

        if self.get_value(test_index) == 0:
            self.heading = heading
            self.curr_index = test_index
        else:
            self.curr_index = (
                self.curr_index[0] + self.heading[0]
                , self.curr_index[1]+ self.heading[1]
                )

        '''get the sum of all the neighbour grid points
        surrounding this grid point'''
        value = sum([self.get_value(corrd) for
                     corrd in get_neighbours(self.curr_index)])

        '''side effect, store this value'''
        self.grid[self.curr_index] = value

        return self.curr_index


def part_one(target):
    """Return the answer to part one of this day
        Expectation is that the answer is in the
        correct format"""

    '''presume each spiral is a box, the box has a
    length and a max digit value which would be the
    length ** 2 in the bottom right corner'''
    def box_len(index):
        return 1 + 2*(index - 1)
    def box_max(index):
        return box_len(index) ** 2

    '''find the box which contains our target
    value by counting upward from 1 and checking
    if the output is less than the target. If so,
    then the box we're looking at is not the box
    containing our target. Halt when greater or
    equal to the target'''
    rank_of_box = 1
    while box_max(rank_of_box) < target:
        rank_of_box += 1

    '''get the side length of the box
    containing the target value'''
    length = box_len(rank_of_box)

    '''take the diff of the target and the max
    value of the preceeding box, mod by the length
    of the current box to find the distance from the
    center'''
    distance_to_mid = (target - box_max(rank_of_box - 1)) % length

    '''correct for approaching from the left or right of mid'''
    if distance_to_mid > (length - 1) / 2:
        distance_to_mid -= int((length - 1) / 2)

    return distance_to_mid + rank_of_box


def part_two(target):
    """Return the answer to part two of this day
        Expectation is that the answer is in the
        correct format"""

    #we need the number which is bigger then test number
    index = 3
    current = 1 #first number
    grid = RingGrid()
    while current < target:

        next_index = grid.get_next_index()
        current = grid.get_value(next_index)

        index += 1

    return current


if __name__ == "__main__":

    INPUT = 289326 #input is single variable in this problem

    print("Part 1: {}".format(part_one(INPUT)))
    print("Part 2: {}".format(part_two(INPUT)))
