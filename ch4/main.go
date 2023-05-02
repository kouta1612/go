package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// スライスが基底配列を参照していることを確認
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3]
	fmt.Printf("arr: %v\n", &arr[1])
	fmt.Printf("slice: %v\n", &slice[0])

	caps := make([]int, 2, 5)
	fmt.Println(len(caps), cap(caps), caps)

	employee := Employee{Name: "hagiwara", Salary: 80}
	mutable(&employee, 90)
	fmt.Println(employee)

	w := Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
			Radius: 20,
		},
	}

	bytes, _ := json.MarshalIndent(w, "", " ")
	fmt.Println(string(bytes))
}

func mutable(employee *Employee, salary int) {
	employee.Salary = salary
}
