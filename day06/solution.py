import sys
from operator import mul
from functools import reduce

def wins_possible(time: int, record: int) -> int:
    winCount = 0
    for holdTime in range (1,time):
        speed = holdTime
        traveled = speed * (time - holdTime)
        if traveled > record:
            winCount += 1
    return winCount

def part1(times: list[int], distances: list[int]):
    results = [wins_possible(t,d) for t, d in zip(times, distances)]
    print(results)
    print("Part 1:", reduce(mul, results, 1))

if __name__ == '__main__':
    # expect 2 lines where the first token is a label, the rest are numbers
    with open(sys.argv[1]) as file:
        times = [int(n) for n in file.readline().split()[1:]]
        distances = [int(n) for n in file.readline().split()[1:]]
    part1(times, distances)