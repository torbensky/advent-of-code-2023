
import sys

def diffgen(history: list[int]):
    prev = history[0]
    for v in history[1:]:
        yield v - prev
        prev = v
    if len(history) == 1:
        yield 0

def extrapolate(below, above):
    return above + below

def part1(data: list[list[int]]):
    lastlasts = []
    for history in data:
        lasts = [history[-1]]
        while True:
            history = list(diffgen(history))
            lasts.append(history[-1])
            if history.count(0) == len(history):
                break
        v = 0
        for v2 in lasts[::-1]:
            v = extrapolate(v,v2)
        lastlasts.append(v)
    print(f"Part 1: {sum(lastlasts)}")

if __name__ == '__main__':
    with open(sys.argv[1]) as f:
        lines = f.readlines()
    data = [[int(v) for v in line.split()] for line in lines]
    for tc in [(0,3,3),(3,15,18),(0,1,1),(1,6,7),(7,21,28),(0,2,2),(2,6,8),(8,15,23),(23,45,68),(-1,1,0), (-1,-1,-2), (1,-1,0), (-3,1,-2), (0,2,2), (4,-2,2), (2,-2,0), (1,-1,0), (-8,3,-5), (-4,4,0), (-3,4,1), (0,0,0), (4,-4,0)]:
        got = extrapolate(*tc[0:2])
        assert got == tc[2], f"wanted {tc[2]}, got {got} for {tc}"
    part1(data)