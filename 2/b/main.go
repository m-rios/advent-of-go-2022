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

    outcomeToShape := map[string]map[string]string {
        "X": {
            "A": "C",
            "B": "A",
            "C": "B",
        },
        "Y": {
            "A": "A",
            "B": "B",
            "C": "C",
        },
        "Z": {
            "A": "B",
            "B": "C",
            "C": "A",
        },
    }

    outcomeToPoints := map[string]int {
        "X": LOSE,
        "Y": DRAW,
        "Z": WIN,
    }

    shapesValue := map[string]int {
        "A": 1,
        "B": 2,
        "C": 3,
    }

    outcomeName := map[string]string {
        "X": "loose",
        "Y": "draw",
        "Z": "win",
    }

    shapeName := map[string]string {
        "A": "rock",
        "B": "paper",
        "C": "scissors",
    }

    score := 0
    for _, round := range strategy {
        shapes := strings.Split(round, " ")
        log.Printf("To %s against %s i have to play %s", outcomeName[shapes[1]], shapeName[shapes[0]], shapeName[outcomeToShape[shapes[1]][shapes[0]]] )
        score += shapesValue[outcomeToShape[shapes[1]][shapes[0]]] + outcomeToPoints[shapes[1]]
    }
    log.Println(score)
}