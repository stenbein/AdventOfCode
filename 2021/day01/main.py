
from typing import Iterator, List, Tuple


def part1(measurements: List[int]) -> int:

    count: int = 0
    previous: int = measurements[0]
    for measurement in measurements:
        if measurement > previous:
            count += 1
        previous = measurement

    return count


def part2(measurements: List[int]) -> int:

    windows: Iterator[Tuple[int, int, int]] = zip(
        measurements[:-2],
        measurements[1:-1],
        measurements[2:]
    )

    sums = [
        sum(window)
        for window
        in windows
    ]

    return part1(sums)


with open('input', encoding='utf-8') as f:
    lines: list[int] = [
        int(line)
        for line
        in f.readlines()
    ]

print(part1(lines)) #1713
print(part2(lines)) #1734
