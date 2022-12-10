package main

import (
    "fmt"
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "math"
    "regexp"
    "strconv"
    "strings"
)

var cycles = []int{1}

func noop() {
    drawPixel()
    cycles = append(cycles, cycles[len(cycles) - 1])
}

func addx(x int) {
    noop()
    drawPixel()
    cycles = append(cycles, cycles[len(cycles) - 1] + x)
}

func drawPixel() {
    cycleCount := len(cycles)
    crtPos := (cycleCount - 1) % 40
    spritePos := cycles[len(cycles) - 1]

    // Jump to a new line on cycle 41, 51, 61... but not on cycle 1
    if cycleCount % 40 == 1 && cycleCount != 1 {
        fmt.Print("\n")
    }

    if math.Abs(float64(crtPos) - float64(spritePos)) < 2 {
        fmt.Print("# ")
    } else {
        fmt.Print(". ")
    }
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
        addx(inc)
    }
}