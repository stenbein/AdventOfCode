#!/usr/bin/env python3

'''started cleaning the 2017 files up. I noticed that part 1 
wasn't saved along with part 2 in the day 2 solution. Rewrote
it to include both and switched it to read input from file. I'm
toying with the idea of rewriting all of them to take input from
stdin. Not sure yet if that's the direction I want to go.'''

def part_one(row):

    numbs = [int(x) for x in row.split()]

    return max(numbs) - min(numbs)


def part_two(row):

    numbs = [int(x) for x in row.split()]

    for num in numbs:

        for divisor in numbs:

            if num % divisor == 0 and num != divisor:

                result = num / divisor
                return result



with open('data.txt', 'r') as file:
	lines = file.readlines()

# First part
print(sum([part_one(row) for row in lines]))

# second part
print(sum([part_two(row) for row in lines]))