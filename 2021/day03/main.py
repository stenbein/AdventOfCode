
from collections import Counter
from typing import Iterator, List, Tuple


def transpose(l: List[List]) -> List[List]:

    return list(zip(*l))


def filter_by_pos(l: List, val: str, pos: int) -> List[List[str]]:

    return list(filter(lambda x: x[pos] == val, l))


def bit_critera_o2(l: List[str]) -> int:

    mc = Counter(l).most_common()
    if len(mc) > 1 and mc[0][1] == mc[1][1]:
        return '1'
    return mc[0][0]


def bit_critera_co2(l: List[str]) -> int:

    mc = Counter(l).most_common()
    if len(mc) > 1 and mc[0][1] == mc[1][1]:
        return '0'
    return '0' if mc[0][0] == '1' else '1'


def part1(reports: List[List[str]]) -> int:

    most_common = [
         Counter(v).most_common(1)[0][0]
         for v in transpose(reports)
    ]
    gama_str = "".join(most_common)
    epsilon_str = ''.join(map(lambda x: '1' if x == '0' else '0', gama_str))

    gama = int('0b'+gama_str, 2)
    epsilon = int('0b'+epsilon_str, 2)

    return gama * epsilon
        

def part2(reports: List[List[str]]) -> int:

    i, filtered = 0, reports[:]
    while len(filtered) > 1:
        transposed = transpose(filtered)
        bit = bit_critera_o2(transposed[i])
        filtered = filter_by_pos(filtered, bit, i)
        i += 1

    os_str = ''.join(filtered[0])
    o2 = int('0b'+os_str, 2)

    i, filtered = 0, reports[:]
    while len(filtered) > 1:
        transposed = transpose(filtered)
        bit = bit_critera_co2(transposed[i])
        filtered = filter_by_pos(filtered, bit, i)
        i += 1

    co2_str = ''.join(filtered[0])
    co2 = int('0b'+co2_str, 2)

    return o2 * co2


with open('input') as f:
    lines: List[List[str]] = [
        [char for char in line.strip()] #remove \n
        for line
        in f.readlines()
    ]

print(part1(lines)) #3969000
print(part2(lines)) #4267809
