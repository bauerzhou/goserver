package server

import "net"

type Listener interface{

}

type Server interface{
	Start() error 
	Stop()
}

type Config struct{
	ip   		string
	port 		string
	
	timeout		int
}

type UdpServer struct{
	config 		*Config
	
	running 	bool
	stateChan	chan bool

	Handle(data []byte, c *net.UDPConn, addr *net.UDPAddr) error
}

//func ResolveUDPAddr(net, addr string) (*UDPAddr, error)

func (this *UdpServer) Start(){
	this.running = true
	this.Loop()
}

func (this *UdpServer) Loop(){
	udpAddress, err := ResolveUDPAddr("udp4", config.ip + ":" + config.port)

	conn ,err := net.ListenUDP("udp",udpAddress)

    if err != nil {
            fmt.Println("error listening on UDP port ", address)
            fmt.Println(err)
            return
    }

    var buf []byte = make([]byte, 1500)

    for ; this.running ; {
        n,address, err := conn.ReadFromUDP(buf)

        if err != nil {
            fmt.Println("error reading data from connection")
            fmt.Println(err)
            return
        }

        go this.Handle(buf[:n], conn, address)
    }
}

func (this *UdpServer) StopRunning(){
	this.running = false
}

func (this *UdpServer) Stop(){
	statChan <- true
}

func (this *UdpServer) Wait(){
	for {
		st :=  <- stateChan 
	}
	this.running = false
	return
}















