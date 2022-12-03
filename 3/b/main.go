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

    totalPoints := 0

    var items map[byte]int

    for id, rucksack := range rucksacks {
        if id % 3 == 0 {
            items = make(map[byte]int)
        }

        itemSet := map[byte]bool {}
        for _, c := range rucksack {
            itemSet[byte(c)] = true
        }

        for k, _ := range itemSet {
            items[k] += 1
        }

        if id % 3 == 2 {
            for k, v := range(items) {
                if v == 3 {
                    totalPoints += charValue(k)
                }
            }
        }
    }

    log.Println(totalPoints)
}