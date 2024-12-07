package day_01

func (d Day) secondPart() any {
	var seq1, seq2 = readSequences()

	var count = map[int]int{}

	for _, num := range seq2 {
		count[num]++
	}

	var res int64 = 0

	for _, num := range seq1 {
		res += int64(count[num]) * int64(num)
	}

	return res
}
