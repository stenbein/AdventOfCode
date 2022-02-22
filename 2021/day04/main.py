
from typing import Iterator, List, Tuple


def transpose(l: List[List]) -> List[List]:

    return list(zip(*l))


class Board:

    def __init__(self, board_string):
        '''Example

        22 13 17 11  0
         8  2 23  4 24
        21  9 14 16  7
         6 10  3 18  5
         1 12 20 15 19
        '''
        self.numbers = [
            [
                int(val) for val
                in line.split(' ')
                if val != ''
            ] for line
            in board_string.strip().split('\n')
        ]
        self.marked = []
        self.complete = False

    def check(self):
        for row in self.numbers:
            if all([i in self.marked for i in row]):
                self.complete = True

        transposed = transpose(self.numbers)
        for row in transposed:
            if all([i in self.marked for i in row]):
                self.complete = True

    def mark(self, num: int):
        if num not in self.marked:
            self.marked.append(num)
            self.check()

    def score(self, num) -> int:

        sum_unmarked = 0
        for row in self.numbers:
            sum_unmarked += sum([
                i for i in row
                if i not in self.marked 
            ])

        return num * sum_unmarked

    def __repr__(self):

        return str(self.numbers)


class Bingo:

    def __init__(self, boards: List[str]):

        self._boards = [
            Board(board)
            for board in boards
        ]

        self.score = 0
        self.winner = None

    def call(self, num):
        for board in self._boards:
            board.mark(num)
            if board.complete:
                self.winner = board
                self.score = board.score(num)
                break

    def call_last(self, num):
        for board in self._boards:
            board.mark(num)
            if board.complete:
                self.winner = board
                self.score = board.score(num)

        self._boards = [
            board for board
            in self._boards
            if not board.complete
        ]


def part1(boards: List[str], drawn: List[int]) -> int:

    b = Bingo(boards)
    for num in drawn:
        b.call(num)
        if b.winner:
            return b.score

    raise ValueError('Expected result')
        

def part2(boards: List[str], drawn: List[int]) -> int:

    b = Bingo(boards)
    for num in drawn:
        b.call_last(num)
        if not b._boards:
            return b.score

    raise ValueError('Expected result')



with open('input') as f:
    bingo_input = f.read().split('\n\n')
    drawn = [int(num) for num in bingo_input[0].split(',')]
    boards = bingo_input[1:]

print(part1(boards, drawn)) #
print(part2(boards, drawn)) #
