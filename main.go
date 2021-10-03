package main

import (

	"encoding/json"
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"
	"strconv"
	"log"
)

type Course struct {
	Name string `json:"name,omitempty"`

	CourseID int `json:"courseID,omitempty"`

	Workload int32 `json:"workload,omitempty"`

	StudentSatisfaction int32 `json:"studentSatisfaction,omitempty"`
}



func main() {
	getCourseByID(2)
	postCourse(Course{Name: "Bussemand", CourseID: 2, Workload: 21, StudentSatisfaction: 0})
	getCourseByID(2)
	putCourse(Course{Name: "Ole", CourseID: 2, Workload: 123, StudentSatisfaction: 32})
	getCourseByID(2)
}


   func getCourseByID(id int){
	text := strconv.Itoa(id)
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/course/"+text, nil)
	if err != nil {
	 fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
	 fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	 fmt.Print(err.Error())
	}
	var responseObject Course
	json.Unmarshal(bodyBytes, &responseObject)
	if(responseObject.Name == ""){
		fmt.Printf("No course with this ID exists")
	}else{
	fmt.Printf("API Response as struct %+v\n", responseObject)
	}
   }

   func postCourse(course Course) {
    fmt.Println("2. Performing Http Post...")

    jsonReq, err := json.Marshal(course)
    resp, err := http.Post("http://localhost:8080/course", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)

    // Convert response body to Todo struct
    var courseStruct Course
    json.Unmarshal(bodyBytes, &courseStruct)
    fmt.Printf("%+v\n", courseStruct)
}

func deleteCourseByID(id int) {
	text := strconv.Itoa(id)
    fmt.Println("4. Performing Http Delete...")
	client := &http.Client{}
    req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/courseDel/"+text, nil)
    
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
}

func putCourse(course Course) {
	fmt.Println("2. Performing Http Put...")
	
    jsonReq, err := json.Marshal(course)
    resp, err := http.Post("http://localhost:8080/updateCourse", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)

    // Convert response body to Todo struct
    var courseStruct Course
    json.Unmarshal(bodyBytes, &courseStruct)
    fmt.Printf("%+v\n", courseStruct)
}
