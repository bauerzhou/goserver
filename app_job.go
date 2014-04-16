package server

import (
	"log"
)

type AppJob struct{
	JobName string
	WritePipe  []chan []byte
}

func (app *AppJob) OnDataRequest(buf []byte, len int) {
	log.Println("Recv data len", len)
	log.Printf("Recv data %v", buf[0:len])
	head := DataHeader{}
	head.ReadFromBuf(buf[0:12])
	log.Println("head:", head.flow)
	//log.Println("pipe len:", len(app.WritePipe))
	app.WritePipe[head.flow % 3] <- buf
}

func (app *AppJob) InitPipe(pipe [] chan []byte) error {
	log.Println(len(pipe))
	app.WritePipe = pipe[0: len(pipe)]
	log.Println(len(app.WritePipe))
	return nil
}
