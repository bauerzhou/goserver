package main

import "os"
import "log"
import "fmt"
const (
	L = 1 
	Q
	J
	N
	P
)
func main(){
	f,_ := os.Create("./hello.txt")
	log.SetOutput(f)
	log.SetPrefix("[FATAL]")
	log.Print("hello")
	log.SetPrefix("[DEBUG]")
	log.Print("heihei")
	log.Print(L,Q,J,N,P)
	str := "hehhhh"
	str2 := "hslslsl"
	s3 := fmt.Sprint(str,str2)
	fmt.Println(s3)
	return 
}
