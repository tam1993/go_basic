package mypackage

import (
	"fmt"
)

func Function(a, b int) int {
	if b == 0 {
		return 0
	}
	return a * b
}

func ForLoop() {
	//for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// //do while
	// i := 1
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 10 {
	// 		break
	// 	}
	// }
	// // while loop
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
}

func Array() {
	// //array
	// // var myArray [3]int	//no initial data
	// myArray := [3]int{10, 20, 30} //initial data
	// myArray[0] = 11
	// fmt.Println(myArray)

	// //slice
	// mySlice := []int{}
	// mySlice = append(mySlice, 10, 20)
	// fmt.Println(mySlice)
	// mySubSlice := mySlice[0:1]
	// fmt.Println(mySubSlice)

	//map
	myMap := make(map[string]int)
	myMap["apple"] = 10
	myMap["banana"] = 20
	fmt.Println(myMap)
	//delete
	delete(myMap, "banana")
	fmt.Println(myMap)
	//check key
	val, found := myMap["apple"]
	if found {
		fmt.Println(val)
	}

	for k, v := range myMap {
		fmt.Println(k, "->", v)
	}
}

// struct
type Person struct {
	Name    string
	Address Address
}
type Address struct {
	Province string
}

func Struct() {
	Person := Person{
		Name: "namee",
		Address: Address{
			Province: "cnx",
		},
	}
	fmt.Println(Person)
}

func ForArrayStruct() {
	Persons := []Person{}
	for i := 0; i < 5; i++ {
		Persons = append(Persons, Person{
			Name: fmt.Sprint("name", i),
			Address: Address{
				Province: fmt.Sprint("pv", i),
			},
		})
	}
	fmt.Println(Persons)
}
func FunctionStruct() {
	Person := Person{
		Name: "name",
		Address: Address{
			Province: "cnx",
		},
	}
	fmt.Println(Person.PersonAddress())
}
func (p Person) PersonAddress() string {
	return p.Name + "@" + p.Address.Province
}

// pointer
func Pointer() {
	num := 10
	fmt.Println(num)
	ChangeValue(&num)
	fmt.Println(num)

	num2 := &num
	*num2 = 30
	fmt.Println(num)
}
func ChangeValue(x *int) {
	*x = 20
}

func PointerStruct() {
	Person := Person{
		Name: "name",
		Address: Address{
			Province: "cnx",
		},
	}
	fmt.Println(Person)
	ChangeValueStruct1(&Person)
	fmt.Println(Person)
	Person.ChangeValueStruct2()
	fmt.Println(Person)
}
func ChangeValueStruct1(p *Person) {
	p.Name = "name2"
	p.Address.Province = "cnx2"
}
func (p *Person) ChangeValueStruct2() {
	p.Name = "name3"
	p.Address.Province = "cnx3"
}

// interface
type Dog struct {
	Name string
}

type Human struct {
	Name string
}

func (d Dog) Speak() string {
	return "Wool"
}

func (h Human) Speak() string {
	return "Hello"
}

type Speaker interface {
	Speak() string
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func Interface() {
	Dog := Dog{Name: "Dogg"}
	Human := Human{Name: "Human"}

	makeSound(Dog)
	makeSound(Human)
}
