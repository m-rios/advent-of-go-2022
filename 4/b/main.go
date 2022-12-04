package main

import (
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "strconv"
    "strings"
)

func main() {
    input := utils.ReadFile("input")
    assignements := strings.Split(input, "\n")

    fullyContained := 0

    for _, assignement := range assignements {
        pairs := strings.Split(assignement, ",")

        pair1 := make([]int, 2)
        pair2 := make([]int, 2)

        for idx, el := range strings.Split(pairs[0], "-") {
            elInt, err := strconv.Atoi(el)
            if err != nil {
                log.Fatal(err)
            }
            pair1[idx] = elInt
        }

        for idx, el := range strings.Split(pairs[1], "-") {
            elInt, err := strconv.Atoi(el)
            if err != nil {
                log.Fatal(err)
            }
            pair2[idx] = elInt
        }

        if pair1[0] <= pair2[0] && pair1[1] >= pair2[0] || pair1[0] >= pair2[0] && pair1[0] <= pair2[1] {
            log.Printf("assignement %s is overlapping", assignement)
            fullyContained += 1
        }
    }

    log.Println(fullyContained)
}