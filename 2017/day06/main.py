#!/usr/bin/python3
'''Day 6 of the 2017 advent of code'''


def redistribute(memory):
    '''helper to redistribute the memory'''

    size = len(memory)

    max_index = 0
    max_value = memory[0] #always assumed to be memory
    #find max and index of max
    for i in range(size):
        if memory[i] > max_value:
            max_value = memory[i]
            max_index = i

    #reset memory at max
    memory[max_index] = 0
    next_block = 1
    while max_value:

        memory[(max_index + next_block) % size] += 1

        next_block += 1
        max_value -= 1

    return memory



def part_one(data):
    """Return the answer to part one of this day"""

    states = {}
    count = 0
    while True:

        state = str(data)

        if state not in states:
            states[state] = 1
            count += 1
        else:
            if states[state] == 2:
                break
            else:
                states[state] += 1
        data = redistribute(data)

    return count


def part_two(data):
    """Return the answer to part two of this day"""

    states = {}
    count = 0
    while True:

        state = str(data)

        if state not in states:
            states[state] = 1
        else:
            if states[state] == 2:
                break
            else:
                states[state] += 1
                count += 1
        data = redistribute(data)

    return count


if __name__ == "__main__":

    with open('input', 'r') as file:
        ROWS = file.readlines()

    DATA = [int(numb) for numb in ROWS[0].split()]

    print("Part 1: {}".format(part_one(DATA)))
    print("Part 2: {}".format(part_two(DATA)))
