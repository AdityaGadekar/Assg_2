//Concurrent sum
package main
import "fmt"

func main(){
	sl1:= generatearray(1,10)
	sl2:= generatearray(11,20)
	sl3:= generatearray(21,30)
	sl4:= generatearray(31,40)

	resultChan :=make(chan int)

	go Sum(sl1, resultChan)
	go Sum(sl2, resultChan)
	go Sum(sl3, resultChan)
	go Sum(sl4, resultChan)

	fmt.Println("Total sum is", <-resultChan+<-resultChan+<-resultChan+<-resultChan)


}

func Sum(slice []int, ch chan int){
	sum:=0
	for _, n:=range slice{
		sum+=n
	}
	ch<-sum
	
}

func generatearray(a,b int) []int{
	sl := []int{}
	for i:=a;i<=b;i++{
		sl=append(sl,i)
	}
	return sl
}