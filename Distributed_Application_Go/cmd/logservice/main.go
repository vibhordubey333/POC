package main
import(
	"Distributed_Application_Go/log"
	"Distributed_Application_Go/registry"
	"Distributed_Application_Go/service"
	"context"
	"fmt"
	stLog "log"
)
func main(){
		log.Run("Distributed_Application_Go.log")
		//TODO: Pick these from configuration files
		host,port := "127.0.0.1","7777"
		serviceAddress := fmt.Sprintf("http://%v:%v",host,port)

		var r registry.Registration
		r.ServiceName = registry.LogService
		r.ServiceURL  =	serviceAddress
		fmt.Println("ServiceName: ",r.ServiceName, "ServiceURL: ",r.ServiceURL)
		fmt.Println("Log service starting on host and port",host,":",port)
		ctx,errorResponse := service.Start(
			context.Background(),
			host,
			port,
			r,
			log.RegisterHandlers,
		)
		if errorResponse != nil{
			stLog.Fatalf(errorResponse.Error())
		}
		<-ctx.Done()
		fmt.Println("Shutting down log service.")
}