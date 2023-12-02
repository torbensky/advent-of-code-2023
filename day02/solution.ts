import { readFileSync } from "fs"

interface Round {
    r: number
    g: number
    b: number
}

interface Game {
    id: number
    rounds: Round[]
}

const gameIdRx = /Game (\d+): (.*)/
function parseGame(line: string): Game {
    const [,id, rounds] = gameIdRx.exec(line)!
    return {id: parseInt(id),rounds: rounds.split(';').map(r => parseRound(r.trim()))}
}

function parseRound(round: string){
    const r = {r: 0, b: 0, g: 0}
    round.split(',')
        .map(pair => pair.trim().split(' '))
        .forEach(([num, color]) => {
            switch(color.trim()){
                case 'red':
                    r.r = parseInt(num)
                    break
                case 'green':
                    r.g = parseInt(num)
                    break
                case 'blue':
                    r.b = parseInt(num)
                    break
            }
        })
    return r
}

function part1(lines: string[]){
    const [maxR,maxG,maxB] = [12,13,14]
    let result = lines.map(parseGame)
        .filter(g => !g.rounds.find(r => r.r > maxR || r.g > maxG || r.b > maxB))
        .reduce((acc,cur) => acc + cur.id, 0)
    console.log('part 1:', result)
}

function part2(lines: string[]){
    const result = lines.map(parseGame)
        .map(g =>  g.rounds.reduce((acc, cur) => {
                acc.r = Math.max(cur.r, acc.r)
                acc.g = Math.max(cur.g, acc.g)
                acc.b = Math.max(cur.b, acc.b)
                return acc
            }, {r: 0, g: 0, b: 0})
        )
        .reduce((acc, cur) => {
            return acc + cur.r * cur.g * cur.b
        }, 0)
    console.log('part 2:', result)
}

const lines = readFileSync(process.argv[2]).toString().trim().split('\n')
part1(lines)
part2(lines)