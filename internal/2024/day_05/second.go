package day_05

import "fmt"

func (d Day) secondPart() {
	rules := d.readRules()

	var res int64 = 0
	for _, sequence := range d.readSequences() {
		if d.isValid(sequence, rules) {
			continue
		}

		ordered := d.makeValid(sequence, rules)

		res += int64(ordered[len(ordered)/2])
	}

	fmt.Println(res)
}

func (d Day) makeValid(sequence []int, rules [][]int) []int {
	index := map[int]int{}

	for i, page := range sequence {
		index[page] = i
	}

	// bubble sort using index[] ordering O(n^2)
	done := false
	for !done {
		done = true
		for _, rule := range rules {
			a, b := rule[0], rule[1]
			_, existsA := index[a]
			_, existsB := index[b]

			if !existsA || !existsB || index[a] < index[b] {
				continue
			}

			index[a], index[b] = index[b], index[a]
			done = false
		}
	}

	result := make([]int, len(sequence))

	for page, i := range index {
		result[i] = page
	}

	return result
}
