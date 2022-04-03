package config

import (
	"log"
	"os"
)
var (
	Warn *log.Logger
	Info *log.Logger
	Error *log.Logger
)

func InitLogger(){
	log.SetFlags((log.LstdFlags|log.Lshortfile))
	file,err:=os.OpenFile("logs.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
	if err!=nil{
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This my log message")
	Info=log.New(file,"INFO: ",log.LstdFlags|log.Lshortfile)
	Warn=log.New(file,"WARNING: ",log.LstdFlags|log.Lshortfile)
	Error=log.New(file,"Error: ",log.LstdFlags|log.Lshortfile)
	Info.Println("Logger Initiated......")
}