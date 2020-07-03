#!/usr/bin/python3
'''Day 10 of the 2017 advent of code'''

#from pdb import set_trace as bp #use pb() #as breakpoint


#number 1
class TwistList():
    """A helper class to hold the list buffer and simplify the twist method"""

    def __init__(self, size):

        self.size = size
        self.list = [i for i in range(size)]
        self.current_index = 0
        self.skip_size = 0

    def twist(self, length):
        """twist step of the knot hash"""

        segment = []

        #grab all the items in the list from
        #our current location until the end of length
        mod_start = self.current_index % self.size
        mod_end = (self.current_index + length) % self.size

        #if we wrapped around to the beginning
        if mod_end <= mod_start:
            segment += self.list[mod_start:]
            segment += self.list[:mod_end]
        else:
            segment += self.list[mod_start:mod_end]

        segment = list(reversed(segment))

        #replace the items in our buffer
        for i in range(length):

            mod_index = self.current_index + i
            self.list[mod_index % self.size] = segment[i]

        self.current_index += length + self.skip_size
        self.skip_size += 1


#number 2
def hash_reduce(ascii_array):
    """Convert sparse hash to dense hash"""

    if len(ascii_array) % 16 != 0:
        raise ValueError("Array size not equally divisible by 16.")

    output = []
    for i in range(0, int(len(ascii_array) / 16)):

        val = 0
        for numb in ascii_array[i*16:i*16+16]:
            val ^= numb
        output.append(val)

    return output


def to_hex_str(dec_array):
    """Helper to format the array as hex string"""

    as_hex = [f'{i:02x}' for i in dec_array]

    return ''.join(as_hex)


def part_one(data):
    """Return the answer to part one of this day"""

    tl = TwistList(256)

    for length in data.split(","):
        tl.twist(int(length))

    return tl.list[0] * tl.list[1]


def part_two(data):
    """Return the answer to part two of this day"""

    tl = TwistList(256)

    ascii_codes = [ord(char) for char in data]
    ascii_codes += [17, 31, 73, 47, 23] #add padding

    for _ in range(64):
        for length in ascii_codes:
            tl.twist(int(length))

    return to_hex_str(hash_reduce(tl.list))


if __name__ == "__main__":

    DATA = ""
    with open("input", "r") as f:
        for line in f:
            DATA += line.rstrip() #hidden newline in file input

    print("Part 1: {}".format(part_one(DATA)))
    print("Part 2: {}".format(part_two(DATA)))
