package main

import "testing"

func Test_ComparePackets(t *testing.T) {
    var(
        leftPacket = "[1,1,3,1,1]"
        rightPacket = "[1,1,5,1,1]"
    )

    if comparePackets(leftPacket, rightPacket) >= 0 {
        t.Fatal("list vs list comparisson not working")
    }

    leftPacket = "[[]]"
    rightPacket = "[[]]"

    if comparePackets(leftPacket, rightPacket) >= 0 {
        t.Fatal("empty list vs empty list comparisson not working")
    }

    leftPacket = "[9]"
    rightPacket = "[]"

    if comparePackets(leftPacket, rightPacket) <= 0 {
        t.Fatal("left list longer than right list not working")
    }

    leftPacket = "[4]"
    rightPacket = "[[5]]"

    if comparePackets(leftPacket, rightPacket) >= 0 {
        t.Fatal("left list longer than right list not working")
    }

    t.Log("working")
}

//func Test_ParseNextValue(t *testing.T) {
//
//    packet := "[10,9,[1],10]"
//    expectedOpening := []bool{true, false, false, true, false, false, false, false}
//    expectedClosing := []bool{false, false, false, false, false, true, false, true}
//    expectedValue := []int{0, 10, 9, 0, 1, 0, 10, 0}
//
//    for idx := 0; idx < len(expectedOpening); idx++ {
//        value, openBracket, closeBracket := parseNextValue(&packet)
//        if value != expectedValue[idx] || openBracket != expectedOpening[idx] || closeBracket != expectedClosing[idx] {
//            t.Fatalf("invalid output, expected `(%d, %t, %t)`, got `(%d, %t, %t)`", value, openBracket, closeBracket, expectedValue[idx], expectedOpening[idx], expectedClosing[idx])
//        }
//    }
//
//}