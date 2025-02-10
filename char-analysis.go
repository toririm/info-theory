package main

import (
	"math"
	"sort"
)

type CharAnalysis struct {
	char        rune
	count       int
	probability float64
}

type CharAnalysisList []CharAnalysis

func (c CharAnalysisList) Len() int {
	return len(c)
}

func (c CharAnalysisList) Less(i, j int) bool {
	return c[i].count < c[j].count
}

func (c CharAnalysisList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func getEntropy(c CharAnalysisList) float64 {
	var entropy float64
	for _, c := range c {
		entropy += -c.probability * math.Log2(c.probability)
	}
	return entropy
}

func createCharAnalysisList(content string) (length int, c CharAnalysisList) {
	length = len(content)
	m := map[rune]int{}
	for _, c := range content {
		m[c]++
	}

	for k, v := range m {
		c = append(c, CharAnalysis{
			char:        k,
			count:       v,
			probability: float64(v) / float64(length),
		})
	}
	sort.Sort(c)

	return
}
