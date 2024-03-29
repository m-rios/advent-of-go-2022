package main

import (
    "fmt"
    "github.com/m-rios/advent-of-go-2022/utils"
    "log"
    "regexp"
    "strconv"
    "strings"
)

type Element struct {
    value int
    list  []Element
}

func printElementList(list []Element) string {
    result := fmt.Sprint("[")
    for _, element := range list {
        if isList(element) {
            result += printElementList(element.list)
        } else {
            result += fmt.Sprintf("%d,", element.value)
        }
    }
    if string(result[len(result) - 1]) == "," {
        result = result[:len(result) - 1]
    }
    return result + fmt.Sprint("]")
}

func consumeCharacter(input *string) {
    *input = (*input)[1:]
}
func parseNextValue(packet *string) (value int, openList bool, closeList bool) {
    // the string representation of next int value. For multi-digit numbers each digit
    // will be appended to this string until a ] or a [ is found
    var valueStr string
    for len(*packet) > 0 {
        head := string((*packet)[0])
        switch {
        case head == "," && valueStr == "":
            consumeCharacter(packet)
            continue
        case head == "," && valueStr != "":
            value, err := strconv.Atoi(valueStr)
            if err != nil {
                log.Fatalf("error parsing next value: %v", err)
            }
            return value, false, false
        case head == "[" && valueStr == "":
            consumeCharacter(packet)
            return 0, true, false
        case head == "[" && valueStr != "":
            value, err := strconv.Atoi(valueStr)
            if err != nil {
                log.Fatalf("error parsing next value: %v", err)
            }
            return value, false, false
        case head == "]" && valueStr == "":
            consumeCharacter(packet)
            return 0, false, true
        case head == "]" && valueStr != "":
            value, err := strconv.Atoi(valueStr)
            if err != nil {
                log.Fatalf("error parsing next value: %v", err)
            }
            return value, false, false
        default:
            consumeCharacter(packet)
            valueStr += head
        }
    }
    log.Fatal("parseNextValue reached the end of the string without returning a value. This is a noop and means something is wrong with the input")
    return
}

func treatAsList(packet string, currentValue int) string {
    regexp.MustCompile("")
    return "[" + strconv.Itoa(currentValue) + "]" + packet
}

// returns < 0 if left is smaller, 0 if left and right are equal and > 0 if right is smaller
func comparePackets(leftPacket, rightPacket string) int {
    fmt.Printf("Compare `%s` vs `%s`\n", leftPacket, rightPacket)
    // get the next value in the packets: either a number or a closening or opening bracket.
    // note that parseNextValue modifies the packet, removing the first token
    leftValue, leftOpenList, leftCloseList := parseNextValue(&leftPacket)
    rightValue, rightOpenList, rightCloseList := parseNextValue(&rightPacket)

    // If we reach the end of the left list before the right list, the left list is smaller than the right
    if leftCloseList {
        return -1
    }
    // On the other hand, if we reach the end of the right list before the left, the right list is bigger
    if rightCloseList {
        return 1
    }

    // If left and right open a new list, compare that list
    if leftOpenList && rightOpenList {
        return comparePackets(leftPacket, rightPacket)
    }

    // If left opens a list and right is a number, treat both right and left as lists and compare one list to another
    if leftOpenList {
        // To trick comparePackets to treat the right packet as a list, we prepend the packet with an opening
        // bracket. Note that this means that there will no longer be a matching closing bracket.
        return comparePackets(treatAsList(leftPacket, leftValue), treatAsList(rightPacket, rightValue))
    }

    // If right opens a list and left is a number, treat left as a list and compare both lists
     if rightOpenList {
         // same comment about missing close bracket applies here
         return comparePackets(treatAsList(leftPacket, leftValue), treatAsList(rightPacket, rightValue))
     }

    // if both are integers we are comparing two lists
    // if both numbers are equal advance to the next value
    if leftValue == rightValue {
        return comparePackets(leftPacket, rightPacket)
    }
    // otherwise
    return leftValue - rightValue
}

func isList(element Element) bool {
    return len(element.list) > 0
}

func compareElementLists(leftList, rightList []Element) int {
    fmt.Printf("Compare %s vs %s \n", printElementList(leftList), printElementList(rightList))

    for len(leftList) > 0 && len(rightList) > 0 {
        left, right := leftList[0], rightList[0]
        leftList, rightList = leftList[1:], rightList[1:]

        switch {
        case isList(left) && isList(right):
            if subResult := compareElementLists(left.list, right.list); subResult == 0 {
                continue
            } else {
                return subResult
            }
        case isList(left):
            if subResult := compareElementLists(left.list, []Element{right}); subResult == 0 {
                continue
            } else {
                return subResult
            }
        case isList(right):
            if subResult := compareElementLists([]Element{left}, right.list); subResult == 0 {
                continue
            } else {
                return subResult
            }
        default: // both are ints
            if result := left.value - right.value; result == 0 {
                continue
            } else {
                return result
            }
        }
    }

    if len(leftList) > 0 {
        return 1
    } else {
        return -1
    }
}

func removeWrappingBrackets(packet string) string {
    packet = packet[1:]
    return packet[:len(packet) - 1]
}

func isRightOrder(pair string) bool {
    // we can expect packets to have length 2
    packets := strings.Split(pair, "\n")
    if size := len(packets); size != 2 {
        log.Fatalf("isRightOrder encountered an unexpected packet pair length. Expected 2, got %d", size)
    }

//    return comparePackets(packets[0], packets[1]) < 0
    return compareElementLists(packetToList(&removeWrappingBrackets(packets[0]), []Element{}), packetToList(&packets[1], []Element{})) < 0
}

func packetToList(packet *string, elems []Element) []Element {
    for *packet != "" {
        value, opens, closes := parseNextValue(packet)
        switch {
        case opens:
            subList := packetToList(packet, make([]Element, 0))
            elems = append(elems, Element{list: subList})
        case closes:
            return elems
        default:
            // if it doesn't open or close then it's an int value
            elems = append(elems, Element{value: value})
        }
    }
    return elems
}

func main() {
    input := utils.ReadFile("example_input")
    pairs := strings.Split(input, "\n\n")

    var rightOrderSum int
    for idx, pair := range pairs {
        fmt.Printf("== Pair %d ==\n", idx + 1)
        if isRightOrder(pair) {
            fmt.Printf("pair %d is in the right order\n", idx + 1)
            // idx is 0-based index, but aoc expect 1-based indexes, so we add 1
            rightOrderSum += 1 + idx
        } else {
            fmt.Printf("pair %d is in the wrong order\n", idx + 1)
        }
    }
    fmt.Println(rightOrderSum)
}