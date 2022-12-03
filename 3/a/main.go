package main

import (
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "strings"
)

func charValue(c byte) int {
    asciiValue := int(c)

    if asciiValue < 97 {
        return asciiValue - (65 - 27)
    } else {
        return asciiValue - 96
    }
}

func main() {
    input := utils.ReadFile("input")
    rucksacks := strings.Split(input, "\n")

    var rucksackPoints = make([]int, len(rucksacks))
    totalPoints := 0

    for rId, rucksack := range rucksacks {
        // split rucksack into the two compartments
        compartmentA := map[byte]bool {}
        compartmentB := map[byte]bool {}
        for c := 0; c < len(rucksack); c++ {
            if c < len(rucksack) / 2 {
                compartmentA[rucksack[c]] = true
            } else {
                compartmentB[rucksack[c]] = true
            }
        }

        // find all characters that appear in compartment A and B
        for k := range compartmentA {
            if compartmentB[k] {
                kValue := charValue(k)
                log.Printf("type %s value %d", string(k), kValue)
                rucksackPoints[rId] += kValue
                totalPoints += kValue
            }
        }
    }

    log.Println(totalPoints)
}