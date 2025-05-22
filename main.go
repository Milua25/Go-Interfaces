package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
)

//type englishBot struct {
//}
//
//type spanishBot struct {
//}
//
//type bot interface {
//	getGreeting() string
//}

type logWriter struct{}

func main() {
	//eb := englishBot{}
	//sb := spanishBot{}

	//printGreeting(eb)
	//printGreeting(sb)

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("Errror: %v", err)
		os.Exit(1)
	}

	lw := logWriter{}

	_, err = io.Copy(lw, resp.Body)

	fmt.Println("=========================")

	trArea := triangle{
		base:   5.0,
		height: 34.68,
	}

	sqArea := square{
		sideLength: 50,
	}

	print(trArea)
	print(sqArea)

	fmt.Printf("%s\n", reflect.TypeOf(trArea))
	fmt.Printf("%s\n", reflect.TypeOf(sqArea))

	fmt.Println("=========================")

	filename := os.Args[1]

	openFile(filename)

	//if err != nil {
	//	fmt.Printf("Errror: %v\n", err)
	//	os.Exit(1)
	//}

	// custom Writer

}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("----------------")
	fmt.Printf("Just wrote this many bytes: %v\n", len(bs))

	return len(bs), nil
}

//func printGreeting(b bot) {
//	fmt.Println(b.getGreeting())
//}
//
//func (englishBot) getGreeting() string {
//	// very custom login for english greeting
//	return "Hi there!!!"
//}
//
//func (spanishBot) getGreeting() string {
//	// very custom login for spanish greeting
//	return "Hola!"
//}
