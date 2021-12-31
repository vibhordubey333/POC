package main

import (
	"Distributed_Application_Go/grades"
	"Distributed_Application_Go/registry"
	"Distributed_Application_Go/service"
	"context"
	"fmt"
	"log"
	stlog "Distributed_Application_Go/log"
)

func main(){
	host,port := "localhost","6666"
	serviceAddress := fmt.Sprintf("http://%v:%v",host,port)

	var r registry.Registration
	r.ServiceName = registry.GradeService
	r.ServiceURL  = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx,err := service.Start(
		context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
		if err != nil{
			log.Fatalf("%v",err)
		}
		if logProvider,err := registry.GetProviders(registry.LogService); err !=nil{
			fmt.Println("Logging service found at : %v\n",logProvider)
			stlog.SetClientLogger(logProvider,r.ServiceName)

		}
		<-ctx.Done()
		log.Println("Shutting down")
}