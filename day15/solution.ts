import { readFileSync } from "fs";

// 1. convert each char to ASCII int value
// 2. increase the **current value** by the ASCII code
// 3. **current value** *= 17
// 4. **current value** %= 256
const answer = readFileSync(process.argv[2]).toString().trim().split(',')
.map(s => s.split('')
    .map(c => c.charCodeAt(0))
    .reduce((cv, v) => (cv + v) * 17 % 256,0)
).reduce((acc,cur) => acc+cur,0)
console.log(`Part 1: ${answer}`)