package day_01

import (
	"fmt"
	"math"
	"sort"
)

func (d Day) firstPart() {
	var res = 0
	var seq1, seq2 = readSequences()

	sort.Slice(seq1, func(i, j int) bool {
		return seq1[i] > seq1[j]
	})

	sort.Slice(seq2, func(i, j int) bool {
		return seq2[i] > seq2[j]
	})

	for i := 0; i < len(seq1); i++ {
		res += int(math.Abs(float64(seq1[i] - seq2[i])))
	}

	fmt.Println(res)
}
