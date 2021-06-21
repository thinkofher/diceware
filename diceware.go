package diceware

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
)

// Generate generates diceware password with given number
// of words. Each word starts with capital letter and there
// are no spaces or other whitespace characters between words.
func Generate(wordsNum int) string {
	wordsMapGenerated := wordsMap()
	wordsSlice := make([]string, wordsNum)

	for i := 0; i < wordsNum; i++ {
		wordsSlice[i] = strings.Title(
			wordsMapGenerated[diceStringThrow(numOfDiceSides, throwStringLength)],
		)
	}

	return strings.Join(wordsSlice, "")
}

const throwStringLength = 5
const numOfDiceSides = 6

func wordsMap() map[string]string {
	res := map[string]string{}
	for _, line := range strings.Split(words, "\n") {
		if line != "" {
			key := line[:strings.Index(line, "\t")]
			value := line[strings.Index(line, "\t")+1:]
			res[key] = value
		}
	}
	return res
}

func randRange(min, max int64) int64 {
	res, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return 0
	}
	return res.Int64() + min
}

func throwDice(k int64) int64 {
	if k < 0 {
		return 0
	}
	return randRange(1, k)
}

func diceStringThrow(k int64, length int) string {
	if length <= 0 {
		return ""
	}
	res := ""
	for i := 0; i < length; i++ {
		res += strconv.FormatInt(throwDice(k), 10)
	}
	return res
}
