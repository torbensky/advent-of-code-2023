import { readFileSync } from "fs"
import { Mapping, AlmanacMap, range } from "./types"

function parseSeeds(line: string): number[] {
    return line.split(/\s+/).slice(1).map(s => parseInt(s))
}

function parseMap(line: string): Mapping {
    const [dest, source, length] = line.split(' ').map(v => parseInt(v))
    return new Mapping(source,dest,length)
}


function parseMaps(parts: string[]){
    return parts.map(p => {
        const [label, ...mappings] = p.split('\n')
        return new AlmanacMap(label, mappings.map(parseMap))
    })
}

function part1(parts: string[]){
    const seeds = parseSeeds(parts[0])
    const almanacMaps = parseMaps(parts.slice(1))
    let lowest = Number.MAX_SAFE_INTEGER
    seeds.forEach(seed => {
        let location = seed
        almanacMaps.forEach(am => {
            location = am.map(location)
        })        
        lowest = location < lowest ? location : lowest
    })

    console.log(`Part 1: ${lowest}`)
}

function part2(parts: string[]){
    const seeds = parseSeeds(parts[0])
    const almanacMaps = parseMaps(parts.slice(1))
    
    let lowest = Number.MAX_SAFE_INTEGER
    for(let i = 0; i < seeds.length; i += 2){
        let ranges: range[] = [[seeds[i], seeds[i] + seeds[i+1] - 1]]
        for(let am of almanacMaps){
            ranges = ranges.map(r => am.mapRange(r)).flat()
        }
        ranges.forEach(r => {
            lowest = lowest > r[0] ? r[0] : lowest
        })
    }

    console.log(`Part 2: ${lowest}`)
}

const parts = readFileSync(process.argv[2]).toString().trim().split("\n\n")
part1(parts)
part2(parts)