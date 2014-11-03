package glog

import "log"
import "os"
import "fmt"
import "strconv"

const (
	LOG_TRACE 	= 1 
	LOG_DEBUG 	= 2
	LOG_NORMAL 	= 3
	LOG_INFO 	= 4
	LOG_ERROR  	= 5
	LOG_FATAL 	= 6
)

const (
	STR_LOG_TRACE  = "LOG_TRACE"
	STR_LOG_DEBUG  = "LOG_DEBUG"
	STR_LOG_NORMAL = "LOG_NORMAL"
	STR_LOG_INFO   = "LOG_INFO"
	STR_LOG_ERROR  = "LOG_ERROR"
	STR_LOG_FATAL  = "LOG_FATAL"
	STR_LOG_UNKNOW = "LOG_UNKNOW"
)

type GLoger struct {
	MaxSizePerLog  	int
	MaxFileNum		int

	CurIndex 		int
	CurSize			int

	PrefixName		string
	PrefixPath		string

	loger			*log.Logger
	loglevel  		int
}



func New(prepath, prename string) (*GLoger, error){
	name := fmt.Sprint(prepath, prename, ".log")

	w, err := os.Create(name)
	
	if err != nil {
		return nil, err
	}

//	return nil, nil

	return &GLoger{PrefixPath:prepath, PrefixName:prename,  MaxSizePerLog:1024*1024*4, MaxFileNum: 1000, CurIndex:0,loger:log.New(w, "", log.LstdFlags)},nil
}

var l *GLoger

func (this *GLoger) createWriter() error {
	this.CurIndex = (this.CurIndex + 1) % this.MaxFileNum
	name := this.PrefixPath+this.PrefixName+strconv.Itoa(this.CurIndex)+".log"
	_ , err := os.Create(name)
	if err != nil {
		return err
	}

	//this.loger.SetOutput(w)

	return nil
}

func (this *GLoger) SetLogLevel(l int) {
	this.loglevel = l
}

func (this *GLoger) GetLogLevel() int {
	return this.loglevel 
}


func (this *GLoger) Log(level int, v ...interface{}){
	if level < this.loglevel {
		return
	}
	var prefix string
	switch level {
		case 	LOG_TRACE :
			prefix = STR_LOG_TRACE
		case 	LOG_DEBUG :
			prefix = STR_LOG_DEBUG
		case 	LOG_NORMAL :
			prefix = STR_LOG_NORMAL
		case 	LOG_ERROR :
			prefix = STR_LOG_ERROR
		default:
			prefix = STR_LOG_UNKNOW
	}
	this.loger.SetPrefix(prefix)
	this.loger.Print(v...)
}

func Init(prepath, prename string) error {
	var err error 
	l, err = New(prepath, prename)
	if err != nil {
		return err
	}
	return nil
}

func SetLogLevel(level int) {
	l.SetLogLevel(level)
}

func GetLogLevel() int{
	return l.GetLogLevel()
}

func Log(level int, v ...interface{}){
	l.Log(level, v...)
}


