package main

import (
        "fmt"
        "github.com/abrookins/go_iching"
)

// reportHexagram prints details about a hexagram.
func reportHexagram(hexagram *iching.Hexagram) {
        fmt.Printf("#%v: %v (%v)\n", hexagram.Num, hexagram.Name, hexagram.Character)
        fmt.Println(hexagram.GetWillhelmUrl())
        fmt.Println(hexagram.GetLeggeUrl())
}

func main() {
        var question string
        fmt.Println("What is your question for the oracle?\n")
        fmt.Scanln(&question)
        reading := iching.GetReading(question)

        fmt.Printf("Lines: %v\n", reading.Lines)
        reportHexagram(reading.Hexagram)

        if reading.NextHexagram != reading.Hexagram {
                fmt.Println("\nHexagram is changing. Next hexagram:")
                reportHexagram(reading.NextHexagram)
        }
}
