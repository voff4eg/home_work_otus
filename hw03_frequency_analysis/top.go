package hw03frequencyanalysis

import (
	"math"
	"sort"
	"strings"
)

func Top10(s string) []string {
	ss := strings.Fields(s)

	type wordSumStruct struct {
		Word string
		Sum  int
	}

	wordSum := make(map[string]int)
	words := make([]string, 10)
	i := 0

	for _, w := range ss {
		wordSum[w]++
	}

	wordsSumStruct := make([]wordSumStruct, len(wordSum))

	for word, sum := range wordSum {
		wordsSumStruct[i] = wordSumStruct{
			word,
			sum,
		}
		i++
	}

	sort.Slice(wordsSumStruct, func(i, j int) bool {
		return (wordsSumStruct[i].Sum > wordsSumStruct[j].Sum) ||
			(wordsSumStruct[i].Sum == wordsSumStruct[j].Sum &&
				LexicographicSort(wordsSumStruct[j].Word, wordsSumStruct[i].Word))
	})

	if len(wordsSumStruct) == 0 {
		return make([]string, 0)
	}

	for i, j := range wordsSumStruct {
		if i > 9 {
			break
		}
		words[i] = j.Word
	}

	return words
}

func LexicographicSort(prevStr string, nextStr string) bool {
	prev := []rune(prevStr)
	next := []rune(nextStr)
	r := len(prev) < len(next)
	ln := int(math.Min(float64(len(prev)), float64(len(next))))
	for i := 0; i < ln; i++ {
		if prev[i] == next[i] {
			continue
		}

		r = prev[i] > next[i]
		break
	}

	return r
}
