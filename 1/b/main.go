package main

import (
    utils "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "sort"
    "strconv"
    "strings"
)

func main() {
    inventory := strings.Split(utils.ReadFile("input"), "\n")

    elfCalories := make([]int, 1)

    for _, caloryEntry := range inventory {
        if caloryEntry == "" {
            elfCalories = append(elfCalories, 0)
        } else {
            numericCalory, err := strconv.Atoi(caloryEntry)
            if err != nil {
                log.Fatal(err)
            }
            elfCalories[len(elfCalories) - 1] += numericCalory
        }
    }

    sort.Ints(elfCalories)

    highest3Count := 0
    for i := 0; i < 3; i++ {
        highest3Count += elfCalories[len(elfCalories) - 1 - i]
    }

    log.Println(highest3Count)
}