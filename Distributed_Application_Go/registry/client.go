package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error{
	bufferRessponse := new(bytes.Buffer)
	encodedResponse := json.NewEncoder(bufferRessponse)
	errorResponse := encodedResponse.Encode(r)
	if errorResponse != nil{
		return errorResponse
	}
	responsePost,errorPost := http.Post(ServiceURL,"application/json",bufferRessponse)
	if errorPost != nil{
		return errorPost
	}
	if responsePost.StatusCode != http.StatusOK{
		return fmt.Errorf("Failed to regsiter service. Register service"+"Responded with code %v",responsePost.StatusCode)
	}
	return nil
}

func ShutdownService(serviceURL string) error{
	request,error := http.NewRequest(http.MethodDelete,ServiceURL,bytes.NewBuffer([]byte(serviceURL)))
	if error != nil{
		return error
	}
	request.Header.Add("Content-Type","text/plain")
	response,error := http.DefaultClient.Do(request)
	if response.StatusCode != http.StatusOK{
		return fmt.Errorf("Failed to deregister service. Registry service responded with code %v",response.StatusCode)
	}
	return nil
}
