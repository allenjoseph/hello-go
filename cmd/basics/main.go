package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var (
	name   = "Allen"
	course = "Golang"
	module = 3.2
)

func main() {
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("module is set to", module, "and is of type", reflect.TypeOf(module))

	// using assignment and initialization of  variables
	lastname := "Velasco Arias"
	fmt.Println("\nLastname is set to", lastname)

	// using reference a pointer and de-reference a pointer
	ptr := &module
	fmt.Println("\nMemory address of *module* variable is", ptr, "and the value of *module* is", *ptr)

	// using pointers to change the original value
	changeCourse(&course)
	fmt.Println("You are now watching course", course)

	// using enviroments
	user := os.Getenv("USER")
	fmt.Println("\nHi", user, "you're currently watching")

	// using functions
	otherCourse := "reactive programming"
	otherAutor := "allen joseph"
	fmt.Println(formatCourseAndAutor(otherCourse, otherAutor))

	//using slice arrays as parameter function
	lowerValue := getLower(10, 13, 9, 17, 2, 21)
	fmt.Println("\nThe lower value is", lowerValue)
}

func changeCourse(course *string) string {
	*course = "Javascript"
	fmt.Println("\nTrying to change your course to", *course)
	return *course
}

func formatCourseAndAutor(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)
	return s1, s2
}

func getLower(values ...int) int {
	lower := values[0]
	for _, i := range values {
		if i < lower {
			lower = i
		}
	}
	return lower
}
