package service

import (
	"Distributed_Application_Go/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context,  host string, port string, reg registry.Registration,registerHandlerFunc func()) (context.Context, error) {
	registerHandlerFunc()
	fmt.Println("Starting Service")
	fmt.Println("Start, host:",host," port",port)
	ctx = startService(ctx, reg.ServiceName, port, host)
	err := registry.RegisterService(reg)
	if err != nil{
		return ctx,err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, port, host string) context.Context {

	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	log.Print("Start service:",host, ":", port)

	//hostURL := "http://"+host+":"+port
	srv.Addr = ":"+port

	log.Println("Srv Items: ",srv.Addr)
	//host = ":" + port

	go func() {
		log.Println("Listen and serve: ",srv.ListenAndServe())
		cancel()
	}()
	go func() {
		fmt.Printf("%v started. Press any key to quit.", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprint("http://%v:%v",host,port))
		if err != nil{
			log.Println(err)
		}
		srv.Shutdown(ctx)
	}()
	return ctx
}
