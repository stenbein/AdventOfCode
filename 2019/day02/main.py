#!/usr/bin/python3
'''Day 2 of the 2019 advent of code

Debug notes:

    1. Tried using enum which I had recently seen in a code review

        class Opcode(Enum):
            ADD = 1
            MUL = 2
            HLT = 99

    Surpised to see
        >> 99 == Opcode.HLT
        >> False

    But it makes sense as
        >> Opcode(99) == Opcode.HLT
        >> True
        >> Opcode(7)
        >> raise

    2. Failed to final instruction on part 1 to set the special values

    3. Used an asign instead of a copy in part 2.

'''


from enum import Enum


def _input_help(raw):
    '''Whatever additional steps we need with the raw inputs'''
    return [int(r) for r in raw.split(',')]


class Opcode(Enum):
    ADD = 1
    MUL = 2
    HLT = 99

class Intcode:
    '''The puzzle text suggests we'll need this again later.
    If so I'll improve it then. This will do for now.'''

    def __init__(self, program):
        self.__hard_copy = program
        self._program = program
        self._counter = 0

    def reset(self):
        '''copy not assign'''
        self._program = self.__hard_copy[:]
        self._counter = 0

    def _next(self):
        code = self._program[self._counter]
        self._counter += 1
        return code

    def fetch(self, address):
        return self._program[address]

    def store(self, address, val):
        self._program[address] = val

    def _eval(self, opcode):

        if Opcode(opcode) == Opcode.ADD:
            val1 = self.fetch(self._next())
            val2 = self.fetch(self._next())
            self.store(self._next(), val1 + val2)

        if Opcode(opcode) == Opcode.MUL:
            val1 = self.fetch(self._next())
            val2 = self.fetch(self._next())
            self.store(self._next(), val1 * val2)

    def run(self):

        code = self._next()
        while Opcode(code) != Opcode.HLT:
            self._eval(code)
            code = self._next()


def seek(target, intcode):
    '''Quick and dirty scan through the input space.
    We'll probably need something fancier later.
    '''

    for noun in range(100):
        for verb in range(100):

            intcode.reset()

            intcode.store(address=1, val=noun)
            intcode.store(address=2, val=verb)

            intcode.run()

            if intcode.fetch(0) == target:
                return 100 * noun + verb

    raise ValueError("Not found")


def part_one(raw):
    '''Calculate and format the result of part 1'''

    intcode = Intcode(program=_input_help(raw))

    #special inputs for part 1
    intcode.store(address=1, val=12)
    intcode.store(address=2, val=2)

    intcode.run()

    return intcode.fetch(0)


def part_two(raw):
    '''Calculate and format the result of part 2'''

    intcode = Intcode(program=_input_help(raw))

    return seek(target=19690720, intcode=intcode)


if __name__ == "__main__":

    with open("input", "r") as file:

        RAW = file.read().rstrip()

    print(f'Part 1: {part_one(RAW)}')
    print(f'Part 2: {part_two(RAW)}')
