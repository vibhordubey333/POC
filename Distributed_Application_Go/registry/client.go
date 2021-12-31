package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

func RegisterService(r Registration) error{
	servicesUpdateURL,err := url.Parse(r.ServiceUpdateURL)
	if err !=nil{
		return err
	}
	http.Handle(servicesUpdateURL.Path,&serviceUpdateHandler{})

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

type serviceUpdateHandler struct{}

func( suh serviceUpdateHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request) {
	if request.Method != http.MethodPost{
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(request.Body)
	var p patch
	err := dec.Decode(&p)
	if err != nil{
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	prov.Update(p)
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


type providers struct{
	services map[ServiceName][]string
	mutex *sync.RWMutex
}

var prov = providers{
	services : make(map[ServiceName][]string),
	mutex: new(sync.RWMutex),
}

func(p *providers) Update(pat patch){
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _,patchEntry := range pat.Added{
		if _,ok := p.services[patchEntry.Name]; !ok{
			p.services[patchEntry.Name] = make([]string,0)
		}
		p.services[patchEntry.Name] = append(p.services[patchEntry.Name],patchEntry.URL)
	}
	for _,patchEntry := range pat.Removed{
		if providedURLs,ok := p.services[patchEntry.Name]; ok{
			for i := range providedURLs{
				if providedURLs[i] == patchEntry.URL{
					p.services[patchEntry.Name] = append(providedURLs[:i],providedURLs[i+1:]...)
				}
			}
		}
	}
}

func(p providers) get(name ServiceName)(string,error){
	providers,ok := p.services[name] ; if !ok{
		return "",fmt.Errorf("no providers available for service %v",name)
	}
	idx := int(rand.Float32() * float32(len(providers)))
	return providers[idx],nil
}

func GetProviders(name ServiceName) (string,error){
	return prov.get(name)
}