import { readFileSync } from "fs";

interface Card {num: number, winning: string[], have: string[]}

function parseLine(line: string, num: number): Card{
    // Example: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
    let [left, right] = line.split(': ')[1].split(' | ')
    const winning = left.split(' ').map(v => v.trim()).filter(v => v)
    const have = right.split(' ').map(v => v.trim()).filter(v => v)
    
    return {num, winning, have}
}

function part1(lines: string[]){
    const result = lines.map(parseLine).reduce((acc, cur) => {
        const numWon = cur.have.filter(h => cur.winning.includes(h)).length
        return acc + (numWon > 0 ? 2 ** (numWon - 1) : 0)
    }, 0)

    console.log(`Part 1: ${result}`)
}

const lines = readFileSync(process.argv[2]).toString().trim().split("\n")
part1(lines)