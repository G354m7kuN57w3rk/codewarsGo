package growthofapopulation

//NbYear returns the growth of population
func NbYear(p0 int, percent float64, aug int, p int) int {
	count := 0
	for p > p0 {
		p0 += int(float64(p0)*percent/100) + aug
		count++
	}
	return count
}
