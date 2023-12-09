import { readFileSync } from "fs";

function parseInput(lines: string[]): number[][] {
    return lines.map(l => l.split(' ').map(v => parseInt(v)))
}

function solveRow(data: number[]){
    const vals:number[] = [data[data.length-1]]
    do{
        let prev = data[0]
        data = data.slice(1).map(next => {
            const tmp = next - prev
            prev = next
            return tmp
        })
        vals.push(data.length > 0 ? data[data.length-1] : 0)

    }while(!data.every(v => v === 0))
    return vals.reduce((a,c) => a + c, 0)
}

function solve(data: number[][]){
    return data.map(solveRow).reduce((acc, cur) => acc+cur, 0)
}

function part1(data: number[][]){
    console.log(`Part 1: ${solve(data)}`)
}

const lines = readFileSync(process.argv[2]).toString().split('\n')
part1(parseInput(lines))