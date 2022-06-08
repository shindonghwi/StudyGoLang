package main

type person struct {
	name string
	age  int
	job  string
}

func main() {

	var myList []person
	println(myList)

	s1 := person{name: "DongHwi1", age: 31, job: "Developer1"}
	s2 := person{name: "DongHwi2", age: 32, job: "Developer2"}
	s3 := person{name: "DongHwi3", age: 33, job: "Developer3"}

	myList = append(myList, s1)
	myList = append(myList, s2)
	myList = append(myList, s3)

	for i, p := range myList {
		println("hello: ", i, p.name)
	}
}
