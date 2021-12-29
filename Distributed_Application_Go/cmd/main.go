package main
import(
	"Distributed_Application_Go/log"
	"Distributed_Application_Go/service"
	"context"
	"fmt"
	stLog "log"
)
func main(){
		log.Run("../:Distributed_Application_Go.log")
		host,port := "localhost","4000"
		fmt.Println("Log service starting on host and port",host,":",port)
		ctx,errorResponse := service.Start(context.Background(),
			"Log Service",
			host,
			port,
			log.RegisterHandlers,
		)
		if errorResponse != nil{
			stLog.Fatalf(errorResponse.Error())
		}
		<-ctx.Done()
		fmt.Println("Shutting down log service.")
}