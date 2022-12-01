package main

import (
    utils "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "strconv"
    "strings"
)

func updateCaloryCount(elfCalories []int, currentHighestCaloryCount int, currentHighestCaloryIndex int) (int, int) {
    lastElfCaloriesIndex := len(elfCalories) - 1
    lastElfCalories := elfCalories[lastElfCaloriesIndex]
    highestCaloryCount := currentHighestCaloryCount
    highestCaloryIndex := currentHighestCaloryIndex
    if lastElfCalories > currentHighestCaloryCount {
        highestCaloryCount = lastElfCalories
        highestCaloryIndex = lastElfCaloriesIndex
    }

    return highestCaloryCount, highestCaloryIndex
}

func main() {
    inventory := strings.Split(utils.ReadFile("input"), "\n")

    elfCalories := make([]int, 1)
    highestCaloryCount := 0
    var highestCaloryIndex int

    for _, caloryEntry := range inventory {
        if caloryEntry == "" {
            highestCaloryCount, highestCaloryIndex = updateCaloryCount(elfCalories, highestCaloryCount, highestCaloryIndex)
            elfCalories = append(elfCalories, 0)
        } else {
            numericCalory, err := strconv.Atoi(caloryEntry)
            if err != nil {
                log.Fatal(err)
            }
            elfCalories[len(elfCalories) - 1] += numericCalory
        }
    }

    highestCaloryCount, highestCaloryIndex = updateCaloryCount(elfCalories, highestCaloryCount, highestCaloryIndex)

    log.Printf("There are %d elfs", len(elfCalories))

    log.Printf("Elf number %d had the most calories, with a total of %d calories", highestCaloryIndex + 1, highestCaloryCount)
}