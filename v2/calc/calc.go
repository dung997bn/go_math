package calc

//Add func
func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s

}

//Sub func
func Sub(a, b int) int {
	return a - b
}
