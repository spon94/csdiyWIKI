// package main;

// import "fmt"

// type Address struct{
// 	street string
// 	city string
// 	state string
// 	zip int
// }

// type VCard struct{
// 	address *Address
// 	name string
// 	birth string
// 	image string
// }

// func main() {
// 	xiaoMing := VCard{
// 		address: &Address{
// 			street: "123 Main Street",
// 			city: "Anytown",
// 			state: "CA",
// 			zip: 12345,
// 		},
// 		name: "Xiao Ming",
// 		birth: "1990-01-01",
// 		image: "https://example.com/xiao-ming.jpg",
// 	}
// 	fmt.Printf("the information of xiaoMing %v",xiaoMing)
// }


// package main

// import "fmt"

// type anonymous struct{
// 	a float32
// 	int
// 	string
// }

// func main(){
// 	res := anonymous{1.34,24,"test"}
// 	fmt.Printf("the information: %v",res)
// }


// package main

// import "fmt"

// type Car struct{
// 	wheelCount int
// }

// type Mercedes struct{
// 	Car
// }

// func (c *Car) numberOfWheels() int {
// 	return c.wheelCount
// }

// func (m *Mercedes) sayHiToMerkel() int{
// 	return m.numberOfWheels()*1000
// }

// func main(){
// 	var carOne = Mercedes{Car{wheelCount:4}}
// 	fmt.Printf("the wheel count: %d\n",carOne.numberOfWheels())
// 	fmt.Printf("the output: %d\n",carOne.sayHiToMerkel())
// }


// package main

// import "fmt"

// type Base struct{
// 	id int
// }

// type Person struct{
// 	Base
// 	FirstName string
// 	LastName string
// }

// type Employee struct{
// 	Person
// 	Salary int
// }

// func (b *Base) Id() int {
// 	return b.id
// }

// func (b *Base) SetId(num int) {
// 	b.id = num
// }

// func main(){
// 	var employeeOne = Employee{Person{Base{1},"John","Doe"},1000}
// 	fmt.Printf("the id is: %d", employeeOne.Id())
// }


// package main

// import "fmt"

// type Integer int

// func (p Integer) get() int{
// 	return int(p)
// }

// func main(){
// 	var i Integer = 1
// 	fmt.Printf("the value is: %d",i.get())
// }

// package main

// import "fmt"

// type Integer struct{
// 	n int
// }

// func (p Integer) get() int {
// 	return p.n
// }

// func main(){
// 	var i = Integer{1}
// 	fmt.Printf("the value is: %d",i.get())
// }


package main

import "fmt"
import "strconv"

type T struct {
    a int
    b float64
    c string
}

func (t *T) String() string{
	return (strconv.Itoa(t.a) + "/" + strconv.FormatFloat(t.b, 'f', 2, 32) + "/" + t.c)
}

func main(){
	var test = T{7,-2.35,"abc\\tdef"}
	fmt.Printf("%v",test)
}