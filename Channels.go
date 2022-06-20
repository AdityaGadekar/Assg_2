package main

import "fmt"

func main(){
	ch:=GetChannel()
	for i := 0; i < 10; i++ {
		ch <- i + 1
		fmt.Println("Produced", i+1)
	}
	close(ch)	
}

func GetChannel() chan int {
	ch := make(chan int)
	go func() {
		for v:=range ch{
			fmt.Println("Consumed",v)
		}
	}()
	return ch
}

