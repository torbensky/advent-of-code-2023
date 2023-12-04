import { readFileSync } from "fs";

interface Card {num: number, winning: string[], have: string[]}

const countWins = (c: Card) => c.have.filter(h => c.winning.includes(h)).length

function parseLine(line: string, num: number): Card{
    // Example: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
    let [left, right] = line.split(': ')[1].split(' | ')
    const winning = left.split(' ').map(v => v.trim()).filter(v => v)
    const have = right.split(' ').map(v => v.trim()).filter(v => v)
    
    return {num, winning, have}
}

function part2(lines: string[]){
    const cards = lines.map(parseLine)

    // fill a cache of the cards that are won by each card
    const cardWinCache = new Map<number, Card[]>()
    cards.map(countWins).forEach((numWon, cardNum) => cardWinCache.set(cardNum, cards.slice(cardNum+1, cardNum+numWon+1)))
    let result = 0
    while(cards.length > 0){
        result++
        const c = cards.pop()!
        cards.push(...cardWinCache.get(c.num)!)
    }
    console.log(`Part 2: ${result}`)
}

function part1(lines: string[]){
    const result = lines.map(parseLine).reduce((acc, cur) => {
        const numWon = countWins(cur)
        return acc + (numWon > 0 ? 2 ** (numWon - 1) : 0)
    }, 0)

    console.log(`Part 1: ${result}`)
}

const lines = readFileSync(process.argv[2]).toString().trim().split("\n")
part1(lines)
part2(lines)