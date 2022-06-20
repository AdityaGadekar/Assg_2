package main

import "fmt"

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go Prod1(ch1)
	go Prod2(ch2)
	go combine(ch1,ch2,ch3)
    for i:=range ch3{
		fmt.Println("Comsumed",i)
	}
}
func Prod1(ch1 chan int) {
	for i := 0; i < 11; i++ {
		ch1 <- i + 1
		fmt.Println("Produced",i+1)
	}
	close(ch1)
}

func Prod2(ch2 chan int) {
	for i := 11; i < 21; i++ {
		ch2 <- i + 1
		fmt.Println("Produced",i+1)
	}
	close(ch2)
}
func combine(ch1, ch2, ch3 chan int){
	for v:=range ch1{
		ch3 <- v
	}
	for v1:=range ch2{
		ch3 <- v1
	}
	close(ch3)
}