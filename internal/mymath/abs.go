package mymath


type Float float64

func (f Float) Square() Float {
	return f * f
}

func dummy(x int) int {
	return x
}


func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
