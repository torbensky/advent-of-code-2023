import sys
import re

def part1(lines: list[str]):
    result = 0
    for line in lines:
        digits = [c for c in line if c.isnumeric()]
        result += int(digits[0] + digits[-1])
    print('part 1:', result)

digitnames = ['one','two','three','four','five','six','seven','eight','nine']
def translate(v: str) -> str:
    return v if v.isdigit() else str(digitnames.index(v)+1)

def part2(lines: list[str]):
    result = 0
    for line in lines:
        nums = re.findall(r'(?=(one|two|three|four|five|six|seven|eight|nine|\d))', line)
        result += int(translate(nums[0]) + translate(nums[-1]))
    print('part 2:', result)

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = [line.rstrip() for line in file]        
    part1(lines)
    part2(lines)