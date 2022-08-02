package main

func f(ch chan<- int, v int) {
	ch <- v
}

func main() {
	// send-only channel
	ch := make(chan<- int)

	go f(ch, 42)
	go f(ch, 41)
	go f(ch, 40)

	// 它是 只写channel，不能读数据
	// fmt.Println(<- ch)
}
