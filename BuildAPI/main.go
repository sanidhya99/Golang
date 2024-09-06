package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// models - file
type Course struct {
	CourseId    int     `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// fake DB
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

var (
	coursess     []Course
	lastCourseID int
	idMutex      sync.Mutex
)

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

func main() {
	fmt.Println("Welcome to model making tutorial")

	//seeding

	courses = append(courses, Course{CourseId: 1, CourseName: "GoLang", CoursePrice: 1999, Author: &Author{Name: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: 2, CourseName: "Django-Rest-Framework", CoursePrice: 2999, Author: &Author{Name: "Sanidhya Gupta", Website: "https://veergatha.netlify.app"}})
	//routing

	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/get/courses/", getCourses).Methods("GET")
	r.HandleFunc("/get/course/{id}/", getCourse).Methods("GET")
	r.HandleFunc("/post/course/", createCourse).Methods("POST")
	r.HandleFunc("/update/course/{id}/", updateCourse).Methods("PUT")
	r.HandleFunc("/delete/course/{id}/", deleteCourse).Methods("DELETE")

	//porting
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello GoLang Backend Developers</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if strconv.Itoa(course.CourseId) == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found for id" + params["id"])
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Name is empty")
		return
	}
	idMutex.Unlock()
	lastCourseID++
	course.CourseId = lastCourseID
	idMutex.Lock()
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if strconv.Itoa(course.CourseId) == params["id"] {
			var updateCourse Course
			json.NewDecoder(r.Body).Decode(&updateCourse)
			courses[i] = updateCourse
			json.NewEncoder(w).Encode(courses[i])
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this id")
	return
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if strconv.Itoa(course.CourseId) == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "No course found with this id"})
	return
}
