package main

import "net"
import "log"
import "time"

func request(){
    conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		// handle error
		log.Println("error: ",err)
	}
	buf := make([]byte, 5)
	buf[0] = 1
	buf[1] = 2
	for i := range buf {
		if i > 1 {
			buf[i] = buf[i - 1] + 1
		}
	}

    hbuf := []byte("hello svr, test ....xkkdkdk ssssssss")
	conn.Write(hbuf)
	//log.Println("write finish")
	buf2 := make([]byte, 256)
	len,_ := conn.Read(buf2)
	log.Printf("buf %s len %d \n", buf2[:len], len)
	time.Sleep(1*1e9)
}

func main(){
	request()
}
