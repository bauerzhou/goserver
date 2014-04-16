package server

import (
		"log"
		"net"
		//"errors"
	)

const MAX_SVR_CHAN_NUM uint32 = 4
const MAX_CONNECT_NUM  uint32 = 4
var EOF SvrError

type SvrError struct{
	err string
}

type Server struct{
	port string
	svrIp string
	connectType uint8

	connFlow [MAX_CONNECT_NUM]ConnInfo
	SvrPipe  [] chan []byte 

	flow  uint32
	job  Job
}

type ConnInfo struct{
	activeTime 	uint
	conn		*net.Conn
	status      int
}

func (c *ConnInfo) Close(){
	activeTime = 0
	status = UNUSED
	if c.conn != nil {
		c.conn.Close()
	}
}

type Job interface{
	InitPipe([] chan []byte) error
	OnDataRequest([]byte, int)
}

type Session struct{
	id 				uint64
	flow 			uint32
	connectTime 	uint64
	activeTime 		uint64
	clientIp 		uint32
	port 			uint16
}

const (
	USING = 1
	UNUSED = 2
)


func init(){
	EOF = SvrError{"EOF"} 
}


func (server *Server) Run(){
	log.Println("start run")
	l, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	go server.onWriteData()
	log.Println("do for loop")
	for {
		log.Println("wait...for...connection")
		//wait for connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("new conn")
		//handle the connection
		go server.handleConnection(&conn)

	}

}

func (server *Server) handleConnection(c *net.Conn) {
	if server.flow >= MAX_CONNECT_NUM {
		log.Println("connection over flow, refuse")
		(*c).Close()
		return
	}

	ci := ConnInfo{conn:c, activeTimeStamp: time.Now().Unix()}
	flow := server.flow   
	server.flow ++ //need mutex lock

	server.connFlow[flow] = ci
	buf := make([]byte, 1024)
	head := DataHeader{}
	head.flow = flow
	head.cmd = 0x01
	head.WriteToBuf(buf)

	var length int

	for {
		length, err := (*c).Read(buf[12:])
		if err != nil {
			if err.Error() == "EOF" {
				continue
			}
			log.Println("handleConnection err: ", err)
			break
		}
		server.connFlow[flow].activeTimeStamp = time.Now().Unix()

		server.job.OnDataRequest(buf[0:length + 12], length + 12)
	}
}

func (this *Server) SetWorker(worker Job){
	this.job = worker
	this.job.InitPipe(this.SvrPipe)
}

func (this *Server) onWriteData(){
	log.Println("onWriteData")
	end := make(chan int)
	for i := range this.SvrPipe  {
		go func(ch chan []byte, e chan int, index int){
			log.Println("select chan ", index)
			for{
				data := <-ch
				log.Printf("index: %d, write data: %s", index, data)
				this.HandleWrite(data)
				
			}
			e <- 1
		}(this.SvrPipe[i], end, i)
	}
	count := 0
	for{
		<-end
		count += 1
		if count > 4 {
				break
		}
	}
	return
}

func (this *Server) HandleWrite(buf []byte){
	head := DataHeader{}
	head.ReadFromBuf(buf[0:12])

	flow := head.flow
	len, err := (*this.connFlow[flow]).Write(buf[12:])
	if err != nil {
		log.Println("HandleWrite write fail", err)
		return
	}
	log.Println("HandleWrite len ", len)
}

func (this *Server) HandleConnectionTimeout(){
	for c := range this.connFlow {
		if c.status == USING && abs(time.Now().Unix() - c.activeTimeStamp) >= CONN_TIME_OUT {
			(&c).Close()
		}
	}
}
