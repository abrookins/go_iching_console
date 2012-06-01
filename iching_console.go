package main

import (
	"bufio"
	"fmt"
	"github.com/abrookins/go_iching"
	"os"
)

// reportHexagram prints details about a hexagram.
func reportHexagram(hexagram *iching.Hexagram) {
	fmt.Printf("Result: %v (%v, #%v)\n\nDescription:\n%v\n\nTranslations:\n",
		hexagram.Name, hexagram.Character, hexagram.Num,
		hexagram.Description)

	for _, tr := range hexagram.TranslationUrls {
		fmt.Println(tr)
	}
}

func main() {
	fmt.Println("What is your question for the oracle?\n")
	
	in := bufio.NewReader(os.Stdin)
	question, _ := in.ReadString('\n')
	reading := iching.GetReading(question)

	fmt.Printf("----------------------------------\n\nLines: %v\n\n", reading.Lines)
	reportHexagram(reading.Hexagram)

	if reading.NextHexagram.Num != reading.Hexagram.Num {
		fmt.Println("\nHexagram is changing. Next hexagram:")
		reportHexagram(reading.NextHexagram)
	}
}
