package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thinkofher/diceware"
)

func main() {
	wordsNum := func() int {
		if len(os.Args) < 2 {
			return 5
		}

		res, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return 5
		}

		return res
	}()
	fmt.Println(diceware.Generate(wordsNum))
}
