import { readFileSync } from "fs";

const NODE_RE = /(\w+)\W+(\w+), (\w+)/
type Move = 'L'|'R'
interface Network {
    [key: string]: {
        L: string
        R: string
    }
}

// gcd,lcm,lcmAll take from:
// https://stackoverflow.com/questions/31302054/how-to-find-the-least-common-multiple-of-a-range-of-numbers
const gcd = (a:number, b:number): number => b == 0 ? a : gcd(b, a % b)
const lcm = (a:number, b:number) =>  a / gcd(a, b) * b
const lcmAll = (ns:number[]) => ns.reduce(lcm, 1)

function navigate(nodes: string[], target: string, moves: Move[], network: Network){
    let steps = 0
    const stepsToCycle:number[] = []
    while(nodes.length > 0){
        const m = moves[steps % moves.length]
        nodes.forEach((n, i) => {
            nodes[i] = network[n][m]
        })
        steps += 1
        // We know that all of start nodes repeat a cycle of a fixed length.
        // So we remove any nodes that reached their destination as there is
        // nothing new to be learned
        const next = nodes.filter(s => !s.endsWith(target))
        if(next.length != nodes.length){
            stepsToCycle.push(steps)
        }
        nodes = next
    }
    // The total number of steps for everything is when all the cycles sync up
    // so that each node is at the target.
    // That is the lowest common multiple!
    return lcmAll(stepsToCycle)
}

function part1(moves: Move[], network: Network){
    const answer = navigate(['AAA'],'ZZZ', moves, network)
    console.log(`Part 1: ${answer}`)
}

function part2(moves: Move[], network: Network){
    const starts = Object.keys(network).filter(n => n.endsWith('A'))
    const answer = navigate(starts, 'Z', moves, network)
    console.log(`Part 2: ${answer}`)
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
part2(moves, network)