package part1

func RecevieSlice(a *[]int) {
	for i := 0; i < len(*a); i++ {
		(*a)[i] *= 2
	}
}
