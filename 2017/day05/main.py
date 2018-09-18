#!/usr/bin/python3
'''Day 5 of the 2017 advent of code'''


def process_commands(cmds, program_counter):
    """helper for part 2"""

    jmp_offset = cmds[program_counter]

    if jmp_offset > 2:
        cmds[program_counter] -= 1
    else:
        cmds[program_counter] += 1

    return jmp_offset + program_counter


def part_one(rows):
    """Return the answer to part one of this day"""

    cmds = [int(cmd) for cmd in rows]

    count = 0
    prg_counter = 0

    while True:

        try:
            offset = cmds[prg_counter]
            cmds[prg_counter] += 1
            prg_counter = prg_counter + offset
            count += 1
        except IndexError:
            break

    return count


def part_two(rows):
    """Return the answer to part two of this day"""

    cmds = [int(cmd) for cmd in rows]

    count = 0
    next_counter = 0

    while True:

        try:
            next_counter = process_commands(cmds, next_counter)
            count += 1
        except IndexError:
            break

    return count


if __name__ == "__main__":

    with open('input', 'r') as file:
        ROWS = file.readlines()

    print("Part 1: {}".format(part_one(ROWS)))
    print("Part 2: {}".format(part_two(ROWS)))

