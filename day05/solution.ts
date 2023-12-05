import { readFileSync } from "fs"

function parseSeeds(line: string): number[] {
    return line.split(/\s+/).slice(1).map(s => parseInt(s))
}

function parseMap(line: string): Mapping {
    const [dest, source, length] = line.split(' ').map(v => parseInt(v))
    return new Mapping(source,dest,length)
}

class AlmanacMap {
    name: string
    mappings: Mapping[]
    constructor(name: string, mappings: Mapping[]){
        this.name = name
        this.mappings = mappings
    }
    map(source: number): number {
        for(const m of this.mappings){
            const v = m.map(source)
            if(v !== undefined){
                return v
            }
        }
        return source
    }
}

class Mapping {
    source: number
    dest: number
    length: number
    constructor(source: number,dest: number,length: number){
        this.source = source
        this.dest = dest
        this.length = length
    }

    map(source: number): number | undefined {
        const diff = source - this.source
        if(diff >= 0 && diff < this.length){
            return this.dest + diff
        }        
    }
}

function followMap(seed: number, almanac: AlmanacMap[]): number {
    let value = seed
    almanac.forEach(am => {
        value = am.map(value)
    })
    return value
}

function parseMaps(parts: string[]){
    return parts.map(p => {
        const [label, ...mappings] = p.split('\n')
        return new AlmanacMap(label, mappings.map(parseMap))
    })
}

function part1(lines: string[]){
    const seeds = parseSeeds(lines[0])
    const almanacMaps = parseMaps(lines.slice(1))
    let lowest = Number.MAX_SAFE_INTEGER
    seeds.forEach(seed => {
        const location = followMap(seed,almanacMaps)
        lowest = location < lowest ? location : lowest
    })

    console.log(`Part 1: ${lowest}`)
}

const parts = readFileSync(process.argv[2]).toString().trim().split("\n\n")
part1(parts)