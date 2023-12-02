import sys

sum = 0
with open(sys.argv[1]) as f:
    for line in f:
        digits = [c for c in line if c.isnumeric()]
        sum += int("{}{}".format(digits[0],digits[-1]))

print('part 1:', sum)