import {readFileSync} from 'fs'


function part1(lines: string[]){
    const sum = lines
        .map(l => l.split('').filter(c => c >= '0' && c <= '9'))
        .map(d => parseInt(d[0] + d[d.length-1]))
        .reduce((acc, cur) => acc += cur, 0)
    console.log('part 1:', sum)
}

function part2(lines: string[]){
    const wordToNums: Record<string,string> = {'one':'1','two':'2','three':'3','four':'4','five':'5','six':'6','seven':'7','eight':'8','nine':'9'}
    
    const mapLine = (line: string) => {
        const nums = Array.from(line.matchAll(/(?=(one|two|three|four|five|six|seven|eight|nine|\d))/g))
            .flat()
            .filter(v => v !== '')
            .map(v => wordToNums[v] ?? v)
        return parseInt(nums[0] + nums[nums.length - 1])
    }
    
    const sum = lines.map(mapLine).reduce((acc, cur) => acc += cur, 0)
    console.log('part 2:', sum)
}

const lines = readFileSync(process.argv[2]).toString().trim().split("\n");
part1(lines)
part2(lines)