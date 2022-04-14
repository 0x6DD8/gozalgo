package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func randInt(min int, max int) (x int) {
	x = rand.Intn(max-min) + min
	return x
}

func zalgo() (s string) {
	s = fmt.Sprintf("%c", randInt(768, 879))
	return s
}

func zalgoText(text string, curseCount int) {
	var zalgoedText strings.Builder
	zalgoedText.WriteString(text[0:1])

	for i := 0; i < len([]rune(text)); i++ {

		var textForZalgo strings.Builder
		textForZalgo.WriteString(text[i : i+1])

		for j := 0; j < curseCount; j++ {
			textForZalgo.WriteString(zalgo())
		}

		fmt.Print(textForZalgo.String())
	}
	fmt.Println()
}

func helpText() {
	fmt.Println("Usage: gozalgo <text> <number of Cursed Characters> \n\n", "-h   : this")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	arg := os.Args

	legthOfArgs := len(arg)

	switch legthOfArgs {
	case 1:
		fmt.Println("missing args, type 'gozalgo -h' for help")
		break
	case 2:
		if strings.Contains(arg[1], "-h") {
			helpText()
		} else {
			zalgoText(arg[1], 5)
		}

		break
	case 3:
		x, err := strconv.Atoi(arg[2])
		if err != nil {
			fmt.Println("error, type 'gozalgo -h' for help")
			os.Exit(2)
		}
		zalgoText(arg[1], x)
		break
	default:
		if legthOfArgs > 3 {
			fmt.Println("error, type 'gozalgo -h' for help")
		}
		break
	}

}
