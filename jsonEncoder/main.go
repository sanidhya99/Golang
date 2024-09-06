package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name        string   `json:"course_name"` //you can mention the name with which the field will be addressed in json format
	Views       int      `json:"course_views"`
	Description string   `json:"course_description"`
	Password    string   `json:"course_password"`
	Tags        []string `json:"course_tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Encoding tutorial")
	// jsonEncoder()
	jsonDecoder()
}

func jsonEncoder() {
	lcoCourses := []course{
		{"Django", 120, "Good tutorial of all concepts of Django", "009@NIThmsecoder", []string{"Django", "Django Rest Framework", "API Development", "REST API"}},
		{"Docker", 135, "Good tutorial of all concepts of Docker", "009@NIThmsecoder", []string{"Docker", "DevOps", "Containerization", "Virtualization", "CI/CD/CD"}},
		{"AWS", 150, "Good tutorial of all concepts of AWS", "009@NIThmsecoder", []string{"AWS", "Cloud Computing", "DevOps"}},
		{"Django", 120, "Good tutorial of all concepts of Django", "009@NIThmsecoder", nil}}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //in this second is prefix which has no use and 3rd is break point for indent but most of time in this we have to use '\t' only
	// finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s/n", finalJson)
}

func jsonDecoder() {
	jsdt := []byte(`
	{
			"course_name": "Django",
			"course_views": 120,
			"course_description": "Good tutorial of all concepts of Django",
			"course_password": "009@NIThmsecoder",
			"course_tags":["Django","Django Rest Framework"]
	}`)

	// var onlineCourse course
	// is_valid := json.Valid(jsdt)
	// if is_valid {
	// 	fmt.Println("JSON is valid")
	// 	json.Unmarshal(jsdt, &onlineCourse)
	// 	fmt.Printf("%#v\n", onlineCourse)
	// } else {
	// 	fmt.Println("JSON is not valid")
	// }

	var onlineCoursedt map[string]interface{}
	json.Unmarshal(jsdt, &onlineCoursedt)
	fmt.Printf("%#v\n", onlineCoursedt)
	for k, v := range onlineCoursedt {
		fmt.Printf("Key is %v & Value is %v\n", k, v)
	}
}
