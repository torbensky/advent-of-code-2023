from collections import defaultdict
import sys

Coord = tuple[int,int]

def parsegrid(lines: list[str]) -> tuple[dict[Coord,int],defaultdict[Coord,str]]:
    nums = dict()
    symbols = defaultdict(str)
    for y, line in enumerate(lines):
        num = ''
        for x, c in enumerate(line):
            if c.isdigit():
                num = num + c
            elif num:
                nums[(x-len(num),y)] = int(num)
                num = ''
            
            if c != '.' and not c.isdigit():
                symbols[(x,y)] = c
        if num:
            nums[(len(line)-len(num),y)] = int(num)
            num = ''
    return (nums,symbols)

def findparts(nums: dict[Coord,int], symbols: defaultdict[Coord,str]) -> list[int]:
    parts = []
    for x,y in nums:
        v = nums[(x,y)]
        # search for symbols around it
        for i in range(-1, len(str(v))+1):
            if symbols[(x + i, y - 1)]:
                parts.append(v)
            elif symbols[(x + i, y)]:
                parts.append(v)
            elif symbols[(x + i, y + 1)]:
                parts.append(v)
    return parts
    

def collides(p: Coord, min: Coord, max: Coord) -> bool:
    return p[0] >= min[0] and p[1] >= min[1] and p[0] <= max[0] and p[1] <= max[1]

def part1(lines: list[str]):
    nums,symbols = parsegrid(lines)
    parts = findparts(nums,symbols)
    print('part 1:', sum(parts))

def part2(lines: list[str]):
    result = 0
    nums,symbols = parsegrid(lines)
    parts = findparts(nums,symbols)
    parts = {k: v for k, v in nums.items() if v in parts}
    symbols = {k:v for k,v in symbols.items() if v == '*'}
    for spos in symbols:
        adjacent = []
        for ppos in parts:
            partnum = nums[ppos]
            if collides(spos, (ppos[0]-1,ppos[1]-1), (ppos[0]+len(str(partnum)), ppos[1] + 1)):
                adjacent.append(partnum)
        if len(adjacent) == 2:
            result += adjacent[0] * adjacent[1]
                
    print('part 2:', result)

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = [line.rstrip() for line in file]  
    part1(lines)
    part2(lines)