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


package main

import "fmt"

type anonymous struct{
	a float32
	int
	string
}

func main(){
	res := anonymous{1.34,24,"test"}
	fmt.Printf("the information: %v",res)
}