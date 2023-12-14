
import sys

def score(grid: list[list[str]]):
    return sum([(len(grid) - i) * row.count('O') for (i, row) in enumerate(grid)])

def fall(grid):
    for col in range(len(grid[0])):
        anchor = 0
        for row in range(len(grid)):
            c = grid[row][col]
            if c == '.':
                continue
            if c == 'O':
                if anchor != row:
                    grid[anchor][col] = 'O'
                    grid[row][col] = '.'
                anchor += 1
            if c == '#':
                anchor = row + 1

def cycle(grid):
    for i in range(4):
        fall(grid)
        grid = [list(e) for e in zip(*grid[::-1])]
    return grid

def part1(grid):
    fall(grid)
    print(score(grid))

def part2(grid):
    cache = dict()
    i = 0
    target = 1000000000
    skipped = False
    while i < target:
        grid = cycle(grid)
        i += 1
        key = "".join(["".join(r) for r in grid])
        if not skipped:
            if key in cache:
                # We saw this before! Cycle detected!
                period = i - cache[key]
                # FAST FORWAAAARD!
                i = target - ((target - i) % period)
                skipped = True
            else:
                cache[key] = i
    print(f'Part 2: {score(grid)}')
    
if __name__ == '__main__':
    with open(sys.argv[1]) as f:
        data = f.read()
    grid = [list(line) for line in data.split('\n')]
    part1([g.copy() for g in grid])
    part2(grid)