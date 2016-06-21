package main

const (
	c = 5000
	n = 1000
)

var x = 0

func adder() int {

	for i := 0; i < c; i++ {
		for j := 0; j < n; j++ {
			x++
		}

	}
	return x
}
func main() {
	sum := adder()
	println(sum)

}
