package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"unicode/utf8"
)

var (
	name   = "Allen"
	course = "Golang"
	module = 3.2
)

type user struct {
	name    string
	surname string
}

func main() {
	sampleMap()
}

func sampleMap() {
	users := map[string]user{
		"Carlos":  {"Carlos", "Roca"},
		"Anthony": {"Anthony", "Rojas"},
		"Keven":   {"Keven", "Saldaña"},
	}
	for key, val := range users {
		fmt.Println(key, val)
	}

	delete(users, "Keven")
	for key, val := range users {
		fmt.Println(key, val)
	}

	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}

func sampleCopySlice() {
	// using value semantic form of the for range.
	// output:
	// [Carlos]
	// [Keven]
	// [Anthony]
	// [Checho]
	friends := []string{"Carlos", "Keven", "Anthony", "Checho"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("[%s]\n", v)
	}

	// using the pointer semantic form of the for range.
	// output:
	// panic: runtime error: index out of range [2] with length 2
	friends = []string{"Carlos", "Keven", "Anthony", "Checho"}
	for i := range friends {
		friends = friends[:2]
		fmt.Printf("[%s]\n", friends[i])
	}
}

func sampleRangeString() {
	foo := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	for i, r := range foo {

		// Capture the number of bytes for this rune.
		rl := utf8.RuneLen(r)

		si := i + rl

		copy(buf[:], foo[i:si])

		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}

func sampleFunctions() {
	// using functions
	otherCourse := "reactive programming"
	otherAutor := "allen joseph"
	fmt.Println(formatCourseAndAutor(otherCourse, otherAutor))

	lowerValue := getLower(10, 13, 9, 17, 2, 21)
	fmt.Println("\nThe lower value is", lowerValue)
}

func sampleEnv() {
	// using enviroments
	user := os.Getenv("USER")
	fmt.Println("\nHi", user, "you're currently watching")
}

func samplePointer() {
	// using reference a pointer and de-reference a pointer
	ptr := &module
	fmt.Println("\nMemory address of *module* variable is", ptr, "and the value of *module* is", *ptr)

	// using pointers to change the original value
	changeCourse(&course)
	fmt.Println("You are now watching course", course)
}

func sampleAssignment() {
	// using assignment and initialization of  variables
	lastname := "Velasco Arias"
	fmt.Println("\nLastname is set to", lastname)
}

func sampleHello() {
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("module is set to", module, "and is of type", reflect.TypeOf(module))
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
