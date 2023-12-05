import { AlmanacMap, Mapping } from "./types"

describe("mappings work", () => {
    describe("single mapping", () => {
        const mapping = new Mapping(4, 8, 4)
        it("handles a range outside, to the left properly", () => {
            const result = mapping.mapRange([0,3])
            expect(result).toHaveLength(3)
            expect(result).toEqual([[0,3],,])
        })
        it("handles a range outside to the right properly", () => {
            const result = mapping.mapRange([9,12])
            expect(result).toHaveLength(3)
            expect(result).toEqual([,,[9,12]])
        })
        it("handles a range completely inside properly", () => {
            let result = mapping.mapRange([4,5])
            expect(result).toHaveLength(3)
            expect(result).toEqual([,[8,9],])

            result = mapping.mapRange([6,7])
            expect(result).toHaveLength(3)
            expect(result).toEqual([,[10,11],])

            result = mapping.mapRange([4,7])
            expect(result).toHaveLength(3)
            expect(result).toEqual([,[8,11],])
        })
        it('handles a range overlapping on both sides properly', () => {
            let result = mapping.mapRange([0, 20])
            expect(result).toHaveLength(3)
            expect(result[0]).toEqual([0,3])
            expect(result[1]).toEqual([8,11])
            expect(result[2]).toEqual([8,20])
        })
    })
    describe("multiple mappings", () => {
        const am = new AlmanacMap('testmap', [
            new Mapping(10, 20, 5),
            new Mapping(30,40,5)
        ])
        it('handles a range to the left of everything properyl', () => {
            let result = am.mapRange([0,0])
            expect(result).toEqual([[0,0]])

            result = am.mapRange([9,9])
            expect(result).toEqual([[9,9]])

            result = am.mapRange([0,9])
            expect(result).toEqual([[0,9]])
        })

        it('handles a range that overlaps with the second range properly', () => {
            let result = am.mapRange([30,34])
            expect(result).toEqual([[40,44]])
        })

        it('handles a range that overlaps everything properly', () => {
            let result = am.mapRange([0,50])
            expect(result).toHaveLength(5)
            expect(result).toContainEqual([0,9])
            expect(result).toContainEqual([20,24])
            expect(result).toContainEqual([15,29])
            expect(result).toContainEqual([40,44])
            expect(result).toContainEqual([35,50])
        })
    })
})