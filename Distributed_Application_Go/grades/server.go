package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type studentsHandler struct{}

var(
	studentsMutex sync.Mutex
)


func RegisterHandlers(){
	handler  := new(studentsHandler)
	http.Handle("/students",handler)
	http.Handle("/students/",handler)
}
/*
	Entire class
	students/{id} - a single student's record
	sudents/{id}/grades - a single strudent's grades
*/
func (sh studentsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	pathSegments := strings.Split(request.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		sh.getAll(writer, request)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(writer, request, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(writer, request, id)
	default:
		writer.WriteHeader(http.StatusNotFound)

	}
}

func (sh studentsHandler) getAll(writer http.ResponseWriter,request *http.Request){
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data,err := sh.toJson(students)
	if err !=nil{
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	writer.Header().Add("Content-Type","application/json")
	writer.Write(data)
}

func(sh studentsHandler) toJson(object interface{}) ([]byte,error){
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(object)
	if err !=nil{
		return nil,fmt.Errorf("failed to serialize students: %q",err )
	}
	return b.Bytes(),nil
}

func (sh studentsHandler) getOne(writer http.ResponseWriter, request *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student,err := students.GetByID(id)
	if err != nil{
		writer.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	data,err := sh.toJson(student)
	if err != nil{
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("Failed to serialize students: %q",err))
		return
	}
	//Sending response back to the clients
	writer.Header().Add("Content-Type","application/json")
	writer.Write(data)
}

func (sh studentsHandler) addGrade(writer http.ResponseWriter, request *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student,err := students.GetByID(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	var g Grade
	dec := json.NewDecoder(request.Body)
	err = dec.Decode(&g)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	student.Grades = append(student.Grades,g)
	writer.WriteHeader(http.StatusCreated)

	data,err := sh.toJson(g)
	if err != nil{
		log.Println(err)
		return
	}
	writer.Header().Add("content-type","application/json")
	writer.Write(data)
}

