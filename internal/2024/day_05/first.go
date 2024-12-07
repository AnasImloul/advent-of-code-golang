package day_05

func (d Day) firstPart() any {
	rules := d.readRules()

	var res int64 = 0
	for _, sequence := range d.readSequences() {
		if d.isValid(sequence, rules) {
			res += int64(sequence[len(sequence)/2])
		}
	}
	return res
}

func (d Day) isValid(sequence []int, rules [][]int) bool {
	index := map[int]int{}

	for i, page := range sequence {
		index[page] = i
	}

	for _, rule := range rules {
		a, b := rule[0], rule[1]

		_, existsA := index[a]
		_, existsB := index[b]

		if !existsA || !existsB {
			continue
		}

		if index[a] > index[b] {
			return false
		}
	}

	return true
}
