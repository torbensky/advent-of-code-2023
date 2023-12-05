import sys

class Mapping:
    def __init__(self, source: int, dest: int, length: int):
        self.source = source
        self.dest = dest
        self.length = length        
    
    def map(self, source: int) -> int | None:
        diff = source - self.source
        if diff >= 0 and diff < self.length:
            return self.dest + diff
        return None

class MappingList:
    def __init__(self, mappings: list[Mapping]):
        self.mappings = mappings
    
    def map(self, source: int) -> int:
        for m in self.mappings:
            v = m.map(source)
            if v is not None:
                return v        
        return source

def parse_seeds(line: str) -> list[int]:
    return [int(s) for s in line.split()[1:]]

def parse_mapping(line: str) -> Mapping:
    dest, source, length = line.split()
    return Mapping(int(source), int(dest), int(length))

def parse_mappings(mstr: str) -> MappingList:
    return MappingList([parse_mapping(m) for m in mstr.split('\n')[1:]]) # skip first line, it's text

def parse_input(parts: list[str]):
    seeds, *mappings = parts
    seeds = parse_seeds(seeds)
    return seeds, [parse_mappings(m) for m in mappings]

def part1(parts: list[str]):
    seeds, mappings = parse_input(parts)
    locations = []
    for seed in seeds:
        val = seed
        for m in mappings:
            val = m.map(val)
        locations.append(val)
    
    print('Part 1:', min(locations))
                
        

if __name__ == '__main__':
    with open(sys.argv[1]) as file:
        parts = file.read().split('\n\n')
    part1(parts)