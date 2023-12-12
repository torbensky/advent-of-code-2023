import { readFileSync } from "fs";

function findEmptyColumns(universe: string[]){
    const cols: number[] = []
    for(let col = 0; col < universe[0].length; col++){
        let empty = true
        for(let row of universe){
            if(row[col] === '#'){
                empty = false
                break
            }
        }
        if(empty){
            cols.push(col)
        }
    }
    return cols
}
interface galaxy{ x: number; y: number; }
function solve(galaxies: galaxy[], expansionRows: number[], expansionCols: number[], eFactor: number) {
    expansionRows.forEach((row,i) => {
        galaxies = galaxies.map(g => {
            if(g.y > row+i*eFactor){
                return {...g, y: g.y + eFactor}
            }
            return g
        })
    })
    expansionCols.forEach((col,i) => {
        galaxies = galaxies.map(g => {
            if(g.x > col+i*eFactor){
                return {...g, x: g.x+eFactor}
            }
            return g
        })
    })
    let answer = 0
    galaxies.forEach((g1,i) => {
        galaxies.slice(i+1).forEach(g2 => {
            answer += Math.abs(g1.x-g2.x) + Math.abs(g1.y-g2.y)
        })
    })
    return answer
}

const rows = readFileSync(process.argv[2]).toString().split('\n')
const expansionRows = rows.reduce<number[]>((acc,cur,i) => {
    return cur.includes('#') ? acc : [...acc,i]
}, [])
const expansionCols = findEmptyColumns(rows)
const galaxies = rows.flatMap((row,y) => {
    return row.split('')
        .map((v,x) => ({v,x,y}))
        .filter(t => t.v === '#')
})

console.log(`Part 1: ${solve(galaxies, expansionRows, expansionCols, 1)}`)
console.log(`Part 2: ${solve(galaxies, expansionRows, expansionCols, 1000000-1)}`)