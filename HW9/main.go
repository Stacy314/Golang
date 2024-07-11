/*Create a web server for viewing information about the school class.
The user must be able to receive general information about the class (list of students, name of the class).
Additional requirements:
• information about students must be stored in RAM and be available during each request;
• obtaining information about the student (for example, the average score in subjects) should be carried out by the 
GET method at the address "/student/{id}", where {id} is the student's unique identifier;
• data can only be retrieved if the user is a teacher in this class.*/


package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Student struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Grades    map[string]float64 `json:"grades"`
}

type Class struct {
	Name    string    `json:"name"`
	Students []Student `json:"students"`
	Teacher  string    `json:"teacher"`
}

var classData Class

func init() {
	classData = Class{
		Name:   "10-Б",
		Teacher: "teacher123",
		Students: []Student{
			{ID: "1", Name: "Ivan Ivanov", Grades: map[string]float64{"Math": 4.5, "Physics": 3.8}},
			{ID: "2", Name: "Petro Petrov", Grades: map[string]float64{"Math": 3.2, "Physics": 4.1}},
		},
	}
}

func isTeacher(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	return authHeader == "Bearer "+classData.Teacher
}

func classInfoHandler(w http.ResponseWriter, r *http.Request) {
	if !isTeacher(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classData)
}

func studentInfoHandler(w http.ResponseWriter, r *http.Request) {
	if !isTeacher(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/student/")
	for _, student := range classData.Students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/class", classInfoHandler)
	http.HandleFunc("/student/", studentInfoHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}