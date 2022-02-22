
from dataclasses import dataclass
from typing import Iterator, List, Tuple



@dataclass(eq=True, frozen=True)
class Vent:
    x: int
    y: int


class VentLine:

    def __init__(self, raw):

        #raw 189,59 -> 512,382
        raw = raw.split(' -> ')
        start, end = raw[0].split(','), raw[1].split(',')
        self.start = Vent(int(start[0]), int(start[1]))
        self.end = Vent(int(end[0]), int(end[1]))


    def is_vertical(self):

        return self.start.x == self.end.x

    def is_horizontal(self):

        return self.start.y == self.end.y

    def is_diagonal(self):

        return not self.is_horizontal and not self.is_vertical

    def walk(self):

        return points_between(self.start, self.end)


class VentField:

    def __init__(self, raw):

        self.vents = [
            VentLine(line)
            for line in raw
        ]

    def straight_lines(self):
        return [
            ventline
            for ventline in self.vents
            if ventline.is_vertical() or ventline.is_horizontal()
        ]

    def count_crossed(self, diag=False):

        coverage = {}
        ventlines = self.vents if diag else self.straight_lines()

        for ventline in ventlines:
            for vent in ventline.walk():
                coverage[vent] = coverage.get(vent, 0) + 1

        return sum([1 for v in coverage.values() if v > 1])


def points_between(v1: Vent, v2: Vent) -> List[Vent]:

    if v1.x == v2.x:
        #vertical case
        if v1.y <= v2.y:
            return [Vent(v1.x, y) for y in range(v1.y, v2.y+1)] #+1 to cover last vent
        else:
            return [Vent(v1.x, y) for y in range(v2.y, v1.y+1)]

    if v1.y == v2.y:
        #vertical case
        if v1.x <= v2.x:
            return [Vent(x, v1.y) for x in range(v1.x, v2.x+1)]
        else:
            return [Vent(x, v1.y) for x in range(v2.x, v1.x+1)]

    if v1.x <= v2.x:
        if v1.y <= v2.y:
            return [Vent(v1.x+i, v1.y+i) for i in range((v2.x-v1.x)+1)]
        return [Vent(v1.x+i, v1.y-i) for i in range((v2.x-v1.x)+1)]
    else:
        if v1.y <= v2.y:
            return [Vent(v1.x-i, v1.y+i) for i in range((v1.x-v2.x)+1)]
        return [Vent(v1.x-i, v1.y-i) for i in range((v1.x-v2.x)+1)]



def part1(vents: List[str]) -> int:

    field = VentField(vents)
    return field.count_crossed()
        

def part2(vents: List[str]) -> int:
    
    field = VentField(vents)
    return field.count_crossed(diag=True)


with open('input') as f:
    vents = f.read().strip().split('\n')
    

print(part1(vents)) #6856
print(part2(vents)) #20666
