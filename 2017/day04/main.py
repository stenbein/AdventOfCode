#!/usr/bin/python3
'''Day 4 of the 2017 advent of code'''


def is_anagram(phrase1, phrase2):
    '''inputs are sorted strings'''

    if len(phrase1) != len(phrase2):
        return False

    for i in range(len(phrase1)):
        if phrase1[i] != phrase2[i]:
            return False

    return True


def part_one(rows):
    """Return the answer to part one of this day"""

    count = 0
    for phrase in rows:

        if not phrase.rstrip():
            continue

        tokens = phrase.split()
        if len(tokens) == len(set(tokens)):
            count += 1

    return count


def part_two(rows):
    """Return the answer to part two of this day"""

    count = 0
    for phrase in rows:

        invalid = False
        tokens = phrase.split()
        tokens = [''.join(sorted(token)) for token in tokens]
        token_count = len(tokens)
        
        for i in range(token_count-1):
            for j in range(i+1, token_count):
                if is_anagram(tokens[i], tokens[j]):
                    invalid = True
                    break
            if invalid:
                break

        if not invalid:
            count += 1

    return count


if __name__ == "__main__":

    with open('input', 'r') as file:
        ROWS = file.readlines()

    print("Part 1: {}".format(part_one(ROWS)))
    print("Part 2: {}".format(part_two(ROWS)))

