package main

import (
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "strings"
)

const WIN = 6
const DRAW = 3
const LOSE = 0

const ROCK = 1
const PAPER = 2
const SCISSORS = 3

func main() {
	strategy := strings.Split(utils.ReadFile("input"), "\n")

    outcomes := map[string]map[string]int{
        "A": { // rock
            "X": DRAW, // rock
            "Y": WIN,  // paper
            "Z": LOSE, // scissors
        },
        "B": { // paper
            "X": LOSE, // rock
            "Y": DRAW, // paper
            "Z": WIN,  // scissors
        },
        "C": { // scissors
            "X": WIN,  // rock
            "Y": LOSE, // paper
            "Z": DRAW, // scissors
        },
    }

    shapesValue := map[string]int {
        "A": 1,
        "B": 2,
        "C": 3,
        "X": 1,
        "Y": 2,
        "Z": 3,
    }

    score := 0
    for _, round := range strategy {
        shapes := strings.Split(round, " ")
        score += outcomes[shapes[0]][shapes[1]] + shapesValue[shapes[1]]
    }
    log.Println(score)
}