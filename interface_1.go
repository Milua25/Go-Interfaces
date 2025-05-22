package main

import (
	"fmt"
	"io"
	"os"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func print(sh shape) {
	fmt.Printf("The area of the shape is %v\n", sh.getArea())
}

func openFile(filename string) {
	data, err := os.Open(filename)
	if err != nil {
		fmt.Printf("unable to open filename %v", filename)
		os.Exit(1)
	}

	//io.Copy(os.Stdout, data)
	bs, _ := io.ReadAll(data)
	fmt.Println(string(bs))

}
