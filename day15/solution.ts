import { readFileSync } from "fs";

const ops = readFileSync(process.argv[2]).toString().trim().split(',')

// 1. convert each char to ASCII int value
// 2. increase the **current value** by the ASCII code
// 3. **current value** *= 17
// 4. **current value** %= 256
const hash = (s: string) => s.split('').map(c => c.charCodeAt(0)).reduce((cv, v) => (cv + v) * 17 % 256,0)

const answer = ops.map(hash).reduce((acc,cur) => acc+cur,0)
console.log(`Part 1: ${answer}`)

interface Lens {
    label: string
    focalLength: number
}
const boxes: Lens[][] = Array(256).fill(null).map(() => [])
const printBoxes = () => boxes.forEach((b,i) => {
    if(b.length > 0){
        console.log(`Box ${i}: ${b.map(l => `[${l.label} ${l.focalLength}]`).join(' ')}`)
    }
})
ops.forEach(op => {
    let opType = '-'
    if (!op.includes('-')) {
        opType = '='
    }
    const [label,focalLength] = op.split(opType)
    const boxNum = hash(label)
    const box = boxes[boxNum]
    if(opType === '='){
        const existingLens = box.find(l => l.label === label)
        if(existingLens){
            existingLens.label = label
            existingLens.focalLength = Number(focalLength)
        }else{
            box.push({label, focalLength: Number(focalLength)})
        }
    }else{
        boxes[boxNum] = box.filter(l => l.label !== label)
    }
    // console.log(`After "${op}"`)
    // printBoxes()
})

function confirmInstallation(){
    return boxes.map((box,boxNum) => {
        return box.map((l,slotNum) => {
            return (boxNum+1) * (slotNum + 1) * l.focalLength
        }).reduce((acc,cur) => acc+cur, 0)
    }).reduce((acc,cur) => acc+cur, 0)
}

console.log(`Part 2: ${confirmInstallation()}`)