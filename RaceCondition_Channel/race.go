package main

const (
	n = 1000
)

var sum = 0
var done = make(chan bool)
var receiver = make(chan int)

func adder(msg chan int) {
	var i int
	for {
		for i = 0; i < n; i++ {
			i++
		}
		receiver <- i
		done <- true
	}
}

func main() {
	sender := make(chan int)
	for i := 0; i < n; i++ {
		go adder(sender)
	}
	for i := 0; i < n; i++ {
		z := <-receiver
		sum = sum + z
		<-done
	}
	println(sum)
}
