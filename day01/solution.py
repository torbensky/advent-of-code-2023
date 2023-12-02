import sys
import re

def part1(lines: list[str]):
    result = 0
    for line in lines:
        digits = [c for c in line if c.isnumeric()]
        result += int("{}{}".format(digits[0],digits[-1]))
    print('part 1:', result)

def translate(v: str) -> str:
    match v:
        case 'one':
            return '1'
        case 'two':
            return '2'
        case 'three':
            return '3'
        case 'four':
            return '4'
        case 'five':
            return '5'
        case 'six':
            return '6'
        case 'seven':
            return '7'
        case 'eight':
            return '8'
        case 'nine':
            return '9'
        case other:
            return v

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