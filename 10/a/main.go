package main

import (
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "regexp"
    "strconv"
    "strings"
)

var cycles = []int{1}

func noop() {
    cycles = append(cycles, cycles[len(cycles) - 1])
}

func addx(x int) {
    cycles = append(cycles, cycles[len(cycles) - 1] + x)
}

func printCycles() {
    for idx, cycle := range cycles {
        log.Printf("cycle %d: %d", idx, cycle)
    }
}

func main() {
    input := utils.ReadFile("input")
    lines := strings.Split(input, "\n")

    for _, instruction := range lines {
        noopr := regexp.MustCompile(`^noop`)
        addxr := regexp.MustCompile(`^addx\s`)

        if noopr.MatchString(instruction) {
            noop()
            continue
        }

        incString := addxr.ReplaceAllString(instruction, "")
        inc, err := strconv.Atoi(incString)
        if err != nil {
            log.Fatalf("unable to convert `%s` to int", incString)
        }
        noop()
        addx(inc)
    }

    printCycles()

    relevantCycles := []int{20, 60, 100, 140, 180, 220}
    var sumSignal int
    for _, cycleIdx := range relevantCycles {
        log.Printf("cycle %d signal strength: %d", cycleIdx, cycleIdx * cycles[cycleIdx])
        sumSignal += cycleIdx * cycles[cycleIdx]
    }

    log.Println(sumSignal)
}