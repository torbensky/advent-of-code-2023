import re, sys, operator
from functools import reduce

linere = re.compile(r'Game (\d+): (.*)')
roundre = re.compile(r'(\d+) (red|blue|green)')

class Game:
    def __init__(self, line):
        matched = linere.match(line).groups()
        self.gameid = int(matched[0])
        rounds = matched[1].split(';')
        self.rounds = [parseRound(r) for r in rounds]
    
    def max(self, color: str) -> int:
        return max(*[v[color] for v in self.rounds if v.get(color, 0)], 0)
    
    def maxall(self) -> tuple[int,int,int]:
        return (self.max('red'),self.max('green'),self.max('blue'))

def parseLine(line: str) -> Game:
    return Game(line)

def parseRound(round: str) -> dict[str,int]:
    d = {}
    for r in roundre.findall(round):
        d[r[1]] = int(r[0])
    return d


def part1(lines: list[str]):
    maxr,maxg,maxb = 12,13,14
    result = 0
    for line in lines:
        g = parseLine(line)
        if g.max('red') > maxr or g.max('blue') > maxb or g.max('green') > maxg:
            continue
        result += g.gameid
        
    print('part 1:', result)

def part2(lines: list[str]):
    result = 0
    for line in lines:
        g = parseLine(line)
        result += reduce(operator.mul, g.maxall())
    print ('part 2:', result)

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = [line.rstrip() for line in file]  
    part1(lines)
    part2(lines)