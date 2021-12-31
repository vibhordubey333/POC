package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"
const ServiceURL = "http://localhost" + ServerPort + "/services"

type registry struct{
	registrations []Registration
	mutex 		*sync.RWMutex
}

func(r *registry) add(reg Registration) error{
	r.mutex.Lock()
	r.registrations = append(r.registrations,reg)
	r.mutex.Unlock()
	err := r.sendRequiredServices(reg)
	if err != nil{
		return err
	}
	return nil
}


func(r *registry) remove(url string) error{
	for i := range r.registrations{
		if r.registrations[i].ServiceURL == url{
			r.mutex.Lock()
			r.registrations = append(r.registrations[:i],r.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("Service at URL %v not found",url)
}

func (r *registry) sendRequiredServices(registration Registration) error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var p patch
	for _,serviceReg := range r.registrations{
		for _,reqService := range registration.RequiredServices{
			if serviceReg.ServiceName == reqService{
				p.Added = append(p.Added, patchEntry{
					Name: serviceReg.ServiceName,
					URL:  serviceReg.ServiceURL,
				})
			}
		}
	}
	err := r.sendPatch(p,registration.ServiceUpdateURL)
	if err != nil{
		return err
	}
	return nil
}

func (r *registry) sendPatch(p patch, url string) error {
	d,err := json.Marshal(p)
	if err != nil{
		return err
	}
	_,err = http.Post(url,"application/json",bytes.NewBuffer(d))
	if err != nil{
		return err
	}
	return nil

}

var reg = registry{
	registrations: make([]Registration,0),
	mutex : new(sync.RWMutex),
}

type RegistryService struct{}

func(s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Println("Request received by RegistryService")
	switch r.Method{
	case http.MethodPost:
			log.Println("Inside ServeHTTP Method Post")
			dec := json.NewDecoder(r.Body)
			var r Registration
			err := dec.Decode(&r)
			if err != nil{
				log.Println("RegsitryService | Error while decoding service.")
				w.WriteHeader(http.StatusBadRequest)
				return
		}
		log.Printf("Adding service : %v with URL: %v \n",r.ServiceName,r.ServiceURL)
			err = reg.add(r)
		if err != nil{
			log.Println("RegsitryService | Add | ErrorResponse |",err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		log.Println("Inside ServeHTTP Method Delete")
		payload,err := ioutil.ReadAll(r.Body)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(payload)
		log.Printf("Removing service at URL: %v",url)
		err = reg.remove(url)
		if err != nil{
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}