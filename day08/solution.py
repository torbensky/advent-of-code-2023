
import sys, re

nodere = re.compile(r'(\w+) = \((\w+), (\w+)\)')
network = dict[str,tuple[str,str]]

def parselines(lines: list[str]) -> tuple[list[str], network]:
    # R's and L's, like RLRR...
    movements = [c for c in lines[0]]
    # nodes are formatted like "CDA = (ACF, SDX)"
    # they are all 3 letters
    nodes = [nodere.match(line).groups() for line in lines[2:]]
    nodes = {node[0]: node[1:] for node in nodes}
    return (movements, nodes)

def navigate(node: str, target: str, movements: list[str], nodes: network):
    steps, moves = 0, len(movements)

    while node != target:
        i = steps % moves
        # move to the next node
        m = movements[i]
        node = nodes[node][0] if m == 'L' else nodes[node][1]
        steps += 1
    return steps


def part1(movements: list[str], nodes: network):    
    # IMPORTANT: this only reliably works in Python 3.7+
    # https://docs.python.org/3.7/library/stdtypes.html#dict.values
    steps = navigate('AAA', 'ZZZ', movements, nodes)
    
    print(f"Part 1: {steps}")
        

if __name__ == '__main__':
    with open(sys.argv[1]) as f:
        lines = f.read().splitlines()
    part1(*parselines(lines))
    