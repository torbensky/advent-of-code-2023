
import copy
import sys

def solve(expansion_factor, galaxies, expand_rows, expand_cols):
    for (i,row) in enumerate(expand_rows):
        for g in galaxies:
            if g[1] > (row + i * expansion_factor):
                g[1] += expansion_factor
    for (i,col) in enumerate(expand_cols):
        for g in galaxies:
            if g[0] > col + i * expansion_factor:
                g[0] += expansion_factor
    answer = 0
    for (i, g1) in enumerate(galaxies):
        for g2 in galaxies[i+1:]:
            answer += abs(g1[0]-g2[0]) + abs(g1[1] - g2[1])
    return answer

if __name__ == '__main__':
    with open(sys.argv[1]) as f:
        data = f.read()
    grid = [list(line) for line in data.splitlines()]

    galaxies = []
    for (y,row) in enumerate(grid):
        galaxies.extend([[x,y] for (x,c) in enumerate(row) if c == '#'])

    expand_rows = []
    for i in range(0, len(grid)-1):
        if '#' not in grid[i]:
            expand_rows.append(i)

    expand_cols = []
    grid = list(zip(*grid))
    for i in range(0, len(grid)-1):
        if '#' not in grid[i]:
            expand_cols.append(i)
    
    print(f"Part 1: {solve(1,copy.deepcopy(galaxies),expand_rows,expand_cols)}")
    print(f"Part 2: {solve(1000000-1,copy.deepcopy(galaxies),expand_rows,expand_cols)}")
    