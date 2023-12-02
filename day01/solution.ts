import {readFileSync} from 'fs'


function part1(lines: string[]){
    const sum = lines
        .map(l => l.split('').filter(c => c >= '0' && c <= '9'))
        .map(d => parseInt(d[0] + d[d.length-1]))
        .reduce((acc, cur) => acc += cur, 0)
    console.log(sum)
}

const lines = readFileSync(process.argv[2]).toString().split("\n");
part1(lines)