
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

def part1(grid):
    fall(grid)
    print(score(grid))
   
if __name__ == '__main__':
    with open(sys.argv[1]) as f:
        data = f.read()
    grid = [list(line) for line in data.split('\n')]
    part1([g.copy() for g in grid])