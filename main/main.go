package main

import "server"
import "log"
import "time"

func main(){
	appjob := &server.AppJob{JobName:"test"}
	sr := server.Server{}
	
	var pipeArray [4]chan []byte
	for i := range pipeArray {
		pipeArray[i] = make(chan []byte)
	}
	sr.SvrPipe = pipeArray[:]
    //appjob.WritePipe = pipeArray[:]
   // go write(&sr)

	sr.SetWorker(appjob)
	sr.Run()	
	log.Println(appjob)
}

func write(svr *server.Server){
	time.Sleep(10*1e9)
	//svr.SvrPipe[0] <- 0x1
}