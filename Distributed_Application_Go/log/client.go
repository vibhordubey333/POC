package log

import (
	"Distributed_Application_Go/registry"
	"bytes"
	"fmt"
	stlog "log"
	"net/http"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName){
	stlog.SetPrefix(fmt.Sprintf("[%v] - ",clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}



type clientLogger struct{
	url string
}

func(cl clientLogger) Write(data []byte)(int,error){
	b :=  bytes.NewBuffer([]byte(data))
	res,err := http.Post(cl.url+"/log","text/plain",b)
	if err != nil{
		return 0,err
	}
	if res.StatusCode != http.StatusOK{
		return 0,fmt.Errorf("failed to send log message. Service responded with error code.%v",err)
	}
	return len(data),nil
}