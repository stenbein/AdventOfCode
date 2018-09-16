#!/usr/bin/python3
'''Day 1 of the 2017 advent of code'''

def filter_on(seq, offset):
    """taverse the sequence of digits in the captcha,
    if the current digit is equal to the digit at the
    offset, keep the digit for summation. The offset
    can wrap"""

    to_keep = []
    arg_len = len(seq)

    if arg_len < 2:
        ValueError("Invalid series")

    for i in range(arg_len):

        if seq[i] == seq[int((i + offset) % arg_len)]:
            to_keep.append(int(seq[i]))


    return to_keep


def part_one(raw):
    """Return the answer to part one of this day
        Expectation is that the answer is in the
        correct format"""

    return sum(filter_on(raw, 1))


def part_two(raw):
    """Return the answer to part two of this day
        Expectation is that the answer is in the
        correct format"""

    offset = len(raw) / 2
    return sum(filter_on(raw, offset))


if __name__ == "__main__":

    with open("input", "r") as file:

        RAW = file.read().rstrip()
    
    print("Part 1: {}".format(part_one(RAW)))
    print("Part 2: {}".format(part_two(RAW)))
