#!/usr/bin/python3
'''Day 1 of the 2019 advent of code'''


def _input_help(raw):
    '''Whatever additional steps we need with the raw inputs'''
    return [int(r) for r in raw.split()]


def _required_fuel(mass):
    '''integer div'''

    fuel = (mass // 3) - 2

    if fuel > 0:
        return fuel

    return 0


def _tyranny_fuel(mass):
    '''Damn you Tsiolkovsky!'''

    mass_total = 0
    mass_new = _required_fuel(mass)
    while mass_new > 0:
        mass_total += mass_new
        mass_new = _required_fuel(mass_new)

    return mass_total


def part_one(raw):
    '''Calculate and format the result of part 1'''

    modules = _input_help(raw)

    return sum(map(_required_fuel, modules))


def part_two(raw):
    '''Calculate and format the result of part 2'''

    modules = _input_help(raw)

    return sum(map(_tyranny_fuel, modules))


if __name__ == "__main__":

    with open("input", "r") as file:

        RAW = file.read().rstrip()
    
    print(f'Part 1: {part_one(RAW)}')
    print(f'Part 2: {part_two(RAW)}')
