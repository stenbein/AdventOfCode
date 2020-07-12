#!/usr/bin/python3
'''Day 9 of the 2017 advent of code'''


def process_garbage(stream, index):
    """Traverse stream. Break on '>' as end of garbage,
    return total as size of garbage and the new index
    """

    total = 0
    length = len(stream)
    while index < length:

        if stream[index] == ">":
            break

        elif stream[index] == "!":
            index += 1 #skip one character

        elif stream[index] == "<":
            pass #ignore

        else:
            total += 1
        
        index += 1

    return total, index


def process_stream(stream, index, rank=0):
    """if we hit garbage switch to the helper function
    else recurse down the stream incrementing the rank
    for each sub group

    return the total garbage, and group count
    """

    score = 0
    garbage = 0

    length = len(stream)
    while index < length:

        if stream[index] == "<":
            new_garbage, index = process_garbage(stream, index+1)
            garbage += new_garbage

        elif stream[index] == "{":
            new_garbage, new_score, index = process_stream(stream, index+1, rank+1)
            garbage += new_garbage
            score += new_score

        elif stream[index] == "}":
            break

        elif stream[index] == "!":
            index += 1 #skip one character

        index += 1

    return garbage, score+rank, index


def part_one(data):
    """Return the answer to part one of this day"""

    return process_stream(data, 0)[1] #sum of score


def part_two(data):
    """Return the answer to part two of this day"""

    return process_stream(data, 0)[0] #sum of non cancelled garbage


if __name__ == "__main__":

    DATA = ""
    with open("input", "r") as f:
        for line in f:
            DATA += line.rstrip() #hidden newline in file input

    print("Part 1: {}".format(part_one(DATA)))
    print("Part 2: {}".format(part_two(DATA)))
