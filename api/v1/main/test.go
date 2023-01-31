package main

import "fmt"

type Person struct {
	name string
	age  int
	car  Car
}
type Car struct {
	name string
}

var personMap map[string]Person

func setName(per Person, name string) {
	per.name = name
}
func setAge(per *Person, age int) {
	per.age = age
}
func setCar(per *Person, car Car) {
	per.car = car

}
func (per Person) print() {
	fmt.Println(per.car)
}

func main() {
	per := Person{}
	per.name = "www"
	println(per.name)
	setName(per, "sss")
	println(per.name)
	per.car = Car{"bbw"}
	setCar(&per, Car{"bc"})
	per.print()
	personMap = make(map[string]Person)
	personMap["per1"] = per
	per.name = "aaa"
	for _, value := range personMap {
		fmt.Println(value)
	}
}
