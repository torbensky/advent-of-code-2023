import { readFileSync } from "fs";

const NODE_RE = /(\w+)\W+(\w+), (\w+)/
type Move = 'L'|'R'
interface Network {
    [key: string]: {
        L: string
        R: string
    }
}

function navigate(starts: string[], target: string, moves: Move[], network: Network){
    let steps = 0
    while(!starts.every(n => n.endsWith(target))){
        const m = moves[steps % moves.length]
        starts.forEach((n, i) => {
            starts[i] = network[n][m]
        })
        steps += 1
    }
    return steps
}

function part1(moves: Move[], network: Network){
    const answer = navigate(['AAA'],'ZZZ', moves, network)
    console.log(`Part 1: ${answer}`)
}

function parseLines(lines: string[]){
    const network: Network = {}
    lines.splice(2).forEach(l => {
        const matches = NODE_RE.exec(l)!
        network[matches[1]] = {L: matches[2], R: matches[3]}
    })
    return {moves: lines[0].split('') as Move[], network}
}

const lines = readFileSync(process.argv[2]).toString().split('\n')
const {moves, network} = parseLines(lines)
part1(moves, network)