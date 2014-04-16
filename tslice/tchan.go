package main

import (
	"log"
	"time"
)

func main(){
	test2()
}

func test2(){
	var chArray [4] chan []byte

	for i := range chArray {
		chArray[i] = make(chan []byte)
	}
	//ch := make(chan []byte)

	go writeByte(chArray[:])

	select {
		case b := <- chArray[1]:
			log.Println("rcv bytes len: ", len(b))
			log.Println(b)
	}
}
func writeByte(ch []chan []byte){
	b := make([]byte, 10)
	b[0] = 1
	b[1] = 2
	b[2] = 3
	ch[1] <- b
}
func test(){
	var ch1 [4]chan int
	ch1[1] = make(chan int)
	ch1[0] = make(chan int)
	log.Println(len(ch1))
	for _, ch := range ch1 {
		go read(ch)	
	}

	go write(ch1[:])
	time.Sleep(3*1e9)
}

func write(ch []chan int){
	ch[1] <- 3
	log.Println(<-ch[1])
}

func read(ch chan int){
	log.Println(<-ch)
}


