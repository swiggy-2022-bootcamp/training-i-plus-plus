package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"strconv"
	"log"
	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct{
	CourseId string `json:"courseid"`
	Coursename string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:author`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middleware,helper - file

func (c *Course)IsEmpty() bool{
	return c.CourseId=="" && c.Coursename==""
}

func main(){
    r:=mux.NewRouter()
	//seeding
	courses=append(courses,Course{CourseId:"2",Coursename:"ReactJS",CoursePrice: 299,	Author:&Author{Fullname:"Hitesh Chaudhary",Website:"lco.dev"}})
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("POST")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":4000",r))

}


// controllers - file

// serve home route

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1> Welcome to API by LearningCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}


func getOneCourse(w http.ResponseWriter,r *http.Request){
	 fmt.Println("Get One course")
	 w.Header().Set("Content-Type","application/json")
	 params:=mux.Vars(r)

	 for _,course := range courses{
		 if course.CourseId==params["id"]{
			 json.NewEncoder(w).Encode(course)
			 return
		 }
	 }
	 json.NewEncoder(w).Encode("no course found")
	 return
}


func createOneCourse(w http.ResponseWriter,r *http.Request){

	fmt.Println("Create One course")
	w.Header().Set("Content-Type","application/json")

	if r.Body==nil{
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses=append(courses,course)
	json.NewEncoder(w).Encode(course)
	return
}


func updateOneCourse(w http.ResponseWriter, r*http.Request){
	fmt.Println("Get One course")
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	// loop, id, reove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
}


func deleteOneCourse(w http.ResponseWriter, r*http.Request){
	fmt.Println("Delete One course")
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	// loop, id, remove (index,index+1)
	for index,course:=range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			break
		}
	}
}