import { readFileSync } from "fs";

interface Coord {
    x: number
    y: number
}
interface GridNode<T>{
    value: T
    pos: Coord
}
class Grid<T> {
    data = new Map<string,GridNode<T>>()
    add(pos: Coord, val: T){
        this.data.set(`${pos.x},${pos.y}`, {value: val, pos})
    }
    get(pos:Coord){
        this.data.get(`${pos.x},${pos.y}`)
    }
    forEach(cb: (v: GridNode<T>, k: string) => void){
        this.data.forEach(cb)
    }
}

const boundary = (c: Coord, width: number) => ({min: {x: c.x-1, y:c.y-1}, max: {x: c.x + width, y: c.y+1}})
const within = (c: Coord, bounds: {min: Coord, max: Coord}) => c.x >= bounds.min.x && c.y >= bounds.min.y && c.x <= bounds.max.x && c.y <= bounds.max.y

const inRange = (n: number, min: number, max: number) => n >= min && n <= max
const isDigit = (ch: string) => inRange(ch.charCodeAt(0) - '0'.charCodeAt(0), 0, 9)

function parseGrid(lines: string[]){
    const nums = new Grid<number>()
    const symbols = new Grid<string>()
    lines.forEach((line, y) => {
        let numStr = '', numX = -1
        for(let x = 0; x < line.length; x++){
            const ch = line.charAt(x)
            if(ch === '.'){
                // empty
            }else if(isDigit(ch)){
                numStr += ch
                numX = numX === -1 ? x : numX
                if(x < line.length - 1){
                    continue
                }
            }else{
                symbols.add({x,y}, ch)
            }
            
            if(numStr.length > 0){
                nums.add({x: numX,y}, parseInt(numStr))
                numStr = ''
                numX = -1
            }
        }
    })
    return {nums, symbols}
}

function part1(lines: string[]){
    const {nums, symbols} = parseGrid(lines)
    const result = findParts(nums, symbols).reduce((acc,cur) => acc + cur, 0)
    console.log(`Part 1: ${result}`)
}

function findParts(nums: Grid<number>, symbols: Grid<string>) {
    const parts: number[] = [];
    nums.forEach((n) => {
        symbols.forEach(s => {
            if (within(s.pos, boundary(n.pos, n.value.toString().length))) {
                parts.push(n.value);
            }
        });
    });
    return parts
}

function part2(lines: string[]){
    let result = 0
    const {nums, symbols} = parseGrid(lines)
    const parts = findParts(nums, symbols)
    symbols.forEach(s => {
        const adjacent: number[] = []
        nums.forEach(n => {
            if(parts.includes(n.value)){
                if(within(s.pos, boundary(n.pos, n.value.toString().length))){
                    adjacent.push(n.value)
                }
            }
        })
        if(adjacent.length === 2){
            result += adjacent[0] * adjacent[1]
        }
    })
    console.log(`Part 2: ${result}`)
}

const lines = readFileSync(process.argv[2]).toString().trim().split('\n')
part1(lines)
part2(lines)