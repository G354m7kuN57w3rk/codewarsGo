package complementarydna

import "strings"

//DNAStrand returns dna complement strings
func DNAStrand(dna string) string {
	dnaSplit := strings.Split(dna, "")
	var ans string
	for _, v := range dnaSplit {
		switch {
		case v == "A":
			ans += "T"
		case v == "T":
			ans += "A"
		case v == "G":
			ans += "C"
		case v == "C":
			ans += "G"
		}
	}
	return ans
}
