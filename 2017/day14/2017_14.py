# -*- coding: utf-8 -*-
"""
--- Day 14: Disk Defragmentation ---

Cut and pasted the knot hash script from day 10. When I first tackled this
problem I omitted some details from the day 10 implementation and ran into
errors. It was the intent of the contest designer to use the prior method
including the inputs and padding. This implementation returned hashes as
strings. This leads to problems.

The first challenge asks that we return the number of used memory cells. In
this representation the used cells are the 1's in the binary representation of
the hash as a string. So the answer is just the count of the occurance of 1's
in each string for each row. But in part two we needed to treat the memory banks
as a traversible grid to get the neighbours of each cell. Since the
representation was a string the leading zeros were truncated. ie. '000101010...'
becomes '101010...'. Secondly, I was imagining an array of digits when really I
have a string. So grid[i][j] = region won't work. To fix this I cast to an
array and left everything else as it was.

"""


class TwistList(object):

    def __init__(self, size):

        self.size = size
        self.list = [x for x in range(size)]
        self.current_index = 0
        self.skip_size = 0

    def twist(self, length):
        
        list_slice = []

        for i in range(length):

            mod_index = self.current_index + i
            list_slice.append(self.list[mod_index % self.size])

        list_slice = list(reversed(list_slice))

        for i in range(length):

            mod_index = self.current_index + i
            self.list[mod_index % self.size] = list_slice[i]


        self.current_index += length + self.skip_size
        self.skip_size += 1


def hash_reduce(ascii_array):

    if len(ascii_array) % 16 != 0:
        raise ValueError("Array size not equally divisible by 16.")
    
    output = []
    for i in range(0,int(len(ascii_array) / 16)):

        val = 0
        for numb in ascii_array[i*16:i*16+16]:
            val ^= numb
        output.append(val)

    return output

def to_hex_str(dec_array):

    output = ""
    for dec in dec_array:
        val = str(hex(dec)[2:])
        if len(val) == 1:
            val = '0' + val
        output += val
    
    return output

#the binary string was chopping off leading zeros
def to_bin_str(hex_str, size):

    output = bin(int(hex_str, 16))[2:]
    while len(output) < size:
        output = '0' + output
    
    return output

def knot_hash(str_input):

    PADDING = [17, 31, 73, 47, 23] #padding is carried over from problem 10
    
    tl = TwistList(256)
    line = [ord(char) for char in str_input]
    [line.append(numb) for numb in PADDING]
    for itter in range(64):
        for length in line:
            tl.twist(int(length))
    return to_hex_str(hash_reduce(tl.list))


def grid_generate(input_key):

    key_list = [input_key + "-" + str(i) for i in range(128)]
    hash_list = [knot_hash(key) for key in key_list]
    binary_list = [to_bin_str(hash_value, 128) for hash_value in hash_list]

    return binary_list

def solve_part_1():
    
    binary_list = grid_generate("amgozmfv")
        
    count = 0
    for row in binary_list:
        for char in row:
            if char == '1':
                
                count += 1
                
    print("Total number of active cells: {}".format(count))


def paint_regions(i, j, region, grid, region_map):

    if grid[i][j] == 1:
        grid[i][j] = region
        region_map[(i,j)] = region

        if i > 0:
            paint_regions(i-1, j, region, grid, region_map)
        if i < len(grid)-1:
            paint_regions(i+1, j, region, grid, region_map)
        if j > 0:
            paint_regions(i, j-1, region, grid, region_map)
        if j < len(grid)-1:
            paint_regions(i, j+1, region, grid, region_map)

    else:
        return
    

def solve_part_2():

    grid = grid_generate("amgozmfv")
    grid = [[int(char) for char in row] for row in grid]
    region_map = {}
    count = 1 #regions start at number 2, we'll share the data structure

    size = len(grid)
    
    for i in range(size):
        for j in range(size):
            if grid[i][j] == 1:
                if (i,j) not in region_map:

                    #new region
                    count += 1
                    region_map[(i,j)] = count

                    paint_regions(i, j, count, grid, region_map)

    print("Total number of regions: {}".format(count-1)) #offset for starting at 2 for region indexing
    
solve_part_1()
solve_part_2()



