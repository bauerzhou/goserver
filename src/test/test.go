package main

//import "os"
//import "log"
import "fmt"
import "glog"
const (
	L = 1 
	Q
	J
	N
	P
)
func main(){
	//f,_ := os.Create("./hello.txt")
	//log.SetOutput(f)
	// log.SetPrefix("[FATAL]")
	// log.Print("hello")
	// log.SetPrefix("[DEBUG]")
	// log.Print("heihei")
	// log.Print(L,Q,J,N,P)
	str := "hehhhh"
	str2 := "hslslsl"
	s3 := fmt.Sprint(str,str2)
	fmt.Println(s3)
	glog.Init("./", "test")
	glog.Log(2,"hello log 2")
	glog.Log(3,"hello log 2")
	return 
}
