
from collections import Counter
from typing import Iterator, List, Tuple

class School:

    def __init__(self, fish: str):

        self._school: List[int] = [
            int(fishy) for fishy
            in fish.split(',')
        ]

    def tick(self):

        ready_to_spawn: int = 0
        has_spawned: int = 6
        baby_fish: int = 8

        new_fish: List[int] = [
            baby_fish for fish
            in self._school
            if fish == ready_to_spawn
        ]

        self._school = [
            has_spawned
            if fish == ready_to_spawn
            else fish - 1
            for fish in self._school
        ]

        self._school.extend(new_fish)

    def simulate(self, ticks: int):

        for _ in range(ticks):
            self.tick()

        return len(self._school)


class FastSchool:

    def __init__(self, fish: str):

        self._school = {i:0 for i in range(9)}
        for fishy in fish.split(','):
            self._school[int(fishy)] += 1

    def tick(self):

        ready_to_spawn: int = 0
        has_spawned: int = 6
        new_fish: int = 8

        will_spawn: int = self._school[ready_to_spawn]
        #shift all the items down one day in the school
        for i in range(8):
            self._school[i] = self._school[i+1]

        self._school[has_spawned] += will_spawn
        self._school[new_fish] = will_spawn

    def simulate(self, ticks: int):

        for _ in range(ticks):
            self.tick()

        return sum(self._school.values())




def part1(fish: List[int]) -> int:

    return School(fish).simulate(ticks=80)
        

def part2(fish: List[str]) -> int:
    
    return FastSchool(fish).simulate(ticks=256)


with open('input') as f:
    fish = f.read().strip()


print(part1(fish)) #
print(part2(fish)) #

