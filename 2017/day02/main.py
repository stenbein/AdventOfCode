#!/usr/bin/python3
'''Day 2 of the 2017 advent of code

 Started cleaning the 2017 files up. I noticed that part 1
wasn't saved along with part 2 in the day 2 solution. Rewrote
it to include both and switched it to read input from file. I'm
toying with the idea of rewriting all of them to take input from
stdin. Not sure yet if that's the direction I want to go.'''

def max_min_extract(row):
    """returns the difference between the max
    and min of an input list of integers"""

    return max(row) - min(row)


def divisor_extract(row):
    """returns the quotient of the only evenly
    divisible numbers in the list. This assumes
    that they exist"""
    for num in row:
        for divisor in row:

            if num % divisor == 0 and num != divisor:
                return num / divisor

    raise ValueError("Row does not have even divisor pair")


def part_one(rows):
    """Return the answer to part one of this day
        Expectation is that the answer is in the
        correct format"""

    return sum([max_min_extract(row) for row in rows])


def part_two(rows):
    """Return the answer to part two of this day
        Expectation is that the answer is in the
        correct format"""

    return sum([divisor_extract(row) for row in rows])


if __name__ == "__main__":

    with open('input', 'r') as file:
        SHEET = file.readlines()
        ROWS = [[int(x) for x in row.split()]
                for row in SHEET]

    print("Part 1: {}".format(part_one(ROWS)))
    #int is used here to truncate
    print("Part 2: {}".format(int(part_two(ROWS))))
