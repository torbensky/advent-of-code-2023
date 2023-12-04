import sys

Card = tuple[int, set[str], set[str]]

def parseLine(line: str) -> Card:
    cardno, info = line.split(': ')
    cardno = int(cardno.split()[1])
    winning, have = info.split(' | ')
    return (cardno, set(winning.split()), set(have.split()))

def part2(lines: list[str]):
    result = 0
    cards = []
    for line in lines:
        cards.append(parseLine(line))
    
    # cache the set of cards that each card produces so we don't have to
    # compute this at every iteration
    mapping = {}
    for cardnum, card in enumerate(cards):
        won = card[1].intersection(card[2])
        mapping[cardnum + 1] = []
        for i in range(1,len(won)+1):
            mapping[cardnum+1].append(cards[cardnum + i])
    
    while len(cards) > 0:
        c = cards.pop()
        result += 1
        cards += mapping[c[0]]
    print('Part 2:', result)



def part1(lines: list[str]):
    result = 0
    for line in lines:
        card = parseLine(line)
        won = card[1].intersection(card[2])
        if won:
            points = 2 ** (len(won)-1)
            result += points
    print('Part 1:', result)

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = [line.rstrip() for line in file]  
    part1(lines)
    part2(lines)