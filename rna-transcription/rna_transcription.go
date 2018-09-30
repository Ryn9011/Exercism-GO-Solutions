package strand

import (
	"strings"
)

func ToRNA(dna string) string {
	var rna strings.Builder

	strandMap := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	for i := range dna {
		rna.WriteString(strandMap[string(dna[i])])
	}
	return rna.String()
}
