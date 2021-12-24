package log

import (
	"fmt"
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func(fl fileLog) Write(data []byte) (int,error){
	//Writing log to a file.
	f,errorResponse := os.OpenFile(string(fl),os.O_CREATE | os.O_WRONLY | os.O_APPEND,0600 )
	defer f.Close()
	if errorResponse != nil{
		return 0, errorResponse
	}
	return f.Write(data)
}

func Run(destination string){
	log = stlog.New(fileLog(destination),"",stlog.LstdFlags)
}

func RegisterHandlers(){
	fmt.Println("Inside RegisterHandlers")
	http.HandleFunc("/log",func(w http.ResponseWriter,r *http.Request){
		fmt.Println("Starting log service")
		msgRecieved,errorResponse := ioutil.ReadAll(r.Body)
		if errorResponse != nil || len(msgRecieved) == 0{
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msgRecieved))
	})
}

func write(message string){
	fmt.Println("Inside write")
	log.Printf("%v\n",message)
}