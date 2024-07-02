// Chapter 6.6
// package main

// import "fmt"

// func main() {

// 	for i := 0; i <= 10; i++ {
// 		pos, val := fibonaci(i)
// 		fmt.Printf("index: %d,val: %d\n", pos, val)
// 	}
// }

// func fibonaci(num int) (pos int, val int) {
// 	if num <= 1 {
// 		return num, 1
// 	} else {
// 		_, tempNum_1 := fibonaci(num - 1)
// 		_, tempNum_2 := fibonaci(num - 2)
// 		return num, tempNum_1 + tempNum_2
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	numPrint(10)
// }

// func numPrint(num int) {
// 	if num == 1 {
// 		fmt.Printf("%d\n", num)
// 		return
// 	} else{
// 		numPrint(num - 1)
// 		fmt.Printf("%d\n", num)
// 	}
// }


// package main

// import "fmt"

// func main(){
// 	for i := 2; i <= 10; i++{
// 		fmt.Println("%d",multi(i))
// 	}
// }

// func multi (num int) int{
// 	if num <= 1{
// 		return 1
// 	} else{
// 		return num * multi(num - 1)
// 	}
// }

// package main

// import "fmt"

// func main() {
//     f()
// }
// func f() {
//     for i := 0; i < 4; i++ {
//         g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
//         g(i)
//         fmt.Printf(" - g is of type %T and has value %v\n", g, g)
//     }
// }


package main

import "fmt"

func main(){
	res := fibonaci()
	for i := 0; i <= 10; i ++{
		fmt.Println("index: %d, value: %d", i, res(i))
	}
}

func fibonaci(num int) func(int) int {
	var sum int
	return func(delta int) int{
		sum += delta
		if delta <= 0{
			return 1
		} else{
			return sum
		}
	}
}