package main

import "sort"

type EncodingNode struct {
	probability float64
	chars       []rune
}
type EncodingNodeList []EncodingNode

func (e EncodingNodeList) Len() int {
	return len(e)
}
func (e EncodingNodeList) Less(i, j int) bool {
	return e[i].probability < e[j].probability
}
func (e EncodingNodeList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func createEncodingHuffman(c CharAnalysisList) (nodes EncodingNodeList) {
	for _, c := range c {
		nodes = append(nodes, EncodingNode{
			probability: c.probability,
			chars:       []rune{c.char},
		})
	}
	return
}

func (nodes EncodingNodeList) coding() (codes map[rune]string) {
	codes = make(map[rune]string)
	for len(nodes) > 1 {
		sort.Sort(nodes)

		newNode := EncodingNode{
			probability: nodes[0].probability + nodes[1].probability,
			chars:       append(nodes[0].chars, nodes[1].chars...),
		}
		for _, c := range nodes[0].chars {
			codes[c] += "0"
		}
		for _, c := range nodes[1].chars {
			codes[c] += "1"
		}
		nodes = append(nodes[2:], newNode)
	}
	return codes
}
