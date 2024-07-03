// package main;

// if err != nil {
//   fmt.Println(err)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var orig string = "ABC"
// 	// var an int
// 	var newS string
// 	// var err error

// 	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
// 	// anInt, err = strconv.Atoi(origStr)
// 	an, err := strconv.Atoi(orig)
// 	if err != nil {
// 		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
// 		return
// 	}
// 	fmt.Printf("The integer is %d\n", an)
// 	an = an + 5
// 	newS = strconv.Itoa(an)
// 	if err != nil {
// 		fmt.Printf("The new string is: %s\n", newS)
// }
// }

// package main

// import(
// 	"fmt"
// )

// func main(){
// 	k := 6
// 	switch k {
// 		case 4: fmt.Println("was <= 4"); fallthrough;
// 		case 5: fmt.Println("was <= 5"); fallthrough;
// 		case 6: fmt.Println("was <= 6"); fallthrough;
// 		case 7: fmt.Println("was <= 7"); fallthrough;
// 		case 8: fmt.Println("was <= 8"); fallthrough;
// 		default: fmt.Println("default case")
// 	}
// }