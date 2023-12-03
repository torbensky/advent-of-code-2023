from collections import defaultdict
import sys

def part1(lines: list[str]):
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
    
    result = 0
    for x,y in nums:
        v = nums[(x,y)]
        # search for symbols around it
        for i in range(-1, len(str(v))+1):
            if symbols[(x + i, y - 1)]:
                result+=v
            elif symbols[(x + i, y)]:
                result+=v
            elif symbols[(x + i, y + 1)]:
                result+=v
    print('part 1:', result)


if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = [line.rstrip() for line in file]  
    part1(lines)