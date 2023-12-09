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

function part2(data: number[][]){
    // We can take advantage of an interesting property of the number pyramid to
    // solve part 2 and simply reverse the rows to get the answer.
    // 
    // In part 1:
    // For a row of numbers v1, v2... vn, the last number of the next row is
    // given by vn - v(n-1) and then the next number for that row is v(n+1) = (vn - v(n-1) + vn).
    // 
    // In part 2: 
    // For the same row of numbers, the first number of the next row is given by
    // (v2 - v1) and then the previous number for that row is v0 = v1 - (v2 - v1)
    // When we reverse, v0 = v(n+1), v1 = vn, v2 = v(n-1) which means
    // v0 = v1 - (v2 - v1)
    // = v(n+1) = vn - (v(n-1) - vn)
    // = vn - v(n-1) + vn <--------- gee, doesn't that look familiar!? ;)
    data.forEach(v => v.reverse())
    console.log(`Part 2: ${solve(data)}`)
}

const lines = readFileSync(process.argv[2]).toString().split('\n')
const data = parseInput(lines)
part1(data)
part2(data)