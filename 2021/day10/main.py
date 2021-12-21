

from typing import Iterator, List, Optional, Tuple


_LEGAL_PAIRS = {
    '(': ')',
    '[': ']',
    '{': '}',
    '<': '>'
}


def first_corrupted(line: str) -> Optional[str]:

    stack = []
    for char in line:
        if char in _LEGAL_PAIRS:
            stack.append(_LEGAL_PAIRS[char]) #push what we want to match
            continue

        if char != stack.pop():
            return char

    return None


def autocomplete(line: str) -> str:
    '''
    Assume only incomplete lines are passed
    whatever is left in the stack is the missing
    elements to complete the chunk
    '''
    stack = []
    for char in line:
        if char in _LEGAL_PAIRS:
            stack.append(_LEGAL_PAIRS[char]) #push what we want to match
            continue

        stack.pop()

    return ''.join(stack[::-1]) #reversed


def score_completed(line: str) -> int:

    total: int = 0
    for char in line:

        total *= 5
        match char:
            case ')': total += 1
            case ']': total += 2
            case '}': total += 3
            case '>': total += 4
            case _: raise ValueError(f'Invalid char: {char}')

    return total


def part1(subsystem: List[str]) -> int:

    points = {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137
    }

    score: int = 0
    for line in subsystem:
        if (corr_char := first_corrupted(line)):
            score += points[corr_char]

    return score


def part2(subsystem: List[str]) -> int:

    incomplete = [
        line for line in subsystem
        if first_corrupted(line) is None
    ]

    completed = [autocomplete(line) for line in incomplete]
    scores = sorted([score_completed(line) for line in completed])

    return scores[(len(scores) // 2)] #always an odd number, but indexed at 0


with open('input') as f:
    subsystem: List[str] = f.read().strip().split('\n')

print(part1(subsystem)) #311895
print(part2(subsystem)) #2904180541
