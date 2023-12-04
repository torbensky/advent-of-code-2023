import sys

def parseLine(line: str):
    info = line.split(': ')[1]
    winning, have = info.split(' | ')
    winning = winning.split()
    have = have.split()
    return (set(winning), set(have))

def part1(lines: list[str]):
    result = 0
    for line in lines:
        winning, have = parseLine(line)
        won = winning.intersection(have)
        if won:
            points = 2 ** (len(won)-1)
            result += points
    print('Part 1:', result)

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = [line.rstrip() for line in file]  
    part1(lines)