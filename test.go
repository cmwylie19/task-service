package main

import (
	"fmt"
	//"time"
)

type Request struct {
	hostname string
	path string
//	timestamp time
}

type Target struct {
	address string
	state uint
}
func NetworkLoadBalancer(req Request) []Request {
	targets := make([]Target, 3)
	for i, target := range targets {
		target = req
		fmt.Println(target," ",i)
	} 
	return targets
}


func main() {
	req := Request{
		"localhost","/healthz",
	}
	fmt.Println("Targets:",NetworkLoadBalancer(req))
//	var a [5]int
//	for i:=0;i<len(a);i++ {
//		a[i]=i
//		fmt.Println(a)
//	}
}
