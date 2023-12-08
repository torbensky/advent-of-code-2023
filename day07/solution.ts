import { readFileSync } from "fs";

const DEFAULT_CARD_ORDER = ['2','3','4','5','6','7','8','9','T','J','Q','K','A']
type Round = [string,number]
const rankCard = (c: string) => DEFAULT_CARD_ORDER.indexOf(c)

function rankHandType(hand: string): number {
    const counts = new Map<string,number>()
    for(let i = 0; i < hand.length; i++){
        const card = hand.charAt(i)
        counts.set(card, (counts.get(card) ?? 0) + 1)
    }
    const numGroups = counts.size
    const orderedCounts = Array.from(counts).sort((a,b) => b[1] - a[1])

    switch(numGroups){
        case 1:
            // five of a kind (highest)
            return 7
        case 2:
            if(orderedCounts[0][1] === 4){
                // four of a kind
                return 6
            }
            // full house
            return 5
        case 3:
            if(orderedCounts[0][1] === 3 && orderedCounts[1][1] == 1){
                // three of a kind
                return 4
            }
            // two pair
            return 3
        case 4:
            // one pair
            return 2
        default:
            // high card (lowest)
            return 1;
    }
}

function compareRounds(a: Round, b: Round): number {
    // first order by hand type rank
    const [atr, btr] = [rankHandType(a[0]), rankHandType(b[0])]
    if(atr !== btr){
        return atr - btr
    }

    // tie breaker is card rank
    for(let i = 0; i < a[0].length; i++){
        const [aRank, bRank] = [rankCard(a[0].charAt(i)), rankCard(b[0].charAt(i))]
        if(aRank !== bRank){
            return aRank - bRank
        }
    }
    
    return 0
}

function parseLines(lines: string[]): Round[]{
    return lines.map(l => {
        let [hand, bid] = l.split(' ')
        return [hand, parseInt(bid)]
    })
}

function part1(rounds: Round[]){
    rounds = rounds.sort(compareRounds)
    const answer = rounds.reduce((acc, cur, i) => {
        return acc + (i+1) * cur[1]
    }, 0)
    console.log(`Part 1: ${answer}`)
}

const lines = readFileSync(process.argv[2]).toString().split('\n')
const rounds = parseLines(lines)
part1(rounds)