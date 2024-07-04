package main

import "fmt"

type Simple struct{
	num int
}

func (s *Simple) Set(numChange int) {
	s.num = numChange
}

func (s *Simple) Get() int{
	return s.num
}

type Simpler interface{
	Set(numChange int)
	Get() int
}

func main(){
	var s = Simple{num: 42}
	var er Simpler = &s
	fmt.Println(er.Get())
	er.Set(24)
	fmt.Println(er.Get())
}