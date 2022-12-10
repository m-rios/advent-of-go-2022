package main

import (
    stack "github.com/m-rios/advent-of-go-2022"
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "regexp"
    "strconv"
    "strings"
)

func parseStacks(lines []string) (moveListStart int, stacks []*stack.PointerStack) {
    stacks = []*stack.PointerStack{}

    for idx, line := range lines {
        if len(line) > 1 && string(line[1]) == "1" {
            moveListStart = idx + 2
            break
        }

        for x := 1; x < len(line); x += 4 {
            crate := string(line[x])
            stackIndex := x / 4

            if len(stacks) <= stackIndex {
                stacks = append(stacks, new(stack.PointerStack))
            }

            if crate != " " {
                stacks[stackIndex].Push(crate)
            }
        }
    }

    for _, s := range stacks {
        s.Reverse()
    }

    return moveListStart, stacks
}

func parseMove(move string) (quant, from, to int) {
    r := regexp.MustCompile(`\s[0-9]+`)
    moveArguments := r.FindAllString(move, 3)

    if len(moveArguments) != 3 {
        log.Fatalf("unable to parse '%s': expected to have 3 arguments but only found %d", move, len(moveArguments))
    }

    args := make([]int, 3)

    for idx, arg := range moveArguments {
        var err error
        args[idx], err = strconv.Atoi(strings.Trim(arg, " "))
        if err != nil {
            log.Fatalf("unable to convert `%s` to int", arg)
        }
    }

    return args[0], args[1], args[2]
}

func main() {
    input := utils.ReadFile("input")
    diagram := strings.Split(input, "\n")

    moveListStart, stacks := parseStacks(diagram)

    // run move strategy
    for idx := moveListStart; idx < len(diagram); idx++ {
        quant, from, to := parseMove(diagram[idx])
        for x := 0; x < quant; x++ {
            crate, err := stacks[from - 1].Pop()
            if err != nil {
                log.Fatal(err)
            }
            stacks[to - 1].Push(crate)
        }
    }

    // print result
    result := make([]string, len(stacks))
    for idx, s := range stacks {
        var err error
        result[idx], err = s.Peek()
        if err != nil {
            log.Fatal(err)
        }
    }
    log.Println(strings.Join(result, ""))
}