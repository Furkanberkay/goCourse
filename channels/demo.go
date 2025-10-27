package channels

func GetData() string {
	chanel1 := make(chan string)
	go producer(chanel1)
	data := <-chanel1
	return data
}

func producer(ch chan<- string) {
	ch <- "selam"
}
