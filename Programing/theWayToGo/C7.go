// package main

// import "fmt"

// func main() {
// 	var fibos [51]int
// 	fibos[1] = 1
// 	fmt.Printf("index: %d, value: %d\n", 1, fibos[1])
// 	for i := 2; i < 51; i++ {
// 		fibos[i] = fibos[i-1] + fibos[i-2]
// 		fmt.Printf("index: %d, value: %d\n", i, fibos[i])
// 	}
// }


package main

import "fmt"

func main(){
	var fibos []int = make([]int, 51)
	calcu(fibos)
	for i := 0; i < 51; i++ {
		fmt.Printf("index: %d, value: %d\n",i,fibos[i])
	}
}

func calcu(num []int){
	num[0],num[1]=1,1
	for i:=2; i < 51; i++{
		num[i] = num[i-1] + num[i-2]
	}
	return
}