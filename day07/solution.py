import sys
from collections import Counter

# in order of rank, lowest -> highest
p1cardorder = ['2','3','4','5','6','7','8','9','T','J','Q','K','A']
p2cardorder = ['J','2','3','4','5','6','7','8','9','T','Q','K','A']

def typerank(hand: str, wildcards: bool) -> int:
    if(hand.count('J') == 5):
        return 7
    
    if wildcards:
        c = Counter(hand.replace('J', ''))
        hand = hand.replace('J', c.most_common()[0][0])
    c = Counter(hand)
    numgroups = len(c)

    # print(hand)
    # STRONGEST
    if numgroups == 1:
        # Five of a 
        # print('five of kind')
        return 7
    if numgroups == 2:
        if c.most_common(1)[0][1] == 4:
            # Four of a kind
            # print('four of kind')
            return 6
        # Full house
        # print('full house')
        return 5
    if numgroups == 3:
        mc = c.most_common(3)
        if mc[0][1] == 3 and mc[1][1] == 1 and mc[2][1] == 1:
            # Three of a kind
            # print('three of kind')
            return 4
        if mc[0][1] == 2 and mc[1][1] == 2 and mc[2][1] == 1:
            # Two pair
            # print('two pair')
            return 3
    if numgroups == 4:
        # One pair
        # print('one pair')
        return 2
    # High card
    # print('high card')
    return 1

def handrank(hand: str, wildcards: bool):
    order = p2cardorder if wildcards else p1cardorder
    cardranks = [order.index(c) for c in hand]
    cardranks.insert(0, typerank(hand, wildcards))
    return tuple(cardranks)


def part1(rounds: list[tuple[str,str]]):
    rounds = [(handrank(h, False),b) for (h,b) in rounds]
    rounds.sort()
    result = sum([(i+1) * int(bid) for i,(h,bid) in enumerate(rounds)])
    print(f"Part 1: {result}")

def part2(rounds: list[tuple[str,str]]):
    rounds = [(handrank(h, True),b) for (h,b) in rounds]
    rounds.sort()
    result = sum([(i+1) * int(bid) for i,(h,bid) in enumerate(rounds)])
    print(f"Part 2: {result}")

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        lines = file.readlines()
    data = [line.split() for line in lines]
    part1(data)
    part2(data)