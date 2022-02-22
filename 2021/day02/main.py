
from typing import Iterator, List, Tuple


def part1(commands: List[Tuple[str, int]]) -> int:

    horz_pos: int = 0
    depth: int = 0

    for command, unit in commands:
        if command == 'forward':
            horz_pos += unit
        if command == 'down':
            depth += unit
        if command == 'up':
            depth -= unit

    return horz_pos * depth
        

def part2(commands: List[Tuple[str, int]]) -> int:

    horz_pos: int = 0
    depth: int = 0
    aim: int = 0

    for command, unit in commands:
        if command == 'forward':
            depth += aim * unit
            horz_pos += unit
        if command == 'down':
            aim += unit
        if command == 'up':
            aim -= unit

    return horz_pos * depth



with open('input') as f:
    lines: list[Tuple[str, int]] = [
        (part[0], int(part[1]))
        for part in
        [
            line.split(' ')
            for line
            in f.readlines()
        ]
    ]

print(part1(lines)) #1561344
print(part2(lines)) #1848454425
