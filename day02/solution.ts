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
    console.log(result)
}

const lines = readFileSync(process.argv[2]).toString().trim().split('\n')
part1(lines)